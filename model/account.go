package model

import (
	"main/request"
	"time"
)

type Account struct {
	request.Account
	// The opening date for the account
	OpenDate time.Time `json:"openDate" example:"2022-11-26 11:59:38"`
} //	@name	Account
