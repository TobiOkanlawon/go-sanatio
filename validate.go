package sanatio

type StringValidator struct {
	value string
	errors []error
	hasUserSetValue bool
}

func NewStringValidator() *StringValidator {
	return &StringValidator{}
}

func (s *StringValidator) SetValue(value string) (*StringValidator) {
	s.value = value
	s.hasUserSetValue = true
	return s
}

func (s *StringValidator) GetValue() (string, error) {
	if s.hasUserSetValue == false {
		return s.value, ErrValueNotProvided
	}
	return s.value, nil
}

func (s *StringValidator) Required() (*StringValidator) {
	if s.value == "" {
		s.errors = append(s.errors, ErrRequiredValueNotProvided)
	}
	
	return s
}

func (s *StringValidator) GetErrors() ([]error) {
	return s.errors
}
