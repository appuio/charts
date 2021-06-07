package test

import (
	"testing"

	batchv1 "k8s.io/api/batch/v1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
)

var tplPostInstallJob = []string{"templates/post-install-job.yaml"}

func Test_PostInstallJob_GivenReplicaCount_WhenZero_ThenDoNotRenderJob(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"stardog.adminPassword":     adminPassword,
			"replicaCount":	"0",
		},
	}

	output, err := helm.RenderTemplateE(t, options, helmChartPath, releaseName, tplPostInstallJob)

	assert.Error(t, err)
	assert.Equal(t, "", output)
}

func Test_PostInstallJob_GivenReplicaCount_WhenGreaterThanZero_ThenRenderJob(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"stardog.adminPassword":     adminPassword,
			"replicaCount":	"1",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplPostInstallJob)

	job := batchv1.Job{}
	helm.UnmarshalK8SYaml(t, output, &job)

	assert.NotEmpty(t, job.Spec)
}
