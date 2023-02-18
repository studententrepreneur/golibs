package domain

type Profiles struct {
	Network  string `json:"network"`
	Username string `json:"username"`
	URL      string `json:"url"`
}

type Location struct {
	Id              string `json:"id" example:"asdasdsa1"`
	Address         string `json:"address" example:"Acme corp"`
	PostalCode      string `json:"postalCode" example:"123456"`
	CountryCode     string `json:"countryCode" example:"LD"`
	Region          string `json:"region" example:"LD"`
	City            string `json:"city" example:"palakkad"`
	VisaSponsorship bool   `json:"visaSponsorship"`
}
