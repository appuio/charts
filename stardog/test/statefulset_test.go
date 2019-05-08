package test

import (
	"testing"

	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"

	"github.com/stretchr/testify/assert"

	"github.com/gruntwork-io/terratest/modules/helm"
)

var tplStatefulset = []string{"templates/statefulset.yaml"}
var tplConfigMap = []string{"templates/configmap.yaml"}

func Test_StatefulSet_Should_Include_StorageClass_IfGiven(t *testing.T) {
	storageClassName := "example"
	options := &helm.Options{
		SetValues: map[string]string{
			"persistence.storageClass": storageClassName,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, tplStatefulset)

	statefulset := appv1.StatefulSet{}
	helm.UnmarshalK8SYaml(t, output, &statefulset)

	volumeClaimTemplates := statefulset.Spec.VolumeClaimTemplates
	assert.Equal(t, *volumeClaimTemplates[0].Spec.StorageClassName, storageClassName, "StatefulSet does not use configured StorageClass")
}

func Test_StatefulSet_Should_UseExistingLicenseSecret(t *testing.T) {
	licenseName := "my-custom-license"
	options := &helm.Options{
		SetValues: map[string]string{
			"stardog.existingLicenseSecret": licenseName,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, tplStatefulset)

	statefulset := appv1.StatefulSet{}
	helm.UnmarshalK8SYaml(t, output, &statefulset)

	licenseVol := statefulset.Spec.Template.Spec.Volumes[0]
	assert.Equal(t, licenseVol.Secret.SecretName, licenseName, "StatefulSet does not use existing license secret")
}

func Test_StatefulSet_Should_UseExistingPullSecret_IfGiven(t *testing.T) {
	secretName := "my-custom-registry"
	options := &helm.Options{
		SetValues: map[string]string{
			"image.existingPullSecret": secretName,
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, tplStatefulset)

	statefulset := appv1.StatefulSet{}
	helm.UnmarshalK8SYaml(t, output, &statefulset)

	imageSecret := statefulset.Spec.Template.Spec.ImagePullSecrets[0]
	assert.Equal(t, imageSecret.Name, secretName, "StatefulSet does not use existing image pull secret")
}

func Test_StatefulSet_Should_Contain_ZooKeeper_ConnectionString_InClusterMode(t *testing.T) {
	options := &helm.Options{
		SetValues: map[string]string{
			"zookeeper.enabled": "true",
		},
	}

	output := helm.RenderTemplate(t, options, helmChartPath, tplConfigMap)

	configMap := &corev1.ConfigMap{}
	helm.UnmarshalK8SYaml(t, output, &configMap)
	assert.Contains(t, configMap.Data["stardog.properties"], "pack.enabled=true", "Stardog config doesn't enable ZooKeeper")
	assert.Contains(t, configMap.Data["stardog.properties"], "pack.zookeeper.address", "Stardog config doesn't contain the ZooKeeper connection string")
}
