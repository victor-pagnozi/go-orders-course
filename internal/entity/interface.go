package entity

type OrderRepositoryInterface interface {
	Save(order *Order) error
	GetTotal() (total int, err error)
}
