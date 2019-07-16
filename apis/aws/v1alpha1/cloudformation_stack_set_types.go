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

type CloudformationStackSet struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CloudformationStackSetSpec   `json:"spec,omitempty"`
	Status            CloudformationStackSetStatus `json:"status,omitempty"`
}

type CloudformationStackSetSpec struct {
	AdministrationRoleArn string `json:"administration_role_arn"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Capabilities []string `json:"capabilities,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	ExecutionRoleName string `json:"execution_role_name,omitempty"`
	Name              string `json:"name"`
	// +optional
	Parameters map[string]string `json:"parameters,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	TemplateUrl string `json:"template_url,omitempty"`
}

type CloudformationStackSetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
