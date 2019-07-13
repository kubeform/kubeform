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

type AzurermNetworkWatcher struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermNetworkWatcherSpec   `json:"spec,omitempty"`
	Status            AzurermNetworkWatcherStatus `json:"status,omitempty"`
}

type AzurermNetworkWatcherSpec struct {
	Name              string            `json:"name"`
	ResourceGroupName string            `json:"resource_group_name"`
	Location          string            `json:"location"`
	Tags              map[string]string `json:"tags"`
}



type AzurermNetworkWatcherStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermNetworkWatcherList is a list of AzurermNetworkWatchers
type AzurermNetworkWatcherList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermNetworkWatcher CRD objects
	Items []AzurermNetworkWatcher `json:"items,omitempty"`
}