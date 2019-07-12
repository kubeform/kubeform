package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermStorageAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermStorageAccountSpec   `json:"spec,omitempty"`
	Status            AzurermStorageAccountStatus `json:"status,omitempty"`
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

type AzurermStorageAccountSpecIdentity struct {
	Type        string `json:"type"`
	PrincipalId string `json:"principal_id"`
	TenantId    string `json:"tenant_id"`
}

type AzurermStorageAccountSpec struct {
	Name                          string                      `json:"name"`
	Location                      string                      `json:"location"`
	SecondaryLocation             string                      `json:"secondary_location"`
	SecondaryFileHost             string                      `json:"secondary_file_host"`
	AccountKind                   string                      `json:"account_kind"`
	AccountReplicationType        string                      `json:"account_replication_type"`
	AccountEncryptionSource       string                      `json:"account_encryption_source"`
	EnableBlobEncryption          bool                        `json:"enable_blob_encryption"`
	PrimaryBlobEndpoint           string                      `json:"primary_blob_endpoint"`
	PrimaryWebHost                string                      `json:"primary_web_host"`
	AccessTier                    string                      `json:"access_tier"`
	PrimaryTableEndpoint          string                      `json:"primary_table_endpoint"`
	PrimaryDfsHost                string                      `json:"primary_dfs_host"`
	SecondaryFileEndpoint         string                      `json:"secondary_file_endpoint"`
	PrimaryConnectionString       string                      `json:"primary_connection_string"`
	SecondaryBlobConnectionString string                      `json:"secondary_blob_connection_string"`
	ResourceGroupName             string                      `json:"resource_group_name"`
	EnableHttpsTrafficOnly        bool                        `json:"enable_https_traffic_only"`
	PrimaryBlobHost               string                      `json:"primary_blob_host"`
	AccountType                   string                      `json:"account_type"`
	NetworkRules                  []AzurermStorageAccountSpec `json:"network_rules"`
	PrimaryLocation               string                      `json:"primary_location"`
	SecondaryBlobEndpoint         string                      `json:"secondary_blob_endpoint"`
	SecondaryBlobHost             string                      `json:"secondary_blob_host"`
	PrimaryQueueHost              string                      `json:"primary_queue_host"`
	PrimaryTableHost              string                      `json:"primary_table_host"`
	SecondaryWebHost              string                      `json:"secondary_web_host"`
	PrimaryDfsEndpoint            string                      `json:"primary_dfs_endpoint"`
	PrimaryAccessKey              string                      `json:"primary_access_key"`
	SecondaryConnectionString     string                      `json:"secondary_connection_string"`
	Tags                          map[string]string           `json:"tags"`
	CustomDomain                  []AzurermStorageAccountSpec `json:"custom_domain"`
	EnableFileEncryption          bool                        `json:"enable_file_encryption"`
	IsHnsEnabled                  bool                        `json:"is_hns_enabled"`
	PrimaryQueueEndpoint          string                      `json:"primary_queue_endpoint"`
	SecondaryQueueEndpoint        string                      `json:"secondary_queue_endpoint"`
	SecondaryTableHost            string                      `json:"secondary_table_host"`
	PrimaryWebEndpoint            string                      `json:"primary_web_endpoint"`
	SecondaryDfsEndpoint          string                      `json:"secondary_dfs_endpoint"`
	SecondaryDfsHost              string                      `json:"secondary_dfs_host"`
	PrimaryFileHost               string                      `json:"primary_file_host"`
	Identity                      []AzurermStorageAccountSpec `json:"identity"`
	SecondaryTableEndpoint        string                      `json:"secondary_table_endpoint"`
	SecondaryWebEndpoint          string                      `json:"secondary_web_endpoint"`
	PrimaryFileEndpoint           string                      `json:"primary_file_endpoint"`
	SecondaryAccessKey            string                      `json:"secondary_access_key"`
	PrimaryBlobConnectionString   string                      `json:"primary_blob_connection_string"`
	AccountTier                   string                      `json:"account_tier"`
	SecondaryQueueHost            string                      `json:"secondary_queue_host"`
}



type AzurermStorageAccountStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermStorageAccountList is a list of AzurermStorageAccounts
type AzurermStorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermStorageAccount CRD objects
	Items []AzurermStorageAccount `json:"items,omitempty"`
}