package sanatio

import (
	"errors"
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

	t.Run("returns an error if the length of a string is more than the set max length", func(t *testing.T) {
		validator := NewStringValidator()
		value := "string"
		errors := validator.SetValue(value).MaxLength(4).GetErrors()

		if len(errors) != 1 {
			t.Fatal("expected an error to be thrown when we are in contravention of the maximum langth validator")
		}

		if errors[0] != ErrGreaterThanMaximumLength {
			t.Errorf("expected %s, but got %s", ErrGreaterThanMaximumLength, errors[0])
		}
	})

	t.Run("returns an error if the length of a string is less than the min set length", func (t *testing.T) {
		validator := NewStringValidator()
		value := "string"
		errors := validator.SetValue(value).MinLength(7).GetErrors()

		if len(errors) != 1 {
			t.Fatal("expected an error to be thrown when we are in contravention of the minimum langth validator")
		}

		if errors[0] != ErrLessThanMinimumLength {
			t.Errorf("expected %s, but got %s", ErrLessThanMinimumLength, errors[0])
		}
	})

	t.Run("allows the user to add a custom validator", func(t *testing.T) {
		validator := NewStringValidator()
		value := "email@domain.com"
		customValidator := func(value string) (error) {
			return nil
		}
		validationErrors := validator.SetValue(value).Required().AddCustomValidator(customValidator).GetErrors()

		if len(validationErrors) != 0 {
			t.Errorf("shouldn't return an error while trying to use a custom validator")
		}
	})

	t.Run("adds the custom validation's error to the errors map", func(t *testing.T) {
		validator := NewStringValidator()
		value := "non-registered-email@domain.com"
		customValidator := func(value string) (error) {
			return errors.New("email is not registered")
		}

		validationErrors := validator.SetValue(value).Required().AddCustomValidator(customValidator).GetErrors()

		if len(validationErrors) != 1 {
			t.Errorf("should return an error when a custom validation fails")
		}
	})

	t.Run("gets the value and returns a ValidationError if there is one", func(t *testing.T) {
		validator := NewStringValidator()
		validator.SetValue("value")
		customValidator := func(value string) (error) {
			return errors.New("custom validator always fails")
		}

		value, err := validator.AddCustomValidator(customValidator).GetValueAndError()

		if err == nil {
			t.Fatalf("expected a ValidationError, got nil error")
		}

		if value != "value" {
			t.Errorf("expected to get the correct value: value got %s", value)
		}
	})

	t.Run("gets the value and nil error if there's no error", func(t *testing.T) {
		validator := NewStringValidator()
		validator.SetValue("value")

		value, err := validator.GetValueAndError()

		if err != nil {
			t.Errorf("expected no error, but got %s", err)
		}

		if value != "value" {
			t.Errorf("expected to get the correct value: value got %s", value)
		}
	})
}
