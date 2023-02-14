package models

type Accounts struct {
	ID           int
	AccessToken  string
	RefreshToken string
	Expires      int
}
