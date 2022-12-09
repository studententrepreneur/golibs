package domain

import (
	"time"
)

type Works struct {
	Id                string      `firestore:"id" json:"id"`
	Logs              AbstractLog `firestore:"logs" json:"logs"`
	Name              string      `firestore:"name" json:"name"`
	Position          string      `firestore:"position" json:"position"`
	Website           string      `firestore:"Website" json:"website"`
	StartDate         time.Time   `firestore:"startDate" json:"startDate"`
	EndDate           time.Time   `firestore:"endDate" json:"endDate"`
	Summary           string      `firestore:"summary" json:"summary"`
	Highlights        []string    `firestore:"highlights" json:"highlights"`
	Location          Location    `firestore:"location" json:"location"`
	IsCurrentPosition bool        `firestore:"isCurrentPosition" json:"isCurrentPosition,omitempty"`
	UserId            string      `firestore:"userId" json:"userId"`
}

type Volunteers struct {
	Id               string      `firestore:"id" json:"id"`
	Logs             AbstractLog `firestore:"logs" json:"logs"`
	Organization     string      `firestore:"organization" json:"organization"`
	Position         string      `firestore:"position" json:"position"`
	Website          string      `firestore:"Website" json:"website"`
	StartDate        time.Time   `firestore:"startDate" json:"startDate"`
	EndDate          time.Time   `firestore:"endDate" json:"endDate"`
	Summary          string      `firestore:"summary" json:"summary"`
	Highlights       []string    `firestore:"highlights" json:"highlights"`
	CurrentlyServing bool        `firestore:"currentlyServing" json:"currentlyServing,omitempty"`
	UserId           string      `firestore:"userId" json:"userId"`
}

type Educations struct {
	Id                string      `firestore:"id" json:"id"`
	Logs              AbstractLog `firestore:"logs" json:"logs"`
	Institution       string      `firestore:"institution" json:"institution"`
	Website           string      `firestore:"Website" json:"website"`
	Area              string      `firestore:"area" json:"area"`
	StudyType         string      `firestore:"studyType" json:"studyType"`
	StartDate         time.Time   `firestore:"startDate" json:"startDate"`
	EndDate           time.Time   `firestore:"endDate" json:"endDate"`
	GPA               int         `firestore:"gpa" json:"gpa"`
	Courses           []string    `firestore:"courses" json:"courses"`
	CurrentlyPursuing bool        `firestore:"currentlyPursuing" json:"currentlyPursuing,omitempty"`
	UserId            string      `firestore:"userId" json:"userId"`
	Summary           string      `firestore:"summary" json:"summary"`
}

type Awards struct {
	Id      string      `firestore:"id" json:"id"`
	Logs    AbstractLog `firestore:"logs" json:"logs"`
	Title   string      `firestore:"title" json:"title"`
	Date    time.Time   `firestore:"date" json:"date"`
	Awarder string      `firestore:"awarder" json:"awarder"`
	Summary string      `firestore:"summary" json:"summary"`
	UserId  string      `firestore:"userId" json:"userId"`
}

type Publications struct {
	Id          string      `firestore:"id" json:"id"`
	Logs        AbstractLog `firestore:"logs" json:"logs"`
	Name        string      `firestore:"name" json:"name"`
	Publisher   string      `firestore:"publisher" json:"publisher"`
	ReleaseDate time.Time   `firestore:"releaseDate" json:"releaseDate"`
	Website     string      `firestore:"Website" json:"website"`
	Summary     string      `firestore:"summary" json:"summary"`
	UserId      string      `firestore:"userId" json:"userId"`
}

type Skills struct {
	Id       string      `firestore:"id" json:"id"`
	Logs     AbstractLog `firestore:"logs" json:"logs"`
	Name     string      `firestore:"name" json:"name"`
	Level    string      `firestore:"level" json:"level"`
	Value    int         `firestore:"value" json:"value"`
	Keywords []string    `firestore:"keywords" json:"keywords"`
	UserId   string      `firestore:"userId" json:"userId"`
}

type Languages struct {
	Id       string      `firestore:"id" json:"id"`
	Logs     AbstractLog `firestore:"logs" json:"logs"`
	Language string      `firestore:"language" json:"language"`
	Fluency  string      `firestore:"fluency" json:"fluency"`
	Value    int         `firestore:"value" json:"value"`
	UserId   string      `firestore:"userId" json:"userId"`
}

type Interests struct {
	Id       string      `firestore:"id" json:"id"`
	Logs     AbstractLog `firestore:"logs" json:"logs"`
	Name     string      `firestore:"name" json:"name"`
	GifUrl   string      `firestore:"gifUrl" json:"gifUrl"`
	Keywords []string    `firestore:"keywords" json:"keywords"`
	UserId   string      `firestore:"userId" json:"userId"`
}

type References struct {
	Id        string      `firestore:"id" json:"id"`
	Logs      AbstractLog `firestore:"logs" json:"logs"`
	Name      string      `firestore:"name" json:"name"`
	Reference string      `firestore:"reference" json:"reference"`
	UserId    string      `firestore:"userId" json:"userId"`
}

type Projects struct {
	Id          string      `firestore:"id" json:"id"`
	Logs        AbstractLog `firestore:"logs" json:"logs"`
	Name        string      `firestore:"name" json:"name"`
	Description string      `firestore:"description" json:"description"`
	Highlights  []string    `firestore:"highlights" json:"highlights"`
	Keywords    []string    `firestore:"keywords" json:"keywords"`
	StartDate   time.Time   `firestore:"startDate" json:"startDate"`
	EndDate     time.Time   `firestore:"endDate" json:"endDate"`
	Website     string      `firestore:"link" json:"link"`
	Images      []string    `firestore:"images" json:"images"`
	Roles       []string    `firestore:"roles" json:"roles"`
	Entity      string      `firestore:"entity" json:"entity"`
	Type        string      `firestore:"type" json:"type"`
	UserId      string      `firestore:"userId" json:"userId"`
}

type Certificates struct {
	Id     string      `firestore:"id" json:"id"`
	Logs   AbstractLog `firestore:"logs" json:"logs"`
	Name   string      `firestore:"name" json:"name"`
	Date   time.Time   `firestore:"date" json:"date"`
	Issuer string      `firestore:"issuer" json:"issuer"`
	URL    string      `firestore:"url" json:"url"`
	UserId string      `firestore:"userId" json:"userId"`
	Images string      `firestore:"image" json:"image"`
}
