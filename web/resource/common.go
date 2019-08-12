package resource

import (
	"errors"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"net/http"
	"strconv"
	"strings"
)

//--
// Error response payloads & renderers
//--

// ErrResponse renderer type for handling all sorts of errors.
//
// In the best case scenario, the excellent github.com/pkg/errors package
// helps reveal information on the error, setting it on Err, and in the Render()
// method, using it to set the application-specific error code in AppCode.
type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}

var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}

type Filters struct {
	page    int
	sort    []string
	order   []string
	embed   []string
	expand  []string
	filters []Filter
}

type Filter struct {
	match  string
	field  string
	values []string
}

func ParamInt(r *http.Request, key string) int {
	val, _ := strconv.Atoi(chi.URLParam(r, key))
	return val
}

func ParamArray(r *http.Request, key string) []string {
	return strings.Split(chi.URLParam(r, key), ",")
}

var matchTypes = []string{"ne", "gte", "lte", "like"}

func FiltersFromRequest(r *http.Request) (*Filters, error) {
	filters := &Filters{
		page:   ParamInt(r, "_page"),
		sort:   ParamArray(r, "_sort"),
		order:  ParamArray(r, "_order"),
		embed:  ParamArray(r, "_embed"),
		expand: ParamArray(r, "_expand"),
	}

	for field, values := range r.URL.Query() {
		parts := strings.Split(field, "_")
		if len(parts) == 1 {
			filters.filters = append(filters.filters, Filter{
				field:  field,
				values: values,
				match:  "eq",
			})
			continue
		}

		if len(parts) != 2 {
			return nil, errors.New("invalid param" + field)
		}

		prefix, suffix := parts[0], parts[1]

		if len(prefix) == 0 {
			continue
		}

		if !IsValidMatchType(suffix) {
			return nil, errors.New("invalid param" + field)
		}

		filters.filters = append(filters.filters, Filter{
			field:  prefix,
			values: values,
			match:  suffix,
		})
	}
	return filters, nil
}

func IsValidMatchType(k string) bool {
	return contains(matchTypes, k)
}

// Contains tells whether a contains x.
func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
