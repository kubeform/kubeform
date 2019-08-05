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
	PrincipalID string `json:"principalID" tf:"principal_id"`
	TenantID    string `json:"tenantID" tf:"tenant_id"`
	Type        string `json:"type" tf:"type"`
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
	NetworkRules                  []StorageAccountSpecNetworkRules `json:"networkRules,omitempty" tf:"network_rules,omitempty"`
	PrimaryAccessKey              string                           `json:"-" sensitive:"true" tf:"primary_access_key"`
	PrimaryBlobConnectionString   string                           `json:"-" sensitive:"true" tf:"primary_blob_connection_string"`
	PrimaryBlobEndpoint           string                           `json:"primaryBlobEndpoint" tf:"primary_blob_endpoint"`
	PrimaryBlobHost               string                           `json:"primaryBlobHost" tf:"primary_blob_host"`
	PrimaryConnectionString       string                           `json:"-" sensitive:"true" tf:"primary_connection_string"`
	PrimaryDfsEndpoint            string                           `json:"primaryDfsEndpoint" tf:"primary_dfs_endpoint"`
	PrimaryDfsHost                string                           `json:"primaryDfsHost" tf:"primary_dfs_host"`
	PrimaryFileEndpoint           string                           `json:"primaryFileEndpoint" tf:"primary_file_endpoint"`
	PrimaryFileHost               string                           `json:"primaryFileHost" tf:"primary_file_host"`
	PrimaryLocation               string                           `json:"primaryLocation" tf:"primary_location"`
	PrimaryQueueEndpoint          string                           `json:"primaryQueueEndpoint" tf:"primary_queue_endpoint"`
	PrimaryQueueHost              string                           `json:"primaryQueueHost" tf:"primary_queue_host"`
	PrimaryTableEndpoint          string                           `json:"primaryTableEndpoint" tf:"primary_table_endpoint"`
	PrimaryTableHost              string                           `json:"primaryTableHost" tf:"primary_table_host"`
	PrimaryWebEndpoint            string                           `json:"primaryWebEndpoint" tf:"primary_web_endpoint"`
	PrimaryWebHost                string                           `json:"primaryWebHost" tf:"primary_web_host"`
	ResourceGroupName             string                           `json:"resourceGroupName" tf:"resource_group_name"`
	SecondaryAccessKey            string                           `json:"-" sensitive:"true" tf:"secondary_access_key"`
	SecondaryBlobConnectionString string                           `json:"-" sensitive:"true" tf:"secondary_blob_connection_string"`
	SecondaryBlobEndpoint         string                           `json:"secondaryBlobEndpoint" tf:"secondary_blob_endpoint"`
	SecondaryBlobHost             string                           `json:"secondaryBlobHost" tf:"secondary_blob_host"`
	SecondaryConnectionString     string                           `json:"-" sensitive:"true" tf:"secondary_connection_string"`
	SecondaryDfsEndpoint          string                           `json:"secondaryDfsEndpoint" tf:"secondary_dfs_endpoint"`
	SecondaryDfsHost              string                           `json:"secondaryDfsHost" tf:"secondary_dfs_host"`
	SecondaryFileEndpoint         string                           `json:"secondaryFileEndpoint" tf:"secondary_file_endpoint"`
	SecondaryFileHost             string                           `json:"secondaryFileHost" tf:"secondary_file_host"`
	SecondaryLocation             string                           `json:"secondaryLocation" tf:"secondary_location"`
	SecondaryQueueEndpoint        string                           `json:"secondaryQueueEndpoint" tf:"secondary_queue_endpoint"`
	SecondaryQueueHost            string                           `json:"secondaryQueueHost" tf:"secondary_queue_host"`
	SecondaryTableEndpoint        string                           `json:"secondaryTableEndpoint" tf:"secondary_table_endpoint"`
	SecondaryTableHost            string                           `json:"secondaryTableHost" tf:"secondary_table_host"`
	SecondaryWebEndpoint          string                           `json:"secondaryWebEndpoint" tf:"secondary_web_endpoint"`
	SecondaryWebHost              string                           `json:"secondaryWebHost" tf:"secondary_web_host"`
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
