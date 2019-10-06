package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
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
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	KeyID   string `json:"keyID" tf:"key_id"`
	KeyType string `json:"keyType" tf:"key_type"`
	// +optional
	Name        string `json:"name,omitempty" tf:"name,omitempty"`
	UsagePlanID string `json:"usagePlanID" tf:"usage_plan_id"`
	// +optional
	Value string `json:"value,omitempty" tf:"value,omitempty"`
}

type ApiGatewayUsagePlanKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayUsagePlanKeySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
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
