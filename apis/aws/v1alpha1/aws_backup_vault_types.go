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

type AwsBackupVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBackupVaultSpec   `json:"spec,omitempty"`
	Status            AwsBackupVaultStatus `json:"status,omitempty"`
}

type AwsBackupVaultSpec struct {
	Name           string            `json:"name"`
	Tags           map[string]string `json:"tags"`
	KmsKeyArn      string            `json:"kms_key_arn"`
	Arn            string            `json:"arn"`
	RecoveryPoints int               `json:"recovery_points"`
}



type AwsBackupVaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsBackupVaultList is a list of AwsBackupVaults
type AwsBackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBackupVault CRD objects
	Items []AwsBackupVault `json:"items,omitempty"`
}