package dependency_injection

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Vanessa")

	resultado := buffer.String()
	esperado := "Olá, Vanessa"

	if resultado != esperado {
		t.Errorf("resultado %s, esperaod %s", resultado, esperado)
	}
}
