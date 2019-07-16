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

type ApiManagementSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSubscriptionSpec   `json:"spec,omitempty"`
	Status            ApiManagementSubscriptionStatus `json:"status,omitempty"`
}

type ApiManagementSubscriptionSpec struct {
	ApiManagementName string `json:"api_management_name"`
	DisplayName       string `json:"display_name"`
	ProductId         string `json:"product_id"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	State  string `json:"state,omitempty"`
	UserId string `json:"user_id"`
}

type ApiManagementSubscriptionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
