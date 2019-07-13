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

type AzurermIotDpsCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AzurermIotDpsCertificateSpec   `json:"spec,omitempty"`
	Status            AzurermIotDpsCertificateStatus `json:"status,omitempty"`
}

type AzurermIotDpsCertificateSpec struct {
	IotDpsName         string `json:"iot_dps_name"`
	CertificateContent string `json:"certificate_content"`
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
}



type AzurermIotDpsCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AzurermIotDpsCertificateList is a list of AzurermIotDpsCertificates
type AzurermIotDpsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AzurermIotDpsCertificate CRD objects
	Items []AzurermIotDpsCertificate `json:"items,omitempty"`
}