package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type AwsElastictranscoderPreset struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AwsElastictranscoderPresetSpec   `json:"spec,omitempty"`
	Status            AwsElastictranscoderPresetStatus `json:"status,omitempty"`
}

type AwsElastictranscoderPresetSpecVideo struct {
	MaxWidth           string `json:"max_width"`
	SizingPolicy       string `json:"sizing_policy"`
	Codec              string `json:"codec"`
	DisplayAspectRatio string `json:"display_aspect_ratio"`
	FixedGop           string `json:"fixed_gop"`
	FrameRate          string `json:"frame_rate"`
	KeyframesMaxDist   string `json:"keyframes_max_dist"`
	MaxFrameRate       string `json:"max_frame_rate"`
	MaxHeight          string `json:"max_height"`
	PaddingPolicy      string `json:"padding_policy"`
	AspectRatio        string `json:"aspect_ratio"`
	BitRate            string `json:"bit_rate"`
	Resolution         string `json:"resolution"`
}

type AwsElastictranscoderPresetSpecAudio struct {
	AudioPackingMode string `json:"audio_packing_mode"`
	BitRate          string `json:"bit_rate"`
	Channels         string `json:"channels"`
	Codec            string `json:"codec"`
	SampleRate       string `json:"sample_rate"`
}

type AwsElastictranscoderPresetSpecAudioCodecOptions struct {
	BitDepth string `json:"bit_depth"`
	BitOrder string `json:"bit_order"`
	Profile  string `json:"profile"`
	Signed   string `json:"signed"`
}

type AwsElastictranscoderPresetSpecVideoWatermarks struct {
	Id               string `json:"id"`
	MaxHeight        string `json:"max_height"`
	MaxWidth         string `json:"max_width"`
	Opacity          string `json:"opacity"`
	Target           string `json:"target"`
	VerticalAlign    string `json:"vertical_align"`
	HorizontalAlign  string `json:"horizontal_align"`
	HorizontalOffset string `json:"horizontal_offset"`
	SizingPolicy     string `json:"sizing_policy"`
	VerticalOffset   string `json:"vertical_offset"`
}

type AwsElastictranscoderPresetSpecThumbnails struct {
	PaddingPolicy string `json:"padding_policy"`
	Resolution    string `json:"resolution"`
	SizingPolicy  string `json:"sizing_policy"`
	AspectRatio   string `json:"aspect_ratio"`
	Format        string `json:"format"`
	Interval      string `json:"interval"`
	MaxHeight     string `json:"max_height"`
	MaxWidth      string `json:"max_width"`
}

type AwsElastictranscoderPresetSpec struct {
	Video             []AwsElastictranscoderPresetSpec `json:"video"`
	Audio             []AwsElastictranscoderPresetSpec `json:"audio"`
	AudioCodecOptions []AwsElastictranscoderPresetSpec `json:"audio_codec_options"`
	Description       string                           `json:"description"`
	Name              string                           `json:"name"`
	VideoWatermarks   []AwsElastictranscoderPresetSpec `json:"video_watermarks"`
	VideoCodecOptions map[string]string                `json:"video_codec_options"`
	Arn               string                           `json:"arn"`
	Container         string                           `json:"container"`
	Thumbnails        []AwsElastictranscoderPresetSpec `json:"thumbnails"`
	Type              string                           `json:"type"`
}

type AwsElastictranscoderPresetStatus struct {
	Output *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AwsElastictranscoderPresetList is a list of AwsElastictranscoderPresets
type AwsElastictranscoderPresetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AwsElastictranscoderPreset CRD objects
	Items []AwsElastictranscoderPreset `json:"items,omitempty"`
}
