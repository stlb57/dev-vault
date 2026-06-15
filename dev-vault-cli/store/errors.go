package store

type StoreError struct {
	Code    string
	Message string
	Err     error
}

var ErrNotFound = &StoreError{
	Code:    "NOT_FOUND",
	Message: "snippet not found",
}

func (s StoreError) Error() string {
	return s.Message
}

func (s StoreError) Unwrap() error {
	return s.Err
}
