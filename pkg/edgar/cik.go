package edgar

type CikResponse struct {
	Fields []string        `json:"fields"`
	Data   [][]interface{} `json:"data"`
}

/*
https://www.sec.gov/files/company_tickers_exchange.json
data array "fields":["cik","name","ticker","exchange"],"data":[[320193,"Apple Inc.","AAPL","Nasdaq"], ...
*/
