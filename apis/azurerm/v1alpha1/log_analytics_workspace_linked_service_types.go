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

type LogAnalyticsWorkspaceLinkedService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsWorkspaceLinkedServiceSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsWorkspaceLinkedServiceStatus `json:"status,omitempty"`
}

type LogAnalyticsWorkspaceLinkedServiceSpecLinkedServiceProperties struct {
	ResourceID string `json:"resourceID" tf:"resource_id"`
}

type LogAnalyticsWorkspaceLinkedServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	LinkedServiceName string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`
	// +optional
	LinkedServiceProperties []LogAnalyticsWorkspaceLinkedServiceSpecLinkedServiceProperties `json:"linkedServiceProperties,omitempty" tf:"linked_service_properties,omitempty"`
	ResourceGroupName       string                                                          `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceID string `json:"resourceID,omitempty" tf:"resource_id,omitempty"`
	// +optional
	Tags          map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	WorkspaceName string            `json:"workspaceName" tf:"workspace_name"`
}

type LogAnalyticsWorkspaceLinkedServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LogAnalyticsWorkspaceLinkedServiceSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogAnalyticsWorkspaceLinkedServiceList is a list of LogAnalyticsWorkspaceLinkedServices
type LogAnalyticsWorkspaceLinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogAnalyticsWorkspaceLinkedService CRD objects
	Items []LogAnalyticsWorkspaceLinkedService `json:"items,omitempty"`
}
