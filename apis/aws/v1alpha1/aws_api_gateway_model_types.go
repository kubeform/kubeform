package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayModel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayModelSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayModelStatus `json:"status,omitempty"`
}

type AwsApiGatewayModelSpec struct {
	Schema      string `json:"schema"`
	ContentType string `json:"content_type"`
	RestApiId   string `json:"rest_api_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}



type AwsApiGatewayModelStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayModelList is a list of AwsApiGatewayModels
type AwsApiGatewayModelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayModel CRD objects
	Items []AwsApiGatewayModel `json:"items,omitempty"`
}