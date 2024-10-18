// validate.go
package main

import (
	"strconv"
	"strings"
)

// Função para validar CPF
func IsValidCPF(cpf string) bool {
	// Remover caracteres não numéricos
	cpf = strings.ReplaceAll(cpf, ".", "")
	cpf = strings.ReplaceAll(cpf, "-", "")

	// Verificar se tem 11 dígitos
	if len(cpf) != 11 {
		return false
	}

	// Verificar se todos os dígitos são iguais
	allEqual := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			allEqual = false
			break
		}
	}
	if allEqual {
		return false
	}

	// Calcular o primeiro dígito verificador
	sum := 0
	for i := 0; i < 9; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (10 - i)
	}
	firstCheckDigit := (sum * 10) % 11
	if firstCheckDigit == 10 {
		firstCheckDigit = 0
	}

	// Verificar o primeiro dígito
	if firstCheckDigit != int(cpf[9]-'0') {
		return false
	}

	// Calcular o segundo dígito verificador
	sum = 0
	for i := 0; i < 10; i++ {
		num, _ := strconv.Atoi(string(cpf[i]))
		sum += num * (11 - i)
	}
	secondCheckDigit := (sum * 10) % 11
	if secondCheckDigit == 10 {
		secondCheckDigit = 0
	}

	// Verificar o segundo dígito
	return secondCheckDigit == int(cpf[10]-'0')
}

// Função para validar CNPJ
func IsValidCNPJ(cnpj string) bool {
	// Remover caracteres não numéricos
	cnpj = strings.ReplaceAll(cnpj, ".", "")
	cnpj = strings.ReplaceAll(cnpj, "-", "")
	cnpj = strings.ReplaceAll(cnpj, "/", "")

	// Verificar se tem 14 dígitos
	if len(cnpj) != 14 {
		return false
	}

	// Definir os multiplicadores
	multipliers1 := []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	multipliers2 := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

	// Calcular o primeiro dígito verificador
	sum := 0
	for i := 0; i < 12; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * multipliers1[i]
	}
	firstCheckDigit := (sum % 11)
	if firstCheckDigit < 2 {
		firstCheckDigit = 0
	} else {
		firstCheckDigit = 11 - firstCheckDigit
	}

	// Verificar o primeiro dígito
	if firstCheckDigit != int(cnpj[12]-'0') {
		return false
	}

	// Calcular o segundo dígito verificador
	sum = 0
	for i := 0; i < 13; i++ {
		num, _ := strconv.Atoi(string(cnpj[i]))
		sum += num * multipliers2[i]
	}
	secondCheckDigit := (sum % 11)
	if secondCheckDigit < 2 {
		secondCheckDigit = 0
	} else {
		secondCheckDigit = 11 - secondCheckDigit
	}

	// Verificar o segundo dígito
	return secondCheckDigit == int(cnpj[13]-'0')
}

// Função para validar se o documento é CPF ou CNPJ
func ValidateDocument(doc string) (bool, string) {
	doc = strings.ReplaceAll(doc, " ", "")
	if len(doc) == 11 && IsValidCPF(doc) {
		return true, "CPF válido"
	} else if len(doc) == 14 && IsValidCNPJ(doc) {
		return true, "CNPJ válido"
	}
	return false, "Documento inválido"
}
