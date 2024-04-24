package sanatio_test

import (
	"testing"
	"time"

	"github.com/TobiOkanlawon/go-sanatio"
)

func TestSanatioStringValidation(t *testing.T) {
	t.Run("happy path works", func(t *testing.T){		
		happy := sanatio.NewStringValidator().SetValue("cow").Required()

		val, err := happy.GetValue()
		
		if val != "cow" {
			t.Errorf("returned %s", val)
		}

		if err != nil {
			t.Errorf("returned err %s", err)
		}
	})

	t.Run("returns errors when there are errors", func(t *testing.T) {
		fail := sanatio.NewStringValidator().Required()

		if len(fail.GetErrors()) == 0 {
			t.Errorf("errors map has length %d", len(fail.GetErrors()))
		}

		_, err := fail.GetValue()

		if err == nil {
			t.Error("doesn't return an error when trying to retrieve empty value")
		}
	})

	t.Run("custom validation works properly, even in time", func(t *testing.T) {
		validator := func(value string) error {
			time.Sleep(20 * time.Millisecond)
			return nil
		}
		
		pass := sanatio.NewStringValidator().SetValue("email@domain.com").Required().AddCustomValidator(validator).GetErrors()

		if len(pass) != 0 {
			t.Error("fails woefully while trying to use custom validators that spend a bit of time doing their stuff")
		}
	})
}
