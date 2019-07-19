package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
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

type SsmDocumentSpecPermissions struct {
	AccountIDS string `json:"accountIDS" tf:"account_ids"`
	Type       string `json:"type" tf:"type"`
}

type SsmDocumentSpec struct {
	Content string `json:"content" tf:"content"`
	// +optional
	DocumentFormat string `json:"documentFormat,omitempty" tf:"document_format,omitempty"`
	DocumentType   string `json:"documentType" tf:"document_type"`
	Name           string `json:"name" tf:"name"`
	// +optional
	Permissions map[string]SsmDocumentSpecPermissions `json:"permissions,omitempty" tf:"permissions,omitempty"`
	// +optional
	Tags        map[string]string         `json:"tags,omitempty" tf:"tags,omitempty"`
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type SsmDocumentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
