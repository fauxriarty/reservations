package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

//creates a custom form struct, embeds a url.Values object
type Form struct {
	url.Values
	Errors errors // errors from errors.go
}

// initializes a form struct
func New(data url.Values) *Form {
	return &Form{
		data,
		errors(map[string][]string{}),
	}
}

// Has checks if a form field is in the post and not empty
// returns true if the field is found, false otherwise
func (f *Form) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {

		return false
	}
	return true
}

// valid checks if the form has any errors by checking length of the errors map
// returns true if it doesn't
func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}

// required checks if the form field is in post and not empty
// the ... in the argument means it can take any number of arguments i.e. variatic function
func (f *Form) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)

		//if the value is empty, add an error message
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}

	}
}

// MinLength checks if a form field is a minimum length
func (f *Form) MinLength(field string, length int, r *http.Request) bool {
	x := r.Form.Get(field)
	if len(x) < length {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", length))
		return false
	}
	return true
}

// IsEmail checks if a form field is a valid email address
func (f *Form) IsEmail(field string) {
	if !govalidator.IsEmail(f.Get(field)) {
		f.Errors.Add(field, "Invalid email address")
	}
}
