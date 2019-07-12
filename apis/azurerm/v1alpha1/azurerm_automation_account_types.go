package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAutomationAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationAccountSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationAccountStatus `json:"status,omitempty"`
}

type AzurermAutomationAccountSpecSku struct {
	Name string `json:"name"`
}

type AzurermAutomationAccountSpec struct {
	DscPrimaryAccessKey   string                         `json:"dsc_primary_access_key"`
	DscSecondaryAccessKey string                         `json:"dsc_secondary_access_key"`
	Name                  string                         `json:"name"`
	Sku                   []AzurermAutomationAccountSpec `json:"sku"`
	Tags                  map[string]string              `json:"tags"`
	DscServerEndpoint     string                         `json:"dsc_server_endpoint"`
	Location              string                         `json:"location"`
	ResourceGroupName     string                         `json:"resource_group_name"`
	SkuName               string                         `json:"sku_name"`
}

type AzurermAutomationAccountStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAutomationAccountList is a list of AzurermAutomationAccounts
type AzurermAutomationAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationAccount CRD objects
	Items []AzurermAutomationAccount `json:"items,omitempty"`
}
