package sanatio

type StringValidator struct {
	value string
	errors []error
}

func NewStringValidator() *StringValidator {
	return &StringValidator{}
}

func (s *StringValidator) SetValue(value string) (*StringValidator) {
	s.value = value
	return s
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
