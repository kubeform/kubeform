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

type AwsApiGatewayAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayAccountSpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayAccountStatus `json:"status,omitempty"`
}

type AwsApiGatewayAccountSpecThrottleSettings struct {
	BurstLimit int     `json:"burst_limit"`
	RateLimit  float64 `json:"rate_limit"`
}

type AwsApiGatewayAccountSpec struct {
	CloudwatchRoleArn string                     `json:"cloudwatch_role_arn"`
	ThrottleSettings  []AwsApiGatewayAccountSpec `json:"throttle_settings"`
}



type AwsApiGatewayAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayAccountList is a list of AwsApiGatewayAccounts
type AwsApiGatewayAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayAccount CRD objects
	Items []AwsApiGatewayAccount `json:"items,omitempty"`
}