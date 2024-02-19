package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

type Symbol int8

const (
	Empty Symbol = iota
	Escape
	Letter
	Digit
	EscapedSymbol
)

var ErrInvalidString = errors.New("invalid string")

const escapeChar rune = '\\'

func Unpack(input string) (string, error) { //nolint: gocognit
	var (
		outputString                      strings.Builder
		currentSymbolType, prevSymbolType Symbol
		runeToPrint                       rune
		err                               error
	)

	for i, runeValue := range input {
		// определяем тип текущего символа
		if currentSymbolType, err = getRuneType(runeValue); err != nil {
			return "", err
		}

		// возращаем ошибку, если строка начинается с неэкранированной цифры
		if prevSymbolType == Empty && currentSymbolType == Digit {
			return "", ErrInvalidString
		}

		// обрабатыем символы на основе их типов
		switch currentSymbolType { //nolint: exhaustive
		case Escape:
			if prevSymbolType == Letter || prevSymbolType == EscapedSymbol {
				outputString.WriteRune(runeToPrint)
			} else if prevSymbolType == Escape {
				runeToPrint = runeValue
				currentSymbolType = EscapedSymbol
			}
		case Letter:
			if prevSymbolType == Escape {
				return "", ErrInvalidString
			}
			if prevSymbolType == Letter || prevSymbolType == EscapedSymbol {
				outputString.WriteRune(runeToPrint)
			}
			runeToPrint = runeValue
		case Digit:
			if prevSymbolType == Digit {
				return "", ErrInvalidString
			}
			if prevSymbolType == Letter || prevSymbolType == EscapedSymbol {
				outputString.WriteString(strings.Repeat(string(runeToPrint), int(runeValue-'0')))
			} else if prevSymbolType == Escape {
				runeToPrint = runeValue
				currentSymbolType = EscapedSymbol
			}
		default:
			return "", ErrInvalidString
		}

		// обрабатываем последний символ при неоходимости
		if i == len(input)-1 {
			if isPrintLastSymbol, err := checkLastSymbol(currentSymbolType); err != nil {
				return "", err
			} else if isPrintLastSymbol {
				outputString.WriteRune(runeValue)
			}
		}

		// сохраняем тип текущего символа для обработки следующего
		prevSymbolType = currentSymbolType
	}

	return outputString.String(), nil
}

func getRuneType(runeValue int32) (Symbol, error) {
	switch {
	case runeValue == escapeChar:
		return Escape, nil
	case unicode.IsLetter(runeValue):
		return Letter, nil
	case unicode.IsDigit(runeValue):
		return Digit, nil
	default:
		return Empty, ErrInvalidString
	}
}

func checkLastSymbol(currentSymbolType Symbol) (bool, error) {
	switch currentSymbolType { //nolint: exhaustive
	case Letter, EscapedSymbol:
		return true, nil
	case Escape:
		return false, ErrInvalidString
	default:
		return false, nil
	}
}
