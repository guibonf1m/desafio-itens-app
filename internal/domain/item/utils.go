package item

import (
	"errors"
	"math/rand"
	"strings"
)

func GenerateItemCode(nome string) (string, error) {

	if len(nome) < 3 {
		return nome, errors.New("A palavra deve ter no mínimo 3 letras")
	}

	prefixo := strings.ToUpper(nome[:3])

	const numeros = "0123456789"

	n := 8
	id := make([]byte, n)

	for i := 0; i < n; i++ {
		id[i] = numeros[rand.Intn(len(numeros))]
	}
	return prefixo + string(id), nil
}
