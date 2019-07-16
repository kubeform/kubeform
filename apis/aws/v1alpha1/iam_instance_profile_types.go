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

type IamInstanceProfile struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamInstanceProfileSpec   `json:"spec,omitempty"`
	Status            IamInstanceProfileStatus `json:"status,omitempty"`
}

type IamInstanceProfileSpec struct {
	// +optional
	NamePrefix string `json:"name_prefix,omitempty"`
	// +optional
	Path string `json:"path,omitempty"`
}

type IamInstanceProfileStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamInstanceProfileList is a list of IamInstanceProfiles
type IamInstanceProfileList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamInstanceProfile CRD objects
	Items []IamInstanceProfile `json:"items,omitempty"`
}
