package util

type Err struct {
	detail string
}

func (err *Err) Error() string {
	return err.detail
}

func NewError(detail string) error {
	return &Err{
		detail: detail,
	}
}
