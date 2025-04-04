package context

import (
	"errors"

	"github.com/TotallyNotLirgo/back-seat-boys/models"
)

func (c *Context) WriteErrorResponse(err error) {
    code := getStatusCode(err)
    msg, err := getErrorMessage(err)
    if err != nil {
        c.WriteJSONMessage(500, err.Error())
        return
    }
    c.WriteJSONMessage(code, msg)
}

func getStatusCode(err error) int {
	switch {
	case errors.Is(err, models.ErrUnauthorized):
		return 401
	case errors.Is(err, models.ErrForbidden):
		return 403
	case errors.Is(err, models.ErrNotFound):
		return 404
	case errors.Is(err, models.ErrConflict):
		return 409
	case errors.Is(err, models.ErrBadRequest):
		return 422
	case errors.Is(err, models.ErrServerError):
		return 500
	}
	return 500
}

func getErrorMessage(err error) (string, error) {
	errs, ok := err.(interface{ Unwrap() []error })
	if !ok {
		return "", errors.New("could not unwrap error messages")
	}
	for _, e := range errs.Unwrap() {
		if !errors.Is(e, models.ErrServiceError) {
			return e.Error(), nil
		}
	}
	return "", errors.New("no non-service error found")
}
