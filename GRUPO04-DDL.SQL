CREATE DATABASE ordem_cronologica;
USE ordem_cronologica;

CREATE TABLE Pessoa (
	idPessoa INT AUTO_INCREMENT,
	nomePessoa VARCHAR(500),
	PRIMARY KEY (idPessoa)
);

CREATE TABLE Contrato (
	idContrato INT AUTO_INCREMENT,
	idPessoa INT,
	nomeContrato VARCHAR(500),
	numeroLicitacao VARCHAR(255),
	PRIMARY KEY (idContrato),
	FOREIGN KEY (idPessoa)
	REFERENCES Pessoa
);

CREATE TABLE Empenho (
	idEmpenho INT AUTO_INCREMENT,
	idContrato INT,
	anoDocumento INT,
	numeroEmpenho INT,
	numeroOrdemFila INT,
	motivoQuebra VARCHAR(500),
	PRIMARY KEY (idEmpenho),
	FOREIGN KEY (idContrato)
	REFERENCES Contrato
);

CREATE TABLE Fonte (
	idFonte INT AUTO_INCREMENT,
	nomeFonteRecurso VARCHAR(500),
	PRIMARY KEY (idFonte)
);

CREATE TABLE EmpenhoFonte (
	idFonte INT,
	idEmpenho INT,
	PRIMARY KEY (idFonte, idEmpenho),
	FOREIGN KEY (idFonte) REFERENCES Fonte,
	FOREIGN KEY (idEmpenho) REFERENCES Empenho
);

CREATE TABLE Liquidacao (
	idLiquidacao INT AUTO_INCREMENT,
	idEmpenho INT,
	numeroLiquidacao INT,
	dataLiquidacao DATETIME,
	dataLiquidacaoVencimento DATETIME,
	PRIMARY KEY (idLiquidacao),
	FOREIGN KEY (idEmpenho)
	REFERENCES Empenho
);

CREATE TABLE Despesa (
	idDespesa INT AUTO_INCREMENT,
	idLiquidacao INT,
	saldoLiquidacaoFinal FLOAT,
	numeroPagamentoOrdem INT,
	valorLiquidacao FLOAT,
	PRIMARY KEY (idDespesa),
	FOREIGN KEY (idLiquidacao)
	REFERENCES Liquidacao
);