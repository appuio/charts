package test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	appv1 "k8s.io/api/apps/v1"
)

var (
	tplDeployment = []string{"templates/deployment.yaml"}
)

func Test_Deployment_ShouldRender_DefaultServiceAccount(t *testing.T) {
	want := releaseName + "-stardog-userrole-operator"
	options := &helm.Options{}

	got := renderDeployment(t, options, false)
	serviceName := got.Spec.Template.Spec.ServiceAccountName
	assert.Equal(t, want, serviceName, "Deployment does not render configured serviceName")
}

func Test_Deployment_ShouldRender_CustomServiceAccount(t *testing.T) {
	want := "test"
	options := &helm.Options{
		SetValues: map[string]string{
			"serviceAccount.name": want,
		},
	}

	got := renderDeployment(t, options, false)

	serviceName := got.Spec.Template.Spec.ServiceAccountName
	assert.Equal(t, want, serviceName, "Deployment does not render configured serviceName")
}

func Test_Deployment_ShouldRender_Resources(t *testing.T) {
	want := "1Gi"
	options := &helm.Options{
		SetValues: map[string]string{
			"resources.limits.memory": want,
		},
	}

	got := renderDeployment(t, options, false)
	resources := got.Spec.Template.Spec.Containers[0].Resources
	assert.Equal(t, want, resources.Limits.Memory().String(), "Deployment does not render configured memory limit")
}

func Test_Deployment_ShouldRender_Labels(t *testing.T) {
	options := &helm.Options{}

	got := renderDeployment(t, options, false)

	selector := got.Spec.Selector.MatchLabels
	matchLabels := got.Spec.Template.Labels
	assert.Equal(t, selector, matchLabels, "Deployment does not render matching labels")
}

func Test_Deployment_ShouldRender_ImagePullSecrets(t *testing.T) {
	expectedImageSecret := "secret"
	options := &helm.Options{
		SetValues: map[string]string{
			"imagePullSecrets[0].name": expectedImageSecret,
		},
	}

	got := renderDeployment(t, options, false)

	assert.Equal(t, expectedImageSecret, got.Spec.Template.Spec.ImagePullSecrets[0].Name, "Deployment does not render image pull secrets")
}

func Test_Deployment_ShouldRender_CorrectImage(t *testing.T) {
	expectedImage := "ghcr.io/vshn/stardog-userrole-operator:v0.2.1"
	options := &helm.Options{}

	got := renderDeployment(t, options, false)

	assert.Equal(t, expectedImage, got.Spec.Template.Spec.Containers[0].Image, "Deployment does not render image location")
}

func Test_Deployment_WhenProxyEnabled_ThenEnableRBACContainer(t *testing.T) {
	expectedBaseImage := "ghcr.io/vshn/stardog-userrole-operator:v0.2.0"
	expectedProxyImage := "gcr.io/kubebuilder/kube-rbac-proxy:v0.12.0"
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.proxy.enabled": "true",
		},
	}

	got := renderDeployment(t, options, false)

	assert.Equal(t, 2, len(got.Spec.Template.Spec.Containers), "Deployment does not have the correct amount of containers")
	assert.Equal(t, expectedBaseImage, got.Spec.Template.Spec.Containers[0].Image, "Deployment has a wrong image configured")
	assert.Equal(t, expectedProxyImage, got.Spec.Template.Spec.Containers[1].Image, "Deployment does not render proxy image location")
}

func renderDeployment(t *testing.T, options *helm.Options, wantErr bool) *appv1.Deployment {
	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplDeployment)
	if wantErr {
		require.Error(t, err)
		return nil
	}
	require.NoError(t, err)
	deployment := appv1.Deployment{}
	helm.UnmarshalK8SYaml(t, output, &deployment)
	return &deployment
}
