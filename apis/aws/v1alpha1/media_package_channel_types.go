package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type MediaPackageChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MediaPackageChannelSpec   `json:"spec,omitempty"`
	Status            MediaPackageChannelStatus `json:"status,omitempty"`
}

type MediaPackageChannelSpecHlsIngestIngestEndpoints struct {
	Password string `json:"-" sensitive:"true" tf:"password"`
	Url      string `json:"url" tf:"url"`
	Username string `json:"username" tf:"username"`
}

type MediaPackageChannelSpecHlsIngest struct {
	IngestEndpoints []MediaPackageChannelSpecHlsIngestIngestEndpoints `json:"ingestEndpoints" tf:"ingest_endpoints"`
}

type MediaPackageChannelSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	KubeFormSecret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	Arn       string `json:"arn" tf:"arn"`
	ChannelID string `json:"channelID" tf:"channel_id"`
	// +optional
	Description string                             `json:"description,omitempty" tf:"description,omitempty"`
	HlsIngest   []MediaPackageChannelSpecHlsIngest `json:"hlsIngest" tf:"hls_ingest"`
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
	State *apis.State `json:"state,omitempty"`
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
