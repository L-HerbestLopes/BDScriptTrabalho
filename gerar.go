package main

import (
	"os"
	"log"
	"strconv"
)

func GerarScript(dados DadosNormalizados) {
	arquivo, err := os.OpenFile("GRUPO04-DML.SQL", os.O_WRONLY|os.O_CREATE|os.O_TRUNC|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer arquivo.Close()

	GerarDML(arquivo, dados)
}

func GerarDML(arquivo *os.File, dados DadosNormalizados) {
	arquivo.WriteString("-- Pessoa\n")
	for _, pessoa := range(dados.Pessoas) {
		values := "(\"" + pessoa.NomePessoa + "\");\n"
		arquivo.WriteString("INSERT INTO Pessoa (nomePessoa) VALUES " + values)
	}

	arquivo.WriteString("\n-- Contrato\n")
	for _, contrato := range(dados.Contratos) {
		values := "(" + strconv.Itoa(contrato.IdPessoa) + ", \"" + contrato.NomeContrato + "\", \"" + contrato.NumLicitacao + "\");\n"
		arquivo.WriteString("INSERT INTO Contrato (idPessoa, nomeContrato, numeroLicitacao) VALUES " + values)
	}

	arquivo.WriteString("\n-- Empenho\n")
	for _, empenho := range(dados.Empenhos) {
		values := "(" + strconv.Itoa(empenho.IdContrato) + ", " + strconv.Itoa(empenho.AnoDocumento) + ", " + strconv.Itoa(empenho.NumEmpenho) + ", " + strconv.Itoa(empenho.NumOrdemFila) + ", \"" + empenho.MotivoQuebra + "\");\n"
		arquivo.WriteString("INSERT INTO Empenho (idContrato, anoDocumento, numeroEmpenho, numeroOrdemFila, motivoQuebra) VALUES " + values)
	}

	arquivo.WriteString("\n-- Fonte\n")
	for _, fonte := range(dados.Fontes) {
		values := "(\"" + fonte.NomFonteRecurso + "\");\n"
		arquivo.WriteString("INSERT INTO Fonte (nomeFonteRecurso) VALUES " + values)
	}

	arquivo.WriteString("\n-- EmpenhoFonte\n")
	for _, empenhoFonte := range(dados.EmpenhoFontes) {
		values := "(" + strconv.Itoa(empenhoFonte.IdFonte) + ", " + strconv.Itoa(empenhoFonte.IdEmpenho) + ");\n"
		arquivo.WriteString("INSERT INTO EmpenhoFonte (idFonte, idEmpenho) VALUES " + values)
	}

	arquivo.WriteString("\n-- Liquidacao\n")
	for _, liquidacao := range(dados.Liquidacoes) {
		values := "(" + strconv.Itoa(liquidacao.IdEmpenho) + ", " + strconv.Itoa(liquidacao.NumLiquidacao) + ", \"" + liquidacao.DataLiquidacao + "\", \"" + liquidacao.DataLiquidacaoVencimento + "\");\n"
		arquivo.WriteString("INSERT INTO Liquidacao (idEmpenho, numeroLiquidacao, dataLiquidacao, dataLiquidacaoVencimento) VALUES " + values)
	}

	arquivo.WriteString("\n-- Despesa\n")
	for _, despesa := range(dados.Despesas) {
		values := "(" + strconv.Itoa(despesa.IdLiquidacao) + ", " + strconv.FormatFloat(despesa.SaldoLiquidacaoFinal, 'f', -1, 64) + ", " + strconv.Itoa(despesa.NumPagamentoOrdem) + ", " + strconv.FormatFloat(despesa.ValorLiquidacao, 'f', -1, 64) + ");\n"
		arquivo.WriteString("INSERT INTO Despesa (idLiquidacao, saldoLiquidacaoFinal, numeroPagamentoOrdem, valorLiquidacao) VALUES " + values)
	}
}