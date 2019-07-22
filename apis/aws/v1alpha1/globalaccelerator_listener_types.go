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
	ProviderRef    core.LocalObjectReference `json:"providerRef" tf:"-"`
	AcceleratorArn string                    `json:"acceleratorArn" tf:"accelerator_arn"`
	// +optional
	ClientAffinity string `json:"clientAffinity,omitempty" tf:"client_affinity,omitempty"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	PortRange []GlobalacceleratorListenerSpecPortRange `json:"portRange" tf:"port_range"`
	Protocol  string                                   `json:"protocol" tf:"protocol"`
}

type GlobalacceleratorListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
