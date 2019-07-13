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

type AzurermStorageAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageAccountSpec   `json:"spec,omitempty"`
	Status            AzurermStorageAccountStatus `json:"status,omitempty"`
}

type AzurermStorageAccountSpecNetworkRules struct {
	IpRules                 []string `json:"ip_rules"`
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids"`
	DefaultAction           string   `json:"default_action"`
	Bypass                  []string `json:"bypass"`
}

type AzurermStorageAccountSpecCustomDomain struct {
	Name         string `json:"name"`
	UseSubdomain bool   `json:"use_subdomain"`
}

type AzurermStorageAccountSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermStorageAccountSpec struct {
	SecondaryDfsEndpoint          string                      `json:"secondary_dfs_endpoint"`
	SecondaryDfsHost              string                      `json:"secondary_dfs_host"`
	PrimaryFileEndpoint           string                      `json:"primary_file_endpoint"`
	SecondaryFileHost             string                      `json:"secondary_file_host"`
	PrimaryBlobHost               string                      `json:"primary_blob_host"`
	SecondaryBlobEndpoint         string                      `json:"secondary_blob_endpoint"`
	SecondaryQueueEndpoint        string                      `json:"secondary_queue_endpoint"`
	NetworkRules                  []AzurermStorageAccountSpec `json:"network_rules"`
	SecondaryQueueHost            string                      `json:"secondary_queue_host"`
	PrimaryWebEndpoint            string                      `json:"primary_web_endpoint"`
	PrimaryDfsEndpoint            string                      `json:"primary_dfs_endpoint"`
	PrimaryBlobConnectionString   string                      `json:"primary_blob_connection_string"`
	Location                      string                      `json:"location"`
	CustomDomain                  []AzurermStorageAccountSpec `json:"custom_domain"`
	EnableHttpsTrafficOnly        bool                        `json:"enable_https_traffic_only"`
	Identity                      []AzurermStorageAccountSpec `json:"identity"`
	PrimaryQueueEndpoint          string                      `json:"primary_queue_endpoint"`
	PrimaryTableEndpoint          string                      `json:"primary_table_endpoint"`
	PrimaryConnectionString       string                      `json:"primary_connection_string"`
	SecondaryBlobConnectionString string                      `json:"secondary_blob_connection_string"`
	ResourceGroupName             string                      `json:"resource_group_name"`
	AccountType                   string                      `json:"account_type"`
	EnableBlobEncryption          bool                        `json:"enable_blob_encryption"`
	PrimaryDfsHost                string                      `json:"primary_dfs_host"`
	PrimaryFileHost               string                      `json:"primary_file_host"`
	SecondaryFileEndpoint         string                      `json:"secondary_file_endpoint"`
	PrimaryWebHost                string                      `json:"primary_web_host"`
	SecondaryWebHost              string                      `json:"secondary_web_host"`
	PrimaryAccessKey              string                      `json:"primary_access_key"`
	Name                          string                      `json:"name"`
	AccountKind                   string                      `json:"account_kind"`
	AccountEncryptionSource       string                      `json:"account_encryption_source"`
	PrimaryTableHost              string                      `json:"primary_table_host"`
	SecondaryWebEndpoint          string                      `json:"secondary_web_endpoint"`
	Tags                          map[string]string           `json:"tags"`
	AccessTier                    string                      `json:"access_tier"`
	IsHnsEnabled                  bool                        `json:"is_hns_enabled"`
	PrimaryBlobEndpoint           string                      `json:"primary_blob_endpoint"`
	SecondaryTableEndpoint        string                      `json:"secondary_table_endpoint"`
	SecondaryTableHost            string                      `json:"secondary_table_host"`
	SecondaryAccessKey            string                      `json:"secondary_access_key"`
	SecondaryConnectionString     string                      `json:"secondary_connection_string"`
	AccountReplicationType        string                      `json:"account_replication_type"`
	EnableFileEncryption          bool                        `json:"enable_file_encryption"`
	PrimaryQueueHost              string                      `json:"primary_queue_host"`
	SecondaryBlobHost             string                      `json:"secondary_blob_host"`
	AccountTier                   string                      `json:"account_tier"`
	PrimaryLocation               string                      `json:"primary_location"`
	SecondaryLocation             string                      `json:"secondary_location"`
}



type AzurermStorageAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermStorageAccountList is a list of AzurermStorageAccounts
type AzurermStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageAccount CRD objects
	Items []AzurermStorageAccount `json:"items,omitempty"`
}