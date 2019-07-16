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

type BatchCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              BatchCertificateSpec   `json:"spec,omitempty"`
	Status            BatchCertificateStatus `json:"status,omitempty"`
}

type BatchCertificateSpec struct {
	AccountName string `json:"account_name"`
	Certificate string `json:"certificate"`
	Format      string `json:"format"`
	// +optional
	Password            string `json:"password,omitempty"`
	ResourceGroupName   string `json:"resource_group_name"`
	Thumbprint          string `json:"thumbprint"`
	ThumbprintAlgorithm string `json:"thumbprint_algorithm"`
}

type BatchCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// BatchCertificateList is a list of BatchCertificates
type BatchCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of BatchCertificate CRD objects
	Items []BatchCertificate `json:"items,omitempty"`
}
