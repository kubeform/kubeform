package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElasticacheParameterGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticacheParameterGroupSpec   `json:"spec,omitempty"`
	Status            AwsElasticacheParameterGroupStatus `json:"status,omitempty"`
}

type AwsElasticacheParameterGroupSpecParameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type AwsElasticacheParameterGroupSpec struct {
	Parameter   []AwsElasticacheParameterGroupSpec `json:"parameter"`
	Name        string                             `json:"name"`
	Family      string                             `json:"family"`
	Description string                             `json:"description"`
}

type AwsElasticacheParameterGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElasticacheParameterGroupList is a list of AwsElasticacheParameterGroups
type AwsElasticacheParameterGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheParameterGroup CRD objects
	Items []AwsElasticacheParameterGroup `json:"items,omitempty"`
}
