package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-yaml/yaml"
)

const (
	KUSTOMIZATIONFILE = "kustomization.yaml"
	PRODUCTION        = "production"
	STAGING           = "staging"
	LATEST            = "latest"
)

type Kustomization struct {
	Images []Image
	Bases  []string
}

type Image struct {
	Name   string `yaml:"name"`
	NewTag string `yaml:"newTag"`
}

func ReadOnStruct(fileBuffer []byte) (Kustomization, error) {
	data := Kustomization{}
	err := yaml.Unmarshal(fileBuffer, &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func parse(filePath string) *Kustomization {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// []byte を []Test に変換します。
	data, err := ReadOnStruct(buf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &data
}

func Lint() {
	productionPaths := make([]string, 0)
	productionFiles, _ := filepath.Glob("*/production/kustomization.yaml")
	for _, f := range productionFiles {
		if strings.HasSuffix(f, KUSTOMIZATIONFILE) {
			productionPaths = append(productionPaths, f)
		}
	}
	stagingPaths := make([]string, 0)
	stagingFiles, _ := filepath.Glob("*/staging/kustomization.yaml")
	for _, f := range stagingFiles {
		if strings.HasSuffix(f, KUSTOMIZATIONFILE) {
			stagingPaths = append(stagingPaths, f)
		}
	}
	var hasError bool
	if !validateProduction(productionPaths) {
		fmt.Println("invalid yaml in production")
		hasError = true
	}
	if !validateStaging(stagingPaths) {
		fmt.Println("invalid yaml in staging")
		hasError = true
	}
	if hasError {
		os.Exit(1)
	}
}

func validateProduction(paths []string) bool {
	for _, path := range paths {
		data := parse(path)
		for _, tag := range data.Images {
			if !strings.Contains(tag.NewTag, PRODUCTION) && !strings.Contains(tag.NewTag, LATEST) {
				return false
			}
		}
	}
	return true
}

func validateStaging(paths []string) bool {
	for _, path := range paths {
		data := parse(path)
		for _, tag := range data.Images {
			if !strings.Contains(tag.NewTag, STAGING) && !strings.Contains(tag.NewTag, LATEST) {
				return false
			}
		}
	}
	return true
}
