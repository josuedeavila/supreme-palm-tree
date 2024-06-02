package fake

type UseCases struct {
	ResetFunc func() error
}

func (u *UseCases) Reset() error {
	return u.ResetFunc()
}

func NewUseCases(resetFunc func() error) *UseCases {
	return &UseCases{
		ResetFunc: resetFunc,
	}
}
