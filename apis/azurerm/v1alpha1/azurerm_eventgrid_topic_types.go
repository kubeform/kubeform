package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AzurermEventgridTopic struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermEventgridTopicSpec   `json:"spec,omitempty"`
	Status            AzurermEventgridTopicStatus `json:"status,omitempty"`
}

type AzurermEventgridTopicSpec struct {
	Name               string            `json:"name"`
	Location           string            `json:"location"`
	ResourceGroupName  string            `json:"resource_group_name"`
	Tags               map[string]string `json:"tags"`
	Endpoint           string            `json:"endpoint"`
	PrimaryAccessKey   string            `json:"primary_access_key"`
	SecondaryAccessKey string            `json:"secondary_access_key"`
}

type AzurermEventgridTopicStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AzurermEventgridTopicList is a list of AzurermEventgridTopics
type AzurermEventgridTopicList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermEventgridTopic CRD objects
	Items []AzurermEventgridTopic `json:"items,omitempty"`
}
