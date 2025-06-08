package model

type URLRequest struct {
	URL string `json:"url" binding:"required"`
}

type URLResponse struct {
	ShortURL string `json:"short_url"`
}
