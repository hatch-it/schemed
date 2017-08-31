package models

// Venue represents a place where events can occur.
type Venue struct {
	Model
	GoogleID     string `json:"googleId" form:"googleId"`
	FoursquareID string `json:"foursquareId" form:"foursquareId"`
	YelpID       string `json:"yelpId" form:"yelpId"`
}
