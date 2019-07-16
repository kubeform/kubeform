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

type DataLakeStore struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeStoreSpec   `json:"spec,omitempty"`
	Status            DataLakeStoreStatus `json:"status,omitempty"`
}

type DataLakeStoreSpec struct {
	// +optional
	EncryptionState string `json:"encryption_state,omitempty"`
	// +optional
	FirewallAllowAzureIps string `json:"firewall_allow_azure_ips,omitempty"`
	// +optional
	FirewallState     string `json:"firewall_state,omitempty"`
	Location          string `json:"location"`
	Name              string `json:"name"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	Tier string `json:"tier,omitempty"`
}

type DataLakeStoreStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataLakeStoreList is a list of DataLakeStores
type DataLakeStoreList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeStore CRD objects
	Items []DataLakeStore `json:"items,omitempty"`
}
