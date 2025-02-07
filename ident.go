package golang

import (
	"fmt"
	"strings"
)

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

func (i *Ident) String(imports *ImportMap) string {
	if i.From == "" {
		return i.Name
	}
	localName := imports.AddImport(i.From)
	return fmt.Sprintf("%s.%s", localName, i.Name)
}
