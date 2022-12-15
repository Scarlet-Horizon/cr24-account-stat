package model

type Account struct {
	// User UUID
	PK string `json:"userID" binding:"required" example:"6204037c-30e6-408b-8aaa-dd8219860b4b"`
	// Account UUID
	SK string `json:"accountID" binding:"required" example:"09130407-1f81-4ac5-be85-6557683462d0"`
	// The opening date for the account
	Type string `json:"type" binding:"required" example:"checking" enums:"checking,saving"`
	//The opening date for the account
	OpenDate string `json:"openDate" binding:"required" example:"2022-11-26 11:59:38"`
} //	@name	Account
