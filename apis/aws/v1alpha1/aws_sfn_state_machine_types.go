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

type AwsSfnStateMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSfnStateMachineSpec   `json:"spec,omitempty"`
	Status            AwsSfnStateMachineStatus `json:"status,omitempty"`
}

type AwsSfnStateMachineSpec struct {
	RoleArn      string            `json:"role_arn"`
	CreationDate string            `json:"creation_date"`
	Status       string            `json:"status"`
	Tags         map[string]string `json:"tags"`
	Definition   string            `json:"definition"`
	Name         string            `json:"name"`
}



type AwsSfnStateMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSfnStateMachineList is a list of AwsSfnStateMachines
type AwsSfnStateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSfnStateMachine CRD objects
	Items []AwsSfnStateMachine `json:"items,omitempty"`
}