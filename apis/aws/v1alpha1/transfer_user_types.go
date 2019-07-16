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

type TransferUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TransferUserSpec   `json:"spec,omitempty"`
	Status            TransferUserStatus `json:"status,omitempty"`
}

type TransferUserSpec struct {
	// +optional
	HomeDirectory string `json:"home_directory,omitempty"`
	// +optional
	Policy   string `json:"policy,omitempty"`
	Role     string `json:"role"`
	ServerId string `json:"server_id"`
	// +optional
	Tags     map[string]string `json:"tags,omitempty"`
	UserName string            `json:"user_name"`
}

type TransferUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// TransferUserList is a list of TransferUsers
type TransferUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of TransferUser CRD objects
	Items []TransferUser `json:"items,omitempty"`
}
