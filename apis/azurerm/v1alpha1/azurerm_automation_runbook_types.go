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

type AzurermAutomationRunbook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAutomationRunbookSpec   `json:"spec,omitempty"`
	Status            AzurermAutomationRunbookStatus `json:"status,omitempty"`
}

type AzurermAutomationRunbookSpecPublishContentLinkHash struct {
	Algorithm string `json:"algorithm"`
	Value     string `json:"value"`
}

type AzurermAutomationRunbookSpecPublishContentLink struct {
	Uri     string                                           `json:"uri"`
	Version string                                           `json:"version"`
	Hash    []AzurermAutomationRunbookSpecPublishContentLink `json:"hash"`
}

type AzurermAutomationRunbookSpec struct {
	LogProgress        bool                           `json:"log_progress"`
	LogVerbose         bool                           `json:"log_verbose"`
	Description        string                         `json:"description"`
	ResourceGroupName  string                         `json:"resource_group_name"`
	RunbookType        string                         `json:"runbook_type"`
	Location           string                         `json:"location"`
	Content            string                         `json:"content"`
	PublishContentLink []AzurermAutomationRunbookSpec `json:"publish_content_link"`
	Tags               map[string]string              `json:"tags"`
	Name               string                         `json:"name"`
	AccountName        string                         `json:"account_name"`
}



type AzurermAutomationRunbookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermAutomationRunbookList is a list of AzurermAutomationRunbooks
type AzurermAutomationRunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationRunbook CRD objects
	Items []AzurermAutomationRunbook `json:"items,omitempty"`
}