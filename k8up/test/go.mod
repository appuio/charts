module test

go 1.15

require (
	github.com/gruntwork-io/terratest v0.32.3
	github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.45.0
	github.com/stretchr/testify v1.7.0
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
)

replace k8s.io/client-go => k8s.io/client-go v0.20.2
