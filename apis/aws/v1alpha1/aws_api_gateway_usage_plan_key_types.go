package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsApiGatewayUsagePlanKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsApiGatewayUsagePlanKeySpec   `json:"spec,omitempty"`
	Status            AwsApiGatewayUsagePlanKeyStatus `json:"status,omitempty"`
}

type AwsApiGatewayUsagePlanKeySpec struct {
	KeyId       string `json:"key_id"`
	KeyType     string `json:"key_type"`
	UsagePlanId string `json:"usage_plan_id"`
	Name        string `json:"name"`
	Value       string `json:"value"`
}



type AwsApiGatewayUsagePlanKeyStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsApiGatewayUsagePlanKeyList is a list of AwsApiGatewayUsagePlanKeys
type AwsApiGatewayUsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsApiGatewayUsagePlanKey CRD objects
	Items []AwsApiGatewayUsagePlanKey `json:"items,omitempty"`
}