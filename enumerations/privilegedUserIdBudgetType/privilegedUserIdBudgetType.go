package privilegedUserIdBudgetType

// Enum is PrivilegedUserId BudgetType enum
type Enum int

func (e Enum) String() string {
	switch e {
	case Default:
		return "Default"
	case RunningAsBackgroundLoad:
		return "RunningAsBackgroundLoad"
	case Unthrottled:
		return "Unthrottled"
	}
	return ""
}

const (
	// Default for interactive, charge against a copy of target mailbox budget.
	Default Enum = iota
	// RunningAsBackgroundLoad for running as background load
	RunningAsBackgroundLoad
	// Unthrottled for unthrottled budget
	Unthrottled
)
