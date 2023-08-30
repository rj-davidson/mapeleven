package serializer

type TeamSerializer struct {
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	Badge   string `json:"badge"`
	Country string `json:"country"`
}
