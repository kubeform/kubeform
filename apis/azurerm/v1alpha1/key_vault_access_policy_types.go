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

type KeyVaultAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KeyVaultAccessPolicySpec   `json:"spec,omitempty"`
	Status            KeyVaultAccessPolicyStatus `json:"status,omitempty"`
}

type KeyVaultAccessPolicySpec struct {
	// +optional
	ApplicationId string `json:"application_id,omitempty"`
	// +optional
	CertificatePermissions []string `json:"certificate_permissions,omitempty"`
	// +optional
	KeyPermissions []string `json:"key_permissions,omitempty"`
	ObjectId       string   `json:"object_id"`
	// +optional
	SecretPermissions []string `json:"secret_permissions,omitempty"`
	// +optional
	StoragePermissions []string `json:"storage_permissions,omitempty"`
	TenantId           string   `json:"tenant_id"`
}

type KeyVaultAccessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KeyVaultAccessPolicyList is a list of KeyVaultAccessPolicys
type KeyVaultAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KeyVaultAccessPolicy CRD objects
	Items []KeyVaultAccessPolicy `json:"items,omitempty"`
}
