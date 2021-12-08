package xmlWriter

import (
	"bytes"

	"github.com/shabbyrobe/xmlwriter"

	"github.com/anfilat/go-ews/internal/enumerations/xmlNamespace"
	"github.com/anfilat/go-ews/internal/errors"
	"github.com/anfilat/go-ews/internal/utils"
)

type writer struct {
	IsTimeZoneHeaderEmitted bool

	b  *bytes.Buffer
	w  *xmlwriter.Writer
	ec *xmlwriter.ErrCollector
}

func New() *writer {
	b := &bytes.Buffer{}
	w := xmlwriter.Open(b, xmlwriter.WithIndent())
	ec := &xmlwriter.ErrCollector{}

	return &writer{
		b:  b,
		w:  w,
		ec: ec,
	}
}

func (w *writer) Err() error {
	if w.ec.Err != nil {
		return errors.NewWriteXMLError(w.ec)
	}
	return nil
}

func (w *writer) Bytes() []byte {
	return w.b.Bytes()
}

func (w *writer) Flush() {
	w.ec.Do(w.w.Flush())
}

func (w *writer) WriteDoc() {
	w.ec.Do(w.w.StartDoc(xmlwriter.Doc{}))
}

func (w *writer) WriteEndDoc() {
	w.ec.Do(w.w.EndDoc())
}

func (w *writer) WriteStartElement(namespace xmlNamespace.Enum, localName string) {
	w.ec.Do(w.w.StartElem(xmlwriter.Elem{
		Prefix: utils.GetNamespacePrefix(namespace),
		Name:   localName,
		URI:    utils.GetNamespaceUri(namespace),
	}))
}

func (w *writer) WriteEndElement() {
	w.ec.Do(w.w.EndElem())
}

func (w *writer) WriteAttributeValueBool(localName string, alwaysWriteEmptyString bool, value interface{}) {
	val, ok := w.tryConvertObjectToString(value)
	if !ok {
		w.ec.Do(errors.NewSerializationError(value, localName))
	}
	if alwaysWriteEmptyString || val != "" {
		w.writeAttributeString("", localName, val)
	}
}

func (w *writer) WriteAttributeValue(namespacePrefix, localName string, value interface{}) {
	val, ok := w.tryConvertObjectToString(value)
	if !ok {
		w.ec.Do(errors.NewSerializationError(value, localName))
	}
	if val != "" {
		w.writeAttributeString(namespacePrefix, localName, val)
	}
}

func (w *writer) writeAttributeString(namespacePrefix, localName, value string) {
	w.ec.Do(w.w.WriteAttr(xmlwriter.Attr{
		Prefix: namespacePrefix,
		URI:    "",
		Name:   localName,
		Value:  value,
	}))
}

func (w *writer) tryConvertObjectToString(value interface{}) (string, bool) {
	switch val := value.(type) {
	case string:
		return val, true
	}
	return "", false
}
