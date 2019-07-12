package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementProduct struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementProductSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementProductStatus `json:"status,omitempty"`
}

type AzurermApiManagementProductSpec struct {
	ApiManagementName    string `json:"api_management_name"`
	DisplayName          string `json:"display_name"`
	SubscriptionRequired bool   `json:"subscription_required"`
	Description          string `json:"description"`
	Terms                string `json:"terms"`
	ProductId            string `json:"product_id"`
	ResourceGroupName    string `json:"resource_group_name"`
	Published            bool   `json:"published"`
	ApprovalRequired     bool   `json:"approval_required"`
	SubscriptionsLimit   int    `json:"subscriptions_limit"`
}



type AzurermApiManagementProductStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementProductList is a list of AzurermApiManagementProducts
type AzurermApiManagementProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementProduct CRD objects
	Items []AzurermApiManagementProduct `json:"items,omitempty"`
}