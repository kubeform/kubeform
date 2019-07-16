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

type ApiManagement struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiManagementSpec   `json:"spec,omitempty"`
	Status            ApiManagementStatus `json:"status,omitempty"`
}

type ApiManagementSpecAdditionalLocation struct {
	Location string `json:"location"`
}

type ApiManagementSpecCertificate struct {
	CertificatePassword string `json:"certificate_password"`
	EncodedCertificate  string `json:"encoded_certificate"`
	StoreName           string `json:"store_name"`
}

type ApiManagementSpecIdentity struct {
	Type string `json:"type"`
}

type ApiManagementSpecSku struct {
	Capacity int    `json:"capacity"`
	Name     string `json:"name"`
}

type ApiManagementSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	AdditionalLocation *[]ApiManagementSpec `json:"additional_location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Certificate *[]ApiManagementSpec `json:"certificate,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity          *[]ApiManagementSpec `json:"identity,omitempty"`
	Location          string               `json:"location"`
	Name              string               `json:"name"`
	PublisherEmail    string               `json:"publisher_email"`
	PublisherName     string               `json:"publisher_name"`
	ResourceGroupName string               `json:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Sku []ApiManagementSpec `json:"sku"`
}

type ApiManagementStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiManagementList is a list of ApiManagements
type ApiManagementList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiManagement CRD objects
	Items []ApiManagement `json:"items,omitempty"`
}
