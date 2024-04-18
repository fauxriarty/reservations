package forms

type errors map[string][]string

// Add adds an error message for a given form field
func (e errors) Add(field, message string) {
	//slice because one field of the form , say email, may have more than one errors possible
	e[field] = append(e[field], message)
}

// Get returns the first error message for a given field
func (e errors) Get(field string) string {
	es := e[field]
	if len(es) == 0 {
		return ""
	}
	return es[0]
}
