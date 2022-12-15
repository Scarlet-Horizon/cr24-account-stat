package response

type ErrorResponse struct {
	// Error description.
	Error string `json:"error" example:"invalid endpoint"`
} //@name ErrorResponse

type EndpointLast struct {
	Name    string `json:"name" example:"api/v1/account"`
	Visited string `json:"visited" example:"2022-12-15 13:17:25"`
}

type EndpointStat struct {
	Name  string `json:"name" example:"api/v1/account"`
	Count int    `json:"count" example:"5"`
}
