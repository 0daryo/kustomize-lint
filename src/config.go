package src

import "strings"

import "fmt"

type Yaml struct {
	Files []File `yaml:"files"`
}

type File struct {
	Name      string     `yaml:"name"`
	Sentences []Sentence `yaml:"sentences"`
}

type Sentence struct {
	Name    string   `yaml:"name"`
	Include []string `yaml:"include"`
}

func (s *Sentence) hasError(kustomization *Kustomization) bool {
	var hasError bool
	switch s.Name {
	case "newTag":
		errorImage := make([]string, 0)
		for _, image := range kustomization.Images {
			var included bool
			for _, inc := range s.Include {
				if strings.Contains(image.NewTag, inc) {
					included = true
				}
			}
			if !included {
				errorImage = append(errorImage, image.Name)
			}
		}
		for _, e := range errorImage {
			fmt.Printf("%s.newTag does not contains %s\n", e, s.Include)
			hasError = true
		}
	case "name":
		errorImage := make([]string, 0)
		for _, image := range kustomization.Images {
			var included bool
			for _, inc := range s.Include {
				if strings.Contains(image.Name, inc) {
					included = true
				}
			}
			if !included {
				errorImage = append(errorImage, image.Name)
			}
		}
		for _, e := range errorImage {
			fmt.Printf("%s does not contains %s\n", e, s.Include)
			hasError = true
		}
	}
	return hasError
}
