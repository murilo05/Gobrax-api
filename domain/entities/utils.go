package entities

type Error struct {
	ErrorMessage string `json:"errorMessage"`
}

type Env struct {
	HOST        string
	PORT        string
	DBHOST      string
	DBPORT      int
	DBUSER      string
	DBPASSAWORD string
	DBNAME      string
	TIMEOUT     int
}
