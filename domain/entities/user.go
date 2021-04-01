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
	UserNameField  string = "NAME"
	UserEmailField string = "EMAIL"
	UserAgeField   string = "AGE"
)

type UserOrder struct {
	Direction *OrderDirection
	Field     *string
}
