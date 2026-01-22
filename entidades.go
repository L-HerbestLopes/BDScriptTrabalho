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
	SDL_LIQUIDACAO_ORDEM_FINAL      int
	VLR_LIQUIDACAO                  float64
	NDA_LICITACAO                   string
	MOTIVO_QUEBRA_ORDEM_CRONOLOGICA string
}

