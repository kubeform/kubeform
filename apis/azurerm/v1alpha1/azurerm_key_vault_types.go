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

type AzurermKeyVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKeyVaultSpec   `json:"spec,omitempty"`
	Status            AzurermKeyVaultStatus `json:"status,omitempty"`
}

type AzurermKeyVaultSpecAccessPolicy struct {
	StoragePermissions     []string `json:"storage_permissions"`
	TenantId               string   `json:"tenant_id"`
	ObjectId               string   `json:"object_id"`
	ApplicationId          string   `json:"application_id"`
	CertificatePermissions []string `json:"certificate_permissions"`
	KeyPermissions         []string `json:"key_permissions"`
	SecretPermissions      []string `json:"secret_permissions"`
}

type AzurermKeyVaultSpecSku struct {
	Name string `json:"name"`
}

type AzurermKeyVaultSpecNetworkAcls struct {
	DefaultAction           string   `json:"default_action"`
	Bypass                  string   `json:"bypass"`
	IpRules                 []string `json:"ip_rules"`
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids"`
}

type AzurermKeyVaultSpec struct {
	AccessPolicy                 []AzurermKeyVaultSpec `json:"access_policy"`
	EnabledForDeployment         bool                  `json:"enabled_for_deployment"`
	ResourceGroupName            string                `json:"resource_group_name"`
	Sku                          []AzurermKeyVaultSpec `json:"sku"`
	SkuName                      string                `json:"sku_name"`
	VaultUri                     string                `json:"vault_uri"`
	TenantId                     string                `json:"tenant_id"`
	EnabledForDiskEncryption     bool                  `json:"enabled_for_disk_encryption"`
	EnabledForTemplateDeployment bool                  `json:"enabled_for_template_deployment"`
	NetworkAcls                  []AzurermKeyVaultSpec `json:"network_acls"`
	Name                         string                `json:"name"`
	Location                     string                `json:"location"`
	Tags                         map[string]string     `json:"tags"`
}



type AzurermKeyVaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermKeyVaultList is a list of AzurermKeyVaults
type AzurermKeyVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKeyVault CRD objects
	Items []AzurermKeyVault `json:"items,omitempty"`
}