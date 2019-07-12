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

type AzurermApiManagementProduct struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementProductSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementProductStatus `json:"status,omitempty"`
}

type AzurermApiManagementProductSpec struct {
	ResourceGroupName    string `json:"resource_group_name"`
	SubscriptionRequired bool   `json:"subscription_required"`
	Published            bool   `json:"published"`
	ApprovalRequired     bool   `json:"approval_required"`
	Description          string `json:"description"`
	ProductId            string `json:"product_id"`
	DisplayName          string `json:"display_name"`
	Terms                string `json:"terms"`
	SubscriptionsLimit   int    `json:"subscriptions_limit"`
	ApiManagementName    string `json:"api_management_name"`
}

type AzurermApiManagementProductStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermApiManagementProductList is a list of AzurermApiManagementProducts
type AzurermApiManagementProductList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementProduct CRD objects
	Items []AzurermApiManagementProduct `json:"items,omitempty"`
}
