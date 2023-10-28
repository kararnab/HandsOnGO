package models

/**
Use Stringer - golang.org/x/tools/cmd/stringer
*/

//go:generate stringer -type=Gender
type Gender int

const (
	Male Gender = iota
	Female
	Transgender
	PreferNotToSay
	Others = Transgender
)
