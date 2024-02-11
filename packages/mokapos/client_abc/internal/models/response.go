package models

// Response represents the overall structure of the API response.
type Response struct {
	Data Data `json:"data"`
	Meta Meta `json:"meta"`
}

// Data holds the primary information returned by the API, typically an array of items.
type Data struct {
	Items      []Item `json:"items"`
	TotalPages int    `json:"total_pages"`
	TotalCount int    `json:"total_count"`
}

// Meta provides metadata about the API response, such as status codes or potential errors.
type Meta struct {
	Code   int               `json:"code"`
	Errors map[string]string `json:"errors,omitempty"`
}
