package main

// Venue represents a place where events can occur.
type Venue struct {
	Model
	Google     string `json:"google" form:"google"`
	Foursquare string `json:"foursquare" form:"foursquare"`
	Yelp       string `json:"yelp" form:"yelp"`
}
