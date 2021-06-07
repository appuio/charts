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
	tplRoleBinding = []string{"templates/rolebinding.yaml"}
)

func Test_RoleBinding_GivenElectionLeader_ThenRenderTemplate(t *testing.T) {
	expectedName := releaseName + "-stardog-userrole-operator-leader-election"
	expectedRoleRef := releaseName + "-stardog-userrole-operator-leader-election"
	expectedSAName := releaseName + "-stardog-userrole-operator"
	options := &helm.Options{
		SetValues: map[string]string{
			"leaderElection.enabled": fmt.Sprintf("%s", "true"),
		},
	}

	roleBinding := renderRoleBinding(t, options, false)

	assert.Equal(t, expectedName, roleBinding.Name, "RoleBinding does not use configured name")
	assert.Equal(t, expectedRoleRef, roleBinding.RoleRef.Name, "RoleRef does not use configured name")
	assert.Equal(t, expectedSAName, roleBinding.Subjects[0].Name, "ServiceAccount does not use configured name")
}

func Test_RoleBinding_GivenDefaultValues_ThenRenderNothing(t *testing.T) {
	renderRoleBinding(t, &helm.Options{}, true)
}

func renderRoleBinding(t *testing.T, options *helm.Options, wantErr bool) *v1.RoleBinding {
	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplRoleBinding)
	if wantErr {
		require.Error(t, err)
		return nil
	}
	require.NoError(t, err)
	roleBinding := v1.RoleBinding{}
	helm.UnmarshalK8SYaml(t, output, &roleBinding)
	return &roleBinding
}
