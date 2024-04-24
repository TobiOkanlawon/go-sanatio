package sanatio

import "errors"

var ErrRequiredValueNotProvided = errors.New("a value that is required has not been provided")
var ErrValueNotProvided = errors.New("value has not been set but attempted to be retrieved")
