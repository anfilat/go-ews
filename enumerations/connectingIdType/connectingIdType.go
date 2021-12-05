package connectingIdType

// Enum defines the type of Id of a ConnectingId object.
type Enum int

const (
	// PrincipalName for a principal name
	PrincipalName Enum = iota
	// SID for an SID
	SID
	// SmtpAddress for an SMTP address
	SmtpAddress
)
