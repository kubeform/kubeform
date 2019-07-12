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

type AwsSsmDocument struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsSsmDocumentSpec   `json:"spec,omitempty"`
	Status            AwsSsmDocumentStatus `json:"status,omitempty"`
}

type AwsSsmDocumentSpecParameter struct {
	Name         string `json:"name"`
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
	Type         string `json:"type"`
}

type AwsSsmDocumentSpecPermissions struct {
	Type       string `json:"type"`
	AccountIds string `json:"account_ids"`
}

type AwsSsmDocumentSpec struct {
	Content        string               `json:"content"`
	Status         string               `json:"status"`
	Tags           map[string]string    `json:"tags"`
	Parameter      []AwsSsmDocumentSpec `json:"parameter"`
	Arn            string               `json:"arn"`
	Name           string               `json:"name"`
	Description    string               `json:"description"`
	Hash           string               `json:"hash"`
	HashType       string               `json:"hash_type"`
	LatestVersion  string               `json:"latest_version"`
	Owner          string               `json:"owner"`
	DocumentType   string               `json:"document_type"`
	SchemaVersion  string               `json:"schema_version"`
	DefaultVersion string               `json:"default_version"`
	PlatformTypes  []string             `json:"platform_types"`
	Permissions    map[string]string    `json:"permissions"`
	DocumentFormat string               `json:"document_format"`
	CreatedDate    string               `json:"created_date"`
}

type AwsSsmDocumentStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsSsmDocumentList is a list of AwsSsmDocuments
type AwsSsmDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmDocument CRD objects
	Items []AwsSsmDocument `json:"items,omitempty"`
}
