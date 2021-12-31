package dto

type RewardRes struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainRA(res int64) RewardRes {
	return RewardRes{RowsAffected: res}
}
