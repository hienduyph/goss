package httpx

import (
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/hienduyph/goss/errorx"
)

type HandleFunc func(r *http.Request) (interface{}, error)

func Handle(fn HandleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp, e := fn(r)
		if e != nil {
			s := statusFromErr(e)
			render.Status(r, s)
			render.JSON(w, r, map[string]interface{}{"error": e.Error()})
			return
		}
		render.JSON(w, r, resp)
	}
}

func statusFromErr(err error) int {
	if errors.Is(err, errorx.ErrBadInput) {
		return http.StatusBadRequest
	}

	if errors.Is(err, errorx.ErrNotFound) {
		return http.StatusNotFound
	}

	return http.StatusInternalServerError
}
