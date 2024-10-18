package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/jackc/pgx/v4"
)

func main() {
	// Função para abrir e ler os arquivos do txt
	file, err := os.Open("base_teste.txt") 
	if err != nil {
		fmt.Println("Erro ao abrir o arquivo:", err)
		return
	}
	defer file.Close() 

	// Conectar ao banco de dados PostgreSQL
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:33720013@localhost:5433/challenge_data") 
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background()) 

	// Uso de uma biblioteca de scanner para ler o arquivo linha por linha
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		linha := strings.TrimSpace(scanner.Text()) // Remove espaços no começo e no final da linha
		if linha == "" {
			continue
		}

		// Separar os valores da linha em um slice, removendo múltiplos espaços
		valores := strings.Fields(linha)

		// Verificar se o número de valores está correto (8 valores esperados)
		if len(valores) != 8 {
			fmt.Printf("Erro: número inesperado de valores na linha (esperado 8, obtido %d): %s\n", len(valores), linha)
			continue 
		}

		// Corrigir o formato dos valores numéricos (trocar vírgula por ponto)
		valores[4] = strings.Replace(valores[4], ",", ".", 1) 
		valores[5] = strings.Replace(valores[5], ",", ".", 1) 

		ticketMedio, err := strconv.ParseFloat(valores[4], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter ticket_medio: %v\n", err)
			continue
		}
		ticketUltimaCompra, err := strconv.ParseFloat(valores[5], 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Erro ao converter ticket_da_ultima_compra: %v\n", err)
			continue
		}

		// Preparar a consulta SQL para inserir os dados na tabela
		query := `INSERT INTO usuario (cpf, private, incompleto, data_da_ultima_compra, ticket_medio, ticket_da_ultima_compra, loja_mais_frequente, loja_da_ultima_compra) 
		          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

		_, err = conn.Exec(context.Background(), query, valores[0], valores[1], valores[2], valores[3], ticketMedio, ticketUltimaCompra, valores[6], valores[7])
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