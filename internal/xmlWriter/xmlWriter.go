package xmlWriter

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/shabbyrobe/xmlwriter"

	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/utils"
)

type Writer struct {
	b  *bytes.Buffer
	w  *xmlwriter.Writer
	ec *xmlwriter.ErrCollector

	rootLevel bool
	rootUris  []string
}

func New() *Writer {
	b := &bytes.Buffer{}
	w := xmlwriter.Open(b, xmlwriter.WithIndent())
	ec := &xmlwriter.ErrCollector{}

	return &Writer{
		b:  b,
		w:  w,
		ec: ec,
	}
}

func (w *Writer) Err() error {
	if w.ec.Err != nil {
		return errors.NewWriteXMLError(w.ec)
	}
	return nil
}

func (w *Writer) Bytes() []byte {
	return w.b.Bytes()
}

func (w *Writer) Flush() {
	w.ec.Do(w.w.Flush())
}

func (w *Writer) WriteDoc() {
	w.ec.Do(w.w.StartDoc(xmlwriter.Doc{}))
}

func (w *Writer) WriteEndDoc() {
	w.ec.Do(w.w.EndDoc())
}

func (w *Writer) WriteStartElement(namespace xmlNamespace.Enum, localName string) {
	prefix := utils.GetNamespacePrefix(namespace)
	uri := utils.GetNamespaceUri(namespace)

	if w.checkRootUri(prefix, uri) {
		uri = ""
	} else {
		w.pushUris(prefix, uri)
	}

	w.ec.Do(w.w.StartElem(xmlwriter.Elem{
		Prefix: prefix,
		Name:   localName,
		URI:    uri,
	}))
}

func (w *Writer) WriteEndElement() {
	w.ec.Do(w.w.EndElemFull())
}

func (w *Writer) WriteElementValue(namespace xmlNamespace.Enum, localName string, value interface{}) {
	if utils.IsNil(value) {
		return
	}

	val, ok := w.tryConvertObjectToString(value)
	if !ok {
		w.ec.Do(errors.NewValueSerializationError(value, localName))
		return
	}

	w.WriteStartElement(namespace, localName)
	w.WriteValue(val)
	w.WriteEndElement()
}

func (w *Writer) WriteValue(value string) {
	w.ec.Do(w.w.WriteText(value))
}

func (w *Writer) WriteAttributeValueBool(localName string, alwaysWriteEmptyString bool, value interface{}) {
	val, ok := w.tryConvertObjectToString(value)
	if !ok {
		w.ec.Do(errors.NewAttrSerializationError(value, localName))
		return
	}
	if alwaysWriteEmptyString || val != "" {
		w.WriteAttributeString("", localName, val)
	}
}

func (w *Writer) WriteAttributeValue(namespacePrefix, localName string, value interface{}) {
	val, ok := w.tryConvertObjectToString(value)
	if !ok {
		w.ec.Do(errors.NewAttrSerializationError(value, localName))
		return
	}
	if val != "" {
		w.WriteAttributeString(namespacePrefix, localName, val)
	}
}

func (w *Writer) WriteAttributeString(namespacePrefix, localName, value string) {
	w.ec.Do(w.w.WriteAttr(xmlwriter.Attr{
		Prefix: namespacePrefix,
		URI:    "",
		Name:   localName,
		Value:  value,
	}))

	w.pushUris(localName, value)
}

func (w *Writer) StartRoot() {
	w.rootLevel = true
}

func (w *Writer) EndRoot() {
	w.rootLevel = false
}

func (w *Writer) pushUris(prefix string, uri string) {
	if w.rootLevel {
		w.rootUris = append(w.rootUris, prefix+":"+uri)
	}
}

func (w *Writer) checkRootUri(prefix string, uri string) bool {
	for _, val := range w.rootUris {
		if val == prefix+":"+uri {
			return true
		}
	}
	return false
}

func (w *Writer) tryConvertObjectToString(value interface{}) (string, bool) {
	switch val := value.(type) {
	case fmt.Stringer:
		return val.String(), true
	case bool:
		if val {
			return "true", true
		}
		return "false", true
	case int:
		return strconv.FormatInt(int64(val), 10), true
	case int8:
		return strconv.FormatInt(int64(val), 10), true
	case int16:
		return strconv.FormatInt(int64(val), 10), true
	case int32:
		return strconv.FormatInt(int64(val), 10), true
	case int64:
		return strconv.FormatInt(val, 10), true
	case uint:
		return strconv.FormatUint(uint64(val), 10), true
	case uint8:
		return strconv.FormatUint(uint64(val), 10), true
	case uint16:
		return strconv.FormatUint(uint64(val), 10), true
	case uint32:
		return strconv.FormatUint(uint64(val), 10), true
	case uint64:
		return strconv.FormatUint(val, 10), true
	case float32:
		return strconv.FormatFloat(float64(val), 'f', -1, 64), true
	case float64:
		return strconv.FormatFloat(val, 'f', -1, 64), true
	case string:
		return val, true
	}
	return "", false
}
