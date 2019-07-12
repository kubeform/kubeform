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

type AwsDxConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDxConnectionSpec   `json:"spec,omitempty"`
	Status            AwsDxConnectionStatus `json:"status,omitempty"`
}

type AwsDxConnectionSpec struct {
	Bandwidth            string            `json:"bandwidth"`
	Location             string            `json:"location"`
	JumboFrameCapable    bool              `json:"jumbo_frame_capable"`
	Tags                 map[string]string `json:"tags"`
	HasLogicalRedundancy string            `json:"has_logical_redundancy"`
	AwsDevice            string            `json:"aws_device"`
	Arn                  string            `json:"arn"`
	Name                 string            `json:"name"`
}

type AwsDxConnectionStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDxConnectionList is a list of AwsDxConnections
type AwsDxConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDxConnection CRD objects
	Items []AwsDxConnection `json:"items,omitempty"`
}
