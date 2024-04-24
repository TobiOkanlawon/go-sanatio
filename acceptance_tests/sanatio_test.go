package sanatio_test

import (
	"testing"

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
}
