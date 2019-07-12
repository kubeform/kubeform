package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

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
	Description        string                         `json:"description"`
	Content            string                         `json:"content"`
	PublishContentLink []AzurermAutomationRunbookSpec `json:"publish_content_link"`
	Tags               map[string]string              `json:"tags"`
	AccountName        string                         `json:"account_name"`
	Location           string                         `json:"location"`
	ResourceGroupName  string                         `json:"resource_group_name"`
	LogProgress        bool                           `json:"log_progress"`
	Name               string                         `json:"name"`
	RunbookType        string                         `json:"runbook_type"`
	LogVerbose         bool                           `json:"log_verbose"`
}



type AzurermAutomationRunbookStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAutomationRunbookList is a list of AzurermAutomationRunbooks
type AzurermAutomationRunbookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAutomationRunbook CRD objects
	Items []AzurermAutomationRunbook `json:"items,omitempty"`
}