package main

import (
	"time"
)

// Model represents a base database type
type Model struct {
  ID        uint        `json:"id" form:"id" bson:"_id,omitempty"`
  CreatedOn time.Time
  UpdatedOn time.Time
  DeletedOn *time.Time
}
