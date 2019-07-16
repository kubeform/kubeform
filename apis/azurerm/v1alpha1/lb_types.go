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

type Lb struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbSpec   `json:"spec,omitempty"`
	Status            LbStatus `json:"status,omitempty"`
}

type LbSpecFrontendIpConfiguration struct {
	Name string `json:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Zones []string `json:"zones,omitempty"`
}

type LbSpec struct {
	// +optional
	// +kubebuilder:validation:MinItems=1
	FrontendIpConfiguration *[]LbSpec `json:"frontend_ip_configuration,omitempty"`
	Location                string    `json:"location"`
	Name                    string    `json:"name"`
	ResourceGroupName       string    `json:"resource_group_name"`
	// +optional
	Sku string `json:"sku,omitempty"`
}

type LbStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbList is a list of Lbs
type LbList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Lb CRD objects
	Items []Lb `json:"items,omitempty"`
}
