package main

import (
	"flag"
	"fmt"
	"os"
)

type Temperatura struct {
	valor float64
	de    string
	para  string
}

func (t Temperatura) converter() (float64, error) {
	switch t.de {
	case "C":
		switch t.para {
		case "F":
			return (t.valor * 9 / 5) + 32, nil
		case "K":
			return t.valor + 273.15, nil
		case "C":
			return t.valor, nil
		default:
			return 0, fmt.Errorf("unidade de destino inválida: %s", t.para)
		}
	case "F":
		switch t.para {
		case "C":
			return (t.valor - 32) * 5 / 9, nil
		case "K":
			return (t.valor-32)*5/9 + 273.15, nil
		case "F":
			return t.valor, nil
		default:
			return 0, fmt.Errorf("unidade de destino inválida: %s", t.para)
		}
	case "K":
		switch t.para {
		case "C":
			return t.valor - 273.15, nil
		case "F":
			return (t.valor-273.15)*9/5 + 32, nil
		case "K":
			return t.valor, nil
		default:
			return 0, fmt.Errorf("unidade de destino inválida: %s", t.para)
		}
	default:
		return 0, fmt.Errorf("unidade de origem inválida: %s", t.de)
	}
}

func main() {
	// Configuração das flags
	tempValor := flag.Float64("valor", 0, "Valor da temperatura a ser convertida")
	tempDe := flag.String("de", "C", "Unidade de origem (C, F, K)")
	tempPara := flag.String("para", "F", "Unidade de destino (C, F, K)")

	flag.Parse()

	temp := Temperatura{
		valor: *tempValor,
		de:    *tempDe,
		para:  *tempPara,
	}

	// Convertendo temperatura
	resultado, err := temp.converter()
	if err != nil {
		fmt.Printf("Erro na conversão: %v", err)
		os.Exit(1)
	}

	//Exibir resultado
	fmt.Printf("%.2fº%s = %.2fº%s", temp.valor, temp.de, resultado, temp.para)
}
