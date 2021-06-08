package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/helm"
)

var tplSecretRedis = []string{"templates/secret-redis.yaml"}

func Test_Secret_Should_UseValuesRedisPassword(t *testing.T) {
	redisPassword := "xyHpnHNruMGGC9XfsKwZ7ArPqZytxvP468YDMGHFH3rgCJdq87YLGAY47rAAppXr"

	options := &helm.Options{
		SetValues: map[string]string{
			"redis.enabled":           "false",
			"snappass.redis.password": redisPassword,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplSecretRedis)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	data := secret.Data
	assert.Equal(t, redisPassword, string(data["redis-password"]), "Secret does not use the redis password from the values")
}
