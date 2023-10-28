package models

type UserRole int

const (
	Other UserRole = iota + 1
	Admin
	Regular
)

// String - Creating common behavior - give the type a String function
func (r UserRole) String() string {
	return [...]string{"Other", "Admin", "Regular"}[r-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex function
func (r UserRole) EnumIndex() int {
	return int(r)
}
