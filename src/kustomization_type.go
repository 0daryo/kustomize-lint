package src

type Kustomization struct {
	Images []Image `yaml:"images"`
}

type Image struct {
	Name   string `yaml:"name"`
	NewTag string `yaml:"newTag"`
}
