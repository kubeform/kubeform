package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsAppsyncResolver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsAppsyncResolverSpec   `json:"spec,omitempty"`
	Status            AwsAppsyncResolverStatus `json:"status,omitempty"`
}

type AwsAppsyncResolverSpec struct {
	Arn              string `json:"arn"`
	ApiId            string `json:"api_id"`
	Type             string `json:"type"`
	Field            string `json:"field"`
	DataSource       string `json:"data_source"`
	RequestTemplate  string `json:"request_template"`
	ResponseTemplate string `json:"response_template"`
}

type AwsAppsyncResolverStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsAppsyncResolverList is a list of AwsAppsyncResolvers
type AwsAppsyncResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsAppsyncResolver CRD objects
	Items []AwsAppsyncResolver `json:"items,omitempty"`
}
