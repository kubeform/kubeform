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

type AzurermKeyVaultSpecSku struct {
	Name string `json:"name"`
}

type AzurermKeyVaultSpecAccessPolicy struct {
	SecretPermissions      []string `json:"secret_permissions"`
	StoragePermissions     []string `json:"storage_permissions"`
	TenantId               string   `json:"tenant_id"`
	ObjectId               string   `json:"object_id"`
	ApplicationId          string   `json:"application_id"`
	CertificatePermissions []string `json:"certificate_permissions"`
	KeyPermissions         []string `json:"key_permissions"`
}

type AzurermKeyVaultSpecNetworkAcls struct {
	IpRules                 []string `json:"ip_rules"`
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids"`
	DefaultAction           string   `json:"default_action"`
	Bypass                  string   `json:"bypass"`
}

type AzurermKeyVaultSpec struct {
	Location                     string                `json:"location"`
	Sku                          []AzurermKeyVaultSpec `json:"sku"`
	TenantId                     string                `json:"tenant_id"`
	AccessPolicy                 []AzurermKeyVaultSpec `json:"access_policy"`
	EnabledForDeployment         bool                  `json:"enabled_for_deployment"`
	EnabledForTemplateDeployment bool                  `json:"enabled_for_template_deployment"`
	Name                         string                `json:"name"`
	SkuName                      string                `json:"sku_name"`
	VaultUri                     string                `json:"vault_uri"`
	EnabledForDiskEncryption     bool                  `json:"enabled_for_disk_encryption"`
	NetworkAcls                  []AzurermKeyVaultSpec `json:"network_acls"`
	Tags                         map[string]string     `json:"tags"`
	ResourceGroupName            string                `json:"resource_group_name"`
}

type AzurermKeyVaultStatus struct {
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
