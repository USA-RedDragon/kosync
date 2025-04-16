package apimodels

type SyncProgressRequest struct {
	Percentage float64 `json:"percentage" binding:"required"`
	Document   string  `json:"document" binding:"required"`
	Device     string  `json:"device" binding:"required"`
	Progress   string  `json:"progress" binding:"required"`
}

type ProgressResponse struct {
	Data ProgressResponseData `json:"data"`
}

type ProgressResponseData struct {
	Type       string  `json:"type"`
	Document   string  `json:"document"`
	Percentage float64 `json:"percentage"`
	Progress   string  `json:"progress"`
	Device     string  `json:"device"`
}
