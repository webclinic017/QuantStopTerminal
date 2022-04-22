package edgar

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

/*
The company-concept API returns all the XBRL disclosures from a single company (CIK)
and concept (a taxonomy and tag) into a single JSON file, with a separate array of facts
for each units on measure that the company has chosen to disclose
(e.g. net profits reported in U.S. dollars and in Canadian dollars).

https://data.sec.gov/api/xbrl/companyconcept/CIK##########/us-gaap/AccountsPayableCurrent.json
*/

type CompanyConcept struct {
	Cik         int    `json:"cik"`
	Taxonomy    string `json:"taxonomy"`
	Tag         string `json:"tag"`
	Label       string `json:"label"`
	Description string `json:"description"`
	EntityName  string `json:"entityName"`
	Units       struct {
		USD []struct {
			End   string `json:"end"`
			Val   int    `json:"val"`
			Accn  string `json:"accn"`
			Fy    int    `json:"fy"`
			Fp    string `json:"fp"`
			Form  string `json:"form"`
			Filed string `json:"filed"`
			Frame string `json:"frame,omitempty"`
		} `json:"USD"`
	} `json:"units"`
}

func (edgar *Edgar) GetCompanyConcept(cik string) (*CompanyConcept, error) {
	responseObject := &CompanyConcept{}
	url := "https://data.sec.gov/api/xbrl/companyconcept/CIK" + cik + "/us-gaap/AccountsPayableCurrent.json"
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
