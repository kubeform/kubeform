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

type AwsDaxSubnetGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsDaxSubnetGroupSpec   `json:"spec,omitempty"`
	Status            AwsDaxSubnetGroupStatus `json:"status,omitempty"`
}

type AwsDaxSubnetGroupSpec struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	SubnetIds   []string `json:"subnet_ids"`
	VpcId       string   `json:"vpc_id"`
}

type AwsDaxSubnetGroupStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsDaxSubnetGroupList is a list of AwsDaxSubnetGroups
type AwsDaxSubnetGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsDaxSubnetGroup CRD objects
	Items []AwsDaxSubnetGroup `json:"items,omitempty"`
}
