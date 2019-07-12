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

type AzurermStorageAccountSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermStorageAccountSpecNetworkRules struct {
	Bypass                  []string `json:"bypass"`
	IpRules                 []string `json:"ip_rules"`
	VirtualNetworkSubnetIds []string `json:"virtual_network_subnet_ids"`
	DefaultAction           string   `json:"default_action"`
}

type AzurermStorageAccountSpecCustomDomain struct {
	Name         string `json:"name"`
	UseSubdomain bool   `json:"use_subdomain"`
}

type AzurermStorageAccountSpec struct {
	Tags                          map[string]string           `json:"tags"`
	AccountReplicationType        string                      `json:"account_replication_type"`
	PrimaryLocation               string                      `json:"primary_location"`
	SecondaryTableHost            string                      `json:"secondary_table_host"`
	PrimaryWebHost                string                      `json:"primary_web_host"`
	PrimaryDfsHost                string                      `json:"primary_dfs_host"`
	PrimaryFileHost               string                      `json:"primary_file_host"`
	SecondaryWebEndpoint          string                      `json:"secondary_web_endpoint"`
	PrimaryConnectionString       string                      `json:"primary_connection_string"`
	AccessTier                    string                      `json:"access_tier"`
	AccountEncryptionSource       string                      `json:"account_encryption_source"`
	EnableBlobEncryption          bool                        `json:"enable_blob_encryption"`
	SecondaryLocation             string                      `json:"secondary_location"`
	PrimaryBlobEndpoint           string                      `json:"primary_blob_endpoint"`
	PrimaryTableHost              string                      `json:"primary_table_host"`
	Identity                      []AzurermStorageAccountSpec `json:"identity"`
	AccountKind                   string                      `json:"account_kind"`
	EnableHttpsTrafficOnly        bool                        `json:"enable_https_traffic_only"`
	PrimaryQueueHost              string                      `json:"primary_queue_host"`
	SecondaryQueueHost            string                      `json:"secondary_queue_host"`
	SecondaryDfsEndpoint          string                      `json:"secondary_dfs_endpoint"`
	SecondaryFileHost             string                      `json:"secondary_file_host"`
	Location                      string                      `json:"location"`
	PrimaryBlobHost               string                      `json:"primary_blob_host"`
	PrimaryWebEndpoint            string                      `json:"primary_web_endpoint"`
	SecondaryWebHost              string                      `json:"secondary_web_host"`
	SecondaryFileEndpoint         string                      `json:"secondary_file_endpoint"`
	PrimaryBlobConnectionString   string                      `json:"primary_blob_connection_string"`
	EnableFileEncryption          bool                        `json:"enable_file_encryption"`
	SecondaryBlobEndpoint         string                      `json:"secondary_blob_endpoint"`
	SecondaryBlobHost             string                      `json:"secondary_blob_host"`
	SecondaryDfsHost              string                      `json:"secondary_dfs_host"`
	PrimaryAccessKey              string                      `json:"primary_access_key"`
	SecondaryAccessKey            string                      `json:"secondary_access_key"`
	Name                          string                      `json:"name"`
	AccountType                   string                      `json:"account_type"`
	IsHnsEnabled                  bool                        `json:"is_hns_enabled"`
	SecondaryQueueEndpoint        string                      `json:"secondary_queue_endpoint"`
	PrimaryTableEndpoint          string                      `json:"primary_table_endpoint"`
	PrimaryFileEndpoint           string                      `json:"primary_file_endpoint"`
	SecondaryConnectionString     string                      `json:"secondary_connection_string"`
	NetworkRules                  []AzurermStorageAccountSpec `json:"network_rules"`
	SecondaryTableEndpoint        string                      `json:"secondary_table_endpoint"`
	SecondaryBlobConnectionString string                      `json:"secondary_blob_connection_string"`
	ResourceGroupName             string                      `json:"resource_group_name"`
	AccountTier                   string                      `json:"account_tier"`
	CustomDomain                  []AzurermStorageAccountSpec `json:"custom_domain"`
	PrimaryQueueEndpoint          string                      `json:"primary_queue_endpoint"`
	PrimaryDfsEndpoint            string                      `json:"primary_dfs_endpoint"`
}

type AzurermStorageAccountStatus struct {
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
