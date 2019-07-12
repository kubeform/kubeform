package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsEcsCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsEcsClusterSpec   `json:"spec,omitempty"`
	Status            AwsEcsClusterStatus `json:"status,omitempty"`
}

type AwsEcsClusterSpec struct {
	Tags map[string]string `json:"tags"`
	Arn  string            `json:"arn"`
	Name string            `json:"name"`
}



type AwsEcsClusterStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsEcsClusterList is a list of AwsEcsClusters
type AwsEcsClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsEcsCluster CRD objects
	Items []AwsEcsCluster `json:"items,omitempty"`
}