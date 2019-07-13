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

type AwsMediaPackageChannel struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsMediaPackageChannelSpec   `json:"spec,omitempty"`
	Status            AwsMediaPackageChannelStatus `json:"status,omitempty"`
}

type AwsMediaPackageChannelSpecHlsIngestIngestEndpoints struct {
	Password string `json:"password"`
	Url      string `json:"url"`
	Username string `json:"username"`
}

type AwsMediaPackageChannelSpecHlsIngest struct {
	IngestEndpoints []AwsMediaPackageChannelSpecHlsIngest `json:"ingest_endpoints"`
}

type AwsMediaPackageChannelSpec struct {
	Arn         string                       `json:"arn"`
	ChannelId   string                       `json:"channel_id"`
	Description string                       `json:"description"`
	HlsIngest   []AwsMediaPackageChannelSpec `json:"hls_ingest"`
	Tags        map[string]string            `json:"tags"`
}



type AwsMediaPackageChannelStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsMediaPackageChannelList is a list of AwsMediaPackageChannels
type AwsMediaPackageChannelList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsMediaPackageChannel CRD objects
	Items []AwsMediaPackageChannel `json:"items,omitempty"`
}