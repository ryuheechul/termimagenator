package formatter

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"

	"github.com/ryuheechul/termimagenator/pkg/image"
)

var tmpl = template.New("format")

type IdNoColonError struct {
	err string
}

func (e *IdNoColonError) Error() string {
	return fmt.Sprintf("No Colon found at %s", e.err)
}

type defaultFormatter struct{}

func (_ defaultFormatter) Format(img image.Image) (string, error) {
	if !strings.Contains(img.ID, ":") {
		return "", &IdNoColonError{err: img.ID}
	}
	img.ID = strings.Split(img.ID, ":")[1][:12]

	tmpl, err := tmpl.Parse("{{.RepoTag}} {{.ID}}")
	if err != nil {
		return "", err
	}
	var octets bytes.Buffer

	if err := tmpl.Execute(&octets, img); err != nil {
		return "", err
	}

	return octets.String(), nil
}

var DefaultFormatter = defaultFormatter{}
