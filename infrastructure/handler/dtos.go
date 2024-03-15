package handler

type Trucker struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Age         int    `json:"age"`
	Nationality string `json:"nationatily"`
}

type Vehicle struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Brand string `json:"brand"`
	Color string `json:"color"`
	Year  int    `json:"year"`
	Plate string `json:"plate"`
}
