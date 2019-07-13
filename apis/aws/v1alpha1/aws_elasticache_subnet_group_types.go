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

type AwsElasticacheSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElasticacheSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsElasticacheSubnetGroupStatus `json:"status,omitempty"`
}

type AwsElasticacheSubnetGroupSpec struct {
	Description string   `json:"description"`
	Name        string   `json:"name"`
	SubnetIds   []string `json:"subnet_ids"`
}



type AwsElasticacheSubnetGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElasticacheSubnetGroupList is a list of AwsElasticacheSubnetGroups
type AwsElasticacheSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElasticacheSubnetGroup CRD objects
	Items []AwsElasticacheSubnetGroup `json:"items,omitempty"`
}