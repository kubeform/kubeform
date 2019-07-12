package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsSesDomainDkim struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSesDomainDkimSpec   `json:"spec,omitempty"`
	Status            AwsSesDomainDkimStatus `json:"status,omitempty"`
}

type AwsSesDomainDkimSpec struct {
	DkimTokens []string `json:"dkim_tokens"`
	Domain     string   `json:"domain"`
}

type AwsSesDomainDkimStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSesDomainDkimList is a list of AwsSesDomainDkims
type AwsSesDomainDkimList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSesDomainDkim CRD objects
	Items []AwsSesDomainDkim `json:"items,omitempty"`
}
