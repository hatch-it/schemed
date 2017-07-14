package main

// Resource objects appear in a JSON API to represent resources.
// http://jsonapi.org/format/#document-resource-objects
type Resource struct {
	ID            string      `json="user"`
	Type          string      `json="type" binding:"required"`
	Attributes    interface{} `json="attributes"`
	Relationships interface{} `json="relationships"`
	Links         interface{} `json="links"`
	Meta          interface{} `json="meta"`
}
