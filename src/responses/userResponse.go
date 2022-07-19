package responses

type UserResponse struct {
	IsSuccess bool                   `json:"isSuccess"`
	Message   string                 `json:"message"`
	Data      map[string]interface{} `json:"data"`
}
