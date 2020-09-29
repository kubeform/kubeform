module kubeform.dev/kubeform

go 1.14

require (
	github.com/appscode/go v0.0.0-20200928211031-cc0c23082d91
	github.com/dave/jennifer v1.4.0
	github.com/go-openapi/spec v0.19.8
	github.com/gobuffalo/flect v0.2.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/hashicorp/terraform-config-inspect v0.0.0-20200526195750-d43f12b82861
	github.com/hashicorp/terraform-plugin-sdk v1.14.0
	github.com/terraform-providers/terraform-provider-aws v2.32.0+incompatible
	github.com/terraform-providers/terraform-provider-azurerm v1.44.0
	github.com/terraform-providers/terraform-provider-digitalocean v1.20.0
	github.com/terraform-providers/terraform-provider-google v2.17.0+incompatible
	github.com/terraform-providers/terraform-provider-linode v1.12.3
	k8s.io/api v0.18.9
	k8s.io/apiextensions-apiserver v0.18.9 // indirect
	k8s.io/apimachinery v0.18.9
	k8s.io/apiserver v0.18.9 // indirect
	k8s.io/cli-runtime v0.18.9 // indirect
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/cloud-provider v0.18.9 // indirect
	k8s.io/component-base v0.18.9 // indirect
	k8s.io/kube-aggregator v0.18.9 // indirect
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	k8s.io/kubectl v0.18.9 // indirect
	k8s.io/kubernetes v1.18.9 // indirect
	kmodules.xyz/client-go v0.0.0-20200929030759-cce6a3c623c1
	kmodules.xyz/crd-schema-fuzz v0.0.0-20200922204806-c1426cd7fcf4 // indirect
	kmodules.xyz/custom-resources v0.0.0-20200922210108-70f2815a43bb // indirect
	kmodules.xyz/monitoring-agent-api v0.0.0-20200922204510-b8b5610e5385 // indirect
	kmodules.xyz/objectstore-api v0.0.0-20200922210707-59bab27e5d41 // indirect
	kmodules.xyz/offshoot-api v0.0.0-20200922211229-36acc531abab // indirect
	kmodules.xyz/openshift v0.0.0-20200922211657-1ece16d36c18 // indirect
	kmodules.xyz/prober v0.0.0-20200922212142-743a6514664e // indirect
	kmodules.xyz/webhook-runtime v0.0.0-20200922211931-8337935590de // indirect
)

replace github.com/linode/linodego => github.com/linode/linodego v0.19.0

replace github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.25.4

replace github.com/keybase/go-crypto v0.0.0-20190523171820-b785b22cc757 => github.com/keybase/go-crypto v0.0.0-20190416182011-b785b22cc757

replace github.com/terraform-providers/terraform-provider-google v2.17.0+incompatible => github.com/terraform-providers/terraform-provider-google v1.20.1-0.20191008212436-363f2d283518

replace github.com/terraform-providers/terraform-provider-aws v2.32.0+incompatible => github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20191010190908-1261a98537f2

replace github.com/terraform-providers/terraform-provider-random v2.2.1+incompatible => github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f

replace bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d

replace cloud.google.com/go => cloud.google.com/go v0.49.0

replace git.apache.org/thrift.git => github.com/apache/thrift v0.13.0

replace github.com/Azure/azure-sdk-for-go => github.com/Azure/azure-sdk-for-go v35.0.0+incompatible

replace github.com/Azure/go-ansiterm => github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.0+incompatible

replace github.com/Azure/go-autorest/autorest => github.com/Azure/go-autorest/autorest v0.9.0

replace github.com/Azure/go-autorest/autorest/adal => github.com/Azure/go-autorest/autorest/adal v0.5.0

replace github.com/Azure/go-autorest/autorest/azure/auth => github.com/Azure/go-autorest/autorest/azure/auth v0.2.0

replace github.com/Azure/go-autorest/autorest/date => github.com/Azure/go-autorest/autorest/date v0.1.0

replace github.com/Azure/go-autorest/autorest/mocks => github.com/Azure/go-autorest/autorest/mocks v0.2.0

replace github.com/Azure/go-autorest/autorest/to => github.com/Azure/go-autorest/autorest/to v0.2.0

replace github.com/Azure/go-autorest/autorest/validation => github.com/Azure/go-autorest/autorest/validation v0.1.0

replace github.com/Azure/go-autorest/logger => github.com/Azure/go-autorest/logger v0.1.0

replace github.com/Azure/go-autorest/tracing => github.com/Azure/go-autorest/tracing v0.5.0

replace github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.1

replace github.com/golang/protobuf => github.com/golang/protobuf v1.3.2

replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.3.1

replace github.com/imdario/mergo => github.com/imdario/mergo v0.3.5

replace github.com/prometheus-operator/prometheus-operator => github.com/prometheus-operator/prometheus-operator v0.42.0

replace github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring => github.com/prometheus-operator/prometheus-operator/pkg/apis/monitoring v0.42.0

replace github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.7.1

replace go.etcd.io/etcd => go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738

replace google.golang.org/api => google.golang.org/api v0.14.0

replace google.golang.org/genproto => google.golang.org/genproto v0.0.0-20191115194625-c23dd37a84c9

replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

replace k8s.io/api => github.com/kmodules/api v0.18.10-0.20200922195318-d60fe725dea0

replace k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200922195535-0c9a1b86beec

replace k8s.io/apiserver => github.com/kmodules/apiserver v0.18.10-0.20200922195747-1bd1cc8f00d1

replace k8s.io/cli-runtime => k8s.io/cli-runtime v0.18.9

replace k8s.io/client-go => github.com/kmodules/k8s-client-go v0.18.10-0.20200922201634-73fedf3d677e

replace k8s.io/component-base => k8s.io/component-base v0.18.9

replace k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6

replace k8s.io/kubernetes => github.com/kmodules/kubernetes v1.19.0-alpha.0.0.20200922200158-8b13196d8dc4

replace k8s.io/utils => k8s.io/utils v0.0.0-20200324210504-a9aa75ae1b89
