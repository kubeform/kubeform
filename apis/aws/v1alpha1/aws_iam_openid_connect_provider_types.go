package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsIamOpenidConnectProvider struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsIamOpenidConnectProviderSpec   `json:"spec,omitempty"`
	Status            AwsIamOpenidConnectProviderStatus `json:"status,omitempty"`
}

type AwsIamOpenidConnectProviderSpec struct {
	Arn            string   `json:"arn"`
	Url            string   `json:"url"`
	ClientIdList   []string `json:"client_id_list"`
	ThumbprintList []string `json:"thumbprint_list"`
}

type AwsIamOpenidConnectProviderStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsIamOpenidConnectProviderList is a list of AwsIamOpenidConnectProviders
type AwsIamOpenidConnectProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsIamOpenidConnectProvider CRD objects
	Items []AwsIamOpenidConnectProvider `json:"items,omitempty"`
}
