package forms

import (
	"net/http"
	"net/url"
)

type Form struct {
	url.Values
	Errors errors
}

//valid returns true if there are no errors.
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// New creates a new form
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has check if the field is not empty
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		f.Errors.Add(field, "This field cannot be empty")
		return false
	}
	return true
}
