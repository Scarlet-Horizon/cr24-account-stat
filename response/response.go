package response

type ErrorResponse struct {
	// Error description
	Error string `json:"error" example:"invalid endpoint"`
} //	@name	ErrorResponse

type EndpointLast struct {
	// Endpoint path
	Name string `json:"name" example:"api/v1/account"`
	// Access date
	Visited string `json:"visited" example:"2022-12-15 13:17:25"`
} //	@name	EndpointLast

type EndpointStat struct {
	// Endpoint path
	Name string `json:"name" example:"api/v1/account"`
	// How many times was the endpoint accessed
	Count int `json:"count" example:"5"`
} //	@name	EndpointStat
