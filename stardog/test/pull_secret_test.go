package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
)

var tplPullSecret = []string{"templates/pull-secret.yaml"}

func Test_PullSecret(t *testing.T) {
	registryUserName := "superhero"
	registryPassword := "supersecret"

	options := &helm.Options{
		SetValues: map[string]string{
			"image.username": registryUserName,
			"image.password": registryPassword,
			"stardog.adminPassword": adminPassword,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, tplPullSecret)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	assert.NotEmpty(t, secret.Data[".dockerconfigjson"])
}
