package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type DxConnection struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DxConnectionSpec   `json:"spec,omitempty"`
	Status            DxConnectionStatus `json:"status,omitempty"`
}

type DxConnectionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AwsDevice string `json:"awsDevice,omitempty" tf:"aws_device,omitempty"`
	Bandwidth string `json:"bandwidth" tf:"bandwidth"`
	// +optional
	HasLogicalRedundancy string `json:"hasLogicalRedundancy,omitempty" tf:"has_logical_redundancy,omitempty"`
	// +optional
	JumboFrameCapable bool   `json:"jumboFrameCapable,omitempty" tf:"jumbo_frame_capable,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DxConnectionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DxConnectionSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DxConnectionList is a list of DxConnections
type DxConnectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DxConnection CRD objects
	Items []DxConnection `json:"items,omitempty"`
}
