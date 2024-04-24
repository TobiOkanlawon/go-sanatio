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
			t.Fail()
		}

		if err != nil {
			t.Fail()
		}
	})

	t.Run("returns errors when there are errors", func(t *testing.T) {
		fail := sanatio.NewStringValidator().Required()

		if len(fail.GetErrors()) != 0 {
			t.Fail()
		}

		_, err := fail.GetValue()

		if err == nil {
			t.Fail()
		}
	})
}
