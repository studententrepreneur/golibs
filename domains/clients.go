package domain

type Clients struct {
	Id           string       `firestore:"id" json:"id"`
	ClientId     string       `firestore:"clientId" json:"clientId"`
	ClientSecret string       `firestore:"clientSecret" json:"clientSecret"`
	Name         string       `firestore:"name" json:"name"`
	Description  string       `firestore:"description" json:"description"`
	IsEnabled    bool         `firestore:"isEnabled" json:"isEnabled"`
	Application  *Application `firestore:"application" json:"application"`
	Roles        []*Roles     `firestore:"roles" json:"roles"`
}
