package util

import (
	"bytes"
	"text/template"
	"time"
)

var TemplateFunctionMap = template.FuncMap{
	"current_timestamp": getCurrentTimestamp,
}

func RenderTemplate(templateContent string, data interface{}) (string, error) {
	t, err := template.New("template").Funcs(TemplateFunctionMap).Parse(templateContent)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func getCurrentTimestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
