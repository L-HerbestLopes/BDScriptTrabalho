package main

func NormalizarDados(dados_nn []DadosNaoNormalizados) DadosNormalizados {
	// inicializa struct
	dados := DadosNormalizados{}

	// preenche Pessoas
	for _, linha := range dados_nn {
		// verifica se pessoa já existe nos dados
		pessoa_existe := false

		for _, pessoa := range dados.Pessoas {
			if linha.NOM_PESSOA == pessoa.NomePessoa {
				pessoa_existe = true
				break
			}
		}

		if !pessoa_existe {
			pessoa := Pessoa{len(dados.Pessoas), linha.NOM_PESSOA}
			dados.Pessoas = append(dados.Pessoas, pessoa)
		}
	}

	// retorna dados normalizados
	return dados
}
