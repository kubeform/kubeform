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

type AzurermBatchAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermBatchAccountSpec   `json:"spec,omitempty"`
	Status            AzurermBatchAccountStatus `json:"status,omitempty"`
}

type AzurermBatchAccountSpec struct {
	ResourceGroupName  string            `json:"resource_group_name"`
	SecondaryAccessKey string            `json:"secondary_access_key"`
	Tags               map[string]string `json:"tags"`
	Name               string            `json:"name"`
	Location           string            `json:"location"`
	StorageAccountId   string            `json:"storage_account_id"`
	PoolAllocationMode string            `json:"pool_allocation_mode"`
	PrimaryAccessKey   string            `json:"primary_access_key"`
	AccountEndpoint    string            `json:"account_endpoint"`
}



type AzurermBatchAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermBatchAccountList is a list of AzurermBatchAccounts
type AzurermBatchAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermBatchAccount CRD objects
	Items []AzurermBatchAccount `json:"items,omitempty"`
}