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

type DigitaloceanFloatingIpAssignment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DigitaloceanFloatingIpAssignmentSpec   `json:"spec,omitempty"`
	Status            DigitaloceanFloatingIpAssignmentStatus `json:"status,omitempty"`
}

type DigitaloceanFloatingIpAssignmentSpec struct {
	IpAddress string `json:"ip_address"`
	DropletId int    `json:"droplet_id"`
}



type DigitaloceanFloatingIpAssignmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DigitaloceanFloatingIpAssignmentList is a list of DigitaloceanFloatingIpAssignments
type DigitaloceanFloatingIpAssignmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DigitaloceanFloatingIpAssignment CRD objects
	Items []DigitaloceanFloatingIpAssignment `json:"items,omitempty"`
}