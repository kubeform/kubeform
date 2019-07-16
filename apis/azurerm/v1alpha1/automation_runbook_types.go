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

type AutomationRunbook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutomationRunbookSpec   `json:"spec,omitempty"`
	Status            AutomationRunbookStatus `json:"status,omitempty"`
}

type AutomationRunbookSpecPublishContentLinkHash struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type AutomationRunbookSpecPublishContentLink struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Hash *[]AutomationRunbookSpecPublishContentLink `json:"hash,omitempty"`
	Uri  string                                     `json:"uri"`
	// +optional
	Version string `json:"version,omitempty"`
}

type AutomationRunbookSpec struct {
	AccountName string `json:"account_name"`
	// +optional
	Description string `json:"description,omitempty"`
	Location    string `json:"location"`
	LogProgress bool   `json:"log_progress"`
	LogVerbose  bool   `json:"log_verbose"`
	Name        string `json:"name"`
	// +kubebuilder:validation:MaxItems=1
	PublishContentLink []AutomationRunbookSpec `json:"publish_content_link"`
	ResourceGroupName  string                  `json:"resource_group_name"`
	RunbookType        string                  `json:"runbook_type"`
}

type AutomationRunbookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutomationRunbookList is a list of AutomationRunbooks
type AutomationRunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutomationRunbook CRD objects
	Items []AutomationRunbook `json:"items,omitempty"`
}
