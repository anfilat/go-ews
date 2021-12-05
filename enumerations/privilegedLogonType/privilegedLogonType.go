package privilegedLogonType

// Enum defines the type of PrivilegedLogonType.
type Enum int

const (
	// Admin - Logon as Admin
	Admin Enum = iota
	// SystemService - Logon as SystemService
	SystemService
)
