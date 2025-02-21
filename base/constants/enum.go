package constants

// Represents allowed task statuses
type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusInitiated TaskStatus = "initiated"
	StatusCompleted TaskStatus = "completed"
	StatusCancelled TaskStatus = "cancelled"
)

// Checks if the status is a valid enum value
func (instance TaskStatus) IsValid() bool {
	switch instance {
	case StatusPending, StatusInitiated, StatusCompleted, StatusCancelled:
		return true
	default:
		return false
	}
}
