package loader

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/cmepw/221b/encryption"
	"github.com/cmepw/221b/templates"
)

type ChaCha20 struct {
	baseLoader
	encryption.ChaCha20
}

func (a ChaCha20) Load(content, key []byte) ([]byte, error) {
	tmpl, err := template.New("loader").Funcs(template.FuncMap{
		"key": func() string {
			return string(key)
		},
		"base_path": func() string {
			return basepath
		},
		"shellcode": func() string {
			result := []string{}

			for _, b := range content {
				result = append(result, fmt.Sprintf("0x%02x", b))
			}

			return strings.Join(result, ", ") + ","
		},
	}).Parse(templates.ChaCha20Tmpl)
	if err != nil {
		return nil, err
	}

	result := new(bytes.Buffer)
	err = tmpl.Execute(result, nil)

	return result.Bytes(), err
}
