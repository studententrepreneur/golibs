package domain

type Application struct {
	Id                    string                 `firestore:"id" json:"id"`
	Name                  string                 `firestore:"name" json:"name"`
	Domain                string                 `firestore:"domain" json:"domain"`
	Description           string                 `firestore:"description" json:"description"`
	Organization          *Organization          `firestore:"organization" json:"organization"`
	ApplicationProperties ApplicationProperties  `firestore:"applicationPoperties" json:"applicationPoperties"`
	TokenProperties       TokenProperties        `firestore:"tokenProperties" json:"tokenProperties"`
	AdvancedSettings      AdvancedSettings       `firestore:"advancedSettings" json:"advancedSettings"`
	Meta                  map[string]interface{} `firestore:"meta" json:"meta,omitempty"`
	DefaultRole           *Roles                 `firestore:"defaultRole" json:"defaultRole"`
	IsSignupDisabled      bool                   `firestore:"isLoginDisabled" json:"isLoginDisabled"`
	Logs                  AbstractLog            `firestore:"logs" json:"logs"`
	SharedWith            []*Credentials         `firestore:"sharedwith" json:"sharedWith"`
	GoogleOauth           GooglePreference       `firestore:"googleOauth" json:"googleOauth"`
}

type GooglePreference struct {
	ClientId     string `firestore:"clientId" json:"clientId"`
	ClientSecret string `firestore:"clientSecret" json:"clientSecret"`
}

type ApplicationProperties struct {
	LogoURL  string   `firestore:"logourl " json:"logoUrl"`
	Origin   []string `firestore:"origin" json:"origin"`
	Callback string   `firestore:"callbacks" json:"callbacks"`
}

type TokenProperties struct {
	ExpiresPeriod            int    `firestore:"expires" json:"expires"`
	Rotation                 bool   `firestore:"rotation" json:"rotation"`
	RotationPeriod           int    `firestore:"rotationPeriod" json:"rotationPeriod"`
	RefreshTokenExpiry       bool   `firestore:"refreshTokenExpiry" json:"refreshTokenExpiry"`
	RefreshTokenExpiryPeriod bool   `firestore:"refreshTokenExpiryPeriod" json:"refreshTokenExpiryPeriods"`
	Identifier               string `firestore:"identifier" json:"identifier"`
}

type AdvancedSettings struct {
	Meta         map[string]interface{} `firestore:"meta" json:"meta"`
	GrantTypes   []string               `firestore:"grantTypes" json:"grantTypes"`
	PublicKey    []byte                 `firestore:"publicKey" json:"publicKey"`
	PrivateKey   []byte                 `firestore:"privateKey" json:"privateKey"`
	WebhookUrl   string                 `firestore:"webhookUrl" json:"webhookUrl"`
	WebsocketURL string                 `firestore:"websocketURL" json:"websocketURL"`
}

type UserSettings struct {
	DefaultIsAccountNonExpired bool `firestore:"defaultIsAccountNonExpired" json:"defaultIsAccountNonExpired"`
	DefaultIsAccountNonLocked  bool `firestore:"defaultIsAccountNonLocked" json:"defaultIsAccountNonLocked"`
	NeedsActivation            bool `firestore:"needsActivation" json:"needsActivation"`
	MFA                        bool `firestore:"mfa" json:"mfa"`
}
