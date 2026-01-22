package main

import (
	json "encoding/json"
	"os"
	"log"
)

func ImportarDados(caminho_dados string) []DadosNaoNormalizados {
	// inicializa slice vazio
	dados := []DadosNaoNormalizados{}

	// lê arquivo de dados
	content, err := os.ReadFile(caminho_dados)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(content, &dados)
	if err != nil {
		log.Fatal(err)
	}

	return dados
}
