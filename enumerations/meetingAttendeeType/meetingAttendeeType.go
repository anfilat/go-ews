package meetingAttendeeType

// Enum defines the type of a meeting attendee.
type Enum int

func (e Enum) String() string {
	switch e {
	case Organizer:
		return "Organizer"
	case Required:
		return "Required"
	case Optional:
		return "Optional"
	case Room:
		return "Room"
	case Resource:
		return "Resource"
	}
	return ""
}

const (
	// Organizer for the attendee is the organizer of the meeting
	Organizer Enum = iota
	// Required for the attendee is required
	Required
	// Optional the attendee is optional
	Optional
	// Room for the attendee is a room
	Room
	// Resource for the attendee is a resource
	Resource
)
