package dateTimePrecision

// Enum defines the precision for returned DateTime values
type Enum int

func (e Enum) String() string {
	switch e {
	case Default:
		return "Default"
	case Seconds:
		return "Seconds"
	case Milliseconds:
		return "Milliseconds"
	}
	return ""
}

const (
	// Default value.  No SOAP header emitted.
	Default Enum = iota
	Seconds
	Milliseconds
)
