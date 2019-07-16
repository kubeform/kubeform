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

type ApiGatewayUsagePlanKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayUsagePlanKeySpec   `json:"spec,omitempty"`
	Status            ApiGatewayUsagePlanKeyStatus `json:"status,omitempty"`
}

type ApiGatewayUsagePlanKeySpec struct {
	KeyId       string `json:"key_id"`
	KeyType     string `json:"key_type"`
	UsagePlanId string `json:"usage_plan_id"`
}

type ApiGatewayUsagePlanKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayUsagePlanKeyList is a list of ApiGatewayUsagePlanKeys
type ApiGatewayUsagePlanKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayUsagePlanKey CRD objects
	Items []ApiGatewayUsagePlanKey `json:"items,omitempty"`
}
