package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppServiceActiveSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceActiveSlotSpec   `json:"spec,omitempty"`
	Status            AppServiceActiveSlotStatus `json:"status,omitempty"`
}

type AppServiceActiveSlotSpec struct {
	AppServiceName     string                    `json:"appServiceName" tf:"app_service_name"`
	AppServiceSlotName string                    `json:"appServiceSlotName" tf:"app_service_slot_name"`
	ResourceGroupName  string                    `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef        core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type AppServiceActiveSlotStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceActiveSlotList is a list of AppServiceActiveSlots
type AppServiceActiveSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceActiveSlot CRD objects
	Items []AppServiceActiveSlot `json:"items,omitempty"`
}
