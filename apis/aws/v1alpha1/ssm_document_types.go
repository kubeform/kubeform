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

type SsmDocument struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SsmDocumentSpec   `json:"spec,omitempty"`
	Status            SsmDocumentStatus `json:"status,omitempty"`
}

type SsmDocumentSpecPermissions struct {
	AccountIds string `json:"account_ids"`
	Type       string `json:"type"`
}

type SsmDocumentSpec struct {
	Content string `json:"content"`
	// +optional
	DocumentFormat string `json:"document_format,omitempty"`
	DocumentType   string `json:"document_type"`
	Name           string `json:"name"`
	// +optional
	Permissions map[string]SsmDocumentSpecPermissions `json:"permissions,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type SsmDocumentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
