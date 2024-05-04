package sanatio

import "errors"

var ErrRequiredValueNotProvided = errors.New("a value that is required has not been provided")
var ErrValueNotProvided = errors.New("value has not been set but attempted to be retrieved")
var ErrGreaterThanMaximumLength = errors.New("value is greater than the length set as max")
var ErrLessThanMinimumLength = errors.New("value is less than the length set as min")
var ValidationError = errors.New("an error occured during validation")
