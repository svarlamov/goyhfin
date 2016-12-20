package goyhfin

import "errors"

var (
	InvalidYahooFinanceResponseLengthError        error = errors.New("Invalid yahoo finance query result array length")
	InvalidYahooFinanceResponseNotEnoughDataError error = errors.New("Invalid yahoo finance query result not enough data returned")
)
