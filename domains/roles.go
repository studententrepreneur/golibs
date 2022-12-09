package domain

type Roles struct {
	Id           string        `firestore:"id" json:"id"`
	Name         string        `firestore:"name" json:"name"`
	Organization *Organization `firestore:"organizationId" json:"orgId"`
	Description  string        `firestore:"description" json:"description"`
	Permissions  []Permissions `firestore:"permissions" json:"permissions"`
}

type Permissions struct {
	Scope       string `firestore:"scope" json:"scope"`
	Description string `firestore:"description" json:"description"`
}
