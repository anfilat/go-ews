package availabilityData

// Enum defines the type of data that can be requested via GetUserAvailability.
type Enum int

const (
	// FreeBusy only return free/busy data.
	FreeBusy Enum = iota
	// Suggestions only return suggestions.
	Suggestions
	// FreeBusyAndSuggestions return both free/busy data and suggestions.
	FreeBusyAndSuggestions
)
