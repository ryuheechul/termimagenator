package multiselect

import (
	"github.com/charmbracelet/bubbles/table"
)

type TableFocusRotator struct {
	styleFocused *table.Styles
	styleBlurred *table.Styles
}

func NewTableFocusRotator(styleFocused *table.Styles, styleBlurred *table.Styles) TableFocusRotator {
	return TableFocusRotator{
		styleFocused: styleFocused,
		styleBlurred: styleBlurred,
	}
}

func findFocused(tbls []*table.Model) int {
	for i, t := range tbls {

		if t.Focused() {
			return i
		}
	}

	return 0
}

func (tfr *TableFocusRotator) focusNext(tbls []*table.Model) {
	numberOfTables := len(tbls)
	focused := findFocused(tbls)
	prev := tbls[focused]

	if numberOfTables-1 == focused {
		focused = 0
	} else {
		focused += 1
	}
	next := tbls[focused]

	prev.Blur()
	prev.SetStyles(*tfr.styleBlurred)
	next.Focus()
	next.SetStyles(*tfr.styleFocused)
}

func (tfr *TableFocusRotator) focusThis(tbls []*table.Model, this *table.Model) {
	for _, t := range tbls {
		t.Blur()
		t.SetStyles(*tfr.styleBlurred)
	}

	this.SetStyles(*tfr.styleFocused)
	this.Focus()
}
