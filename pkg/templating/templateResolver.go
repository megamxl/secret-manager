package templating

import (
	"bytes"
	"log"
	"text/template"
)

func EvaluateTemplate(content string, data map[string]string) (*bytes.Buffer, error) {

	parse, err := template.New("test").Parse(content)
	if err != nil {
		log.Println(err)
	}

	var buf bytes.Buffer
	err = parse.Execute(&buf, data)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}
