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

type AwsElastictranscoderPresetSpecThumbnails struct {
	Format        string `json:"format"`
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
	SizingPolicy  string `json:"sizing_policy"`
	AspectRatio   string `json:"aspect_ratio"`
}

type AwsElastictranscoderPresetSpecVideoWatermarks struct {
	VerticalOffset   string `json:"vertical_offset"`
	SizingPolicy     string `json:"sizing_policy"`
	Target           string `json:"target"`
	HorizontalAlign  string `json:"horizontal_align"`
	HorizontalOffset string `json:"horizontal_offset"`
	Id               string `json:"id"`
	MaxHeight        string `json:"max_height"`
	MaxWidth         string `json:"max_width"`
	Opacity          string `json:"opacity"`
	VerticalAlign    string `json:"vertical_align"`
}

type AwsElastictranscoderPresetSpecAudio struct {
	SampleRate       string `json:"sample_rate"`
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
	Codec            string `json:"codec"`
}

type AwsElastictranscoderPresetSpecAudioCodecOptions struct {
	BitOrder string `json:"bit_order"`
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
	BitDepth string `json:"bit_depth"`
}

type AwsElastictranscoderPresetSpecVideo struct {
	MaxHeight          string `json:"max_height"`
	MaxWidth           string `json:"max_width"`
	BitRate            string `json:"bit_rate"`
	Codec              string `json:"codec"`
	FixedGop           string `json:"fixed_gop"`
	FrameRate          string `json:"frame_rate"`
	PaddingPolicy      string `json:"padding_policy"`
	Resolution         string `json:"resolution"`
	SizingPolicy       string `json:"sizing_policy"`
	AspectRatio        string `json:"aspect_ratio"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
	MaxFrameRate       string `json:"max_frame_rate"`
}

type AwsElastictranscoderPresetSpec struct {
	Arn               string                           `json:"arn"`
	Name              string                           `json:"name"`
	Thumbnails        []AwsElastictranscoderPresetSpec `json:"thumbnails"`
	VideoWatermarks   []AwsElastictranscoderPresetSpec `json:"video_watermarks"`
	VideoCodecOptions map[string]string                `json:"video_codec_options"`
	Audio             []AwsElastictranscoderPresetSpec `json:"audio"`
	AudioCodecOptions []AwsElastictranscoderPresetSpec `json:"audio_codec_options"`
	Container         string                           `json:"container"`
	Description       string                           `json:"description"`
	Type              string                           `json:"type"`
	Video             []AwsElastictranscoderPresetSpec `json:"video"`
}

type AwsElastictranscoderPresetStatus struct {
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
