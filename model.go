package main

import (
	"time"
	"gopkg.in/mgo.v2/bson"
)

// Model represents a base database type
type Model struct {
  ID        bson.ObjectId `json:"id" form:"id" bson:"_id,omitempty"`
  CreatedOn time.Time     `json:"createdOn" form:"createdOn"`
  UpdatedOn time.Time     `json:"updatedOn" form:"updatedOn"`
  DeletedOn *time.Time    `json:"deletedOn" form:"deletedOn"`
}


// ModelFilters contains all the possible filters for a Model
type ModelFilters struct {
  ID        *bson.ObjectId  `json:"id,omitempty" form:"id,omitempty" bson:"_id,omitempty"`
  CreatedOn *time.Time      `json:"createdOn,omitempty" form:"createdOn,omitempty"`
  UpdatedOn *time.Time      `json:"updatedOn,omitempty" form:"updatedOn,omitempty"`
  DeletedOn *time.Time      `json:"deletedOn,omitempty" form:"deletedOn,omitempty"`
}