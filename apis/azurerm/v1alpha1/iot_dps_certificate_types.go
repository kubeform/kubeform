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

type IotDpsCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotDpsCertificateSpec   `json:"spec,omitempty"`
	Status            IotDpsCertificateStatus `json:"status,omitempty"`
}

type IotDpsCertificateSpec struct {
	CertificateContent string `json:"certificate_content"`
	IotDpsName         string `json:"iot_dps_name"`
	Name               string `json:"name"`
	ResourceGroupName  string `json:"resource_group_name"`
}

type IotDpsCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IotDpsCertificateList is a list of IotDpsCertificates
type IotDpsCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IotDpsCertificate CRD objects
	Items []IotDpsCertificate `json:"items,omitempty"`
}
