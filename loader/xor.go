package loader

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/cmepw/221b/templates"
)

type Xor struct {
	key []byte
	baseLoader
}

func NewXorLoader(key []byte) *Xor {
	return &Xor{key: key}
}

func (x Xor) Load(content []byte) ([]byte, error) {
	tmpl, err := template.New("loader").Funcs(template.FuncMap{
		"key": func() string {
			return string(x.key)
		},
		"shellcode": func() string {
			result := []string{}

			for _, b := range content {
				result = append(result, fmt.Sprintf("0x%02x", b))
			}

			return strings.Join(result, ", ") + ","
		},
	}).Parse(templates.XorTmpl)
	if err != nil {
		return nil, err
	}

	result := new(bytes.Buffer)
	err = tmpl.Execute(result, nil)

	return result.Bytes(), err
}
