package entities

type SocialAccount struct {
	Platform   string `bson:"platform"`
	ProfileURI string `bson:"profile_uri"`
}
