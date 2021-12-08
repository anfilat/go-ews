package requests

type Request interface {
	Validate() error
	GetXmlElementName() string
}
