package test

import (
	"strings"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/gruntwork-io/terratest/modules/k8s"
	"github.com/gruntwork-io/terratest/modules/random"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	corev1 "k8s.io/api/core/v1"
)

func Test_Generic_Should_Render_Pod(t *testing.T) {
	namespaceName := "medieval-" + strings.ToLower(random.UniqueId())
	options := &helm.Options{
		KubectlOptions: k8s.NewKubectlOptions("", "", namespaceName),
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, []string{})

	secret := &corev1.Secret{}
	err := helm.UnmarshalK8SYamlE(t, output, &secret)
	require.NoError(t, err)
	assert.Equal(t, "example-secret", secret.Name)
	assert.Equal(t, "secret", secret.StringData["some"])
	assert.Empty(t, secret.Namespace)
}
