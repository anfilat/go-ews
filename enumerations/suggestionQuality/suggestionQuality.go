package suggestionQuality

// Enum defines the quality of an availability suggestion.
type Enum int

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
