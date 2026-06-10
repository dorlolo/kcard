package dto

type PageMeta struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}
type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
type ErrorBody struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}
type JobAccepted struct {
	Job AIJob `json:"job"`
}
type AIJob struct {
	ID              string   `json:"id"`
	JobType         string   `json:"jobType"`
	Status          string   `json:"status"`
	ProgressPercent int      `json:"progressPercent"`
	CurrentStep     string   `json:"currentStep"`
	ErrorMessage    string   `json:"errorMessage,omitempty"`
	ResultDraftIDs  []string `json:"resultDraftIds,omitempty"`
}
