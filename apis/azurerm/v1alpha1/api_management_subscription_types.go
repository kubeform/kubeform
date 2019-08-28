package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiManagementSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSubscriptionSpec   `json:"spec,omitempty"`
	Status            ApiManagementSubscriptionStatus `json:"status,omitempty"`
}

type ApiManagementSubscriptionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	DisplayName       string `json:"displayName" tf:"display_name"`
	// +optional
	PrimaryKey        string `json:"-" sensitive:"true" tf:"primary_key,omitempty"`
	ProductID         string `json:"productID" tf:"product_id"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryKey string `json:"-" sensitive:"true" tf:"secondary_key,omitempty"`
	// +optional
	State string `json:"state,omitempty" tf:"state,omitempty"`
	// +optional
	SubscriptionID string `json:"subscriptionID,omitempty" tf:"subscription_id,omitempty"`
	UserID         string `json:"userID" tf:"user_id"`
}



type ApiManagementSubscriptionStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ApiManagementSubscriptionSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementSubscriptionList is a list of ApiManagementSubscriptions
type ApiManagementSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementSubscription CRD objects
	Items []ApiManagementSubscription `json:"items,omitempty"`
}