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

type CloudformationStackSetInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSetInstanceSpec   `json:"spec,omitempty"`
	Status            CloudformationStackSetInstanceStatus `json:"status,omitempty"`
}

type CloudformationStackSetInstanceSpec struct {
	// +optional
	ParameterOverrides map[string]string `json:"parameter_overrides,omitempty"`
	// +optional
	RetainStack  bool   `json:"retain_stack,omitempty"`
	StackSetName string `json:"stack_set_name"`
}

type CloudformationStackSetInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudformationStackSetInstanceList is a list of CloudformationStackSetInstances
type CloudformationStackSetInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudformationStackSetInstance CRD objects
	Items []CloudformationStackSetInstance `json:"items,omitempty"`
}
