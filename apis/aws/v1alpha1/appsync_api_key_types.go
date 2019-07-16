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

type AppsyncApiKey struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncApiKeySpec   `json:"spec,omitempty"`
	Status            AppsyncApiKeyStatus `json:"status,omitempty"`
}

type AppsyncApiKeySpec struct {
	ApiId string `json:"api_id"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	Expires string `json:"expires,omitempty"`
}

type AppsyncApiKeyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncApiKeyList is a list of AppsyncApiKeys
type AppsyncApiKeyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncApiKey CRD objects
	Items []AppsyncApiKey `json:"items,omitempty"`
}
