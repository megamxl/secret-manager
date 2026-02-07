package templating

import (
	"bytes"
	"text/template"
)

func EvaluateTemplate(content string, data map[string]string, secretName string) (*bytes.Buffer, error) {

	parse, err := template.New(secretName).Option("missingkey=error").Parse(content)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	err = parse.Execute(&buf, data)
	if err != nil {
		return nil, err
	}

	return &buf, nil
}
