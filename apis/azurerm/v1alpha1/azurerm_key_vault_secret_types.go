package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermKeyVaultSecret struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermKeyVaultSecretSpec   `json:"spec,omitempty"`
	Status            AzurermKeyVaultSecretStatus `json:"status,omitempty"`
}

type AzurermKeyVaultSecretSpec struct {
	Tags        map[string]string `json:"tags"`
	Name        string            `json:"name"`
	KeyVaultId  string            `json:"key_vault_id"`
	VaultUri    string            `json:"vault_uri"`
	Value       string            `json:"value"`
	ContentType string            `json:"content_type"`
	Version     string            `json:"version"`
}



type AzurermKeyVaultSecretStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermKeyVaultSecretList is a list of AzurermKeyVaultSecrets
type AzurermKeyVaultSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermKeyVaultSecret CRD objects
	Items []AzurermKeyVaultSecret `json:"items,omitempty"`
}