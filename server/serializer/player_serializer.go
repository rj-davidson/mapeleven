package serializer

type PlayerSerializer struct {
	slug    string         `json:"slug"`
	Name    string         `json:"name"`
	Age     int            `json:"age"`
	Photo   string         `json:"photo"`
	Country string         `json:"country"`
	Team    TeamSerializer `json:"team"`
}
