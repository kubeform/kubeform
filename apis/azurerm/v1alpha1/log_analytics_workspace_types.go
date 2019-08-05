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

type LogAnalyticsWorkspace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsWorkspaceSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsWorkspaceStatus `json:"status,omitempty"`
}

type LogAnalyticsWorkspaceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Location          string `json:"location" tf:"location"`
	Name              string `json:"name" tf:"name"`
	PortalURL         string `json:"portalURL" tf:"portal_url"`
	PrimarySharedKey  string `json:"-" sensitive:"true" tf:"primary_shared_key"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	RetentionInDays    int    `json:"retentionInDays,omitempty" tf:"retention_in_days,omitempty"`
	SecondarySharedKey string `json:"-" sensitive:"true" tf:"secondary_shared_key"`
	Sku                string `json:"sku" tf:"sku"`
	// +optional
	Tags        map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	WorkspaceID string            `json:"workspaceID" tf:"workspace_id"`
}

type LogAnalyticsWorkspaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LogAnalyticsWorkspaceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogAnalyticsWorkspaceList is a list of LogAnalyticsWorkspaces
type LogAnalyticsWorkspaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogAnalyticsWorkspace CRD objects
	Items []LogAnalyticsWorkspace `json:"items,omitempty"`
}
