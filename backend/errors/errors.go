package apperr

type Kind string

const (
	KindValidation Kind = "validation"
	KindNotFound   Kind = "not_found"
	KindConflict   Kind = "conflict"
	KindInternal   Kind = "internal"
)

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

func NewPublicError(kind Kind, code, msg string, err error) *PublicError {
	return &PublicError{Kind: kind, Code: code, Msg: msg, Err: err}
}

func Validation(code, msg string, err error) *PublicError {
	return NewPublicError(KindValidation, code, msg, err)
}

func NotFound(code, msg string, err error) *PublicError {
	return NewPublicError(KindNotFound, code, msg, err)
}

func Conflict(code, msg string, err error) *PublicError {
	return NewPublicError(KindConflict, code, msg, err)
}

func Internal(code, msg string, err error) *PublicError {
	return NewPublicError(KindInternal, code, msg, err)
}
