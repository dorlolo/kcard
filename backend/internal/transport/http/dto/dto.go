// Package dto 提供HTTP API的数据传输对象定义。
package dto

// PageMeta 表示分页元数据。
type PageMeta struct {
	Page     int   `json:"page"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
}
// ErrorResponse 表示错误响应结构。
type ErrorResponse struct {
	Error ErrorBody `json:"error"`
}
// ErrorBody 表示错误详情体。
type ErrorBody struct {
	Code    string   `json:"code"`
	Message string   `json:"message"`
	Details []string `json:"details,omitempty"`
}
// JobAccepted 表示任务已接受的响应。
type JobAccepted struct {
	Job AIJob `json:"job"`
}
// AIJob 表示AI任务的描述信息。
type AIJob struct {
	ID              string   `json:"id"`
	JobType         string   `json:"jobType"`
	Status          string   `json:"status"`
	ProgressPercent int      `json:"progressPercent"`
	CurrentStep     string   `json:"currentStep"`
	ErrorMessage    string   `json:"errorMessage,omitempty"`
	ResultDraftIDs  []string `json:"resultDraftIds,omitempty"`
}
