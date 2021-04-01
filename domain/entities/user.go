package entities

type User struct {
	Id    *int
	Name  *string
	Pwd   *string
	Email *string
	Age   *int
}

type NewUser struct {
	Name  *string
	Pwd   *string
	Email *string
	Age   *int
}

type UserFilter struct {
	Name  *string
	Email *string
	Age   *int
}

type UserPagination struct {
	First   *int
	Last    *int
	OrderBy *UserOrder
}

const (
	Name  OrderField = "NAME"
	Email OrderField = "EMAIL"
	Age   OrderField = "AGE"
)

type UserOrder struct {
	Direction *OrderDirection
	Field     *OrderField
}
