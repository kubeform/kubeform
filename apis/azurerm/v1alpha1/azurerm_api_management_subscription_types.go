package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermApiManagementSubscription struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermApiManagementSubscriptionSpec   `json:"spec,omitempty"`
	Status            AzurermApiManagementSubscriptionStatus `json:"status,omitempty"`
}

type AzurermApiManagementSubscriptionSpec struct {
	ApiManagementName string `json:"api_management_name"`
	State             string `json:"state"`
	PrimaryKey        string `json:"primary_key"`
	SecondaryKey      string `json:"secondary_key"`
	SubscriptionId    string `json:"subscription_id"`
	UserId            string `json:"user_id"`
	ProductId         string `json:"product_id"`
	ResourceGroupName string `json:"resource_group_name"`
	DisplayName       string `json:"display_name"`
}



type AzurermApiManagementSubscriptionStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermApiManagementSubscriptionList is a list of AzurermApiManagementSubscriptions
type AzurermApiManagementSubscriptionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermApiManagementSubscription CRD objects
	Items []AzurermApiManagementSubscription `json:"items,omitempty"`
}