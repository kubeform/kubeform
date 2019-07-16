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

type AppServicePlan struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServicePlanSpec   `json:"spec,omitempty"`
	Status            AppServicePlanStatus `json:"status,omitempty"`
}

type AppServicePlanSpecSku struct {
	Size string `json:"size"`
	Tier string `json:"tier"`
}

type AppServicePlanSpec struct {
	// +optional
	IsXenon bool `json:"is_xenon,omitempty"`
	// +optional
	Kind              string `json:"kind,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []AppServicePlanSpec `json:"sku"`
}

type AppServicePlanStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
