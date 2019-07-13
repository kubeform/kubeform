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

type AwsDocdbClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDocdbClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsDocdbClusterParameterGroupStatus `json:"status,omitempty"`
}

type AwsDocdbClusterParameterGroupSpecParameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

type AwsDocdbClusterParameterGroupSpec struct {
	Name        string                              `json:"name"`
	NamePrefix  string                              `json:"name_prefix"`
	Family      string                              `json:"family"`
	Description string                              `json:"description"`
	Parameter   []AwsDocdbClusterParameterGroupSpec `json:"parameter"`
	Tags        map[string]string                   `json:"tags"`
	Arn         string                              `json:"arn"`
}



type AwsDocdbClusterParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDocdbClusterParameterGroupList is a list of AwsDocdbClusterParameterGroups
type AwsDocdbClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDocdbClusterParameterGroup CRD objects
	Items []AwsDocdbClusterParameterGroup `json:"items,omitempty"`
}