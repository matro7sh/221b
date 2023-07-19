package loader

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/cmepw/221b/encryption"
	"github.com/cmepw/221b/templates"
)

type Aes struct {
	baseLoader
	encryption.Aes
}

func (a Aes) Load(content, key []byte) ([]byte, error) {
	tmpl, err := template.New("loader").Funcs(template.FuncMap{
		"key": func() string {
			return string(key)
		},
		"shellcode": func() string {
			result := []string{}

			for _, b := range content {
				result = append(result, fmt.Sprintf("0x%02x", b))
			}

			return strings.Join(result, ", ") + ","
		},
	}).Parse(templates.AesTmpl)
	if err != nil {
		return nil, err
	}

	result := new(bytes.Buffer)
	err = tmpl.Execute(result, nil)

	return result.Bytes(), err
}
