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

type AppServiceCustomHostnameBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceCustomHostnameBindingSpec   `json:"spec,omitempty"`
	Status            AppServiceCustomHostnameBindingStatus `json:"status,omitempty"`
}

type AppServiceCustomHostnameBindingSpec struct {
	AppServiceName    string `json:"app_service_name"`
	Hostname          string `json:"hostname"`
	ResourceGroupName string `json:"resource_group_name"`
}

type AppServiceCustomHostnameBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceCustomHostnameBindingList is a list of AppServiceCustomHostnameBindings
type AppServiceCustomHostnameBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceCustomHostnameBinding CRD objects
	Items []AppServiceCustomHostnameBinding `json:"items,omitempty"`
}
