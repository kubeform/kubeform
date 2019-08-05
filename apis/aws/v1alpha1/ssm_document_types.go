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

type SsmDocument struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmDocumentSpec   `json:"spec,omitempty"`
	Status            SsmDocumentStatus `json:"status,omitempty"`
}

type SsmDocumentSpecParameter struct {
	// +optional
	DefaultValue string `json:"defaultValue,omitempty" tf:"default_value,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
}

type SsmDocumentSpecPermissions struct {
	AccountIDS string `json:"accountIDS" tf:"account_ids"`
	Type       string `json:"type" tf:"type"`
}

type SsmDocumentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Arn            string `json:"arn" tf:"arn"`
	Content        string `json:"content" tf:"content"`
	CreatedDate    string `json:"createdDate" tf:"created_date"`
	DefaultVersion string `json:"defaultVersion" tf:"default_version"`
	Description    string `json:"description" tf:"description"`
	// +optional
	DocumentFormat string                     `json:"documentFormat,omitempty" tf:"document_format,omitempty"`
	DocumentType   string                     `json:"documentType" tf:"document_type"`
	Hash           string                     `json:"hash" tf:"hash"`
	HashType       string                     `json:"hashType" tf:"hash_type"`
	LatestVersion  string                     `json:"latestVersion" tf:"latest_version"`
	Name           string                     `json:"name" tf:"name"`
	Owner          string                     `json:"owner" tf:"owner"`
	Parameter      []SsmDocumentSpecParameter `json:"parameter" tf:"parameter"`
	// +optional
	Permissions   map[string]SsmDocumentSpecPermissions `json:"permissions,omitempty" tf:"permissions,omitempty"`
	PlatformTypes []string                              `json:"platformTypes" tf:"platform_types"`
	SchemaVersion string                                `json:"schemaVersion" tf:"schema_version"`
	Status        string                                `json:"status" tf:"status"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type SsmDocumentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SsmDocumentSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SsmDocumentList is a list of SsmDocuments
type SsmDocumentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SsmDocument CRD objects
	Items []SsmDocument `json:"items,omitempty"`
}
