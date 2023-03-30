package binding

type User struct {
	Name *IntegerField[int] `json:"name"`
	Age  *IntegerField[uint]
}

func New() *User {
	return &User{
		Name: Integer("name", true, Min(int(1))),
		Age:  Integer("age", false, Max(uint(100))),
	}

}
