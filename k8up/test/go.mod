module test

go 1.14

require (
	github.com/coreos/prometheus-operator v0.45.0
	github.com/gruntwork-io/terratest v0.27.2
	github.com/stretchr/testify v1.7.0
	k8s.io/api v0.17.4
)

replace k8s.io/client-go => k8s.io/client-go v0.17.4
