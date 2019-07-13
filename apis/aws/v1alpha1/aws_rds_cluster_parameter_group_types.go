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

type AwsRdsClusterParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsRdsClusterParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsRdsClusterParameterGroupStatus `json:"status,omitempty"`
}

type AwsRdsClusterParameterGroupSpecParameter struct {
	Name        string `json:"name"`
	Value       string `json:"value"`
	ApplyMethod string `json:"apply_method"`
}

type AwsRdsClusterParameterGroupSpec struct {
	Family      string                            `json:"family"`
	Description string                            `json:"description"`
	Parameter   []AwsRdsClusterParameterGroupSpec `json:"parameter"`
	Tags        map[string]string                 `json:"tags"`
	Arn         string                            `json:"arn"`
	Name        string                            `json:"name"`
	NamePrefix  string                            `json:"name_prefix"`
}



type AwsRdsClusterParameterGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsRdsClusterParameterGroupList is a list of AwsRdsClusterParameterGroups
type AwsRdsClusterParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsRdsClusterParameterGroup CRD objects
	Items []AwsRdsClusterParameterGroup `json:"items,omitempty"`
}