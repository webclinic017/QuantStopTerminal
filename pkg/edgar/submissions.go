package edgar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

/*
Each entity’s current filing history is available at the following URL:

https://data.sec.gov/submissions/CIK##########.json
Where the ########## is the entity’s 10-digit Central Index Key (CIK), including leading zeros.

This JSON data structure contains metadata such as current name, former name, and stock exchanges
and ticker symbols of publicly-traded companies. The object’s property path contains at least
one year’s of filing or to 1,000 (whichever is more) of the most recent filings in a compact
columnar data array. If the entity has additional filings, files will contain an array of
additional JSON files and the date range for the filings each one contains.
*/

type Submissions struct {
	Cik                               string   `json:"cik"`
	EntityType                        string   `json:"entityType"`
	Sic                               string   `json:"sic"`
	SicDescription                    string   `json:"sicDescription"`
	InsiderTransactionForOwnerExists  int      `json:"insiderTransactionForOwnerExists"`
	InsiderTransactionForIssuerExists int      `json:"insiderTransactionForIssuerExists"`
	Name                              string   `json:"name"`
	Tickers                           []string `json:"tickers"`
	Exchanges                         []string `json:"exchanges"`
	Ein                               string   `json:"ein"`
	Description                       string   `json:"description"`
	Website                           string   `json:"website"`
	InvestorWebsite                   string   `json:"investorWebsite"`
	Category                          string   `json:"category"`
	FiscalYearEnd                     string   `json:"fiscalYearEnd"`
	StateOfIncorporation              string   `json:"stateOfIncorporation"`
	StateOfIncorporationDescription   string   `json:"stateOfIncorporationDescription"`
	Addresses                         struct {
		Mailing struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"mailing"`
		Business struct {
			Street1                   string      `json:"street1"`
			Street2                   interface{} `json:"street2"`
			City                      string      `json:"city"`
			StateOrCountry            string      `json:"stateOrCountry"`
			ZipCode                   string      `json:"zipCode"`
			StateOrCountryDescription string      `json:"stateOrCountryDescription"`
		} `json:"business"`
	} `json:"addresses"`
	Phone       string `json:"phone"`
	Flags       string `json:"flags"`
	FormerNames []struct {
		Name string    `json:"name"`
		From time.Time `json:"from"`
		To   time.Time `json:"to"`
	} `json:"formerNames"`
	Filings struct {
		Recent struct {
			AccessionNumber       []string    `json:"accessionNumber"`
			FilingDate            []string    `json:"filingDate"`
			ReportDate            []string    `json:"reportDate"`
			AcceptanceDateTime    []time.Time `json:"acceptanceDateTime"`
			Act                   []string    `json:"act"`
			Form                  []string    `json:"form"`
			FileNumber            []string    `json:"fileNumber"`
			FilmNumber            []string    `json:"filmNumber"`
			Items                 []string    `json:"items"`
			Size                  []int       `json:"size"`
			IsXBRL                []int       `json:"isXBRL"`
			IsInlineXBRL          []int       `json:"isInlineXBRL"`
			PrimaryDocument       []string    `json:"primaryDocument"`
			PrimaryDocDescription []string    `json:"primaryDocDescription"`
		} `json:"recent"`
		Files []struct {
			Name        string `json:"name"`
			FilingCount int    `json:"filingCount"`
			FilingFrom  string `json:"filingFrom"`
			FilingTo    string `json:"filingTo"`
		} `json:"files"`
	} `json:"filings"`
}

func (edgar *Edgar) GetSubmissions(cik string) (*Submissions, error) {

	responseObject := &Submissions{}
	url := "https://data.sec.gov/submissions/CIK" + cik + ".json"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := edgar.Do(req)
	if err != nil {
		return nil, err
	}

	// Close request per https://pkg.go.dev/net/http
	defer resp.Body.Close()

	// Read all bytes from response body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err.Error())
	}

	// Unmarshal/Decode the JSON to the interface
	err = json.Unmarshal(bodyBytes, &responseObject)
	if err != nil {
		return nil, err
	}

	return responseObject, nil
}
