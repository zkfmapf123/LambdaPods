package dto

type ResponseParams[T any] struct {
	Results T      `json:"results"`
	Status  int    `json:"status"`
	Error   string `json:"error,omitempty"` // 마샬링 시, 값이 없으면 제외
}

func NewResponse[T any](results T, status int, error string) ResponseParams[T] {
	return ResponseParams[T]{
		Results: results,
		Status:  status,
		Error:   error,
	}
}
