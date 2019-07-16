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

type GlobalacceleratorListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorListenerSpec   `json:"spec,omitempty"`
	Status            GlobalacceleratorListenerStatus `json:"status,omitempty"`
}

type GlobalacceleratorListenerSpecPortRange struct {
	// +optional
	FromPort int `json:"from_port,omitempty"`
	// +optional
	ToPort int `json:"to_port,omitempty"`
}

type GlobalacceleratorListenerSpec struct {
	AcceleratorArn string `json:"accelerator_arn"`
	// +optional
	ClientAffinity string `json:"client_affinity,omitempty"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	PortRange []GlobalacceleratorListenerSpec `json:"port_range"`
	Protocol  string                          `json:"protocol"`
}

type GlobalacceleratorListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
