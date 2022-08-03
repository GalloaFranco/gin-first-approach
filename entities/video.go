package entities

type Video struct {
	Title       string `json:"title" binding:"required"`                    // Tag for json serialization
	Description string `json:"description" binding:"required,min=2,max=15"` // Tag for json serialization
	URL         string `json:"URL" binding:"required,url"`                  // Tag for json serialization
}
