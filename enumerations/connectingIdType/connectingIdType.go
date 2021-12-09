package connectingIdType

// Enum defines the type of Id of a ConnectingId object.
type Enum int

func (e Enum) String() string {
	switch e {
	case PrincipalName:
		return "PrincipalName"
	case SID:
		return "SID"
	case SmtpAddress:
		return "SmtpAddress"
	}
	return ""
}

const (
	// PrincipalName for a principal name
	PrincipalName Enum = iota
	// SID for an SID
	SID
	// SmtpAddress for an SMTP address
	SmtpAddress
)
