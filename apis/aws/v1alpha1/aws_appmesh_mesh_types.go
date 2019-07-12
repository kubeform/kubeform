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

type AwsAppmeshMesh struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppmeshMeshSpec   `json:"spec,omitempty"`
	Status            AwsAppmeshMeshStatus `json:"status,omitempty"`
}

type AwsAppmeshMeshSpecSpecEgressFilter struct {
	Type string `json:"type"`
}

type AwsAppmeshMeshSpecSpec struct {
	EgressFilter []AwsAppmeshMeshSpecSpec `json:"egress_filter"`
}

type AwsAppmeshMeshSpec struct {
	Name            string               `json:"name"`
	Spec            []AwsAppmeshMeshSpec `json:"spec"`
	Arn             string               `json:"arn"`
	CreatedDate     string               `json:"created_date"`
	LastUpdatedDate string               `json:"last_updated_date"`
	Tags            map[string]string    `json:"tags"`
}

type AwsAppmeshMeshStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsAppmeshMeshList is a list of AwsAppmeshMeshs
type AwsAppmeshMeshList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppmeshMesh CRD objects
	Items []AwsAppmeshMesh `json:"items,omitempty"`
}
