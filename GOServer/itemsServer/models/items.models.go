package models

type ItemJson struct {
	Price_for_unit float32 `json:"pfu"`
	Id             int64   `json:"id"`
	Category_id    int64   `json:"c_id"`
	Category_name  string  `json:"c_name"`
	Name           string  `json:"name"`
	Descriptions   string  `json:"desc"`
	Dimension      string  `json:"dim"`
	Count          int64   `json:"count"`
}

type CategoryJson struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}

type StockJson struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	Address    string `json:"address"`
	ItemsCount int    `json:"items_count"`
}

type PurchaseJson struct {
	Id        int64      `json:"id"`
	User_id   int64      `json:"user_id"`
	Items     []ItemJson `json:"items"`
	DateStart string     `json:"date_start"`
	DateEnd   string     `json:"date_end"`
	Price     float32    `json:"price"`
}
