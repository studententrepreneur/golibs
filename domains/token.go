package domain

import "time"

type Token struct {
	Id          string    `firestore:"documentId" json:"id"`
	UserId      string    `firestore:"userId" json:"userId"`
	CreatedDate time.Time `firestore:"createdDate" json:"createdDate"`
	UserAgent   string    `firestore:"userAgent" json:"userAgent"`
	Platform    string    `firestore:"platform" json:"platform"`
	IP          string    `firestore:"ip" json:"ip"`
	Client      *Clients  `firestore:"client" json:"client"`
	Expires     time.Time `firestore:"expires" json:"expires"`
	LatLong     string    `firestore:"latlong" json:"latlong"`
	City        string    `firestore:"city" json:"city"`
	Region      string    `firestore:"region" json:"region"`
	Country     string    `firestore:"country" json:"country"`
}
