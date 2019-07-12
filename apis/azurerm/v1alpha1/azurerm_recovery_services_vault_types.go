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

type AzurermRecoveryServicesVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermRecoveryServicesVaultSpec   `json:"spec,omitempty"`
	Status            AzurermRecoveryServicesVaultStatus `json:"status,omitempty"`
}

type AzurermRecoveryServicesVaultSpec struct {
	Tags              map[string]string `json:"tags"`
	Sku               string            `json:"sku"`
	Name              string            `json:"name"`
	Location          string            `json:"location"`
	ResourceGroupName string            `json:"resource_group_name"`
}

type AzurermRecoveryServicesVaultStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermRecoveryServicesVaultList is a list of AzurermRecoveryServicesVaults
type AzurermRecoveryServicesVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermRecoveryServicesVault CRD objects
	Items []AzurermRecoveryServicesVault `json:"items,omitempty"`
}
