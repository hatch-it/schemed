package schemed

import "time"

type Model struct {
	id        string
	createdAt time.Time
	updatedAt time.Time
	deletedAt time.Time
}
