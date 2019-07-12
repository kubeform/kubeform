package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermAppServiceCustomHostnameBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermAppServiceCustomHostnameBindingSpec   `json:"spec,omitempty"`
	Status            AzurermAppServiceCustomHostnameBindingStatus `json:"status,omitempty"`
}

type AzurermAppServiceCustomHostnameBindingSpec struct {
	Hostname          string `json:"hostname"`
	ResourceGroupName string `json:"resource_group_name"`
	AppServiceName    string `json:"app_service_name"`
}



type AzurermAppServiceCustomHostnameBindingStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermAppServiceCustomHostnameBindingList is a list of AzurermAppServiceCustomHostnameBindings
type AzurermAppServiceCustomHostnameBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermAppServiceCustomHostnameBinding CRD objects
	Items []AzurermAppServiceCustomHostnameBinding `json:"items,omitempty"`
}