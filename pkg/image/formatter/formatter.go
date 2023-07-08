package formatter

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
	"text/template"

	"github.com/ryuheechul/termimagenator/pkg/image"
)

var tmpl = template.New("format")

var NoColonError = errors.New("No Colon found")

type IdNoColonError struct {
	err string
}

func noColonError(msg string) error {
	return fmt.Errorf("%w at %s", NoColonError, msg)
}

type defaultFormatter struct{}

func (_ defaultFormatter) Format(img image.Image) (string, error) {
	if !strings.Contains(img.ID, ":") {
		return "", noColonError(img.ID)
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

var DefaultFormatter = &defaultFormatter{}
