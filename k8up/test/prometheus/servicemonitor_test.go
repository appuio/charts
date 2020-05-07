package test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	monitoringv1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/gruntwork-io/terratest/modules/helm"
)

var (
	tplServiceMonitor = []string{"templates/prometheus/servicemonitor.yaml"}
)

func Test_ServiceMonitor_GivenEnabled_WhenIntervalDefined_ThenRenderNewInterval(t *testing.T) {
	expectedInterval := "1m10s"
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.serviceMonitor.enabled": "true",
			"metrics.serviceMonitor.scrapeInterval": expectedInterval,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplServiceMonitor)
	monitor := monitoringv1.ServiceMonitor{}
	helm.UnmarshalK8SYaml(t, output, &monitor)

	assert.Equal(t, expectedInterval, monitor.Spec.Endpoints[0].Interval)
}
