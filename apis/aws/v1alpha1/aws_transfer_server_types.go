package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsTransferServer struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsTransferServerSpec   `json:"spec,omitempty"`
	Status            AwsTransferServerStatus `json:"status,omitempty"`
}

type AwsTransferServerSpecEndpointDetails struct {
	VpcEndpointId string `json:"vpc_endpoint_id"`
}

type AwsTransferServerSpec struct {
	LoggingRole          string                  `json:"logging_role"`
	ForceDestroy         bool                    `json:"force_destroy"`
	Tags                 map[string]string       `json:"tags"`
	Arn                  string                  `json:"arn"`
	Endpoint             string                  `json:"endpoint"`
	EndpointType         string                  `json:"endpoint_type"`
	Url                  string                  `json:"url"`
	EndpointDetails      []AwsTransferServerSpec `json:"endpoint_details"`
	InvocationRole       string                  `json:"invocation_role"`
	IdentityProviderType string                  `json:"identity_provider_type"`
}



type AwsTransferServerStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsTransferServerList is a list of AwsTransferServers
type AwsTransferServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsTransferServer CRD objects
	Items []AwsTransferServer `json:"items,omitempty"`
}