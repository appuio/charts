package test

import (
	"io/fs"
	"path/filepath"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
)

func Test_Chart_can_be_rendered(t *testing.T) {
	options := &helm.Options{
		SetValues: minimalOptionsForAllTemplates,
	}

	_, err := renderChart(t, options)

	assert.Nil(t, err)
}

func Test_Chart_can_be_rendered_with_keep_firing_for(t *testing.T) {
	options := &helm.Options{
		SetValues: Merge(minimalOptionsForAllTemplates, map[string]string{
			"alerts.stardogOpenConnections.keep_firing_for": "1h",
		}),
	}

	_, err := renderChart(t, options)

	assert.Nil(t, err)
}

func Test_Chart_can_be_rendered_with_additional_labels_for_alerts(t *testing.T) {
	test1Value := uuid.NewString()
	test2Value := uuid.NewString()
	options := &helm.Options{
		SetValues: Merge(minimalOptionsForAllTemplates, map[string]string{
			"alerts.javaLowHeapMemory.additionalLabels.test1":      test1Value,
			"alerts.javaLowHeapMemory.additionalLabels.test2":      test2Value,
			"alerts.stardogOpenConnections.additionalLabels.test1": test1Value,
			"alerts.stardogOpenConnections.additionalLabels.test2": test2Value,
			"alerts.stardogLicenseExpire.additionalLabels.test1":   test1Value,
			"alerts.stardogLicenseExpire.additionalLabels.test2":   test2Value,
			"alerts.stardogPodsNotReady.additionalLabels.test1":    test1Value,
			"alerts.stardogPodsNotReady.additionalLabels.test2":    test2Value,
			"alerts.httpCheck.additionalLabels.test1":              test1Value,
			"alerts.httpCheck.additionalLabels.test2":              test2Value,
			"alerts.certExpirySoon.additionalLabels.test1":         test1Value,
			"alerts.certExpirySoon.additionalLabels.test2":         test2Value,
			"alerts.zooKeeperPodsNotReady.additionalLabels.test1":  test1Value,
			"alerts.zooKeeperPodsNotReady.additionalLabels.test2":  test2Value,
		}),
	}

	result, err := renderChart(t, options)

	assert.Nil(t, err)

	assert.Equalf(t, 7, strings.Count(result, "\""+test1Value+"\""), "Expected to to have test1Value a correct amount of times")
	assert.Equalf(t, 7, strings.Count(result, "\""+test2Value+"\""), "Expected to to have test2Value a correct amount of times")
}

func renderChart(t *testing.T, options *helm.Options) (string, error) {
	helmChartPathAbsPath, err := filepath.Abs(helmChartPath)
	if err != nil {
		t.Fatal(err)
	}
	files, err := findYamlFiles(helmChartPathAbsPath+"/templates", helmChartPathAbsPath+"/")
	if err != nil {
		t.Fatal(err)
	}

	result, renderErr := helm.RenderTemplateE(t, options, helmChartPath, releaseName, files)
	return result, renderErr
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
