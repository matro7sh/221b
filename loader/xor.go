package loader

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/cmepw/221b/encryption"
	"github.com/cmepw/221b/templates"
)

type Xor struct {
	baseLoader
	encryption.Xor
}

func (x Xor) Load(content, key []byte) ([]byte, error) {
	tmpl, err := template.New("loader").Funcs(template.FuncMap{
		"key": func() string {
			return string(key)
		},
		"basepath": func() string {
			return basepath
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
