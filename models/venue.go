package models

type Venue struct {
	schemed.Model

	google     string
	foursquare string
	yelp       string
}
