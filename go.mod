module kubeform.dev/kubeform

go 1.12

require (
	github.com/appscode/go v0.0.0-20190621064509-6b292c9166e3
	github.com/dave/jennifer v1.3.0
	github.com/emicklei/go-restful v2.9.6+incompatible // indirect
	github.com/evanphx/json-patch v4.5.0+incompatible // indirect
	github.com/go-critic/go-critic v0.0.0-20181204210945-ee9bf5809ead // indirect
	github.com/go-openapi/spec v0.19.2
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/hashicorp/terraform v0.12.2
	github.com/mattn/go-colorable v0.1.2 // indirect
	github.com/sirupsen/logrus v1.2.0 // indirect
	github.com/spf13/afero v1.2.2 // indirect
	github.com/terraform-providers/terraform-provider-aws v0.0.0-20190711201143-6a59557e3feb
	github.com/terraform-providers/terraform-provider-azurerm v1.31.1-0.20190715192829-94907ca9c29d
	github.com/terraform-providers/terraform-provider-digitalocean v1.5.0
	github.com/terraform-providers/terraform-provider-google v1.20.0
	github.com/terraform-providers/terraform-provider-linode v1.7.0
	google.golang.org/api v0.6.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20190503110853-61630f889b3c // indirect
	k8s.io/apimachinery v0.0.0-20190508063446-a3da69d3723c
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/klog v0.3.1 // indirect
	k8s.io/kube-openapi v0.0.0-00010101000000-000000000000
	k8s.io/utils v0.0.0-20190506122338-8fab8cb257d5 // indirect
)

replace (
	github.com/hashicorp/go-azure-helpers => github.com/hashicorp/go-azure-helpers v0.5.0
	k8s.io/api => k8s.io/api v0.0.0-20190313235455-40a48860b5ab
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.0.0-20190315093550-53c4693659ed
	k8s.io/apimachinery => github.com/kmodules/apimachinery v0.0.0-20190508045248-a52a97a7a2bf
	k8s.io/apiserver => github.com/kmodules/apiserver v0.0.0-20190508082252-8397d761d4b5
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.0.0-20190314001948-2899ed30580f
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.0.0-20190314002645-c892ea32361a
	k8s.io/component-base => k8s.io/component-base v0.0.0-20190314000054-4a91899592f4
	k8s.io/klog => k8s.io/klog v0.3.0
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.0.0-20190314000639-da8327669ac5
	k8s.io/kube-openapi => k8s.io/kube-openapi v0.0.0-20190228160746-b3a7cee44a30
	k8s.io/metrics => k8s.io/metrics v0.0.0-20190314001731-1bd6a4002213
	k8s.io/utils => k8s.io/utils v0.0.0-20190221042446-c2654d5206da
)
