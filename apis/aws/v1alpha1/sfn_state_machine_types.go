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

type SfnStateMachine struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SfnStateMachineSpec   `json:"spec,omitempty"`
	Status            SfnStateMachineStatus `json:"status,omitempty"`
}

type SfnStateMachineSpec struct {
	Definition string `json:"definition"`
	Name       string `json:"name"`
	RoleArn    string `json:"role_arn"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SfnStateMachineStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SfnStateMachineList is a list of SfnStateMachines
type SfnStateMachineList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SfnStateMachine CRD objects
	Items []SfnStateMachine `json:"items,omitempty"`
}
