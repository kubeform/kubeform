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

type CodecommitTrigger struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CodecommitTriggerSpec   `json:"spec,omitempty"`
	Status            CodecommitTriggerStatus `json:"status,omitempty"`
}

type CodecommitTriggerSpecTrigger struct {
	// +optional
	Branches []string `json:"branches,omitempty"`
	// +optional
	CustomData     string   `json:"custom_data,omitempty"`
	DestinationArn string   `json:"destination_arn"`
	Events         []string `json:"events"`
	Name           string   `json:"name"`
}

type CodecommitTriggerSpec struct {
	RepositoryName string `json:"repository_name"`
	// +kubebuilder:validation:MaxItems=10
	// +kubebuilder:validation:UniqueItems=true
	Trigger []CodecommitTriggerSpec `json:"trigger"`
}

type CodecommitTriggerStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CodecommitTriggerList is a list of CodecommitTriggers
type CodecommitTriggerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CodecommitTrigger CRD objects
	Items []CodecommitTrigger `json:"items,omitempty"`
}
