package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAppServiceActiveSlot struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServiceActiveSlotSpec   `json:"spec,omitempty"`
	Status            AzurermAppServiceActiveSlotStatus `json:"status,omitempty"`
}

type AzurermAppServiceActiveSlotSpec struct {
	ResourceGroupName  string `json:"resource_group_name"`
	AppServiceName     string `json:"app_service_name"`
	AppServiceSlotName string `json:"app_service_slot_name"`
}



type AzurermAppServiceActiveSlotStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAppServiceActiveSlotList is a list of AzurermAppServiceActiveSlots
type AzurermAppServiceActiveSlotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppServiceActiveSlot CRD objects
	Items []AzurermAppServiceActiveSlot `json:"items,omitempty"`
}