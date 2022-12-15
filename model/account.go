package model

import (
	"main/request"
	"time"
)

type Account struct {
	request.Account
	// The opening date for the account
	OpenDate time.Time `json:"openDate" example:"2022-11-26T11:59:38+01:00"`
} //@name Account
