package sanatio

import (
	"reflect"
	"testing"
)

func TestStringValidation(t *testing.T) {
	compareTypes := func(t testing.TB, toBeValidated, validatedAgainst any) bool {
		t.Helper()
		if reflect.TypeOf(toBeValidated) != reflect.TypeOf(validatedAgainst) {
			return false
		}
		return true
	}

	t.Run("creates a string validator without errors", func(t *testing.T) {
		validator := NewStringValidator()

		if status := compareTypes(t, validator, &StringValidator{}); status == false {
			t.Errorf("doesn't return the right validator type: %+v. Expected %+v", validator, StringValidator{})
		}
	})

	t.Run("sets and gets values properly", func(t *testing.T) {
		validator := NewStringValidator()

		value := "string"
		validator.SetValue(value)

		got, _ := validator.GetValue()

		if value != got {
			t.Errorf("doesn't return {%s} the value that it was given {%s}", got, value)
		}
	})

	t.Run("when passed a value, it returns a StringValidator type", func(t *testing.T) {
		validator := NewStringValidator()

		value := "string"
		v := validator.SetValue(value)

		if status := compareTypes(t, v, &StringValidator{}); status == false {
			t.Errorf("doesn't return a string validator type: %+v. Expected %+v", validator, StringValidator{})
		}
	})

	t.Run("returns false when we use the required validation on a validator that has no value", func(t *testing.T) {
		validator := NewStringValidator()

		validator.Required()

		if len(validator.GetErrors()) == 0 {
			t.Fatalf("expected %s, but did not get an error", ErrRequiredValueNotProvided)
		}
	})
}
