package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/jackc/pgx/v4"
)

func main() {
	// Abrir e ler o arquivo
	file, err := os.Open("base_teste.txt")
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close()

	// Conectar ao banco de dados PostgreSQL
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:33720013@localhost:5433/postgres")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text()) // Remove espaços em branco ao redor da linha

		// Verificar se a linha está vazia e pular
		if linha == "" {
			continue
		}

		// Separar os valores, supondo que o delimitador seja '|'
		valores := strings.Split(linha, "|")

		// Remover espaços extras de cada valor e tratar valores nulos
		for i, valor := range valores {
			valores[i] = strings.TrimSpace(valor)
			if valor == "" {
				valores[i] = "NULL"
			}
		}

		// Verificar se o slice 'valores' tem o comprimento esperado
		if len(valores) != 8 {
			fmt.Printf("Erro: número inesperado de valores na linha (esperado 8, obtido %d): %s\n", len(valores), linha)
			continue
		}

		// Log dos valores que serão inseridos
		fmt.Println("Inserindo valores:", valores)

		// Preparar a consulta SQL com placeholders para valores nulos
		query := `INSERT INTO usuario (cpf, private, incompleto, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra) 
		          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

		// Tratar valores "NULL" como nil para o banco de dados
		args := make([]interface{}, len(valores))
		for i, v := range valores {
			if v == "NULL" {
				args[i] = nil
			} else {
				args[i] = v
			}
		}

		// Inserir os dados no banco de dados
		_, err = conn.Exec(context.Background(), query, args...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao inserir dados: %v\n", err)
			continue
		}

		fmt.Println("Dados inseridos com sucesso!")
	}

	// Verificar se houve algum erro na leitura do arquivo
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro ao ler o arquivo:", err)
	}
}
