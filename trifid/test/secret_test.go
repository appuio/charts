package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/helm"
)

var tplSecret = []string{"templates/secret.yaml"}

func Test_Secret_Should_UseValuesEndpointCredentials(t *testing.T) {
	username := "admin"
	password := "xyHpnHNruMGGC9XfsKwZ7ArPqZytxvP468YDMGHFH3rgCJdq87YLGAY47rAAppXr"

	options := &helm.Options{
		SetValues: map[string]string{
			"trifid.sparql.endpoint.username": username,
			"trifid.sparql.endpoint.password": password,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, tplSecret)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	data := secret.Data
	assert.Equal(t, username, string(data["username"]), "Secret does not use the endpoint username from the values")
	assert.Equal(t, password, string(data["password"]), "Secret does not use the endpoint password from the values")
}
