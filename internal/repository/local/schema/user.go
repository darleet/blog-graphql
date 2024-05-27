package schema

import "net/url"

type User struct {
	ID           uint64
	Username     string
	PasswordHash string
	Email        string
	AvatarURL    *url.URL
}
