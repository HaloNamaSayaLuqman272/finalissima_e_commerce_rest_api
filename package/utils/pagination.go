package utils

type Pagination struct {
	Limit int `json:"limit,omitempty" query:"limit"`
	Page  int `json:"page,omitempty" query:"page"`
}
