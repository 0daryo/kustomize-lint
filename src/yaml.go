package src

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

const (
	KUSTOMIZATIONFILE = "kustomization.yaml"
	PRODUCTION        = "production"
	STAGING           = "staging"
	LATEST            = "latest"
)

func ReadOnConfig(fileBuffer []byte) (Yaml, error) {
	data := Yaml{}
	err := yaml.Unmarshal(fileBuffer, &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func ReadOnKustomize(fileBuffer []byte) (Kustomization, error) {
	data := Kustomization{}
	err := yaml.Unmarshal(fileBuffer, &data)
	if err != nil {
		fmt.Println(err)
		return data, err
	}
	return data, nil
}

func parseKustomize(filePath string) *Kustomization {
	buf, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	data, err := ReadOnKustomize(buf)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return &data
}
func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func Lint() error {
	var buf []byte
	var err error
	if fileExists("kustomize-lint.yaml") {
		buf, err = ioutil.ReadFile("./kustomize-lint.yaml")
	} else {
		buf, err = ioutil.ReadFile("./kustomize-lint-ex.yaml")
	}
	if err != nil {
		fmt.Println(err)
		return err
	}

	conf, err := ReadOnConfig(buf)
	if err != nil {
		fmt.Println(err)
		return err
	}
	var hasError bool
	for _, file := range conf.Files {
		relatedFiles, _ := filepath.Glob(fmt.Sprintf("%s", file.Name))
		for _, rf := range relatedFiles {
			kustomization := parseKustomize(rf)
			for _, sentence := range file.Sentences {
				if sentence.hasError(kustomization) {
					hasError = true
				}
			}
		}
	}
	if hasError {
		return errors.New("validation error")
	}
	return nil
}
