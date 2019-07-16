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

type IothubSharedAccessPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IothubSharedAccessPolicySpec   `json:"spec,omitempty"`
	Status            IothubSharedAccessPolicyStatus `json:"status,omitempty"`
}

type IothubSharedAccessPolicySpec struct {
	// +optional
	DeviceConnect bool   `json:"device_connect,omitempty"`
	IothubName    string `json:"iothub_name"`
	Name          string `json:"name"`
	// +optional
	RegistryRead bool `json:"registry_read,omitempty"`
	// +optional
	RegistryWrite     bool   `json:"registry_write,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	ServiceConnect bool `json:"service_connect,omitempty"`
}

type IothubSharedAccessPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IothubSharedAccessPolicyList is a list of IothubSharedAccessPolicys
type IothubSharedAccessPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IothubSharedAccessPolicy CRD objects
	Items []IothubSharedAccessPolicy `json:"items,omitempty"`
}
