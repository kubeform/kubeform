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

type ElastictranscoderPreset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ElastictranscoderPresetSpec   `json:"spec,omitempty"`
	Status            ElastictranscoderPresetStatus `json:"status,omitempty"`
}

type ElastictranscoderPresetSpecAudio struct {
	// +optional
	AudioPackingMode string `json:"audioPackingMode,omitempty" tf:"audio_packing_mode,omitempty"`
	// +optional
	BitRate string `json:"bitRate,omitempty" tf:"bit_rate,omitempty"`
	// +optional
	Channels string `json:"channels,omitempty" tf:"channels,omitempty"`
	// +optional
	Codec string `json:"codec,omitempty" tf:"codec,omitempty"`
	// +optional
	SampleRate string `json:"sampleRate,omitempty" tf:"sample_rate,omitempty"`
}

type ElastictranscoderPresetSpecAudioCodecOptions struct {
	// +optional
	BitDepth string `json:"bitDepth,omitempty" tf:"bit_depth,omitempty"`
	// +optional
	BitOrder string `json:"bitOrder,omitempty" tf:"bit_order,omitempty"`
	// +optional
	Profile string `json:"profile,omitempty" tf:"profile,omitempty"`
	// +optional
	Signed string `json:"signed,omitempty" tf:"signed,omitempty"`
}

type ElastictranscoderPresetSpecThumbnails struct {
	// +optional
	AspectRatio string `json:"aspectRatio,omitempty" tf:"aspect_ratio,omitempty"`
	// +optional
	Format string `json:"format,omitempty" tf:"format,omitempty"`
	// +optional
	Interval string `json:"interval,omitempty" tf:"interval,omitempty"`
	// +optional
	MaxHeight string `json:"maxHeight,omitempty" tf:"max_height,omitempty"`
	// +optional
	MaxWidth string `json:"maxWidth,omitempty" tf:"max_width,omitempty"`
	// +optional
	PaddingPolicy string `json:"paddingPolicy,omitempty" tf:"padding_policy,omitempty"`
	// +optional
	Resolution string `json:"resolution,omitempty" tf:"resolution,omitempty"`
	// +optional
	SizingPolicy string `json:"sizingPolicy,omitempty" tf:"sizing_policy,omitempty"`
}

type ElastictranscoderPresetSpecVideo struct {
	// +optional
	AspectRatio string `json:"aspectRatio,omitempty" tf:"aspect_ratio,omitempty"`
	// +optional
	BitRate string `json:"bitRate,omitempty" tf:"bit_rate,omitempty"`
	// +optional
	Codec string `json:"codec,omitempty" tf:"codec,omitempty"`
	// +optional
	DisplayAspectRatio string `json:"displayAspectRatio,omitempty" tf:"display_aspect_ratio,omitempty"`
	// +optional
	FixedGop string `json:"fixedGop,omitempty" tf:"fixed_gop,omitempty"`
	// +optional
	FrameRate string `json:"frameRate,omitempty" tf:"frame_rate,omitempty"`
	// +optional
	KeyframesMaxDist string `json:"keyframesMaxDist,omitempty" tf:"keyframes_max_dist,omitempty"`
	// +optional
	MaxFrameRate string `json:"maxFrameRate,omitempty" tf:"max_frame_rate,omitempty"`
	// +optional
	MaxHeight string `json:"maxHeight,omitempty" tf:"max_height,omitempty"`
	// +optional
	MaxWidth string `json:"maxWidth,omitempty" tf:"max_width,omitempty"`
	// +optional
	PaddingPolicy string `json:"paddingPolicy,omitempty" tf:"padding_policy,omitempty"`
	// +optional
	Resolution string `json:"resolution,omitempty" tf:"resolution,omitempty"`
	// +optional
	SizingPolicy string `json:"sizingPolicy,omitempty" tf:"sizing_policy,omitempty"`
}

type ElastictranscoderPresetSpecVideoWatermarks struct {
	// +optional
	HorizontalAlign string `json:"horizontalAlign,omitempty" tf:"horizontal_align,omitempty"`
	// +optional
	HorizontalOffset string `json:"horizontalOffset,omitempty" tf:"horizontal_offset,omitempty"`
	// +optional
	ID string `json:"ID,omitempty" tf:"id,omitempty"`
	// +optional
	MaxHeight string `json:"maxHeight,omitempty" tf:"max_height,omitempty"`
	// +optional
	MaxWidth string `json:"maxWidth,omitempty" tf:"max_width,omitempty"`
	// +optional
	Opacity string `json:"opacity,omitempty" tf:"opacity,omitempty"`
	// +optional
	SizingPolicy string `json:"sizingPolicy,omitempty" tf:"sizing_policy,omitempty"`
	// +optional
	Target string `json:"target,omitempty" tf:"target,omitempty"`
	// +optional
	VerticalAlign string `json:"verticalAlign,omitempty" tf:"vertical_align,omitempty"`
	// +optional
	VerticalOffset string `json:"verticalOffset,omitempty" tf:"vertical_offset,omitempty"`
}

type ElastictranscoderPresetSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Audio []ElastictranscoderPresetSpecAudio `json:"audio,omitempty" tf:"audio,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	AudioCodecOptions []ElastictranscoderPresetSpecAudioCodecOptions `json:"audioCodecOptions,omitempty" tf:"audio_codec_options,omitempty"`
	Container         string                                         `json:"container" tf:"container"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +optional
	Name string `json:"name,omitempty" tf:"name,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Thumbnails []ElastictranscoderPresetSpecThumbnails `json:"thumbnails,omitempty" tf:"thumbnails,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	// +kubebuilder:validation:UniqueItems=true
	Video []ElastictranscoderPresetSpecVideo `json:"video,omitempty" tf:"video,omitempty"`
	// +optional
	VideoCodecOptions map[string]string `json:"videoCodecOptions,omitempty" tf:"video_codec_options,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	VideoWatermarks []ElastictranscoderPresetSpecVideoWatermarks `json:"videoWatermarks,omitempty" tf:"video_watermarks,omitempty"`
}

type ElastictranscoderPresetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ElastictranscoderPresetSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
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
