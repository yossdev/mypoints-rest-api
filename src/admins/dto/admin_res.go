package dto

type AccountCreated struct {
	RowsAffected int64 `json:"rows_affected"`
}

type AccountUpdated struct {
	RowsAffected int64 `json:"rows_affected"`
}
