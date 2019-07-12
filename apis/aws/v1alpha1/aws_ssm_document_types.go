package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	SchemaVersion  string               `json:"schema_version"`
	Hash           string               `json:"hash"`
	LatestVersion  string               `json:"latest_version"`
	Parameter      []AwsSsmDocumentSpec `json:"parameter"`
	Permissions    map[string]string    `json:"permissions"`
	Arn            string               `json:"arn"`
	DefaultVersion string               `json:"default_version"`
	HashType       string               `json:"hash_type"`
	Tags           map[string]string    `json:"tags"`
	Name           string               `json:"name"`
	Content        string               `json:"content"`
	PlatformTypes  []string             `json:"platform_types"`
	DocumentFormat string               `json:"document_format"`
	DocumentType   string               `json:"document_type"`
	CreatedDate    string               `json:"created_date"`
	Description    string               `json:"description"`
	Owner          string               `json:"owner"`
	Status         string               `json:"status"`
}



type AwsSsmDocumentStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsSsmDocumentList is a list of AwsSsmDocuments
type AwsSsmDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsSsmDocument CRD objects
	Items []AwsSsmDocument `json:"items,omitempty"`
}