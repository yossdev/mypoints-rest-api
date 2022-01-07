package dto

type AccountCreated struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainAC(res int64) AccountCreated {
	return AccountCreated{RowsAffected: res}
}

type AccountUpdated struct {
	RowsAffected int64 `json:"rows_affected"`
}

func FromDomainAU(res int64) AccountUpdated {
	return AccountUpdated{RowsAffected: res}
}
