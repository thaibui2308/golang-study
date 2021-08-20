package models

type Response struct {
	Updated     int64  `json:"updated"`
	Country     string `json:"country"`
	CountryInfo struct {
		ID   int     `json:"_id"`
		Iso2 string  `json:"iso2"`
		Iso3 string  `json:"iso3"`
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
		Flag string  `json:"flag"`
	} `json:"countryInfo"`
	Cases               int `json:"cases"`
	TodayCases          int `json:"todayCases"`
	Deaths              int `json:"deaths"`
	TodayDeaths         int `json:"todayDeaths"`
	Recovered           int `json:"recovered"`
	Active              int `json:"active"`
	Critical            int `json:"critical"`
	CasesPerOneMillion  int `json:"casesPerOneMillion"`
	DeathsPerOneMillion int `json:"deathsPerOneMillion"`
	Tests               int `json:"tests"`
	TestsPerOneMillion  int `json:"testsPerOneMillion"`
}
