package entities

type SocialAccount struct {
	Platform   string `bson:"platform" json:"platform"`
	ProfileURI string `bson:"profile_uri" json:"profile_uri"`
}
