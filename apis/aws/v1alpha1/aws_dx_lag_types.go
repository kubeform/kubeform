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

type AwsDxLag struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxLagSpec   `json:"spec,omitempty"`
	Status            AwsDxLagStatus `json:"status,omitempty"`
}

type AwsDxLagSpec struct {
	Name                 string            `json:"name"`
	ConnectionsBandwidth string            `json:"connections_bandwidth"`
	Location             string            `json:"location"`
	ForceDestroy         bool              `json:"force_destroy"`
	JumboFrameCapable    bool              `json:"jumbo_frame_capable"`
	Tags                 map[string]string `json:"tags"`
	Arn                  string            `json:"arn"`
	NumberOfConnections  int               `json:"number_of_connections"`
	HasLogicalRedundancy string            `json:"has_logical_redundancy"`
}



type AwsDxLagStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxLagList is a list of AwsDxLags
type AwsDxLagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxLag CRD objects
	Items []AwsDxLag `json:"items,omitempty"`
}