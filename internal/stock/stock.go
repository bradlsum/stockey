package stock

type Stock struct {
	StockingId  int    `json:"stocking_id"`
	BodyOfWater string `json:"body_of_water"`
	WaterId     int    `json:"water_id"`
	Region      string `json:"region"`
	ReportDate  string `json:"report_date"`
}
