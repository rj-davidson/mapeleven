package teamstats_serializer

import "mapeleven/db/ent"

type APITSCleanSheet struct {
	Home  int `json:"home"`
	Away  int `json:"away"`
	Total int `json:"total"`
}

func SerializeTSCleanSheet(cleanSheet *ent.TSCleanSheet) APITSCleanSheet {
	return APITSCleanSheet{
		Home:  cleanSheet.Home,
		Away:  cleanSheet.Away,
		Total: cleanSheet.Total,
	}
}
