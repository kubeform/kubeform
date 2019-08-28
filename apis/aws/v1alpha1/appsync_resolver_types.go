package v1alpha1

import (
    "encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	core "k8s.io/api/core/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type AppsyncResolver struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppsyncResolverSpec   `json:"spec,omitempty"`
	Status            AppsyncResolverStatus `json:"status,omitempty"`
}

type AppsyncResolverSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	ApiID string `json:"apiID" tf:"api_id"`
	// +optional
	Arn              string `json:"arn,omitempty" tf:"arn,omitempty"`
	DataSource       string `json:"dataSource" tf:"data_source"`
	Field            string `json:"field" tf:"field"`
	RequestTemplate  string `json:"requestTemplate" tf:"request_template"`
	ResponseTemplate string `json:"responseTemplate" tf:"response_template"`
	Type             string `json:"type" tf:"type"`
}



type AppsyncResolverStatus struct {
    // Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	// +optional
    Output *AppsyncResolverSpec `json:"output,omitempty"`
    // +optional
    State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppsyncResolverList is a list of AppsyncResolvers
type AppsyncResolverList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppsyncResolver CRD objects
	Items []AppsyncResolver `json:"items,omitempty"`
}