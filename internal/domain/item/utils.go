package item

import (
	"crypto/rand" // Gera números seguros
	"errors"      // Cria erros
	"fmt"         // Formatar strings
	"math/big"    // Números grandes
	"strings"     // Manipular texto
	"unicode"     // Classificar caracteres
)

func GenerateItemCode(nome string) (string, error) {
	nome = strings.TrimSpace(nome) // Remove espaços: "  Mouse  " → "Mouse"
	if nome == "" {                // Se vazio após limpeza
		return "", errors.New("O nome não pode ser vazio") // Retorna erro
	}

	letras := ""                // String para guardar só letras
	for _, char := range nome { // Percorre cada caractere
		if unicode.IsLetter(char) { // Se é letra (A-Z, ç, etc)
			letras += string(unicode.ToUpper(char)) // Converte para maiúscula e adiciona
			if len(letras) >= 3 {                   // Se já tem 3 letras
				break // Para de processar
			}
		}
	}

	if len(letras) < 3 {
		return "", errors.New("A palavra deve ter no mínimo 3 letras")
	}

	prefixo := letras[:3]      // Pega primeiros 3: "MOUSE" → "MOU"
	var codigo strings.Builder // Builder eficiente para números
	for i := 0; i < 8; i++ {   // Loop 8 vezes (8 dígitos)
		n, err := rand.Int(rand.Reader, big.NewInt(10)) // Gera número 0-9 seguro
		if err != nil {                                 // Se erro na geração
			return "", fmt.Errorf("erro ao gerar código: %w", err) // Retorna erro formatado
		}
		codigo.WriteString(n.String()) // Adiciona número ao builder
	}

	return prefixo + codigo.String(), nil // Junta prefixo + números
}
