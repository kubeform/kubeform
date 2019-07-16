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

type PublicIp struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PublicIpSpec   `json:"spec,omitempty"`
	Status            PublicIpStatus `json:"status,omitempty"`
}

type PublicIpSpec struct {
	// +optional
	DomainNameLabel string `json:"domain_name_label,omitempty"`
	// +optional
	IdleTimeoutInMinutes int `json:"idle_timeout_in_minutes,omitempty"`
	// +optional
	IpVersion string `json:"ip_version,omitempty"`
	Location  string `json:"location"`
	Name      string `json:"name"`
	// +optional
	PublicIpPrefixId  string `json:"public_ip_prefix_id,omitempty"`
	ResourceGroupName string `json:"resource_group_name"`
	// +optional
	ReverseFqdn string `json:"reverse_fqdn,omitempty"`
	// +optional
	Sku string `json:"sku,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty"`
}

type PublicIpStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PublicIpList is a list of PublicIps
type PublicIpList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PublicIp CRD objects
	Items []PublicIp `json:"items,omitempty"`
}
