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

type CloudformationStackSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSetSpec   `json:"spec,omitempty"`
	Status            CloudformationStackSetStatus `json:"status,omitempty"`
}

type CloudformationStackSetSpec struct {
	AdministrationRoleArn string `json:"administrationRoleArn" tf:"administration_role_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Capabilities []string `json:"capabilities,omitempty" tf:"capabilities,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	ExecutionRoleName string `json:"executionRoleName,omitempty" tf:"execution_role_name,omitempty"`
	Name              string `json:"name" tf:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty" tf:"parameters,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TemplateBody string `json:"templateBody,omitempty" tf:"template_body,omitempty"`
	// +optional
	TemplateURL string                    `json:"templateURL,omitempty" tf:"template_url,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type CloudformationStackSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// CloudformationStackSetList is a list of CloudformationStackSets
type CloudformationStackSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of CloudformationStackSet CRD objects
	Items []CloudformationStackSet `json:"items,omitempty"`
}
