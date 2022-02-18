package testunitario

import (
	"fmt"
	"testing"
)

func TestSuma(t *testing.T) {
	total := Suma(5, 9)

	if total != 12 {
		t.Errorf("Suma incorrecta, tiene %d y se esperaba %d", total, 12)
	} else {
		fmt.Println("Suma correcta")
	}
}
