package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermDataLakeAnalyticsAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataLakeAnalyticsAccountSpec   `json:"spec,omitempty"`
	Status            AzurermDataLakeAnalyticsAccountStatus `json:"status,omitempty"`
}

type AzurermDataLakeAnalyticsAccountSpec struct {
	DefaultStoreAccountName string            `json:"default_store_account_name"`
	Tags                    map[string]string `json:"tags"`
	Name                    string            `json:"name"`
	Location                string            `json:"location"`
	ResourceGroupName       string            `json:"resource_group_name"`
	Tier                    string            `json:"tier"`
}



type AzurermDataLakeAnalyticsAccountStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermDataLakeAnalyticsAccountList is a list of AzurermDataLakeAnalyticsAccounts
type AzurermDataLakeAnalyticsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataLakeAnalyticsAccount CRD objects
	Items []AzurermDataLakeAnalyticsAccount `json:"items,omitempty"`
}