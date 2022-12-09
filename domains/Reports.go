package domain

type Reports struct {
	Id           string          `firestore:"id" json:"id"`
	Contact      *Contact        `firestore:"contact" json:"contact"`
	Award        []*Awards       `firestore:"award" json:"award"`
	Certificates []*Certificates `firestore:"certificates" json:"certificates"`
	Educations   []*Educations   `firestore:"educations" json:"educations"`
	Interests    []*Interests    `firestore:"interests" json:"interests"`
	Languages    []*Languages    `firestore:"languages" json:"languages"`
	Projects     []*Projects     `firestore:"projects" json:"projects"`
	Publications []*Publications `firestore:"publications" json:"publications"`
	References   []*References   `firestore:"references" json:"references"`
	Skills       []*Skills       `firestore:"skills" json:"skills"`
	Volunteers   []*Volunteers   `firestore:"award" json:"volunteers"`
	Works        []*Works        `firestore:"timelines" json:"works"`
}
