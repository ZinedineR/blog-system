package model

type ListReq struct {
	Page   PaginationParam
	Order  OrderParam
	Filter FilterParams
}
