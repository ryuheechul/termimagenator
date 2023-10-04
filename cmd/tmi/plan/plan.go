package plan

import (
	"fmt"
	"strings"
)

type Plan struct {
	Desc   string
	Images []string
	Result string
}

func InitialPlan(images []string) Plan {
	otherCounts := len(images) - 1
	first := images[0]
	//
	desc := fmt.Sprintf("`%s` image", first)
	if otherCounts > 0 {
		desc = fmt.Sprintf("`%s` and %d other images", first, otherCounts)
	}

	return Plan{
		desc,
		images,
		"",
	}
}

func (p *Plan) UpdateResult(untagged []string, deleted []string) {
	var b strings.Builder

	if len(untagged) > 1 {
		fmt.Fprintf(&b, "- `%s` and %d others are untagged\n", untagged[0], len(untagged)-1)
	} else if len(untagged) > 0 {
		fmt.Fprintf(&b, "- `%s` is untagged\n", untagged[0])
	}

	if len(deleted) > 1 {
		fmt.Fprintf(&b, "- `%s` and %d others are deleted\n", deleted[0], len(deleted)-1)
	} else if len(deleted) > 0 {
		fmt.Fprintf(&b, "- `%s` is deleted\n", deleted[0])
	}

	p.Result = b.String()
}
