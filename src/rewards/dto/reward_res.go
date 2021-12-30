package dto

type RowsAffected struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainRA(res int64) RowsAffected {
	return RowsAffected{RowsAffected: res}
}
