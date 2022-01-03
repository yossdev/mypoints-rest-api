package dto

type TransactionRes struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomain(res int64) TransactionRes {
	return TransactionRes{RowsAffected: res}
}
