package test

import (
	"testing"

	corev1 "k8s.io/api/core/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
)

var tplUserPwSecret = []string{"templates/user-pw-secret.yaml"}

func Test_UserPwSecret_ShouldThrowError_IfNotGiven(t *testing.T) {

	options := &helm.Options{}

	_, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplUserPwSecret)
	assert.Error(t, err)
}

func Test_UserPwSecret_ShouldListGivenUsers(t *testing.T) {
	user1 := "user1"
	user2 := "user2"
	pw1 := "password1"
	pw2 := "password2"

	options := &helm.Options{
		SetValues: map[string]string{
			"stardog.users[0].name":     user1,
			"stardog.users[0].password": pw1,
			"stardog.users[1].name":     user2,
			"stardog.users[1].password": pw2,
			"stardog.adminPassword":     adminPassword,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplUserPwSecret)

	secret := corev1.Secret{}
	helm.UnmarshalK8SYaml(t, output, &secret)

	assert.Equal(t, adminPassword, string(secret.Data["admin"]))
	assert.Equal(t, pw1, string(secret.Data[user1]))
	assert.Equal(t, pw2, string(secret.Data[user2]))
}
