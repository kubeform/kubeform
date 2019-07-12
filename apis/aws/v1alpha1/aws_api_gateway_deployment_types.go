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

type AwsApiGatewayDeployment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayDeploymentSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayDeploymentStatus `json:"status,omitempty"`
}

type AwsApiGatewayDeploymentSpec struct {
	RestApiId        string            `json:"rest_api_id"`
	StageName        string            `json:"stage_name"`
	Description      string            `json:"description"`
	StageDescription string            `json:"stage_description"`
	Variables        map[string]string `json:"variables"`
	CreatedDate      string            `json:"created_date"`
	InvokeUrl        string            `json:"invoke_url"`
	ExecutionArn     string            `json:"execution_arn"`
}

type AwsApiGatewayDeploymentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayDeploymentList is a list of AwsApiGatewayDeployments
type AwsApiGatewayDeploymentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayDeployment CRD objects
	Items []AwsApiGatewayDeployment `json:"items,omitempty"`
}
