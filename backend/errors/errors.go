package apperr

type Kind string

const (
	KindValidation Kind = "validation"
	KindNotFound   Kind = "not_found"
	KindConflict   Kind = "conflict"
	KindInternal   Kind = "internal"
)

type PublicData struct {
	Code    string
	Message string
}

type PublicError struct {
	Kind Kind
	Code string
	Msg  string
	Err  error
}

func (e *PublicError) Error() string {
	if e.Err == nil {
		return ""
	}
	return e.Err.Error()
}

func (e *PublicError) Unwrap() error { return e.Err }

func NewPublicError(kind Kind, data PublicData, err error) *PublicError {
	return &PublicError{Kind: kind, Code: data.Code, Msg: data.Message, Err: err}
}

func Validation(data PublicData, err error) *PublicError {
	return NewPublicError(KindValidation, data, err)
}

func NotFound(data PublicData, err error) *PublicError {
	return NewPublicError(KindNotFound, data, err)
}

func Conflict(data PublicData, err error) *PublicError {
	return NewPublicError(KindConflict, data, err)
}

func Internal(data PublicData, err error) *PublicError {
	return NewPublicError(KindInternal, data, err)
}
