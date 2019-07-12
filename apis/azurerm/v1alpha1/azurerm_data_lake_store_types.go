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

type AzurermDataLakeStore struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermDataLakeStoreSpec   `json:"spec,omitempty"`
	Status            AzurermDataLakeStoreStatus `json:"status,omitempty"`
}

type AzurermDataLakeStoreSpec struct {
	FirewallState         string            `json:"firewall_state"`
	FirewallAllowAzureIps string            `json:"firewall_allow_azure_ips"`
	Tags                  map[string]string `json:"tags"`
	EncryptionType        string            `json:"encryption_type"`
	Location              string            `json:"location"`
	ResourceGroupName     string            `json:"resource_group_name"`
	Tier                  string            `json:"tier"`
	EncryptionState       string            `json:"encryption_state"`
	Endpoint              string            `json:"endpoint"`
	Name                  string            `json:"name"`
}

type AzurermDataLakeStoreStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermDataLakeStoreList is a list of AzurermDataLakeStores
type AzurermDataLakeStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermDataLakeStore CRD objects
	Items []AzurermDataLakeStore `json:"items,omitempty"`
}
