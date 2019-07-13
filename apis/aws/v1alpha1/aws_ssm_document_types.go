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
	DefaultValue string `json:"default_value"`
	Description  string `json:"description"`
	Type         string `json:"type"`
	Name         string `json:"name"`
}

type AwsSsmDocumentSpecPermissions struct {
	Type       string `json:"type"`
	AccountIds string `json:"account_ids"`
}

type AwsSsmDocumentSpec struct {
	Arn            string               `json:"arn"`
	SchemaVersion  string               `json:"schema_version"`
	HashType       string               `json:"hash_type"`
	Owner          string               `json:"owner"`
	DocumentType   string               `json:"document_type"`
	CreatedDate    string               `json:"created_date"`
	PlatformTypes  []string             `json:"platform_types"`
	Parameter      []AwsSsmDocumentSpec `json:"parameter"`
	Permissions    map[string]string    `json:"permissions"`
	Name           string               `json:"name"`
	DocumentFormat string               `json:"document_format"`
	Status         string               `json:"status"`
	Content        string               `json:"content"`
	DefaultVersion string               `json:"default_version"`
	Description    string               `json:"description"`
	Hash           string               `json:"hash"`
	LatestVersion  string               `json:"latest_version"`
	Tags           map[string]string    `json:"tags"`
}



type AwsSsmDocumentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

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