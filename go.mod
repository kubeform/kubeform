module kubeform.dev/kubeform

go 1.12

require (
	github.com/appscode/go v0.0.0-20191016085057-e186b6c94a3b
	github.com/dave/jennifer v1.3.0
	github.com/go-openapi/spec v0.19.2
	github.com/gobuffalo/flect v0.1.5
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/hashicorp/terraform v0.12.7-0.20190808211310-979a2fa6d13b
	github.com/hashicorp/terraform-config-inspect v0.0.0-20190327195015-8022a2663a70
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20190510001811-4b894dbf13f6
	github.com/terraform-providers/terraform-provider-azurerm v1.32.1
	github.com/terraform-providers/terraform-provider-digitalocean v1.5.0
	github.com/terraform-providers/terraform-provider-google v1.20.0
	github.com/terraform-providers/terraform-provider-linode v1.8.0
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a // indirect
	k8s.io/api v0.0.0-20190503110853-61630f889b3c
	k8s.io/apimachinery v0.0.0-20190508063446-a3da69d3723c
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v0.3.1 // indirect
	k8s.io/kube-openapi v0.0.0-20190502190224-411b2483e503
	k8s.io/utils v0.0.0-20190506122338-8fab8cb257d5 // indirect
	kmodules.xyz/client-go v0.0.0-20191016085444-3044854ca625
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.2+incompatible
	github.com/hashicorp/hcl => github.com/hashicorp/hcl v0.0.0-20170504190234-a4b07c25de5f
	github.com/ugorji/go => github.com/ugorji/go v0.0.0-20171019201919-bdcc60b419d1
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190508082252-8397d761d4b5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190311093542-50b561225d70
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190314000054-4a91899592f4
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/kubernetes => k8s.io/kubernetes v1.14.0
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
	sigs.k8s.io/structured-merge-diff => sigs.k8s.io/structured-merge-diff v0.0.0-20190302045857-e85c7b244fd2
)
