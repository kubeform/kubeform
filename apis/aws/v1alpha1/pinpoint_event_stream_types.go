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

type PinpointEventStream struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PinpointEventStreamSpec   `json:"spec,omitempty"`
	Status            PinpointEventStreamStatus `json:"status,omitempty"`
}

type PinpointEventStreamSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ApplicationID        string `json:"applicationID" tf:"application_id"`
	DestinationStreamArn string `json:"destinationStreamArn" tf:"destination_stream_arn"`
	RoleArn              string `json:"roleArn" tf:"role_arn"`
}

type PinpointEventStreamStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PinpointEventStreamList is a list of PinpointEventStreams
type PinpointEventStreamList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PinpointEventStream CRD objects
	Items []PinpointEventStream `json:"items,omitempty"`
}
