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

type DataLakeAnalyticsAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeAnalyticsAccountSpec   `json:"spec,omitempty"`
	Status            DataLakeAnalyticsAccountStatus `json:"status,omitempty"`
}

type DataLakeAnalyticsAccountSpec struct {
	DefaultStoreAccountName string `json:"default_store_account_name"`
	Location                string `json:"location"`
	Name                    string `json:"name"`
	ResourceGroupName       string `json:"resource_group_name"`
	// +optional
	Tier string `json:"tier,omitempty"`
}

type DataLakeAnalyticsAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataLakeAnalyticsAccountList is a list of DataLakeAnalyticsAccounts
type DataLakeAnalyticsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeAnalyticsAccount CRD objects
	Items []DataLakeAnalyticsAccount `json:"items,omitempty"`
}
