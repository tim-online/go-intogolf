package intogolf

import (
	"github.com/aodin/date"
	"github.com/cockroachdb/apd"
)

type Debiteur struct {
	Nummer             int
	Naam               string
	Adres1             string
	Adres2             string
	Postcode           string
	Plaats             string
	Land               string
	Valuta             string
	Telefoon           string
	Fax                string
	Contactpersoon     string
	Bankrekeningnummer string
	BTWNummer          string
	Email              string
	IBAN               string
	BIC                string
}

type Boeking struct {
	Regelnummer     int
	Dagboek         string
	Periode         int
	Jaar            int
	Boekstuknummer  int
	Boekdatum       date.Date
	Grootboeknummer int
	Bedrag          apd.Decimal
	Valuta          string
	FLCCode         string
	FPRCode         string
}

type Factuur struct {
	Dagboek        string
	Periode        int
	Jaar           int
	Factuurnummer  string
	Omschrijving   string
	Factuurdatum   date.Date
	Debiteurnummer int
	Vervaldatum    date.Date
	Regels         Factuurregels
}

type Factuurregels []Factuurregel

type Factuurregel struct {
	Regelnummer     int
	Grootboeknummer int
	Bedrag          apd.Decimal
	Valuta          string
	BTW             apd.Decimal
}
