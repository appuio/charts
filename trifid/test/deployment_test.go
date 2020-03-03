package test

import (
	"testing"

	appv1 "k8s.io/api/apps/v1"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/helm"
)

var tplDeployment = []string{"templates/deployment.yaml"}

func Test_Deployment_Should_UseValuesEndpointUrl(t *testing.T) {
	url := "https://appuio.ch"
	options := &helm.Options{
		SetValues: map[string]string{
			"trifid.sparql.endpoint.url": url,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplDeployment)

	deployment := appv1.Deployment{}
	helm.UnmarshalK8SYaml(t, output, &deployment)

	env := deployment.Spec.Template.Spec.Containers[0].Env[0]
	assert.Equal(t, env.Value, url, "Deployment does not use the endpoint url from the values")
}

func Test_Deployment_Should_UseChartSecret(t *testing.T) {
	fullname := "trifid-chart-test"
	options := &helm.Options{
		SetValues: map[string]string{
			"fullnameOverride": fullname,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplDeployment)

	deployment := appv1.Deployment{}
	helm.UnmarshalK8SYaml(t, output, &deployment)

	env := deployment.Spec.Template.Spec.Containers[0].Env
	assert.Equal(t, fullname, env[1].ValueFrom.SecretKeyRef.Name, "Deployment does not use the chart secret")
}

func Test_Deployment_Should_UseExistingEndpointSecret(t *testing.T) {
	secretName := "endpoint-secret"
	options := &helm.Options{
		SetValues: map[string]string{
			"trifid.sparql.endpoint.secretName": secretName,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplDeployment)

	deployment := appv1.Deployment{}
	helm.UnmarshalK8SYaml(t, output, &deployment)

	env := deployment.Spec.Template.Spec.Containers[0].Env
	assert.Equal(t, secretName, env[1].ValueFrom.SecretKeyRef.Name, "Deployment does not use the existing endpoint secret")
}
