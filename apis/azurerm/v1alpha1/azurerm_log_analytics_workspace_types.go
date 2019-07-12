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

type AzurermLogAnalyticsWorkspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogAnalyticsWorkspaceSpec   `json:"spec,omitempty"`
	Status            AzurermLogAnalyticsWorkspaceStatus `json:"status,omitempty"`
}

type AzurermLogAnalyticsWorkspaceSpec struct {
	Name               string            `json:"name"`
	Location           string            `json:"location"`
	ResourceGroupName  string            `json:"resource_group_name"`
	RetentionInDays    int               `json:"retention_in_days"`
	WorkspaceId        string            `json:"workspace_id"`
	SecondarySharedKey string            `json:"secondary_shared_key"`
	Tags               map[string]string `json:"tags"`
	Sku                string            `json:"sku"`
	PortalUrl          string            `json:"portal_url"`
	PrimarySharedKey   string            `json:"primary_shared_key"`
}

type AzurermLogAnalyticsWorkspaceStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermLogAnalyticsWorkspaceList is a list of AzurermLogAnalyticsWorkspaces
type AzurermLogAnalyticsWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogAnalyticsWorkspace CRD objects
	Items []AzurermLogAnalyticsWorkspace `json:"items,omitempty"`
}
