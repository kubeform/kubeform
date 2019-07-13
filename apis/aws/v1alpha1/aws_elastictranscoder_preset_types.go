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

type AwsElastictranscoderPreset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElastictranscoderPresetSpec   `json:"spec,omitempty"`
	Status            AwsElastictranscoderPresetStatus `json:"status,omitempty"`
}

type AwsElastictranscoderPresetSpecVideo struct {
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
	MaxFrameRate       string `json:"max_frame_rate"`
	MaxHeight          string `json:"max_height"`
	AspectRatio        string `json:"aspect_ratio"`
	BitRate            string `json:"bit_rate"`
	Codec              string `json:"codec"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	FrameRate          string `json:"frame_rate"`
	MaxWidth           string `json:"max_width"`
	PaddingPolicy      string `json:"padding_policy"`
	FixedGop           string `json:"fixed_gop"`
	Resolution         string `json:"resolution"`
	SizingPolicy       string `json:"sizing_policy"`
}

type AwsElastictranscoderPresetSpecVideoWatermarks struct {
	Target           string `json:"target"`
	VerticalAlign    string `json:"vertical_align"`
	HorizontalOffset string `json:"horizontal_offset"`
	MaxHeight        string `json:"max_height"`
	SizingPolicy     string `json:"sizing_policy"`
	Opacity          string `json:"opacity"`
	VerticalOffset   string `json:"vertical_offset"`
	HorizontalAlign  string `json:"horizontal_align"`
	Id               string `json:"id"`
	MaxWidth         string `json:"max_width"`
}

type AwsElastictranscoderPresetSpecThumbnails struct {
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
	SizingPolicy  string `json:"sizing_policy"`
	AspectRatio   string `json:"aspect_ratio"`
	Format        string `json:"format"`
}

type AwsElastictranscoderPresetSpecAudio struct {
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
	Codec            string `json:"codec"`
	SampleRate       string `json:"sample_rate"`
}

type AwsElastictranscoderPresetSpecAudioCodecOptions struct {
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
	BitDepth string `json:"bit_depth"`
	BitOrder string `json:"bit_order"`
}

type AwsElastictranscoderPresetSpec struct {
	Video             []AwsElastictranscoderPresetSpec `json:"video"`
	VideoWatermarks   []AwsElastictranscoderPresetSpec `json:"video_watermarks"`
	VideoCodecOptions map[string]string                `json:"video_codec_options"`
	Container         string                           `json:"container"`
	Thumbnails        []AwsElastictranscoderPresetSpec `json:"thumbnails"`
	Type              string                           `json:"type"`
	Description       string                           `json:"description"`
	Name              string                           `json:"name"`
	Arn               string                           `json:"arn"`
	Audio             []AwsElastictranscoderPresetSpec `json:"audio"`
	AudioCodecOptions []AwsElastictranscoderPresetSpec `json:"audio_codec_options"`
}



type AwsElastictranscoderPresetStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AwsElastictranscoderPresetList is a list of AwsElastictranscoderPresets
type AwsElastictranscoderPresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElastictranscoderPreset CRD objects
	Items []AwsElastictranscoderPreset `json:"items,omitempty"`
}