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

type IotDpsCertificate struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IotDpsCertificateSpec   `json:"spec,omitempty"`
	Status            IotDpsCertificateStatus `json:"status,omitempty"`
}

type IotDpsCertificateSpec struct {
	// Sensitive Data. Provide secret name which contains one value only
	CertificateContent *core.LocalObjectReference `json:"certificateContent" tf:"certificate_content"`
	IotDpsName         string                     `json:"iotDpsName" tf:"iot_dps_name"`
	Name               string                     `json:"name" tf:"name"`
	ResourceGroupName  string                     `json:"resourceGroupName" tf:"resource_group_name"`
	ProviderRef        core.LocalObjectReference  `json:"providerRef" tf:"-"`
}

type IotDpsCertificateStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
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
