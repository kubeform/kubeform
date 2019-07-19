package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
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
	CustomerDNSIPS   []string `json:"customerDNSIPS" tf:"customer_dns_ips"`
	CustomerUsername string   `json:"customerUsername" tf:"customer_username"`
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	VpcID     string   `json:"vpcID" tf:"vpc_id"`
}

type DirectoryServiceDirectorySpecVpcSettings struct {
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	VpcID     string   `json:"vpcID" tf:"vpc_id"`
}

type DirectoryServiceDirectorySpec struct {
	// +optional
	Alias string `json:"alias,omitempty" tf:"alias,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConnectSettings []DirectoryServiceDirectorySpecConnectSettings `json:"connectSettings,omitempty" tf:"connect_settings,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Edition string `json:"edition,omitempty" tf:"edition,omitempty"`
	// +optional
	EnableSso bool   `json:"enableSso,omitempty" tf:"enable_sso,omitempty"`
	Name      string `json:"name" tf:"name"`
	// Sensitive Data. Provide secret name which contains one value only
	Password core.LocalObjectReference `json:"password" tf:"password"`
	// +optional
	ShortName string `json:"shortName,omitempty" tf:"short_name,omitempty"`
	// +optional
	Size string `json:"size,omitempty" tf:"size,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VpcSettings []DirectoryServiceDirectorySpecVpcSettings `json:"vpcSettings,omitempty" tf:"vpc_settings,omitempty"`
	ProviderRef core.LocalObjectReference                  `json:"providerRef" tf:"-"`
}

type DirectoryServiceDirectoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     *runtime.RawExtension `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
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
