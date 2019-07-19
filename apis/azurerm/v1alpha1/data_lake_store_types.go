package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type DataLakeStore struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeStoreSpec   `json:"spec,omitempty"`
	Status            DataLakeStoreStatus `json:"status,omitempty"`
}

type DataLakeStoreSpec struct {
	// +optional
	EncryptionState string `json:"encryptionState,omitempty" tf:"encryption_state,omitempty"`
	// +optional
	EncryptionType string `json:"encryptionType,omitempty" tf:"encryption_type,omitempty"`
	// +optional
	FirewallAllowAzureIPS string `json:"firewallAllowAzureIPS,omitempty" tf:"firewall_allow_azure_ips,omitempty"`
	// +optional
	FirewallState     string `json:"firewallState,omitempty" tf:"firewall_state,omitempty"`
	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Tier        string                    `json:"tier,omitempty" tf:"tier,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type DataLakeStoreStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataLakeStoreList is a list of DataLakeStores
type DataLakeStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeStore CRD objects
	Items []DataLakeStore `json:"items,omitempty"`
}
