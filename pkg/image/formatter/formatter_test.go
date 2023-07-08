package formatter

import (
	"testing"

	"github.com/ryuheechul/termimagenator/pkg/image"
	"github.com/samber/mo"
)

func Test_defaultFormatter_Format(t *testing.T) {

	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		d    defaultFormatter
		args args
		want mo.Either[string, error]
	}{
		{
			name: "happy",
			args: args{
				img: image.Image{
					RepoTag: "image:tag",
					ID:      "sha256:abcdef1234567890",
				},
			},
			want: mo.Left[string, error]("image:tag abcdef123456"),
		},
		{
			name: "no colon",
			args: args{
				img: image.Image{
					RepoTag: "image:tag",
					ID:      "no-colon",
				},
			},
			want: mo.Right[string, error](
				&IdNoColonError{err: "no-colon"},
			),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := defaultFormatter{}
			tt.want.Match(func(want string) mo.Either[string, error] {
				if got, _ := d.Format(tt.args.img); got != want {
					t.Errorf("defaultFormatter.Format() = %v, want %v", got, want)
				}

				return mo.Left[string, error](want)
			}, func(want error) mo.Either[string, error] {
				if _, got := d.Format(tt.args.img); got.Error() != want.Error() {
					t.Errorf("error for defaultFormatter.Format() = %v, want %v", got, want)
				}

				return mo.Right[string, error](want)
			})
		})
	}
}
