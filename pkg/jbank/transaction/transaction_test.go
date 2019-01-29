package transaction

import (
	"fmt"
	"strings"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	var csvBlob = `"Dato";"";"Tekst";"";"Beløb";"Saldo";"Kommentar";"Afstemt";"Flag";"Konto";"";"Kategori";"Underkategori"
"31.10.2018";"";"DK Forretning A, Trøjborg";"";"-154,35";"3.439,04";"";"nej";"nej";"Fælleskonto";"";"Mad og indkøb";"Dagligvarer"
"30.10.2018";"";"DK Forretning B";"";"-40,90";"3.593,39";"";"nej";"nej";"Fælleskonto";"";"Mad og indkøb";"Dagligvarer"`

	fmt.Println(string(csvBlob[:]))

	transactions, _ := FromCSV(strings.NewReader(csvBlob))
	fmt.Println(transactions)
}
