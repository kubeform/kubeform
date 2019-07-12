package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsBackupVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsBackupVaultSpec   `json:"spec,omitempty"`
	Status            AwsBackupVaultStatus `json:"status,omitempty"`
}

type AwsBackupVaultSpec struct {
	RecoveryPoints int               `json:"recovery_points"`
	Name           string            `json:"name"`
	Tags           map[string]string `json:"tags"`
	KmsKeyArn      string            `json:"kms_key_arn"`
	Arn            string            `json:"arn"`
}



type AwsBackupVaultStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsBackupVaultList is a list of AwsBackupVaults
type AwsBackupVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsBackupVault CRD objects
	Items []AwsBackupVault `json:"items,omitempty"`
}