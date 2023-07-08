package image

type Image struct {
	RepoTag string
	ID      string
}

type Formatter interface {
	Format(img Image) (string, error)
}
