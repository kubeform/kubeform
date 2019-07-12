package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermLogAnalyticsLinkedService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermLogAnalyticsLinkedServiceSpec   `json:"spec,omitempty"`
	Status            AzurermLogAnalyticsLinkedServiceStatus `json:"status,omitempty"`
}

type AzurermLogAnalyticsLinkedServiceSpecLinkedServiceProperties struct {
	ResourceId string `json:"resource_id"`
}

type AzurermLogAnalyticsLinkedServiceSpec struct {
	LinkedServiceName       string                                 `json:"linked_service_name"`
	ResourceId              string                                 `json:"resource_id"`
	LinkedServiceProperties []AzurermLogAnalyticsLinkedServiceSpec `json:"linked_service_properties"`
	Name                    string                                 `json:"name"`
	Tags                    map[string]string                      `json:"tags"`
	ResourceGroupName       string                                 `json:"resource_group_name"`
	WorkspaceName           string                                 `json:"workspace_name"`
}



type AzurermLogAnalyticsLinkedServiceStatus struct {
    Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermLogAnalyticsLinkedServiceList is a list of AzurermLogAnalyticsLinkedServices
type AzurermLogAnalyticsLinkedServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermLogAnalyticsLinkedService CRD objects
	Items []AzurermLogAnalyticsLinkedService `json:"items,omitempty"`
}