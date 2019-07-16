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

type NotificationHubNamespace struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              NotificationHubNamespaceSpec   `json:"spec,omitempty"`
	Status            NotificationHubNamespaceStatus `json:"status,omitempty"`
}

type NotificationHubNamespaceSpec struct {
	// +optional
	Enabled           bool   `json:"enabled,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	NamespaceType     string `json:"namespace_type"`
	ResourceGroupName string `json:"resource_group_name"`
}

type NotificationHubNamespaceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// NotificationHubNamespaceList is a list of NotificationHubNamespaces
type NotificationHubNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of NotificationHubNamespace CRD objects
	Items []NotificationHubNamespace `json:"items,omitempty"`
}
