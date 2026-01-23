package main

import (
	"strconv"
)

func checarIgualdadePessoa(dados_nn DadosNaoNormalizados, pessoa Pessoa) bool {
	return dados_nn.NOM_PESSOA == pessoa.NomePessoa
}

func encontrarPessoa(dados DadosNormalizados, dados_nn DadosNaoNormalizados) int {
	id_pessoa := 0

	for _, pessoa := range dados.Pessoas {
		if checarIgualdadePessoa(dados_nn, pessoa) {
			id_pessoa = pessoa.IdPessoa
		}
	}

	return id_pessoa
}

func checarIgualdadeContrato(dados DadosNormalizados, dados_nn DadosNaoNormalizados, contrato Contrato) bool {
	return dados_nn.NOM_CONTRATO_TIPO == contrato.NomeContrato &&
		dados_nn.NDA_LICITACAO == contrato.NumLicitacao &&
		checarIgualdadePessoa(dados_nn, dados.Pessoas[contrato.IdPessoa])
}

func encontrarContrato(dados DadosNormalizados, dados_nn DadosNaoNormalizados) int {
	id_contrato := 0

	for _, contrato := range dados.Contratos {
		if checarIgualdadeContrato(dados, dados_nn, contrato) {
			id_contrato = contrato.IdContrato
		}
	}

	return id_contrato
}

func checarIgualdadeEmpenho(dados DadosNormalizados, dados_nn DadosNaoNormalizados, empenho Empenho) bool {
	ANO_DOCUMENTO_INT, _ := strconv.Atoi(dados_nn.ANO_DOCUMENTO)
	NRO_EMPENHO_INT, _ := strconv.Atoi(dados_nn.NRO_EMPENHO)

	return ANO_DOCUMENTO_INT == empenho.AnoDocumento &&
		NRO_EMPENHO_INT == empenho.NumEmpenho &&
		dados_nn.NRO_ORDEM_FILA == empenho.NumOrdemFila &&
		dados_nn.MOTIVO_QUEBRA_ORDEM_CRONOLOGICA == empenho.MotivoQuebra &&
		checarIgualdadeContrato(dados, dados_nn, dados.Contratos[empenho.IdContrato])
}

func encontrarEmpenho(dados DadosNormalizados, dados_nn DadosNaoNormalizados) int {
	id_empenho := 0

	for _, empenho := range dados.Empenhos {
		if checarIgualdadeEmpenho(dados, dados_nn, empenho) {
			id_empenho = empenho.IdEmpenho
		}
	}

	return id_empenho
}

func checarIgualdadeFonte(dados_nn DadosNaoNormalizados, fonte Fonte) bool {
	return dados_nn.NOM_FONTE_RECURSO_TCE == fonte.NomFonteRecurso
}

func encontrarFonte(dados DadosNormalizados, dados_nn DadosNaoNormalizados) int {
	id_fonte := 0

	for _, fonte := range dados.Fontes {
		if checarIgualdadeFonte(dados_nn, fonte) {
			id_fonte = fonte.IdFonte
		}
	}

	return id_fonte
}

func checarIgualdadeLiquidacao(dados DadosNormalizados, dados_nn DadosNaoNormalizados, liquidacao Liquidacao) bool {
	NRO_LIQUIDACAO_INT, _ := strconv.Atoi(dados_nn.NRO_LIQUIDACAO)

	return NRO_LIQUIDACAO_INT == liquidacao.NumLiquidacao &&
		dados_nn.DAT_LIQUIDACAO == liquidacao.DataLiquidacao &&
		dados_nn.DAT_LIQUIDACAO_VENCIMENTO == liquidacao.DataLiquidacaoVencimento &&
		checarIgualdadeEmpenho(dados, dados_nn, dados.Empenhos[liquidacao.IdEmpenho])
}

func encontrarLiquidacao(dados DadosNormalizados, dados_nn DadosNaoNormalizados) int {
	id_liquidacao := 0

	for _, liquidacao := range dados.Liquidacoes {
		if checarIgualdadeLiquidacao(dados, dados_nn, liquidacao) {
			id_liquidacao = liquidacao.IdLiquidacao
		}
	}

	return id_liquidacao
}

func checarIgualdadeDespesa(dados DadosNormalizados, dados_nn DadosNaoNormalizados, despesa Despesa) bool {
	NRO_PAGAMENTO_ORDEM_INT, _ := strconv.Atoi(dados_nn.NRO_PAGAMENTO_ORDEM)

	return dados_nn.SDL_LIQUIDACAO_ORDEM_FINAL == despesa.SaldoLiquidacaoFinal &&
		dados_nn.VLR_LIQUIDACAO == despesa.ValorLiquidacao &&
		NRO_PAGAMENTO_ORDEM_INT == despesa.NumPagamentoOrdem &&
		checarIgualdadeLiquidacao(dados, dados_nn, dados.Liquidacoes[despesa.IdLiquidacao])
}

func checarIgualdadeEmpenhoFonte(dados DadosNormalizados, dados_nn DadosNaoNormalizados, empenho_fonte EmpenhoFonte) bool {
	return checarIgualdadeEmpenho(dados, dados_nn, dados.Empenhos[empenho_fonte.IdEmpenho]) &&
		checarIgualdadeFonte(dados_nn, dados.Fontes[empenho_fonte.IdFonte])
}
