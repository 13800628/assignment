package errors

import "errors"

var (
	ErrItemNotFound = errors.New("item not found")
	ErrInvalidInput = errors.New("validation failed")
)

// エラーレスポンスの構造体もここに定義
type AppError struct {
	Error   string   `json:"error"`
	Details []string `json:"details,omitempty"`
}
