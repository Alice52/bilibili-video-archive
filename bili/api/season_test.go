package api

import (
	"fmt"
	"github.com/alice52/archive/bili/util"
	"github.com/alice52/archive/common/global"
	"github.com/wordpress-plus/kit-logger/viperx"
	"github.com/wordpress-plus/kit-logger/zapx"
	"testing"
)

func init() {
	global.VIPER = viperx.Viper(&global.CONFIG, "../config-local.yaml")
	global.LOG = zapx.Zap(global.CONFIG.Zap)
}

func TestSeasonSection(t *testing.T) {
	info, err := logonFunc().SeasonSection("", "729217") //
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(util.MustMarshal(info))
}
