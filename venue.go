package main

// Venue represents a place where events can occur.
type Venue struct {
	Model

	Google     string `json:"google"`
	Foursquare string `json:"foursquare"`
	Yelp       string `json:"yelp"`
}
