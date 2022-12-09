package domain

type Organization struct {
	Id               string                 `firestore:"id" json:"id"`
	Name             string                 `firestore:"name" json:"name"`
	Description      string                 `firestore:"description" json:"description"`
	Meta             map[string]interface{} `firestore:"meta" json:"meta,omitempty"`
	AdvancedSettings AdvancedSettings       `firestore:"advancedSettings" json:"advancedSettings"`
	Branding         Branding               `firestore:"branding" json:"branding"`
	Logs             AbstractLog            `firestore:"logs" json:"logs"`
	SharedWith       []*Credentials         `firestore:"sharedWith" json:"sharedWith"`
}

type Branding struct {
	LogoURL      string                 `firestore:"logourl" json:"logoUrl"`
	PrimaryColor string                 `firestore:"primaryColor" json:"primaryColor"`
	PageColor    string                 `firestore:"pageColor" json:"pageColor"`
	Meta         map[string]interface{} `firestore:"meta" json:"meta,omitempty"`
}
