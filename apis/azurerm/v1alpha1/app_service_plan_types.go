package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppServicePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServicePlanSpec   `json:"spec,omitempty"`
	Status            AppServicePlanStatus `json:"status,omitempty"`
}

type AppServicePlanSpecProperties struct {
	// +optional
	// Deprecated
	AppServiceEnvironmentID string `json:"appServiceEnvironmentID,omitempty" tf:"app_service_environment_id,omitempty"`
	// +optional
	// Deprecated
	PerSiteScaling bool `json:"perSiteScaling,omitempty" tf:"per_site_scaling,omitempty"`
	// +optional
	// Deprecated
	Reserved bool `json:"reserved,omitempty" tf:"reserved,omitempty"`
}

type AppServicePlanSpecSku struct {
	// +optional
	Capacity int    `json:"capacity,omitempty" tf:"capacity,omitempty"`
	Size     string `json:"size" tf:"size"`
	Tier     string `json:"tier" tf:"tier"`
}

type AppServicePlanSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AppServiceEnvironmentID string `json:"appServiceEnvironmentID,omitempty" tf:"app_service_environment_id,omitempty"`
	// +optional
	IsXenon bool `json:"isXenon,omitempty" tf:"is_xenon,omitempty"`
	// +optional
	Kind     string `json:"kind,omitempty" tf:"kind,omitempty"`
	Location string `json:"location" tf:"location"`
	// +optional
	MaximumElasticWorkerCount int `json:"maximumElasticWorkerCount,omitempty" tf:"maximum_elastic_worker_count,omitempty"`
	// +optional
	MaximumNumberOfWorkers int    `json:"maximumNumberOfWorkers,omitempty" tf:"maximum_number_of_workers,omitempty"`
	Name                   string `json:"name" tf:"name"`
	// +optional
	PerSiteScaling bool `json:"perSiteScaling,omitempty" tf:"per_site_scaling,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// Deprecated
	Properties []AppServicePlanSpecProperties `json:"properties,omitempty" tf:"properties,omitempty"`
	// +optional
	Reserved          bool   `json:"reserved,omitempty" tf:"reserved,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []AppServicePlanSpecSku `json:"sku" tf:"sku"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}



type AppServicePlanStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AppServicePlanSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServicePlanList is a list of AppServicePlans
type AppServicePlanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServicePlan CRD objects
	Items []AppServicePlan `json:"items,omitempty"`
}