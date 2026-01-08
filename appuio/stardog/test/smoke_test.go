package test

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
)

func Test_Chart_can_be_rendered(t *testing.T) {
	options := &helm.Options{
		SetValues: minimalOptionsForAllTemplates,
	}

	err := renderChart(t, options)

	assert.Nil(t, err)
}

func Test_Chart_can_be_rendered_with_keep_firing_for(t *testing.T) {
	options := &helm.Options{
		SetValues: Merge(minimalOptionsForAllTemplates, map[string]string{
			"alerts.stardogOpenConnections.keep_firing_for": "1h",
		}),
	}

	err := renderChart(t, options)

	assert.Nil(t, err)
}

func renderChart(t *testing.T, options *helm.Options) error {
	helmChartPathAbsPath, err := filepath.Abs(helmChartPath)
	if err != nil {
		t.Fatal(err)
	}
	files, err := findYamlFiles(helmChartPathAbsPath+"/templates", helmChartPathAbsPath+"/")
	if err != nil {
		t.Fatal(err)
	}

	_, err = helm.RenderTemplateE(t, options, helmChartPath, releaseName, files)
	return err
}

func findYamlFiles(root string, relativeTo string) ([]string, error) {
	var files []string

	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() {
			ext := strings.ToLower(filepath.Ext(path))
			if ext == ".yaml" || ext == ".yml" {
				files = append(files, strings.Replace(path, relativeTo, "", 1))
			}
		}
		return nil
	})

	return files, err
}
