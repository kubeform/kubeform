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

type AzurermMediaServicesAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermMediaServicesAccountSpec   `json:"spec,omitempty"`
	Status            AzurermMediaServicesAccountStatus `json:"status,omitempty"`
}

type AzurermMediaServicesAccountSpecStorageAccount struct {
	Id        string `json:"id"`
	IsPrimary bool   `json:"is_primary"`
}

type AzurermMediaServicesAccountSpec struct {
	Name              string                            `json:"name"`
	Location          string                            `json:"location"`
	ResourceGroupName string                            `json:"resource_group_name"`
	StorageAccount    []AzurermMediaServicesAccountSpec `json:"storage_account"`
}



type AzurermMediaServicesAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermMediaServicesAccountList is a list of AzurermMediaServicesAccounts
type AzurermMediaServicesAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermMediaServicesAccount CRD objects
	Items []AzurermMediaServicesAccount `json:"items,omitempty"`
}