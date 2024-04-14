package app

import (
	"fmt"
	"github.com/pkg/errors"
)

var (
	ErrBannerDoesNotExist = errors.New("banner does not exist")
)

type errorWrapper struct {
	appErr    error
	originErr error
}

func Wrap(applicationError error, originError error) error {
	if originError == nil {
		return nil
	}

	if applicationError == nil {
		return originError
	}

	return errors.WithStack(&errorWrapper{
		appErr:    applicationError,
		originErr: originError,
	})
}

func (e errorWrapper) Error() string {
	return fmt.Sprintf("%s: %s", e.appErr, e.originErr)
}

func (e errorWrapper) Cause() error {
	return e.appErr
}

func (e errorWrapper) Unwrap() error {
	return e.appErr
}
