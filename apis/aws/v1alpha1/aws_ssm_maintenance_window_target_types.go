package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSsmMaintenanceWindowTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmMaintenanceWindowTargetSpec   `json:"spec,omitempty"`
	Status            AwsSsmMaintenanceWindowTargetStatus `json:"status,omitempty"`
}

type AwsSsmMaintenanceWindowTargetSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type AwsSsmMaintenanceWindowTargetSpec struct {
	ResourceType     string                              `json:"resource_type"`
	Targets          []AwsSsmMaintenanceWindowTargetSpec `json:"targets"`
	OwnerInformation string                              `json:"owner_information"`
	WindowId         string                              `json:"window_id"`
}

type AwsSsmMaintenanceWindowTargetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmMaintenanceWindowTargetList is a list of AwsSsmMaintenanceWindowTargets
type AwsSsmMaintenanceWindowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmMaintenanceWindowTarget CRD objects
	Items []AwsSsmMaintenanceWindowTarget `json:"items,omitempty"`
}
