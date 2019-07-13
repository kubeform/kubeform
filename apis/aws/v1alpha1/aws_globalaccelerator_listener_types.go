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

type AwsGlobalacceleratorListener struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsGlobalacceleratorListenerSpec   `json:"spec,omitempty"`
	Status            AwsGlobalacceleratorListenerStatus `json:"status,omitempty"`
}

type AwsGlobalacceleratorListenerSpecPortRange struct {
	FromPort int `json:"from_port"`
	ToPort   int `json:"to_port"`
}

type AwsGlobalacceleratorListenerSpec struct {
	AcceleratorArn string                             `json:"accelerator_arn"`
	ClientAffinity string                             `json:"client_affinity"`
	Protocol       string                             `json:"protocol"`
	PortRange      []AwsGlobalacceleratorListenerSpec `json:"port_range"`
}



type AwsGlobalacceleratorListenerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsGlobalacceleratorListenerList is a list of AwsGlobalacceleratorListeners
type AwsGlobalacceleratorListenerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsGlobalacceleratorListener CRD objects
	Items []AwsGlobalacceleratorListener `json:"items,omitempty"`
}