package main

import (
	"bytes"
	"fmt"
	"html/template"
)

const Email = `
<!DOCTYPE html>
<body>
<p>Hello {{ .User }},

This is a email to notify you that you currently have {{ len .NotifyKeys }} keys are over {{ .NotifyAge }} days old:
{{ range .NotifyKeys }}
	<p>Access Key Id: {{ .Id }}, Days Old: {{ .Age }}.</p>
{{ end }}
These keys will be deleted in {{ .DeleteAge }} days.
</body>
</html>
`

type EmailVars struct {
	User       string
	DeleteAge  int
	NotifyAge  int
	NotifyKeys []*AccessKey
}

type AccessKey struct {
	Age int
	Id  string
}

func main() {
	vars := &EmailVars{
		User:      "Bob",
		DeleteAge: 90,
		NotifyAge: 85,
		NotifyKeys: []*AccessKey{
			&AccessKey{85, "abf3345fcab"},
			&AccessKey{89, "6878ffeccff"},
		},
	}

	t, err := template.New("notification").Parse(Email)
	if err != nil {
		panic(err)
	}

	var buf bytes.Buffer
	err = t.Execute(&buf, vars)
	if err != nil {
		panic(err)
	}

	fmt.Println(buf.String())
}
