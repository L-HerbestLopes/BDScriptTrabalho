package main

import (
	"fmt"
)

func main() {
	caminho_dados := "dados.json"

	// importa os dados do arquivo json para um slice de DadosNaoNormalizados
	lista_dados_nn := ImportarDados(caminho_dados)
	fmt.Printf("Importados %d entidades para a memória principal\n", len(lista_dados_nn))

	// normaliza os dados
	dados_normalizados := NormalizarDados(lista_dados_nn)

	for i := 0; i < len(dados_normalizados.Pessoas); i++ {
		fmt.Println(dados_normalizados.Pessoas[i])
	}
}
