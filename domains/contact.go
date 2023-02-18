package domain

import "time"

type Credentials struct {
	Id                      string                 `firestore:"id" json:"id"`
	Logs                    AbstractLog            `firestore:"logs" json:"logs"`
	Password                string                 `firestore:"password" json:"password"`
	ActivationKey           string                 `firestore:"activationKey" json:"activationKey"`
	ResetKey                string                 `firestore:"resetKey" json:"resetKey"`
	Roles                   []*Roles               `firestore:"roles" json:"roles,omitempty"`
	Meta                    map[string]interface{} `firestore:"meta" json:"meta,omitempty"`
	MFAsecret               string                 `firestore:"mfa" json:"mfa"`
	IsAccountNonExpired     bool                   `firestore:"isAccountNonExpired" json:"isAccountNonExpired"`
	IsAccountNonLocked      bool                   `firestore:"isAccountNonLocked" json:"isAccountNonLocked"`
	IsCredentialsNonExpired bool                   `firestore:"isCredentialsNonExpired" json:"isCredentialsNonExpired"`
	TwoFactor               bool                   `firestore:"twoFactorAuth" json:"twoFactorAuth"`
	Applications            []*Application         `firestore:"application" json:"application"`
	Contact                 *Contact               `firestore:"contact" json:"contact"`
	GoogleAuthId            string                 `firestore:"googleAuthId" json:"googleauthid"`
}

// Contact Model
// @Description this the Contact model, it contains all Contact related fields
type Contact struct {
	Id          string                 `firestore:"id" json:"id" example:"191Aqu5O"`                                                                                                                                                                                                      // This is auto generated from firestore
	LangKey     string                 `firestore:"langKey" json:"langkey" example:"en_US"`                                                                                                                                                                                               // Language selected by the user lang_Country
	FirstName   string                 `firestore:"firstName" json:"firstName" validate:"required,min=2,max=15" example:"john"`                                                                                                                                                           // Firstname of the user
	LastName    string                 `firestore:"lastName" json:"lastName" validate:"min=2,max=15" example:"doe"`                                                                                                                                                                       // Lastname of the user
	Email       string                 `firestore:"email" json:"email" validate:"required,email" example:"john.doe@acme.inc"`                                                                                                                                                             // email of the user
	ImageUrl    string                 `firestore:"imageUrl" json:"image" example:"https://st3.depositphotos.com/6672868/13701/v/600/depositphotos_137014128-stock-illustration-user-profile-icon.jpg"`                                                                                   // profile picture of the user
	CoverPic    string                 `firestore:"coverPic" json:"coverPic" example:"https://www.quotemaker.in/assets/images/default_cover.jpg"`                                                                                                                                         // cover picture of the user
	Label       string                 `firestore:"label" json:"label" example:"Software engineer"`                                                                                                                                                                                       // Label or title of the user
	PhoneNumber string                 `firestore:"phoneNumber" json:"phone" example:"+918899999999"`                                                                                                                                                                                     // Phone number of the user
	Summary     string                 `firestore:"summary" json:"summary" example:"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eget felis semper, volutpat erat elementum, fringilla nisl. Nunc auctor porttitor fermentum. Morbi volutpat vehicula convallis."` // summary of the user
	Location    *Location              `firestore:"location" json:"location"`                                                                                                                                                                                                             // Location of the user
	Profiles    map[string]interface{} `firestore:"profiles" json:"profiles"`                                                                                                                                                                                                             // Social and other public profiles
	Website     string                 `firestore:"website" json:"url" example:"https://johndoe.acme.inc"`                                                                                                                                                                                // Url of the website / portfolio
	IdentifyAs  string                 `firestore:"identifyas" json:"identifyas" example:"They/Them"`                                                                                                                                                                                     // Gender in gender neutral form
	DateOfBirth time.Time              `firestore:"dateOfBirth" json:"dateOfBirth" example:"1658340160"`                                                                                                                                                                                  // Date of birth
	MemberFrom  time.Time              `firestore:"memberFrom" json:"memberFrom" example:"1658340160"`                                                                                                                                                                                    // Signed up date generated by the server
	Logs        AbstractLog            `firestore:"logs" json:"logs,omitempty" xml:"Logs,omitempty"`                                                                                                                                                                                      // Logs of the user, created and updated date
	Meta        map[string]interface{} `firestore:"meta" json:"meta,omitempty"`
	Social      interface{}            `firestore:"sociallinks" json:"social,omitempty"`             // any additional details need about the user will be pushed to this
	Timezone    string                 `firestore:"timezone" json:"timezone" example:"Asia/Kolkata"` // Timezone of the user
}
