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

type AzurermKeyVaultKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKeyVaultKeySpec   `json:"spec,omitempty"`
	Status            AzurermKeyVaultKeyStatus `json:"status,omitempty"`
}

type AzurermKeyVaultKeySpec struct {
	KeyOpts    []string          `json:"key_opts"`
	N          string            `json:"n"`
	X          string            `json:"x"`
	Tags       map[string]string `json:"tags"`
	Name       string            `json:"name"`
	KeyType    string            `json:"key_type"`
	KeySize    int               `json:"key_size"`
	Curve      string            `json:"curve"`
	Version    string            `json:"version"`
	E          string            `json:"e"`
	Y          string            `json:"y"`
	KeyVaultId string            `json:"key_vault_id"`
	VaultUri   string            `json:"vault_uri"`
}

type AzurermKeyVaultKeyStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermKeyVaultKeyList is a list of AzurermKeyVaultKeys
type AzurermKeyVaultKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKeyVaultKey CRD objects
	Items []AzurermKeyVaultKey `json:"items,omitempty"`
}
