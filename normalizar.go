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

			empenho := Empenho{len(dados.Empenhos), id_contrato, ANO_DOCUMENTO_INT,
				NRO_EMPENHO_INT, linha.NRO_ORDEM_FILA, linha.MOTIVO_QUEBRA_ORDEM_CRONOLOGICA}

			dados.Empenhos = append(dados.Empenhos, empenho)
		}
	}

	fmt.Printf("Encontradas %d entidades tipo \"Empenho\".\n", len(dados.Empenhos))

	// preenche Fontes
	for _, linha := range dados_nn {
		// verifica se a fonte já existe
		fonte_existe := false

		for _, fonte := range dados.Fontes {
			if checarIgualdadeFonte(linha, fonte) {
				fonte_existe = true
				break
			}
		}

		if !fonte_existe {
			fonte := Fonte{len(dados.Fontes), linha.NOM_FONTE_RECURSO_TCE}

			dados.Fontes = append(dados.Fontes, fonte)
		}
	}

	fmt.Printf("Encontradas %d entidades tipo \"Fonte\".\n", len(dados.Fontes))

	// preenche Liquidacoes
	for _, linha := range dados_nn {
		// verifica se a liquidacao já existe
		liquidacao_existe := false

		for _, liquidacao := range dados.Liquidacoes {
			if checarIgualdadeLiquidacao(dados, linha, liquidacao) {
				liquidacao_existe = true
				break
			}
		}

		if !liquidacao_existe {
			// encontrar IdEmpenho correto
			id_empenho := encontrarEmpenho(dados, linha)

			NRO_LIQUIDACAO_INT, _ := strconv.Atoi(linha.NRO_LIQUIDACAO)

			liquidacao := Liquidacao{len(dados.Liquidacoes), id_empenho,
				NRO_LIQUIDACAO_INT, linha.DAT_LIQUIDACAO, linha.DAT_LIQUIDACAO_VENCIMENTO}

			dados.Liquidacoes = append(dados.Liquidacoes, liquidacao)
		}
	}

	fmt.Printf("Encontradas %d entidades tipo \"Liquidacao\".\n", len(dados.Liquidacoes))

	// preenche Despesas
	for _, linha := range dados_nn {
		// verifica se despesa já existe
		despesa_existe := false

		for _, despesa := range dados.Despesas {
			if checarIgualdadeDespesa(dados, linha, despesa) {
				despesa_existe = true
				break
			}
		}

		if !despesa_existe {
			// encontrar IdLiquidacacao correto
			id_liquidacao := encontrarLiquidacao(dados, linha)

			NRO_PAGAMENTO_ORDEM_INT, _ := strconv.Atoi(linha.NRO_PAGAMENTO_ORDEM)

			despesa := Despesa{len(dados.Despesas), id_liquidacao,
				linha.SDL_LIQUIDACAO_ORDEM_FINAL, NRO_PAGAMENTO_ORDEM_INT,
				linha.VLR_LIQUIDACAO}

			dados.Despesas = append(dados.Despesas, despesa)
		}
	}

	// preenche EmpenhoFontes
	for _, linha := range dados_nn {
		// verifica se a entidade já existe
		empenho_fonte_existe := false

		for _, empenho_fonte := range dados.EmpenhoFontes {
			if checarIgualdadeEmpenhoFonte(dados, linha, empenho_fonte) {
				empenho_fonte_existe = true
				break
			}
		}

		if !empenho_fonte_existe {
			// encontrar IdEmpenho e IdFonte
			id_empenho := encontrarEmpenho(dados, linha)
			id_fonte := encontrarFonte(dados, linha)

			empenho_fonte := EmpenhoFonte{id_fonte, id_empenho}

			dados.EmpenhoFontes = append(dados.EmpenhoFontes, empenho_fonte)
		}
	}

	fmt.Printf("Encontradas %d entidades associativas tipo \"Empenho Fonte\"\n", len(dados.EmpenhoFontes))

	// retorna dados normalizados
	return dados
}
