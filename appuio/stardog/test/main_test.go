package test

var (
	helmChartPath                 = ".."
	adminPassword                 = "something"
	releaseName                   = "test-release"
	minimalOptionsForAllTemplates = map[string]string{
		"image.password":                   "imagePullSecretPassword",
		"ingress.enabled":                  "true",
		"metrics.enabled":                  "true",
		"metrics.prometheusOperator":       "true",
		"stardog.adminPassword":            adminPassword,
		"stardog.backup.databases[0].name": "test",
		"stardog.backup.s3AccessKey":       "s3AccessKey",
		"stardog.backup.s3BucketName":      "s3BucketName",
		"stardog.backup.s3SecretKey":       "s3SecretKey",
	}
)

func Merge[K comparable, V any](dst map[K]V, src map[K]V) map[K]V {
	for k, v := range src {
		dst[k] = v
	}
	return dst
}
