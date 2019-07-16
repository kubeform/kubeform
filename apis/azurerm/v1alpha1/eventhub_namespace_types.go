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

type EventhubNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              EventhubNamespaceSpec   `json:"spec,omitempty"`
	Status            EventhubNamespaceStatus `json:"status,omitempty"`
}

type EventhubNamespaceSpec struct {
	// +optional
	AutoInflateEnabled bool `json:"auto_inflate_enabled,omitempty"`
	// +optional
	Capacity int `json:"capacity,omitempty"`
	// +optional
	KafkaEnabled      bool   `json:"kafka_enabled,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	Sku               string `json:"sku"`
}

type EventhubNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// EventhubNamespaceList is a list of EventhubNamespaces
type EventhubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of EventhubNamespace CRD objects
	Items []EventhubNamespace `json:"items,omitempty"`
}
