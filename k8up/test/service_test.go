package test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	corev1 "k8s.io/api/core/v1"
)

var (
	tplService = []string{"templates/service.yaml"}
)

func Test_Service_ShouldNotRender_Service_IfMetricsEnabled_WithoutService(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.enabled":         "false",
			"metrics.service.enabled": "true",
		},
	}

	renderService(t, options, true)
}

func Test_Service_ShouldNotRender_Service_IfMetricsEnabled(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.enabled":         "true",
			"metrics.service.enabled": "false",
		},
	}

	renderService(t, options, true)
}

func Test_Service_ShouldRender_Service_IfMetricsAndServiceEnabled(t *testing.T) {
	want := releaseName + "-k8up-metrics"
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.enabled":         "true",
			"metrics.service.enabled": "true",
		},
	}

	got := renderService(t, options, false)

	assert.Equal(t, want, got.Name, "Service does use configured name")
}

func Test_Service_ShouldRender_MatchingLabels_WithDeployment(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.enabled":         "true",
			"metrics.service.enabled": "true",
		},
	}

	service := renderService(t, options, false)
	deployment := renderDeployment(t, options, false)

	assert.Equal(t, service.Spec.Selector, deployment.Spec.Template.Labels, "Service labels do not match with deployment labels")
}

func renderService(t *testing.T, options *helm.Options, wantErr bool) *corev1.Service {
	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplService)
	if wantErr {
		require.Error(t, err)
		return nil
	}
	require.NoError(t, err)
	service := corev1.Service{}
	helm.UnmarshalK8SYaml(t, output, &service)
	return &service
}
