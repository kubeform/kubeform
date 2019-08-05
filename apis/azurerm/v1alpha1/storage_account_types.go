package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type StorageAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountSpec   `json:"spec,omitempty"`
	Status            StorageAccountStatus `json:"status,omitempty"`
}

type StorageAccountSpecCustomDomain struct {
	Name string `json:"name" tf:"name"`
	// +optional
	UseSubdomain bool `json:"useSubdomain,omitempty" tf:"use_subdomain,omitempty"`
}

type StorageAccountSpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type StorageAccountSpecNetworkRules struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Bypass []string `json:"bypass,omitempty" tf:"bypass,omitempty"`
	// +optional
	DefaultAction string `json:"defaultAction,omitempty" tf:"default_action,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	IpRules []string `json:"ipRules,omitempty" tf:"ip_rules,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VirtualNetworkSubnetIDS []string `json:"virtualNetworkSubnetIDS,omitempty" tf:"virtual_network_subnet_ids,omitempty"`
}

type StorageAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	// +optional
	AccessTier string `json:"accessTier,omitempty" tf:"access_tier,omitempty"`
	// +optional
	AccountEncryptionSource string `json:"accountEncryptionSource,omitempty" tf:"account_encryption_source,omitempty"`
	// +optional
	AccountKind            string `json:"accountKind,omitempty" tf:"account_kind,omitempty"`
	AccountReplicationType string `json:"accountReplicationType" tf:"account_replication_type"`
	AccountTier            string `json:"accountTier" tf:"account_tier"`
	// +optional
	// Deprecated
	AccountType string `json:"accountType,omitempty" tf:"account_type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomDomain []StorageAccountSpecCustomDomain `json:"customDomain,omitempty" tf:"custom_domain,omitempty"`
	// +optional
	EnableBlobEncryption bool `json:"enableBlobEncryption,omitempty" tf:"enable_blob_encryption,omitempty"`
	// +optional
	EnableFileEncryption bool `json:"enableFileEncryption,omitempty" tf:"enable_file_encryption,omitempty"`
	// +optional
	EnableHTTPSTrafficOnly bool `json:"enableHTTPSTrafficOnly,omitempty" tf:"enable_https_traffic_only,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity []StorageAccountSpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	// +optional
	IsHnsEnabled bool   `json:"isHnsEnabled,omitempty" tf:"is_hns_enabled,omitempty"`
	Location     string `json:"location" tf:"location"`
	Name         string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRules []StorageAccountSpecNetworkRules `json:"networkRules,omitempty" tf:"network_rules,omitempty"`
	// +optional
	PrimaryAccessKey string `json:"-" sensitive:"true" tf:"primary_access_key,omitempty"`
	// +optional
	PrimaryBlobConnectionString string `json:"-" sensitive:"true" tf:"primary_blob_connection_string,omitempty"`
	// +optional
	PrimaryBlobEndpoint string `json:"primaryBlobEndpoint,omitempty" tf:"primary_blob_endpoint,omitempty"`
	// +optional
	PrimaryBlobHost string `json:"primaryBlobHost,omitempty" tf:"primary_blob_host,omitempty"`
	// +optional
	PrimaryConnectionString string `json:"-" sensitive:"true" tf:"primary_connection_string,omitempty"`
	// +optional
	PrimaryDfsEndpoint string `json:"primaryDfsEndpoint,omitempty" tf:"primary_dfs_endpoint,omitempty"`
	// +optional
	PrimaryDfsHost string `json:"primaryDfsHost,omitempty" tf:"primary_dfs_host,omitempty"`
	// +optional
	PrimaryFileEndpoint string `json:"primaryFileEndpoint,omitempty" tf:"primary_file_endpoint,omitempty"`
	// +optional
	PrimaryFileHost string `json:"primaryFileHost,omitempty" tf:"primary_file_host,omitempty"`
	// +optional
	PrimaryLocation string `json:"primaryLocation,omitempty" tf:"primary_location,omitempty"`
	// +optional
	PrimaryQueueEndpoint string `json:"primaryQueueEndpoint,omitempty" tf:"primary_queue_endpoint,omitempty"`
	// +optional
	PrimaryQueueHost string `json:"primaryQueueHost,omitempty" tf:"primary_queue_host,omitempty"`
	// +optional
	PrimaryTableEndpoint string `json:"primaryTableEndpoint,omitempty" tf:"primary_table_endpoint,omitempty"`
	// +optional
	PrimaryTableHost string `json:"primaryTableHost,omitempty" tf:"primary_table_host,omitempty"`
	// +optional
	PrimaryWebEndpoint string `json:"primaryWebEndpoint,omitempty" tf:"primary_web_endpoint,omitempty"`
	// +optional
	PrimaryWebHost    string `json:"primaryWebHost,omitempty" tf:"primary_web_host,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	SecondaryAccessKey string `json:"-" sensitive:"true" tf:"secondary_access_key,omitempty"`
	// +optional
	SecondaryBlobConnectionString string `json:"-" sensitive:"true" tf:"secondary_blob_connection_string,omitempty"`
	// +optional
	SecondaryBlobEndpoint string `json:"secondaryBlobEndpoint,omitempty" tf:"secondary_blob_endpoint,omitempty"`
	// +optional
	SecondaryBlobHost string `json:"secondaryBlobHost,omitempty" tf:"secondary_blob_host,omitempty"`
	// +optional
	SecondaryConnectionString string `json:"-" sensitive:"true" tf:"secondary_connection_string,omitempty"`
	// +optional
	SecondaryDfsEndpoint string `json:"secondaryDfsEndpoint,omitempty" tf:"secondary_dfs_endpoint,omitempty"`
	// +optional
	SecondaryDfsHost string `json:"secondaryDfsHost,omitempty" tf:"secondary_dfs_host,omitempty"`
	// +optional
	SecondaryFileEndpoint string `json:"secondaryFileEndpoint,omitempty" tf:"secondary_file_endpoint,omitempty"`
	// +optional
	SecondaryFileHost string `json:"secondaryFileHost,omitempty" tf:"secondary_file_host,omitempty"`
	// +optional
	SecondaryLocation string `json:"secondaryLocation,omitempty" tf:"secondary_location,omitempty"`
	// +optional
	SecondaryQueueEndpoint string `json:"secondaryQueueEndpoint,omitempty" tf:"secondary_queue_endpoint,omitempty"`
	// +optional
	SecondaryQueueHost string `json:"secondaryQueueHost,omitempty" tf:"secondary_queue_host,omitempty"`
	// +optional
	SecondaryTableEndpoint string `json:"secondaryTableEndpoint,omitempty" tf:"secondary_table_endpoint,omitempty"`
	// +optional
	SecondaryTableHost string `json:"secondaryTableHost,omitempty" tf:"secondary_table_host,omitempty"`
	// +optional
	SecondaryWebEndpoint string `json:"secondaryWebEndpoint,omitempty" tf:"secondary_web_endpoint,omitempty"`
	// +optional
	SecondaryWebHost string `json:"secondaryWebHost,omitempty" tf:"secondary_web_host,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type StorageAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *StorageAccountSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// StorageAccountList is a list of StorageAccounts
type StorageAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of StorageAccount CRD objects
	Items []StorageAccount `json:"items,omitempty"`
}
