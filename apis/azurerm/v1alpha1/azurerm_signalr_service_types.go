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

type AzurermSignalrService struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermSignalrServiceSpec   `json:"spec,omitempty"`
	Status            AzurermSignalrServiceStatus `json:"status,omitempty"`
}

type AzurermSignalrServiceSpecSku struct {
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}

type AzurermSignalrServiceSpec struct {
	ServerPort                int                         `json:"server_port"`
	PrimaryAccessKey          string                      `json:"primary_access_key"`
	PrimaryConnectionString   string                      `json:"primary_connection_string"`
	SecondaryAccessKey        string                      `json:"secondary_access_key"`
	SecondaryConnectionString string                      `json:"secondary_connection_string"`
	Name                      string                      `json:"name"`
	ResourceGroupName         string                      `json:"resource_group_name"`
	Sku                       []AzurermSignalrServiceSpec `json:"sku"`
	Hostname                  string                      `json:"hostname"`
	IpAddress                 string                      `json:"ip_address"`
	PublicPort                int                         `json:"public_port"`
	Tags                      map[string]string           `json:"tags"`
	Location                  string                      `json:"location"`
}



type AzurermSignalrServiceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermSignalrServiceList is a list of AzurermSignalrServices
type AzurermSignalrServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermSignalrService CRD objects
	Items []AzurermSignalrService `json:"items,omitempty"`
}