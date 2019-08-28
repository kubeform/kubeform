package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	AccountID string `json:"accountID,omitempty" tf:"account_id,omitempty"`
	// +optional
	ParameterOverrides map[string]string `json:"parameterOverrides,omitempty" tf:"parameter_overrides,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	RetainStack bool `json:"retainStack,omitempty" tf:"retain_stack,omitempty"`
	// +optional
	StackID      string `json:"stackID,omitempty" tf:"stack_id,omitempty"`
	StackSetName string `json:"stackSetName" tf:"stack_set_name"`
}



type CloudformationStackSetInstanceStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *CloudformationStackSetInstanceSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
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