package dto

type ProductRes struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainRA(res int64) ProductRes {
	return ProductRes{RowsAffected: res}
}
