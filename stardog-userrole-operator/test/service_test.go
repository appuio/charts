package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/gruntwork-io/terratest/modules/helm"
	corev1 "k8s.io/api/core/v1"
)

var (
	tplService = []string{"templates/service.yaml"}
)

func Test_Service_WhenServicePortOverridden_ThenRenderNewPort(t *testing.T) {
	expectedPort := int32(9090)
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.service.port": fmt.Sprintf("%d", expectedPort),
		},
	}

	service := renderService(t, options, false)

	assert.Equal(t, expectedPort, service.Spec.Ports[0].Port, "Service does not use configured port")
}

func Test_Service_WhenProxyServicePortOverridden_ThenRenderNewPort(t *testing.T) {
	expectedPort := int32(9090)
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.proxy.port": fmt.Sprintf("%d", expectedPort),
			"metrics.proxy.enabled": fmt.Sprintf("%s","true"),
		},
	}

	service := renderService(t, options, false)

	assert.Equal(t, expectedPort, service.Spec.Ports[1].Port, "Service does not use configured proxy port")
}

func Test_Service_WhenProxyServiceEnabled_ThenRenderBothPorts(t *testing.T) {
	expectedPorts := 2
	expectedProxyPort := int32(8443)
	expectedPort := int32(80)
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.proxy.enabled": fmt.Sprintf("%s","true"),
		},
	}

	service := renderService(t, options, false)

	assert.Equal(t, expectedPorts, len(service.Spec.Ports), "Service should have all ports")
	assert.Equal(t, expectedPort, service.Spec.Ports[0].Port, "Service should have correct port")
	assert.Equal(t, expectedProxyPort, service.Spec.Ports[1].Port, "Service should have correct port")
}

func Test_Service_GivenDefaultValues_ThenRenderMatchingLabelsWithDeployment(t *testing.T) {
	expectedName := releaseName + "-stardog-userrole-operator-metrics"
	options := &helm.Options{}

	service := renderService(t, options, false)
	deployment := renderDeployment(t, options, false)

	assert.Equal(t, expectedName, service.Name, "Service does not use configured name")
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
