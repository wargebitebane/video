package format

import (
	"github.com/wargebitebane/video/av/avutil"
	"github.com/wargebitebane/video/format/aac"
	"github.com/wargebitebane/video/format/flv"
	"github.com/wargebitebane/video/format/mp4"
	"github.com/wargebitebane/video/format/rtmp"
	"github.com/wargebitebane/video/format/rtsp"
	"github.com/wargebitebane/video/format/ts"
)

func RegisterAll() {
	avutil.DefaultHandlers.Add(mp4.Handler)
	avutil.DefaultHandlers.Add(ts.Handler)
	avutil.DefaultHandlers.Add(rtmp.Handler)
	avutil.DefaultHandlers.Add(rtsp.Handler)
	avutil.DefaultHandlers.Add(flv.Handler)
	avutil.DefaultHandlers.Add(aac.Handler)
}
