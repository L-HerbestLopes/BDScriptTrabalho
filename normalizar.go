package main

import (
	"fmt"
	"strconv"
)

func NormalizarDados(dados_nn []DadosNaoNormalizados) DadosNormalizados {
	// inicializa struct
	dados := DadosNormalizados{}

	// preenche Pessoas
	for _, linha := range dados_nn {
		// verifica se pessoa já existe nos dados
		pessoa_existe := false

		for _, pessoa := range dados.Pessoas {
			if checarIgualdadePessoa(linha, pessoa) {
				pessoa_existe = true
				break
			}
		}

		if !pessoa_existe {
			pessoa := Pessoa{len(dados.Pessoas), linha.NOM_PESSOA}

			dados.Pessoas = append(dados.Pessoas, pessoa)
		}
	}

	fmt.Printf("Encontradas %d entidades tipo \"Pessoa\".\n", len(dados.Pessoas))

	// preenche Contratos
	for _, linha := range dados_nn {
		// verifica se o contrato já existe nos dados
		contrato_existe := false

		for _, contrato := range dados.Contratos {
			if checarIgualdadeContrato(dados, linha, contrato) {
				contrato_existe = true
				break
			}
		}

		if !contrato_existe {
			// encontrar IdPessoa correto
			id_pessoa := encontrarPessoa(dados, linha)

			contrato := Contrato{len(dados.Contratos), id_pessoa,
				linha.NOM_CONTRATO_TIPO, linha.NDA_LICITACAO}
			
			dados.Contratos = append(dados.Contratos, contrato)
		}
	}

	fmt.Printf("Encontradas %d entidades tipo \"Contrato\".\n", len(dados.Contratos))

	// preenche Empenhos
	for _, linha := range dados_nn {
		// verifica se o empenho já existe
		empenho_existe := false

		for _, empenho := range dados.Empenhos {
			if checarIgualdadeEmpenho(dados, linha, empenho) {
				empenho_existe = true
				break
			}
		}
		
		if !empenho_existe {
			// encontrar IdContrato correto
			id_contrato := encontrarContrato(dados, linha)

			ANO_DOCUMENTO_INT, _ := strconv.Atoi(linha.ANO_DOCUMENTO)
			NRO_EMPENHO_INT, _ := strconv.Atoi(linha.NRO_EMPENHO)

			empenho := Empenho {id_contrato, ANO_DOCUMENTO_INT,
			NRO_EMPENHO_INT, linha.NRO_ORDEM_FILA, linha.MOTIVO_QUEBRA_ORDEM_CRONOLOGICA}

			dados.Empenhos = append(dados.Empenhos, empenho)
		}
	}

	fmt.Printf("Encontradas %d entidades tipo \"Empenho\".\n", len(dados.Empenhos))

	// retorna dados normalizados
	return dados
}
