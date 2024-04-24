package sanatio

type StringValidator struct {
	value           string
	errors          []error
	hasUserSetValue bool
}

/*
To validate strings, you create a NewStringValidator(), make sure
to set the value, or you'll get an error back when you're trying to
retrieve it later and then tack on the validations you want, like:

validator := NewStringValidator.SetValue("x").Required()

To retrieve the errors generated during validation, use the
GetErrors function
*/
func NewStringValidator() *StringValidator {
	errorsSlice := make([]error, 1)
	return &StringValidator{
		errors: errorsSlice,
	}
}

/*
The SetValue function allows us to set the value of to be
validated.  This should be the first function that you call
as there is no validation without it
*/
func (s *StringValidator) SetValue(value string) *StringValidator {
	s.value = value
	s.hasUserSetValue = true
	return s
}

/*
GetValue allows you to retrieve the value that you entered into the
validator in the SetValue function.

It's just an ergonomic function, you can choose to value
before entry, validators should mutate the value at all
*/
func (s *StringValidator) GetValue() (string, error) {
	if s.hasUserSetValue == false {
		return s.value, ErrValueNotProvided
	}
	return s.value, nil
}

/*
The Required() function checks that the string is not empty.
It assumes that you have passed in a value using the SetValue function
*/
func (s *StringValidator) Required() *StringValidator {
	if s.value == "" {
		s.errors = append(s.errors, ErrRequiredValueNotProvided)
	}

	return s
}

/*
The GetErrors function returns a slice containing all the errors
generated during validation
*/
func (s *StringValidator) GetErrors() []error {
	return s.errors
}
