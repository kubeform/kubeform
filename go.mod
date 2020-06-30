module kubeform.dev/kubeform

go 1.14

require (
	github.com/appscode/go v0.0.0-20200323182826-54e98e09185a
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
	k8s.io/api v0.18.3
	k8s.io/apimachinery v0.18.3
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200410145947-61e04a5be9a6
	kmodules.xyz/client-go v0.0.0-20200630053911-20d035822d35
)

replace github.com/linode/linodego => github.com/linode/linodego v0.19.0

replace github.com/aws/aws-sdk-go => github.com/aws/aws-sdk-go v1.25.4

replace github.com/googleapis/gnostic => github.com/googleapis/gnostic v0.2.0

replace git.apache.org/thrift.git => github.com/apache/thrift v0.0.0-20180902110319-2566ecd5d999

replace github.com/keybase/go-crypto v0.0.0-20190523171820-b785b22cc757 => github.com/keybase/go-crypto v0.0.0-20190416182011-b785b22cc757

replace github.com/terraform-providers/terraform-provider-google v2.17.0+incompatible => github.com/terraform-providers/terraform-provider-google v1.20.1-0.20191008212436-363f2d283518

replace github.com/terraform-providers/terraform-provider-aws v2.32.0+incompatible => github.com/terraform-providers/terraform-provider-aws v1.60.1-0.20191010190908-1261a98537f2

replace github.com/terraform-providers/terraform-provider-random v2.2.1+incompatible => github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f

replace (
	bitbucket.org/ww/goautoneg => gomodules.xyz/goautoneg v0.0.0-20120707110453-a547fc61f48d
	github.com/imdario/mergo => github.com/imdario/mergo v0.3.5
	// github.com/prometheus/client_golang => github.com/prometheus/client_golang v1.0.0
	go.etcd.io/etcd => go.etcd.io/etcd v0.5.0-alpha.5.0.20200513171258-e048e166ab9c
	// google.golang.org/genproto => google.golang.org/genproto v0.0.0-20190819201941-24fa4b261c55
	// google.golang.org/grpc => google.golang.org/grpc v1.26.0
	k8s.io/api => github.com/kmodules/api v0.18.4-0.20200524125823-c8bc107809b9
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.19.0-alpha.0.0.20200520235721-10b58e57a423
	k8s.io/apiserver => github.com/kmodules/apiserver v0.18.4-0.20200521000930-14c5f6df9625
	k8s.io/client-go => k8s.io/client-go v0.18.3
)
