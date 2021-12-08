package ewsError

import (
	"errors"
)

var (
	Validate      = errors.New("validation error")
	WriteXML      = errors.New("write XML error")
	Serialization = errors.New("value cannot be serialized")
)
