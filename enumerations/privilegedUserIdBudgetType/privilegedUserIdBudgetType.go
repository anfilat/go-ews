package privilegedUserIdBudgetType

// Enum is PrivilegedUserId BudgetType enum
type Enum int

const (
	// Default for interactive, charge against a copy of target mailbox budget.
	Default Enum = iota
	// RunningAsBackgroundLoad for running as background load
	RunningAsBackgroundLoad
	// Unthrottled for unthrottled budget
	Unthrottled
)
