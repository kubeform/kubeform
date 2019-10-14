package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type LogAnalyticsLinkedService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LogAnalyticsLinkedServiceSpec   `json:"spec,omitempty"`
	Status            LogAnalyticsLinkedServiceStatus `json:"status,omitempty"`
}

type LogAnalyticsLinkedServiceSpecLinkedServiceProperties struct {
	ResourceID string `json:"resourceID" tf:"resource_id"`
}

type LogAnalyticsLinkedServiceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	LinkedServiceName string `json:"linkedServiceName,omitempty" tf:"linked_service_name,omitempty"`
	// +optional
	LinkedServiceProperties []LogAnalyticsLinkedServiceSpecLinkedServiceProperties `json:"linkedServiceProperties,omitempty" tf:"linked_service_properties,omitempty"`
	// +optional
	Name              string `json:"name,omitempty" tf:"name,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	ResourceID string `json:"resourceID,omitempty" tf:"resource_id,omitempty"`
	// +optional
	Tags          map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	WorkspaceName string            `json:"workspaceName" tf:"workspace_name"`
}

type LogAnalyticsLinkedServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LogAnalyticsLinkedServiceSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LogAnalyticsLinkedServiceList is a list of LogAnalyticsLinkedServices
type LogAnalyticsLinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LogAnalyticsLinkedService CRD objects
	Items []LogAnalyticsLinkedService `json:"items,omitempty"`
}
