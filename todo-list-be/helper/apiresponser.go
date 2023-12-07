package helper

type Status string

const (
	Success Status = "success" //2xx
	Fail    Status = "fail"    //4xx
	Error   Status = "error"   //5xx
)

type ApiResponse struct {
	Status  Status `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}