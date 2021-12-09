package suggestionQuality

// Enum defines the quality of an availability suggestion.
type Enum int

func (e Enum) String() string {
	switch e {
	case Excellent:
		return "Excellent"
	case Good:
		return "Good"
	case Fair:
		return "Fair"
	case Poor:
		return "Poor"
	}
	return ""
}

const (
	// Excellent - The suggestion is excellent.
	Excellent Enum = iota
	// Good - The suggestion is good.
	Good
	// Fair - The suggestion is fair.
	Fair
	// Poor - The suggestion is poor.
	Poor
)
