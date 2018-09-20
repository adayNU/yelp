package yelp

// SearchResponse represents a response from the Business Search endpoint.
type SearchResponse struct {
	Total      int
	Businesses []Business
	Region     map[string]map[string]float64
}

// Business represents a Yelp business.
type Business struct {
	ID, // Unique Yelp ID of this business. Example: '4kMBvIEWPxWkWKFN__8SxQ'
	Alias, // Unique Yelp alias of this business. Can contain unicode characters. Example: 'yelp-san-francisco'.
	Name, // Name of this business.
	URL, // URL for business page on Yelp.
	Price, // Price level of the business. Value is one of $, $$, $$$ and $$$$.
	Phone string // Phone number of the business.
	// URL of photo for this business.
	ImageURL string `json:"image_url"`
	// Rating for this business (value ranges from 1, 1.5, ... 4.5, 5).
	Rating float64
	// Number of reviews for this business.
	ReviewCount int `json:"review_count"`
	// A list of category title and alias pairs associated with this business.
	Categories []struct{ Alias, Title string }
	// The coordinates of this business.
	// Keys are "latitude" and "longitude".
	Coordinates map[string]float64
	// The location of this business, including address, city, state, zip code and country.
	Location location
	//A list of Yelp transactions that the business is registered for.
	// Current supported values are "pickup", "delivery", and "restaurant_reservation".
	Transactions []string

	// Only populated by search calls.

	// Distance in meters from the search location. This returns meters regardless of the locale.
	Distance float64

	// Only populated by business look-ups.

	// Whether business has been claimed by a business owner.
	IsClaimed bool `json:"is_claimed"`
	// Whether business has been (permanently) closed.
	IsClosed bool `json:"is_closed"`
	// URLs of up to three photos of the business.
	Photos []string
	// Opening hours of the business.
	Hours []hours
}

// location represents a business' location.
type location struct {
	Address1, // Street address of this business.
	Address2, // Street address of this business, continued.
	Address3, // Street address of this business, continued.
	City, // City of this business.
	Country, // ISO 3166-1 alpha-2 country code of this business.
	State string // ISO 3166-2 (with a few exceptions) state code of this business.
	// Cross streets for this business.
	CrossStreets string `json:"cross_streets"`
	// Zip code of this business.
	ZipCode string `json:"zip_code"`
	// Array of strings that if organized vertically give an address
	// that is in the standard address format for the business's country.
	DisplayAddress []string `json:"display_address"`
}

// hours represents the opening hours of a buiness.
type hours struct {
	// Currently only "REGULAR".
	HoursType string `json:"hours_type"`
	// The detailed opening hours of each day in a week.
	Open []struct {
		// Whether the business opens overnight or not.
		// When this is true, the end time will be lower than the start time.
		IsOvernight bool `json:"is_overnight"`
		// Start/End of the open hours in a day, in 24-hour clock notation.
		// Eg. 1000 means 10 AM.
		End, Start string
		// Day of the week.
		Day int
	}
	// Whether the business is currently open or not.
	IsOpenNow bool `json:"is_open_now"`
}
