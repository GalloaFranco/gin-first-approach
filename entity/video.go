package entity

type Video struct {
	Title       string `json:"title"`       // Tag for json serialization
	Description string `json:"description"` // Tag for json serialization
	URL         string `json:"URL"`         // Tag for json serialization
}
