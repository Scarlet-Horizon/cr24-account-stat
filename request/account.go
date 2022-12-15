package request

type Account struct {
	// User UUID
	PK string `json:"userID" example:"6204037c-30e6-408b-8aaa-dd8219860b4b"`
	// Account UUID
	SK string `json:"accountID" example:"09130407-1f81-4ac5-be85-6557683462d0"`
	// The opening date for the account
	Type string `json:"type" example:"checking" enums:"checking,saving"`
} //	@name	Account