package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAthenaNamedQuery struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAthenaNamedQuerySpec   `json:"spec,omitempty"`
	Status            AwsAthenaNamedQueryStatus `json:"status,omitempty"`
}

type AwsAthenaNamedQuerySpec struct {
	Name        string `json:"name"`
	Query       string `json:"query"`
	Database    string `json:"database"`
	Description string `json:"description"`
}



type AwsAthenaNamedQueryStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAthenaNamedQueryList is a list of AwsAthenaNamedQuerys
type AwsAthenaNamedQueryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAthenaNamedQuery CRD objects
	Items []AwsAthenaNamedQuery `json:"items,omitempty"`
}