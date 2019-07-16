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

type DirectoryServiceDirectory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DirectoryServiceDirectorySpec   `json:"spec,omitempty"`
	Status            DirectoryServiceDirectoryStatus `json:"status,omitempty"`
}

type DirectoryServiceDirectorySpecConnectSettings struct {
	// +kubebuilder:validation:UniqueItems=true
	CustomerDnsIps   []string `json:"customer_dns_ips"`
	CustomerUsername string   `json:"customer_username"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
	VpcId     string   `json:"vpc_id"`
}

type DirectoryServiceDirectorySpecVpcSettings struct {
	// +kubebuilder:validation:UniqueItems=true
	SubnetIds []string `json:"subnet_ids"`
	VpcId     string   `json:"vpc_id"`
}

type DirectoryServiceDirectorySpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConnectSettings *[]DirectoryServiceDirectorySpec `json:"connect_settings,omitempty"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	EnableSso bool   `json:"enable_sso,omitempty"`
	Name      string `json:"name"`
	Password  string `json:"password"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcSettings *[]DirectoryServiceDirectorySpec `json:"vpc_settings,omitempty"`
}

type DirectoryServiceDirectoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DirectoryServiceDirectoryList is a list of DirectoryServiceDirectorys
type DirectoryServiceDirectoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DirectoryServiceDirectory CRD objects
	Items []DirectoryServiceDirectory `json:"items,omitempty"`
}
