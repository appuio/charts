package test

import (
	"testing"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/helm"
)

var tplDeployment = []string{"templates/deployment.yaml"}

func Test_Deployment(t *testing.T) {
	redisName := releaseName + "-redis"
	redisPort := "6379"
	existingSecret := "secret-key-secret"
	options := &helm.Options{
		SetValues: map[string]string{
			"redis.port": redisPort,
			"snappass.existingSecret": existingSecret,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplDeployment)

	env := resolveEnvironmentVariables(t, output)

	variables := []string{"REDIS_HOST", "REDIS_PORT", "REDIS_PASSWORD", "SECRET_KEY"}

	for _, variable := range variables {
		assertVariableExists(t, env, variable)
	}

	assert.Equal(t, redisName, retrieveEnvVariable(env, "REDIS_PASSWORD").ValueFrom.SecretKeyRef.Name, "Deployment does not use the default redis secret")
	assert.Equal(t, redisName+"-master", retrieveEnvVariable(env, "REDIS_HOST").Value, "Deployment does not use the default redis host")
	assert.Equal(t, redisPort, retrieveEnvVariable(env, "REDIS_PORT").Value, "Deployment does not use the configured redis port")

	assert.Equal(t, existingSecret, retrieveEnvVariable(env, "SECRET_KEY").ValueFrom.SecretKeyRef.Name, "Deployment does not use the existing secret")
}

func assertVariableExists(t *testing.T, env []corev1.EnvVar, name string) {
	assert.NotEmpty(t, retrieveEnvVariable(env, name), "Environment variable "+name+" does not exist")
}

func retrieveEnvVariable(env []corev1.EnvVar, name string) corev1.EnvVar {
	for _, element := range env {
		if element.Name == name {
			return element
		}
	}
	return corev1.EnvVar{}
}

func resolveEnvironmentVariables(t *testing.T, output string) []corev1.EnvVar {
	deployment := appv1.Deployment{}
	helm.UnmarshalK8SYaml(t, output, &deployment)

	return deployment.Spec.Template.Spec.Containers[0].Env
}
