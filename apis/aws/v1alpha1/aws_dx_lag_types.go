package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsDxLag struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxLagSpec   `json:"spec,omitempty"`
	Status            AwsDxLagStatus `json:"status,omitempty"`
}

type AwsDxLagSpec struct {
	Name                 string            `json:"name"`
	Location             string            `json:"location"`
	Tags                 map[string]string `json:"tags"`
	Arn                  string            `json:"arn"`
	ConnectionsBandwidth string            `json:"connections_bandwidth"`
	NumberOfConnections  int               `json:"number_of_connections"`
	ForceDestroy         bool              `json:"force_destroy"`
	JumboFrameCapable    bool              `json:"jumbo_frame_capable"`
	HasLogicalRedundancy string            `json:"has_logical_redundancy"`
}



type AwsDxLagStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsDxLagList is a list of AwsDxLags
type AwsDxLagList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxLag CRD objects
	Items []AwsDxLag `json:"items,omitempty"`
}