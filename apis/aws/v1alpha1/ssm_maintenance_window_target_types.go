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

type SsmMaintenanceWindowTarget struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmMaintenanceWindowTargetSpec   `json:"spec,omitempty"`
	Status            SsmMaintenanceWindowTargetStatus `json:"status,omitempty"`
}

type SsmMaintenanceWindowTargetSpecTargets struct {
	Key    string   `json:"key"`
	Values []string `json:"values"`
}

type SsmMaintenanceWindowTargetSpec struct {
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty"`
	// +optional
	OwnerInformation string `json:"owner_information,omitempty"`
	ResourceType     string `json:"resource_type"`
	// +kubebuilder:validation:MaxItems=5
	Targets  []SsmMaintenanceWindowTargetSpec `json:"targets"`
	WindowId string                           `json:"window_id"`
}

type SsmMaintenanceWindowTargetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmMaintenanceWindowTargetList is a list of SsmMaintenanceWindowTargets
type SsmMaintenanceWindowTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmMaintenanceWindowTarget CRD objects
	Items []SsmMaintenanceWindowTarget `json:"items,omitempty"`
}
