package dto

import (
	"time"

	domain "github.com/studententrepreneur/golibs/domains"
)

type Works struct {
	Id                string          `json:"id"`
	Name              string          `json:"name"`
	Position          string          `json:"position"`
	Website           string          `json:"website"`
	StartDate         time.Time       `json:"startDate"`
	EndDate           time.Time       `json:"endDate"`
	Summary           string          `json:"summary"`
	Highlights        []string        `json:"highlights"`
	Location          domain.Location `json:"location"`
	IsCurrentPosition bool            `json:"isCurrentPosition,omitempty"`
	UserId            string          `json:"userId"`
}

type Volunteers struct {
	Id               string    `json:"id"`
	Organization     string    `json:"organization"`
	Position         string    `json:"position"`
	Website          string    `json:"website"`
	StartDate        time.Time `json:"startDate"`
	EndDate          time.Time `json:"endDate"`
	Summary          string    `json:"summary"`
	Highlights       []string  `json:"highlights"`
	CurrentlyServing bool      `json:"currentlyServing,omitempty"`
	UserId           string    `json:"userId"`
}

type Educations struct {
	Id                string    `json:"id"`
	Institution       string    `json:"institution"`
	Website           string    `json:"website"`
	Area              string    `json:"area"`
	StudyType         string    `json:"studyType"`
	StartDate         time.Time `json:"startDate"`
	EndDate           time.Time `json:"endDate"`
	GPA               int       `json:"gpa"`
	Courses           []string  `json:"courses"`
	CurrentlyPursuing bool      `json:"currentlyPursuing,omitempty"`
	UserId            string    `json:"userId"`
	Summary           string    `firestore:"summary" json:"summary"`
}

type Awards struct {
	Id      string    `firestore:"id" json:"id"`
	Title   string    `firestore:"title" json:"title"`
	Date    time.Time `firestore:"date" json:"date"`
	Awarder string    `firestore:"awarder" json:"awarder"`
	Summary string    `firestore:"summary" json:"summary"`
	UserId  string    `firestore:"userId" json:"userId"`
}

type Publications struct {
	Id          string    `firestore:"id" json:"id"`
	Name        string    `firestore:"name" json:"name"`
	Publisher   string    `firestore:"publisher" json:"publisher"`
	ReleaseDate time.Time `firestore:"releaseDate" json:"releaseDate"`
	Website     string    `firestore:"Website" json:"website"`
	Summary     string    `firestore:"summary" json:"summary"`
	UserId      string    `firestore:"userId" json:"userId"`
}

type Skills struct {
	Id       string   `firestore:"id" json:"id"`
	Name     string   `firestore:"name" json:"name"`
	Level    string   `firestore:"level" json:"level"`
	Value    int      `firestore:"value" json:"value"`
	Keywords []string `firestore:"keywords" json:"keywords"`
	UserId   string   `firestore:"userId" json:"userId"`
}

type Languages struct {
	Id       string `firestore:"id" json:"id"`
	Language string `firestore:"language" json:"language"`
	Fluency  string `firestore:"fluency" json:"fluency"`
	Value    int    `firestore:"value" json:"value"`
	UserId   string `firestore:"userId" json:"userId"`
}

type Interests struct {
	Id       string   `firestore:"id" json:"id"`
	Name     string   `firestore:"name" json:"name"`
	GifUrl   string   `firestore:"gifUrl" json:"gifUrl"`
	Keywords []string `firestore:"keywords" json:"keywords"`
	UserId   string   `firestore:"userId" json:"userId"`
}

type References struct {
	Id        string `firestore:"id" json:"id"`
	Name      string `firestore:"name" json:"name"`
	Reference string `firestore:"reference" json:"reference"`
	UserId    string `firestore:"userId" json:"userId"`
}

type Projects struct {
	Id          string    `firestore:"id" json:"id"`
	Name        string    `firestore:"name" json:"name"`
	Description string    `firestore:"description" json:"description"`
	Highlights  []string  `firestore:"highlights" json:"highlights"`
	Keywords    []string  `firestore:"keywords" json:"keywords"`
	StartDate   time.Time `firestore:"startDate" json:"startDate"`
	EndDate     time.Time `firestore:"endDate" json:"endDate"`
	Website     string    `firestore:"link" json:"link"`
	Images      []string  `firestore:"images" json:"images"`
	Roles       []string  `firestore:"roles" json:"roles"`
	Entity      string    `firestore:"entity" json:"entity"`
	Type        string    `firestore:"type" json:"type"`
	UserId      string    `firestore:"userId" json:"userId"`
}

type Certificates struct {
	Id     string    `firestore:"id" json:"id"`
	Name   string    `firestore:"name" json:"name"`
	Date   time.Time `firestore:"date" json:"date"`
	Issuer string    `firestore:"issuer" json:"issuer"`
	URL    string    `firestore:"url" json:"url"`
	UserId string    `firestore:"userId" json:"userId"`
	Images string    `firestore:"image" json:"image"`
}

type Preference struct {
	Theme       string          `json:"theme"`
	ImageHeader string          `json:"imageHeader"`
	Specific    domain.Specific `json:"specific"`
}
