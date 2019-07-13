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

type AwsApiGatewayApiKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayApiKeySpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayApiKeyStatus `json:"status,omitempty"`
}

type AwsApiGatewayApiKeySpecStageKey struct {
	RestApiId string `json:"rest_api_id"`
	StageName string `json:"stage_name"`
}

type AwsApiGatewayApiKeySpec struct {
	Description     string                    `json:"description"`
	Enabled         bool                      `json:"enabled"`
	StageKey        []AwsApiGatewayApiKeySpec `json:"stage_key"`
	CreatedDate     string                    `json:"created_date"`
	LastUpdatedDate string                    `json:"last_updated_date"`
	Value           string                    `json:"value"`
	Name            string                    `json:"name"`
}



type AwsApiGatewayApiKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsApiGatewayApiKeyList is a list of AwsApiGatewayApiKeys
type AwsApiGatewayApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayApiKey CRD objects
	Items []AwsApiGatewayApiKey `json:"items,omitempty"`
}