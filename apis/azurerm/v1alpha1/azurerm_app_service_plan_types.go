package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AzurermAppServicePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServicePlanSpec   `json:"spec,omitempty"`
	Status            AzurermAppServicePlanStatus `json:"status,omitempty"`
}

type AzurermAppServicePlanSpecSku struct {
	Size     string `json:"size"`
	Capacity int    `json:"capacity"`
	Tier     string `json:"tier"`
}

type AzurermAppServicePlanSpecProperties struct {
	AppServiceEnvironmentId string `json:"app_service_environment_id"`
	Reserved                bool   `json:"reserved"`
	PerSiteScaling          bool   `json:"per_site_scaling"`
}

type AzurermAppServicePlanSpec struct {
	Name                      string                      `json:"name"`
	Location                  string                      `json:"location"`
	PerSiteScaling            bool                        `json:"per_site_scaling"`
	MaximumElasticWorkerCount int                         `json:"maximum_elastic_worker_count"`
	MaximumNumberOfWorkers    int                         `json:"maximum_number_of_workers"`
	IsXenon                   bool                        `json:"is_xenon"`
	ResourceGroupName         string                      `json:"resource_group_name"`
	Kind                      string                      `json:"kind"`
	Sku                       []AzurermAppServicePlanSpec `json:"sku"`
	Properties                []AzurermAppServicePlanSpec `json:"properties"`
	AppServiceEnvironmentId   string                      `json:"app_service_environment_id"`
	Reserved                  bool                        `json:"reserved"`
	Tags                      map[string]string           `json:"tags"`
}

type AzurermAppServicePlanStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAppServicePlanList is a list of AzurermAppServicePlans
type AzurermAppServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppServicePlan CRD objects
	Items []AzurermAppServicePlan `json:"items,omitempty"`
}
