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

type KeyVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultSpec   `json:"spec,omitempty"`
	Status            KeyVaultStatus `json:"status,omitempty"`
}

type KeyVaultSpecNetworkAcls struct {
	Bypass        string `json:"bypass"`
	DefaultAction string `json:"default_action"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpRules []string `json:"ip_rules,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids,omitempty"`
}

type KeyVaultSpec struct {
	// +optional
	EnabledForDeployment bool `json:"enabled_for_deployment,omitempty"`
	// +optional
	EnabledForDiskEncryption bool `json:"enabled_for_disk_encryption,omitempty"`
	// +optional
	EnabledForTemplateDeployment bool   `json:"enabled_for_template_deployment,omitempty"`
	Location                     string `json:"location"`
	Name                         string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkAcls       *[]KeyVaultSpec `json:"network_acls,omitempty"`
	ResourceGroupName string          `json:"resource_group_name"`
	TenantId          string          `json:"tenant_id"`
}

type KeyVaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultList is a list of KeyVaults
type KeyVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVault CRD objects
	Items []KeyVault `json:"items,omitempty"`
}
