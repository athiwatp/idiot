package idiot_test

import (
	"testing"
)

func TestIdiot_Helper(t *testing.T) {

	type ValidField struct {
		Name string `idiot:"name,required,min=1,max=5"`
	}

	t.Run("Happy - validate string", func(t *testing.T) {

		//d := idiot.DefaultValidator(ValidField)
		//assert.NotNil(t, result)
		//assert.NoError(t, err)
	})

}
