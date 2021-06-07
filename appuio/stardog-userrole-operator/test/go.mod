module test

go 1.15

require (
	github.com/gruntwork-io/terratest v0.32.18
	github.com/stretchr/testify v1.7.0
	k8s.io/api v0.19.3
)

replace k8s.io/client-go => k8s.io/client-go v0.17.4
