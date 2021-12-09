package availabilityData

// Enum defines the type of data that can be requested via GetUserAvailability.
type Enum int

func (e Enum) String() string {
	switch e {
	case FreeBusy:
		return "FreeBusy"
	case Suggestions:
		return "Suggestions"
	case FreeBusyAndSuggestions:
		return "FreeBusyAndSuggestions"
	}
	return ""
}

const (
	// FreeBusy only return free/busy data.
	FreeBusy Enum = iota
	// Suggestions only return suggestions.
	Suggestions
	// FreeBusyAndSuggestions return both free/busy data and suggestions.
	FreeBusyAndSuggestions
)
