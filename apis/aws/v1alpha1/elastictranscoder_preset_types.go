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

type ElastictranscoderPreset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElastictranscoderPresetSpec   `json:"spec,omitempty"`
	Status            ElastictranscoderPresetStatus `json:"status,omitempty"`
}

type ElastictranscoderPresetSpecAudio struct {
	// +optional
	AudioPackingMode string `json:"audio_packing_mode,omitempty"`
	// +optional
	BitRate string `json:"bit_rate,omitempty"`
	// +optional
	Channels string `json:"channels,omitempty"`
	// +optional
	Codec string `json:"codec,omitempty"`
	// +optional
	SampleRate string `json:"sample_rate,omitempty"`
}

type ElastictranscoderPresetSpecAudioCodecOptions struct {
	// +optional
	BitDepth string `json:"bit_depth,omitempty"`
	// +optional
	BitOrder string `json:"bit_order,omitempty"`
	// +optional
	Profile string `json:"profile,omitempty"`
	// +optional
	Signed string `json:"signed,omitempty"`
}

type ElastictranscoderPresetSpecThumbnails struct {
	// +optional
	AspectRatio string `json:"aspect_ratio,omitempty"`
	// +optional
	Format string `json:"format,omitempty"`
	// +optional
	Interval string `json:"interval,omitempty"`
	// +optional
	MaxHeight string `json:"max_height,omitempty"`
	// +optional
	MaxWidth string `json:"max_width,omitempty"`
	// +optional
	PaddingPolicy string `json:"padding_policy,omitempty"`
	// +optional
	Resolution string `json:"resolution,omitempty"`
	// +optional
	SizingPolicy string `json:"sizing_policy,omitempty"`
}

type ElastictranscoderPresetSpecVideo struct {
	// +optional
	AspectRatio string `json:"aspect_ratio,omitempty"`
	// +optional
	BitRate string `json:"bit_rate,omitempty"`
	// +optional
	Codec string `json:"codec,omitempty"`
	// +optional
	DisplayAspectRatio string `json:"display_aspect_ratio,omitempty"`
	// +optional
	FixedGop string `json:"fixed_gop,omitempty"`
	// +optional
	FrameRate string `json:"frame_rate,omitempty"`
	// +optional
	KeyframesMaxDist string `json:"keyframes_max_dist,omitempty"`
	// +optional
	MaxFrameRate string `json:"max_frame_rate,omitempty"`
	// +optional
	MaxHeight string `json:"max_height,omitempty"`
	// +optional
	MaxWidth string `json:"max_width,omitempty"`
	// +optional
	PaddingPolicy string `json:"padding_policy,omitempty"`
	// +optional
	Resolution string `json:"resolution,omitempty"`
	// +optional
	SizingPolicy string `json:"sizing_policy,omitempty"`
}

type ElastictranscoderPresetSpecVideoWatermarks struct {
	// +optional
	HorizontalAlign string `json:"horizontal_align,omitempty"`
	// +optional
	HorizontalOffset string `json:"horizontal_offset,omitempty"`
	// +optional
	Id string `json:"id,omitempty"`
	// +optional
	MaxHeight string `json:"max_height,omitempty"`
	// +optional
	MaxWidth string `json:"max_width,omitempty"`
	// +optional
	Opacity string `json:"opacity,omitempty"`
	// +optional
	SizingPolicy string `json:"sizing_policy,omitempty"`
	// +optional
	Target string `json:"target,omitempty"`
	// +optional
	VerticalAlign string `json:"vertical_align,omitempty"`
	// +optional
	VerticalOffset string `json:"vertical_offset,omitempty"`
}

type ElastictranscoderPresetSpec struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Audio *[]ElastictranscoderPresetSpec `json:"audio,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	AudioCodecOptions *[]ElastictranscoderPresetSpec `json:"audio_codec_options,omitempty"`
	Container         string                         `json:"container"`
	// +optional
	Description string `json:"description,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Thumbnails *[]ElastictranscoderPresetSpec `json:"thumbnails,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Video *[]ElastictranscoderPresetSpec `json:"video,omitempty"`
	// +optional
	VideoCodecOptions map[string]string `json:"video_codec_options,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VideoWatermarks *[]ElastictranscoderPresetSpec `json:"video_watermarks,omitempty"`
}

type ElastictranscoderPresetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ElastictranscoderPresetList is a list of ElastictranscoderPresets
type ElastictranscoderPresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ElastictranscoderPreset CRD objects
	Items []ElastictranscoderPreset `json:"items,omitempty"`
}
