package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/helm"
)

var tplSecret = []string{"templates/secret.yaml"}

func Test_Secret_Should_UseValuesSecretKey(t *testing.T) {
	secretKey := "xyHpnHNruMGGC9XfsKwZ7ArPqZytxvP468YDMGHFH3rgCJdq87YLGAY47rAAppXr"

	options := &helm.Options{
		SetValues: map[string]string{
			"snappass.secretKey": secretKey,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplSecret)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	data := secret.Data
	assert.Equal(t, secretKey, string(data["secretKey"]), "Secret does not use the secretKey from the values")
}
