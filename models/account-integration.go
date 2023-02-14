package models

type AccountIntegration struct {
	ID                 int
	SecretKey          string
	ClientID           string
	RedirectURL        string
	AuthenticationCode string
}
