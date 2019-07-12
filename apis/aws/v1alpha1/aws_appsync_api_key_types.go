package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncApiKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncApiKeySpec   `json:"spec,omitempty"`
	Status            AwsAppsyncApiKeyStatus `json:"status,omitempty"`
}

type AwsAppsyncApiKeySpec struct {
	Description string `json:"description"`
	ApiId       string `json:"api_id"`
	Expires     string `json:"expires"`
	Key         string `json:"key"`
}



type AwsAppsyncApiKeyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncApiKeyList is a list of AwsAppsyncApiKeys
type AwsAppsyncApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncApiKey CRD objects
	Items []AwsAppsyncApiKey `json:"items,omitempty"`
}