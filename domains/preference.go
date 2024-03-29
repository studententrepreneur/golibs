package domain

type Protfolio struct {
	Id          string   `firestore:"id" json:"id"`
	Report      *Reports `firestore:"report" json:"report"`
	Theme       string   `firestore:"theme" json:"theme"`
	ImageHeader string   `firestore:"imageHeader" json:"imageHeader"`
	Specific    Specific `firestore:"specific" json:"specific"`
}

type Specific struct {
	DreamJob   DreamJob `firestore:"dreamJob" json:"dreamJob"`
	CurrentJob *Works   `firestore:"currentJob" json:"currentJob"`
	IsPrivate  bool     `firestore:"isPrivate" json:"isPrivate"`
}

type DreamJob struct {
	Locations       []Location `firestore:"locations" json:"locations"`
	RemoteFrequency []string   `firestore:"remoteFrequency" json:"remoteFrequency"`
	Type            []string   `firestore:"jobType" json:"jobType"`
	Company         []string   `firestore:"company" json:"company"`
	ExpirienceLevel []string   `firestore:"experience" json:"experience"`
}
