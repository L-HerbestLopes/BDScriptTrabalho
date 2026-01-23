package main

// struct para guardar dados antes da normalização
type DadosNaoNormalizados struct {
	ANO_DOCUMENTO                   string
	NOM_CONTRATO_TIPO               string
	NOM_FONTE_RECURSO_TCE           string
	NRO_ORDEM_FILA                  int
	NRO_LIQUIDACAO                  string
	DAT_LIQUIDACAO                  string
	DAT_LIQUIDACAO_VENCIMENTO       string
	NRO_EMPENHO                     string
	NRO_PAGAMENTO_ORDEM             string
	NOM_PESSOA                      string
	SDL_LIQUIDACAO_ORDEM_FINAL      float64
	VLR_LIQUIDACAO                  float64
	NDA_LICITACAO                   string
	MOTIVO_QUEBRA_ORDEM_CRONOLOGICA string
}

// structs das entidades normalizadas
type Pessoa struct {
	IdPessoa   int
	NomePessoa string
}

type Contrato struct {
	IdContrato   int
	IdPessoa     int
	NomeContrato string
	NumLicitacao string
}

type Empenho struct {
	IdEmpenho    int
	IdContrato   int
	AnoDocumento int
	NumEmpenho   int
	NumOrdemFila int
	MotivoQuebra string
}

type Despesa struct {
	IdDespesa            int
	IdLiquidacao         int
	SaldoLiquidacaoFinal float64
	NumPagamentoOrdem    int
	ValorLiquidacao      float64
}

type Fonte struct {
	IdFonte         int
	NomFonteRecurso string
}

type EmpenhoFonte struct {
	IdFonte   int
	IdEmpenho int
}

type Liquidacao struct {
	IdLiquidacao             int
	IdEmpenho                int
	NumLiquidacao            int
	DataLiquidacao           string
	DataLiquidacaoVencimento string
}

// struct que guarda todas as entidades pós normalizacação
type DadosNormalizados struct {
	Pessoas       []Pessoa
	Contratos     []Contrato
	Empenhos      []Empenho
	Despesas      []Despesa
	Fontes        []Fonte
	EmpenhoFontes []EmpenhoFonte
	Liquidacoes   []Liquidacao
}
