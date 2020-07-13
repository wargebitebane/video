package codec

import (
	"github.com/wargebitebane/video/av"
	"github.com/wargebitebane/video/codec/fake"
	// "log"
	"time"
)

type PCMUCodecData struct {
	typ av.CodecType
}

func (self PCMUCodecData) Type() av.CodecType {
	return self.typ
}

func (self PCMUCodecData) SampleRate() int {
	if self.typ.String() == "PCMU" {
		return 8000
	} else {
		return 48000
	}

}

func (self PCMUCodecData) ChannelLayout() av.ChannelLayout {
	return av.CH_MONO
}

func (self PCMUCodecData) SampleFormat() av.SampleFormat {
	if self.typ.String() == "PCMU" {
		return av.S16
	} else {
		return av.S32
	}
}

func (self PCMUCodecData) PacketDuration(data []byte) (time.Duration, error) {
	return time.Duration(len(data)) * time.Second / time.Duration(8000), nil
}

func NewAACCodecData() av.AudioCodecData {
	return PCMUCodecData{
		typ: av.AAC,
	}
}

// func NewOPUSCodecData() av.AudioCodecData {
// 	return PCMUCodecData{
// 		typ: av.OPUS,
// 	}
// }

func NewPCMMulawCodecData() av.AudioCodecData {
	return PCMUCodecData{
		typ: av.PCM_MULAW,
	}
}

func NewPCMAlawCodecData() av.AudioCodecData {
	return PCMUCodecData{
		typ: av.PCM_ALAW,
	}
}

type SpeexCodecData struct {
	fake.CodecData
}

func (self SpeexCodecData) PacketDuration(data []byte) (time.Duration, error) {
	// libavcodec/libspeexdec.c
	// samples = samplerate/50
	// duration = 0.02s
	return time.Millisecond * 20, nil
}

func NewSpeexCodecData(sr int, cl av.ChannelLayout) SpeexCodecData {
	codec := SpeexCodecData{}
	codec.CodecType_ = av.SPEEX
	codec.SampleFormat_ = av.S16
	codec.SampleRate_ = sr
	codec.ChannelLayout_ = cl
	return codec
}

type OpusCodecData struct {
	fake.CodecData
}

func NewOPUSCodecData(sr int) OpusCodecData {
	codec := OpusCodecData{}
	codec.CodecType_ = av.OPUS
	codec.SampleFormat_ = av.S16
	codec.SampleRate_ = sr
	codec.ChannelLayout_ = 2
	return codec
}
