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

type AppServiceSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceSlotSpec   `json:"spec,omitempty"`
	Status            AppServiceSlotStatus `json:"status,omitempty"`
}

type AppServiceSlotSpec struct {
	AppServiceName   string `json:"app_service_name"`
	AppServicePlanId string `json:"app_service_plan_id"`
	// +optional
	Enabled bool `json:"enabled,omitempty"`
	// +optional
	HttpsOnly         bool   `json:"https_only,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
}

type AppServiceSlotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceSlotList is a list of AppServiceSlots
type AppServiceSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceSlot CRD objects
	Items []AppServiceSlot `json:"items,omitempty"`
}
