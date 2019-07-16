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

type StorageAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              StorageAccountSpec   `json:"spec,omitempty"`
	Status            StorageAccountStatus `json:"status,omitempty"`
}

type StorageAccountSpecCustomDomain struct {
	Name string `json:"name"`
	// +optional
	UseSubdomain bool `json:"use_subdomain,omitempty"`
}

type StorageAccountSpecNetworkRules struct {
	// +optional
	DefaultAction string `json:"default_action,omitempty"`
}

type StorageAccountSpec struct {
	// +optional
	AccountEncryptionSource string `json:"account_encryption_source,omitempty"`
	// +optional
	AccountKind            string `json:"account_kind,omitempty"`
	AccountReplicationType string `json:"account_replication_type"`
	AccountTier            string `json:"account_tier"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	CustomDomain *[]StorageAccountSpec `json:"custom_domain,omitempty"`
	// +optional
	EnableAdvancedThreatProtection bool `json:"enable_advanced_threat_protection,omitempty"`
	// +optional
	EnableBlobEncryption bool `json:"enable_blob_encryption,omitempty"`
	// +optional
	EnableFileEncryption bool `json:"enable_file_encryption,omitempty"`
	// +optional
	EnableHttpsTrafficOnly bool `json:"enable_https_traffic_only,omitempty"`
	// +optional
	IsHnsEnabled bool   `json:"is_hns_enabled,omitempty"`
	Location     string `json:"location"`
	Name         string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	NetworkRules      *[]StorageAccountSpec `json:"network_rules,omitempty"`
	ResourceGroupName string                `json:"resource_group_name"`
}

type StorageAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
