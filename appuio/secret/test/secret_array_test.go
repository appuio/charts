package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

var tplSecretDeprecated = []string{"templates/secret-deprecated.yaml"}

func TestSecret_ArraySyntax(t *testing.T) {
	options := &helm.Options{
		ValuesFiles: []string{"testdata/secret-array.yaml"},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplSecretDeprecated)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	assert.Equal(t, "foo-secret", secret.Name)
	assert.Equal(t, "database", secret.Labels["app"])
	assert.Equal(t, "file", secret.StringData["static"])
}
