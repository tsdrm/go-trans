package format

import (
	"github.com/tangs-drm/go-trans"
	"github.com/tangs-drm/go-trans/format/flv"
	"github.com/tangs-drm/go-trans/format/other"
)

func Init() {
	var formats = go_trans.GetFormats()
	for _, format := range formats {
		go_trans.RegisterPlugin(format, getTransPlugin(format))
	}
}

func getTransPlugin(format string) go_trans.TransPlugin {
	switch format {
	case "flv":
		return flv.Flv{}
	case "other":
		return other.Other{}
	}
	return nil
}
