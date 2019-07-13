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

type AzurermBatchCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermBatchCertificateSpec   `json:"spec,omitempty"`
	Status            AzurermBatchCertificateStatus `json:"status,omitempty"`
}

type AzurermBatchCertificateSpec struct {
	ResourceGroupName   string `json:"resource_group_name"`
	Password            string `json:"password"`
	Thumbprint          string `json:"thumbprint"`
	PublicData          string `json:"public_data"`
	Name                string `json:"name"`
	AccountName         string `json:"account_name"`
	Certificate         string `json:"certificate"`
	Format              string `json:"format"`
	ThumbprintAlgorithm string `json:"thumbprint_algorithm"`
}



type AzurermBatchCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermBatchCertificateList is a list of AzurermBatchCertificates
type AzurermBatchCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermBatchCertificate CRD objects
	Items []AzurermBatchCertificate `json:"items,omitempty"`
}