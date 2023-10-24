package teamstats_serializer

import "capstone-cs.eng.utah.edu/mapeleven/mapeleven/pkg/ent"

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
