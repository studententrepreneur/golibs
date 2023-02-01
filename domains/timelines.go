package domain

import (
	"time"
)

type Timeline interface {
	getId() string
	getUserId() string
}

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

func (t Works) getId() string {
	return t.Id
}
func (t Works) getUserId() string {
	return t.UserId
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

func (t Volunteers) getId() string {
	return t.Id
}
func (t Volunteers) getUserId() string {
	return t.UserId
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

func (t Educations) getId() string {
	return t.Id
}
func (t Educations) getUserId() string {
	return t.UserId
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

func (t Awards) getId() string {
	return t.Id
}
func (t Awards) getUserId() string {
	return t.UserId
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

func (t Publications) getId() string {
	return t.Id
}
func (t Publications) getUserId() string {
	return t.UserId
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

func (t Skills) getId() string {
	return t.Id
}
func (t Skills) getUserId() string {
	return t.UserId
}

type Languages struct {
	Id       string      `firestore:"id" json:"id"`
	Logs     AbstractLog `firestore:"logs" json:"logs"`
	Language string      `firestore:"language" json:"language"`
	Fluency  string      `firestore:"fluency" json:"fluency"`
	Value    int         `firestore:"value" json:"value"`
	UserId   string      `firestore:"userId" json:"userId"`
}

func (t Languages) getId() string {
	return t.Id
}
func (t Languages) getUserId() string {
	return t.UserId
}

type Interests struct {
	Id       string      `firestore:"id" json:"id"`
	Logs     AbstractLog `firestore:"logs" json:"logs"`
	Name     string      `firestore:"name" json:"name"`
	GifUrl   string      `firestore:"gifUrl" json:"gifUrl"`
	Keywords []string    `firestore:"keywords" json:"keywords"`
	UserId   string      `firestore:"userId" json:"userId"`
}

func (t Interests) getId() string {
	return t.Id
}
func (t Interests) getUserId() string {
	return t.UserId
}

type References struct {
	Id        string      `firestore:"id" json:"id"`
	Logs      AbstractLog `firestore:"logs" json:"logs"`
	Name      string      `firestore:"name" json:"name"`
	Reference string      `firestore:"reference" json:"reference"`
	UserId    string      `firestore:"userId" json:"userId"`
}

func (t References) getId() string {
	return t.Id
}
func (t References) getUserId() string {
	return t.UserId
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

func (t Projects) getId() string {
	return t.Id
}
func (t Projects) getUserId() string {
	return t.UserId
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

func (t Certificates) getId() string {
	return t.Id
}
func (t Certificates) getUserId() string {
	return t.UserId
}
