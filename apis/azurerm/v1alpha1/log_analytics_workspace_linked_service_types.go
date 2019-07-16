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

type LogAnalyticsWorkspaceLinkedService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsWorkspaceLinkedServiceSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsWorkspaceLinkedServiceStatus `json:"status,omitempty"`
}

type LogAnalyticsWorkspaceLinkedServiceSpec struct {
	// +optional
	LinkedServiceName string `json:"linked_service_name,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	WorkspaceName     string `json:"workspace_name"`
}

type LogAnalyticsWorkspaceLinkedServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
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
