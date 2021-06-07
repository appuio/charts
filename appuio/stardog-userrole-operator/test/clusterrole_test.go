package test

import (
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/rbac/v1"
	"strings"
	"testing"
)

var (
	tplClusterRole = []string{"templates/clusterrole.yaml"}
)

func Test_ClusterRole_GivenDefaultValues_ThenRenderTemplate(t *testing.T) {
	expectedManagerName := releaseName + "-stardog-userrole-operator-manager"
	expectedMetricsName := releaseName + "-stardog-userrole-operator-metrics-reader"
	options := &helm.Options{}

	clusterRoles := renderClusterRole(t, options, false)

	assert.Equal(t, 2, len(*clusterRoles), "ClusterRole is missing some resources")
	assert.Equal(t, expectedManagerName, (*clusterRoles)[0].Name, "ClusterRole manager has wrong name")
	assert.Equal(t, expectedMetricsName, (*clusterRoles)[1].Name, "ClusterRole metrics reader has a wrong name")
}

func Test_ClusterRole_WhenProxyValues_ThenRenderAllClusterRoles(t *testing.T) {
	expectedManagerName := releaseName + "-stardog-userrole-operator-manager"
	expectedMetricsName := releaseName + "-stardog-userrole-operator-metrics-reader"
	expectedProxyName := releaseName + "-stardog-userrole-operator-proxy"
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.proxy.enabled": "true",
		},
	}

	clusterRoles := renderClusterRole(t, options, false)

	assert.Equal(t, 3, len(*clusterRoles), "ClusterRole is missing some resources")
	assert.Equal(t, expectedManagerName, (*clusterRoles)[0].Name, "ClusterRole manager has a wrong name")
	assert.Equal(t, expectedProxyName, (*clusterRoles)[1].Name, "ClusterRole proxy has a wrong name")
	assert.Equal(t, expectedMetricsName, (*clusterRoles)[2].Name, "ClusterRole metrics reader has a wrong name")
}

func renderClusterRole(t *testing.T, options *helm.Options, wantErr bool) *[]v1.ClusterRole {
	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplClusterRole)
	if wantErr {
		require.Error(t, err)
		return nil
	}
	require.NoError(t, err)
	outputs := strings.Split(output, "---")
	var result []v1.ClusterRole
	for _, out := range outputs {
		if out != "" {
			clusterRole := v1.ClusterRole{}
			helm.UnmarshalK8SYaml(t, out, &clusterRole)
			result = append(result, clusterRole)
		}
	}

	return &result
}
