package models

type ItemJson struct {
	Price_for_unit float32 `json:"pfu"`
	Id             int64   `json:"id"`
	Category_id    int64   `json:"c_id"`
	Category_name  string  `json:"c_name"`
	Name           string  `json:"name"`
	Descriptions   string  `json:"desc"`
	Dimension      string  `json:"dim"`
}

type CategoryJson struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type StockJson struct {
	Id      int64  `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
