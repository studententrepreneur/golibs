package domain

type Profiles struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Location struct {
	Address         string `json:"address" example:"Acme corp"`
	PostalCode      string `json:"postalCode" example:"123456"`
	CountryCode     string `json:"countryCode" example:"LD"`
	Region          string `json:"region" example:"LD"`
	VisaSponsorship bool   `json:"visaSponsorship"`
}
