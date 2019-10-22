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

type GlobalacceleratorListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorListenerSpec   `json:"spec,omitempty"`
	Status            GlobalacceleratorListenerStatus `json:"status,omitempty"`
}

type GlobalacceleratorListenerSpecPortRange struct {
	// +optional
	FromPort int `json:"fromPort,omitempty" tf:"from_port,omitempty"`
	// +optional
	ToPort int `json:"toPort,omitempty" tf:"to_port,omitempty"`
}

type GlobalacceleratorListenerSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AcceleratorArn string `json:"acceleratorArn" tf:"accelerator_arn"`
	// +optional
	ClientAffinity string `json:"clientAffinity,omitempty" tf:"client_affinity,omitempty"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	PortRange []GlobalacceleratorListenerSpecPortRange `json:"portRange" tf:"port_range"`
	Protocol  string                                   `json:"protocol" tf:"protocol"`
}

type GlobalacceleratorListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlobalacceleratorListenerSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalacceleratorListenerList is a list of GlobalacceleratorListeners
type GlobalacceleratorListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalacceleratorListener CRD objects
	Items []GlobalacceleratorListener `json:"items,omitempty"`
}
