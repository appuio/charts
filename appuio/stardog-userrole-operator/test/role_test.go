package test

import (
	"fmt"
	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	v1 "k8s.io/api/rbac/v1"
	"testing"
)

var (
	tplRole = []string{"templates/role.yaml"}
)

func Test_Role_GivenLeaderElectionTrue_ThenRenderTemplate(t *testing.T) {
	expectedName := releaseName + "-stardog-userrole-operator-leader-election"
	options := &helm.Options{
		SetValues: map[string]string{
			"leaderElection.enabled": fmt.Sprintf("%s", "true"),
		},
	}

	role := renderRole(t, options, false)

	assert.Equal(t, expectedName, role.Name, "Role does not use configured name")
}

func Test_Role_GivenDefaultValues_ThenRenderNothing(t *testing.T) {
	renderRole(t, &helm.Options{}, true)
}

func renderRole(t *testing.T, options *helm.Options, wantErr bool) *v1.Role {
	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplRole)
	if wantErr {
		require.Error(t, err)
		return nil
	}
	require.NoError(t, err)
	role := v1.Role{}
	helm.UnmarshalK8SYaml(t, output, &role)
	return &role
}
