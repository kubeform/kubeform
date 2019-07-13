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

type AwsNeptuneParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsNeptuneParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsNeptuneParameterGroupStatus `json:"status,omitempty"`
}

type AwsNeptuneParameterGroupSpecParameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

type AwsNeptuneParameterGroupSpec struct {
	Arn         string                         `json:"arn"`
	Name        string                         `json:"name"`
	Family      string                         `json:"family"`
	Description string                         `json:"description"`
	Parameter   []AwsNeptuneParameterGroupSpec `json:"parameter"`
	Tags        map[string]string              `json:"tags"`
}



type AwsNeptuneParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsNeptuneParameterGroupList is a list of AwsNeptuneParameterGroups
type AwsNeptuneParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsNeptuneParameterGroup CRD objects
	Items []AwsNeptuneParameterGroup `json:"items,omitempty"`
}