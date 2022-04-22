package edgar

/*
The xbrl/frames API aggregates one fact for each reporting entity that is last filed that most closely
fits the calendrical period requested. This API supports for annual, quarterly and instantaneous data:

https://data.sec.gov/api/xbrl/frames/us-gaap/AccountsPayableCurrent/USD/CY2019Q1I.json
Where the units of measure specified in the XBRL contains a numerator and a denominator,
these are separated by “-per-” such as “USD-per-shares”. Note that the default unit in XBRL is “pure”.

The period format is CY#### for annual data (duration 365 days +/- 30 days), CY####Q# for quarterly
data (duration 91 days +/- 30 days), and CY####Q#I for instantaneous data. Because company financial
calendars can start and end on any month or day and even change in length from quarter to quarter
according to the day of the week, the frame data is assembled by the dates that best align with a
calendar quarter or year. Data users should be mindful different reporting start and end dates for
facts contained in a frame.
*/

type AccountsPayableCurrent struct {
	Taxonomy    string `json:"taxonomy"`
	Tag         string `json:"tag"`
	Ccp         string `json:"ccp"`
	Uom         string `json:"uom"`
	Label       string `json:"label"`
	Description string `json:"description"`
	Pts         int    `json:"pts"`
	Data        []struct {
		Accn       string `json:"accn"`
		Cik        int    `json:"cik"`
		EntityName string `json:"entityName"`
		Loc        string `json:"loc"`
		End        string `json:"end"`
		Val        int64  `json:"val"`
	} `json:"data"`
}
