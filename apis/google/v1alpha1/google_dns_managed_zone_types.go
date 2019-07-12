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

type GoogleDnsManagedZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleDnsManagedZoneSpec   `json:"spec,omitempty"`
	Status            GoogleDnsManagedZoneStatus `json:"status,omitempty"`
}

type GoogleDnsManagedZoneSpec struct {
	DnsName     string            `json:"dns_name"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	NameServers []string          `json:"name_servers"`
	Project     string            `json:"project"`
	Labels      map[string]string `json:"labels"`
}

type GoogleDnsManagedZoneStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleDnsManagedZoneList is a list of GoogleDnsManagedZones
type GoogleDnsManagedZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleDnsManagedZone CRD objects
	Items []GoogleDnsManagedZone `json:"items,omitempty"`
}
