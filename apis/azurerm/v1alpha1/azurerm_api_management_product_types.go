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
	SubscriptionsLimit   int    `json:"subscriptions_limit"`
	ResourceGroupName    string `json:"resource_group_name"`
	Published            bool   `json:"published"`
	ApprovalRequired     bool   `json:"approval_required"`
	Description          string `json:"description"`
	Terms                string `json:"terms"`
	ProductId            string `json:"product_id"`
	ApiManagementName    string `json:"api_management_name"`
	DisplayName          string `json:"display_name"`
	SubscriptionRequired bool   `json:"subscription_required"`
}



type AzurermApiManagementProductStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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