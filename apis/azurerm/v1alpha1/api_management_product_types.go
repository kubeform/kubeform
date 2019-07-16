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

type ApiManagementProduct struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementProductSpec   `json:"spec,omitempty"`
	Status            ApiManagementProductStatus `json:"status,omitempty"`
}

type ApiManagementProductSpec struct {
	ApiManagementName string `json:"api_management_name"`
	// +optional
	ApprovalRequired bool `json:"approval_required,omitempty"`
	// +optional
	Description          string `json:"description,omitempty"`
	DisplayName          string `json:"display_name"`
	ProductId            string `json:"product_id"`
	Published            bool   `json:"published"`
	ResourceGroupName    string `json:"resource_group_name"`
	SubscriptionRequired bool   `json:"subscription_required"`
	// +optional
	SubscriptionsLimit int `json:"subscriptions_limit,omitempty"`
	// +optional
	Terms string `json:"terms,omitempty"`
}

type ApiManagementProductStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
