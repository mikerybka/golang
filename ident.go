package golang

import "strings"

func NewIdent(s string) *Ident {
	i := strings.LastIndex(s, ".")
	if i == -1 {
		return &Ident{
			Name: s,
		}
	}
	return &Ident{
		From: s[:i],
		Name: s[i+1:],
	}
}

type Ident struct {
	From string
	Name string
}
