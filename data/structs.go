package data

type Unit struct {
	Id   int64    `json:"id"`
	Name string `json:"name"`
}

type Employee struct {
	Id      int64    `json:"id"`
	Name    string `json:"name"`
	Age     int64    `json:"age"`
	Unit_id int64    `json:"unit_id"`
}
