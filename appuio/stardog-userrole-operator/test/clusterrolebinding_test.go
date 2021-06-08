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
	tplClusterRoleBinding = []string{"templates/clusterrolebinding.yaml"}
)

func Test_ClusterRoleBinding_GivenDefaultValues_ThenRenderTemplate(t *testing.T) {
	expectedManagerName := releaseName + "-stardog-userrole-operator-manager"
	roleRefManagerName := releaseName + "-stardog-userrole-operator-manager"
	serviceAccount := releaseName + "-stardog-userrole-operator"
	options := &helm.Options{}

	clusterRoleBinding := renderClusterRoleBinding(t, options, false)

	assert.Equal(t, 1, len(*clusterRoleBinding), "ClusterRoleBinding does not have all resources")
	assert.Equal(t, expectedManagerName, (*clusterRoleBinding)[0].Name, "ClusterRoleBinding manager does not have correct ame")
	assert.Equal(t, roleRefManagerName, (*clusterRoleBinding)[0].RoleRef.Name, "ClusterRoleBinding proxy does not have correct name")
	assert.Equal(t, serviceAccount, (*clusterRoleBinding)[0].Subjects[0].Name, "ClusterRoleBinding proxy does not have correct name")
}

func Test_ClusterRoleBinding_GivenProxyCRB_ThenRenderBothTemplates(t *testing.T) {
	expectedManagerName := releaseName + "-stardog-userrole-operator-manager"
	roleRefManagerName := releaseName + "-stardog-userrole-operator-manager"
	expectedProxyName := releaseName + "-stardog-userrole-operator-proxy"
	roleRefProxyName := releaseName + "-stardog-userrole-operator-proxy"
	serviceAccount := releaseName + "-stardog-userrole-operator"
	options := &helm.Options{
		SetValues: map[string]string{
			"metrics.proxy.enabled": "true",
		},
	}

	clusterRoleBinding := renderClusterRoleBinding(t, options, false)

	assert.Equal(t, 2, len(*clusterRoleBinding), "ClusterRoleBinding does not have all resources")
	assert.Equal(t, expectedManagerName, (*clusterRoleBinding)[0].Name, "ClusterRoleBinding manager does not have correct ame")
	assert.Equal(t, roleRefManagerName, (*clusterRoleBinding)[0].RoleRef.Name, "ClusterRoleBinding role ref name does not have correct ame")
	assert.Equal(t, serviceAccount, (*clusterRoleBinding)[0].Subjects[0].Name, "ClusterRoleBinding service account name does not have correct ame")
	assert.Equal(t, expectedProxyName, (*clusterRoleBinding)[1].Name, "ClusterRoleBinding proxy does not have correct name")
	assert.Equal(t, roleRefProxyName, (*clusterRoleBinding)[1].RoleRef.Name, "ClusterRoleBinding proxy does not have correct name")
	assert.Equal(t, serviceAccount, (*clusterRoleBinding)[1].Subjects[0].Name, "ClusterRoleBinding service account name does not have correct name")
}

func renderClusterRoleBinding(t *testing.T, options *helm.Options, wantErr bool) *[]v1.ClusterRoleBinding {
	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplClusterRoleBinding)
	if wantErr {
		require.Error(t, err)
		return nil
	}
	require.NoError(t, err)
	outputs := strings.Split(output, "---")
	var result []v1.ClusterRoleBinding
	for _, out := range outputs {
		if out != "" {
			clusterRoleBinding := v1.ClusterRoleBinding{}
			helm.UnmarshalK8SYaml(t, out, &clusterRoleBinding)
			result = append(result, clusterRoleBinding)
		}
	}

	return &result
}
