package testunitario

import (
	"fmt"
	"testing"
)

/* func TestSuma(t *testing.T) {
	total := Suma(5, 9)

	if total != 12 {
		t.Errorf("Suma incorrecta, tiene %d y se esperaba %d", total, 12)
	} else {
		fmt.Println("Suma correcta")
	}
} */

func TestSuma(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 2, 4},
		{25, 25, 50},
	}

	for _, item := range tabla {
		total := Suma(item.a, item.b)

		if total != item.c {
			t.Errorf("La suma no es correcta")
		} else {
			fmt.Println("Suma realizada correctamente")
		}
	}
}

func TestGetMax(t *testing.T) {
	tabla := []struct {
		a int
		b int
		c int
	}{
		{67, 25, 67},
		{9, 3, 9},
		{24, 43, 43},
	}

	for _, item := range tabla {
		max := GetMax(item.a, item.b)

		if max != item.c {
			t.Errorf("La suma no es correcta")
		} else {
			fmt.Println("Suma realizada correctamente")
		}
	}
}

func TestFibo(t *testing.T) {
	tabla := []struct {
		n int
		r int
	}{
		{1, 1},
		{8, 21},
		{50, 12586269025},
	}

	for _, item := range tabla {
		fibo := Fibonacci(item.n)

		if fibo != item.r {
			t.Errorf("La suma no es correcta")
		} else {
			fmt.Println("Suma realizada correctamente")
		}
	}
}
