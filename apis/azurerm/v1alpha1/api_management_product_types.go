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

type ApiManagementProduct struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementProductSpec   `json:"spec,omitempty"`
	Status            ApiManagementProductStatus `json:"status,omitempty"`
}

type ApiManagementProductSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiManagementName string `json:"apiManagementName" tf:"api_management_name"`
	// +optional
	ApprovalRequired bool `json:"approvalRequired,omitempty" tf:"approval_required,omitempty"`
	// +optional
	Description          string `json:"description,omitempty" tf:"description,omitempty"`
	DisplayName          string `json:"displayName" tf:"display_name"`
	ProductID            string `json:"productID" tf:"product_id"`
	Published            bool   `json:"published" tf:"published"`
	ResourceGroupName    string `json:"resourceGroupName" tf:"resource_group_name"`
	SubscriptionRequired bool   `json:"subscriptionRequired" tf:"subscription_required"`
	// +optional
	SubscriptionsLimit int `json:"subscriptionsLimit,omitempty" tf:"subscriptions_limit,omitempty"`
	// +optional
	Terms string `json:"terms,omitempty" tf:"terms,omitempty"`
}



type ApiManagementProductStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *ApiManagementProductSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementProductList is a list of ApiManagementProducts
type ApiManagementProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagementProduct CRD objects
	Items []ApiManagementProduct `json:"items,omitempty"`
}