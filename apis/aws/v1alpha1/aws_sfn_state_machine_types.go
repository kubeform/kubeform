package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSfnStateMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSfnStateMachineSpec   `json:"spec,omitempty"`
	Status            AwsSfnStateMachineStatus `json:"status,omitempty"`
}

type AwsSfnStateMachineSpec struct {
	Name         string            `json:"name"`
	RoleArn      string            `json:"role_arn"`
	CreationDate string            `json:"creation_date"`
	Status       string            `json:"status"`
	Tags         map[string]string `json:"tags"`
	Definition   string            `json:"definition"`
}

type AwsSfnStateMachineStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSfnStateMachineList is a list of AwsSfnStateMachines
type AwsSfnStateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSfnStateMachine CRD objects
	Items []AwsSfnStateMachine `json:"items,omitempty"`
}
