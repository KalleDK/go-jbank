package transaction

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"
)

const (
	dateIndex = iota
	_
	textIndex
	_
	amountIndex
	balanceIndex
	commentIndex
	reconciledIndex
	flagIndex
	accountIndex
	_
	categoryIndex
	subCategoryIndex
)

// Transaction holds the information about a single transaction
type Transaction struct {
	Date        time.Time `csv:"Dato"`
	Text        string    `csv:"Tekst"`
	Amount      int64     `csv:"Beløb"`
	Balance     int64     `csv:"Saldo"`
	Comment     string    `csv:"Kommentar"`
	Reconciled  bool      `csv:"Afstemt"`
	Flag        bool      `csv:"Flag"`
	Account     string    `csv:"Konto"`
	Category    string    `csv:"Kategori"`
	SubCategory string    `csv:"Underkategori"`
}

func parseAmount(s string) (i int64, err error) {
	s = strings.Replace(s, ",", "", -1)
	s = strings.Replace(s, ".", "", -1)
	v, err := strconv.Atoi(s)
	return (int64)(v), err
}

func parseBool(s string) (b bool, err error) {
	if s == "ja" {
		return true, nil
	}
	if s == "nej" {
		return false, nil
	}
	return false, fmt.Errorf("invalid error format: %v", s)
}

// Parse parses a single record line to a Transaction
func Parse(record []string) (t Transaction, err error) {
	date, err := time.Parse("02.01.2006", record[dateIndex])
	if err != nil {
		return t, err
	}

	amount, err := parseAmount(record[amountIndex])
	if err != nil {
		return t, err
	}

	balance, err := parseAmount(record[balanceIndex])
	if err != nil {
		return t, err
	}

	reconciled, err := parseBool(record[reconciledIndex])
	if err != nil {
		return t, err
	}

	flag, err := parseBool(record[flagIndex])
	if err != nil {
		return t, err
	}

	t = Transaction{
		Date:        date,
		Text:        record[textIndex],
		Amount:      amount,
		Balance:     balance,
		Comment:     record[commentIndex],
		Reconciled:  reconciled,
		Flag:        flag,
		Account:     record[accountIndex],
		Category:    record[categoryIndex],
		SubCategory: record[subCategoryIndex],
	}
	return t, nil
}

func FromCSV(cr *csv.Reader) (ts []Transaction, err error) {
	cr.Comma = ';'

	ts = []Transaction{}

	// TODO Remove first line but also verify it
	cr.Read()
	_, err = cr.Read()
	if err == io.EOF {
		return ts, nil
	}

	if err != nil {
		return ts, err
	}

	for {

		record, err := cr.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			return ts, err
		}

		transaction, err := Parse(record)
		if err != nil {
			return ts, err
		}

		ts = append(ts, transaction)
	}
	return ts, nil
}

func FromCSVStream(r io.Reader) (ts []Transaction, err error) {
	cr := csv.NewReader(r)
	return FromCSV(cr)
}
