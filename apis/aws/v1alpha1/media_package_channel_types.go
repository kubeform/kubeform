package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type MediaPackageChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaPackageChannelSpec   `json:"spec,omitempty"`
	Status            MediaPackageChannelStatus `json:"status,omitempty"`
}

type MediaPackageChannelSpecHlsIngestIngestEndpoints struct {
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Url string `json:"url,omitempty" tf:"url,omitempty"`
	// +optional
	Username string `json:"username,omitempty" tf:"username,omitempty"`
}

type MediaPackageChannelSpecHlsIngest struct {
	// +optional
	IngestEndpoints []MediaPackageChannelSpecHlsIngestIngestEndpoints `json:"ingestEndpoints,omitempty" tf:"ingest_endpoints,omitempty"`
}

type MediaPackageChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Arn       string `json:"arn,omitempty" tf:"arn,omitempty"`
	ChannelID string `json:"channelID" tf:"channel_id"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	HlsIngest []MediaPackageChannelSpecHlsIngest `json:"hlsIngest,omitempty" tf:"hls_ingest,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type MediaPackageChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MediaPackageChannelSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MediaPackageChannelList is a list of MediaPackageChannels
type MediaPackageChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MediaPackageChannel CRD objects
	Items []MediaPackageChannel `json:"items,omitempty"`
}
