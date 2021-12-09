package privilegedLogonType

// Enum defines the type of PrivilegedLogonType.
type Enum int

func (e Enum) String() string {
	switch e {
	case Admin:
		return "Admin"
	case SystemService:
		return "SystemService"
	}
	return ""
}

const (
	// Admin - Logon as Admin
	Admin Enum = iota
	// SystemService - Logon as SystemService
	SystemService
)
