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

type GlacierVault struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlacierVaultSpec   `json:"spec,omitempty"`
	Status            GlacierVaultStatus `json:"status,omitempty"`
}

type GlacierVaultSpecNotification struct {
	// +kubebuilder:validation:UniqueItems=true
	Events   []string `json:"events"`
	SnsTopic string   `json:"sns_topic"`
}

type GlacierVaultSpec struct {
	// +optional
	AccessPolicy string `json:"access_policy,omitempty"`
	Name         string `json:"name"`
	// +optional
	Notification *[]GlacierVaultSpec `json:"notification,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty"`
}

type GlacierVaultStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlacierVaultList is a list of GlacierVaults
type GlacierVaultList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlacierVault CRD objects
	Items []GlacierVault `json:"items,omitempty"`
}
