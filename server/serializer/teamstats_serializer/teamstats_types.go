package teamstats_serializer

type Win struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type Lose struct {
	Home string `json:"home"`
	Away string `json:"away"`
}

type MinuteStats struct {
	Total      int    `json:"total"`
	Percentage string `json:"percentage"`
}

type MinuteDetail struct {
	Min0to15    MinuteStats `json:"0-15"`
	Min16to30   MinuteStats `json:"16-30"`
	Min31to45   MinuteStats `json:"31-45"`
	Min46to60   MinuteStats `json:"46-60"`
	Min61to75   MinuteStats `json:"61-75"`
	Min76to90   MinuteStats `json:"76-90"`
	Min91to105  MinuteStats `json:"91-105"`
	Min106to120 MinuteStats `json:"106-120"`
}

type ScoreDetail struct {
	Total   HomeAwayTotalInt    `json:"total"`
	Average HomeAwayTotalString `json:"average"`
	Minute  MinuteDetail        `json:"minute"`
}

type HomeAwayTotalInt struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total,omitempty"`
}

type HomeAwayTotalString struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total,omitempty"`
}

type AverageStats struct {
	Home  string `json:"home"`
	Away  string `json:"away"`
	Total string `json:"total"`
}
