package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	corev1 "k8s.io/api/core/v1"
)

var tplSecret = []string{"templates/secret.yaml"}

func TestSecret_ObjectSyntax(t *testing.T) {
	options := &helm.Options{
		ValuesFiles: []string{"testdata/secret-object.yaml"},
		SetValues: map[string]string{
			"external": "foo",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplSecret)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	assert.Equal(t, "foo-secret", secret.Name)
	assert.Equal(t, "database", secret.Labels["app"])
	assert.Equal(t, "github", secret.Annotations["app.kubernetes.io/source"])
	assert.Equal(t, "file", secret.StringData["static"])
	assert.Equal(t, "App", secret.StringData["dynamic"])
	assert.NotNil(t, secret.Data["base64"])
}
