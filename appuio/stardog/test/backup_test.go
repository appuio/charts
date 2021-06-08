package test

import (
	"testing"

	batchv2alpha1 "k8s.io/api/batch/v2alpha1"

	"github.com/gruntwork-io/terratest/modules/helm"
	"github.com/stretchr/testify/assert"
)

var (
	tplCronJob     = []string{"templates/cronjob.yaml"}
	dbName         = "my-backup-test"
	bucketName     = "test-bucket"
	schedule       = "@daily"
	existingSecret = "aws-credentials"
	backupOptions  = map[string]string{
		"stardog.adminPassword":                      adminPassword,
		"stardog.backup.databases[0].name":           dbName,
		"stardog.backup.databases[0].schedule":       schedule,
		"stardog.backup.s3BucketName":                bucketName,
		"stardog.backup.s3CredentialsExistingSecret": existingSecret,
	}
)

func Test_Backup_CronJob_Should_Exist(t *testing.T) {
	options := &helm.Options{
		SetValues: backupOptions,
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplCronJob)

	cronJob := batchv2alpha1.CronJob{}
	helm.UnmarshalK8SYaml(t, output, &cronJob)

	assert.Equal(t, schedule, cronJob.Spec.Schedule)
	assert.Equal(t, bucketName, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env[1].Value)
	assert.Equal(t, existingSecret, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env[6].ValueFrom.SecretKeyRef.Name)
	assert.Contains(t, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Command[2], dbName)
}

func Test_Backup_CronJob_S3_Endpoint(t *testing.T) {
	s3Endpoint := "sos-ch-dk-2.exo.io"
	backupOptions["stardog.backup.s3Endpoint"] = s3Endpoint
	options := &helm.Options{
		SetValues: backupOptions,
	}

	output := helm.RenderTemplate(t, options, helmChartPath, releaseName, tplCronJob)

	cronJob := batchv2alpha1.CronJob{}
	helm.UnmarshalK8SYaml(t, output, &cronJob)

	assert.Equal(t, s3Endpoint, cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers[0].Env[3].Value)
}
