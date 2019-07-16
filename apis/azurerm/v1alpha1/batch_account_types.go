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

type BatchAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchAccountSpec   `json:"spec,omitempty"`
	Status            BatchAccountStatus `json:"status,omitempty"`
}

type BatchAccountSpecKeyVaultReference struct {
	Id  string `json:"id"`
	Url string `json:"url"`
}

type BatchAccountSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	KeyVaultReference *[]BatchAccountSpec `json:"key_vault_reference,omitempty"`
	Location          string              `json:"location"`
	Name              string              `json:"name"`
	// +optional
	PoolAllocationMode string `json:"pool_allocation_mode,omitempty"`
	ResourceGroupName  string `json:"resource_group_name"`
}

type BatchAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchAccountList is a list of BatchAccounts
type BatchAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchAccount CRD objects
	Items []BatchAccount `json:"items,omitempty"`
}
