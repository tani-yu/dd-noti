package main

import (
	dd "github.com/tani-yu/dd-noti/datadog"
	"github.com/tani-yu/dd-noti/slack"
)

func main() {
	mh := dd.GetMuteHostList()
	mm := dd.GetMuteMonitorList()

	slack.PostMessageHost(mh)
	slack.PostMessageMonitor(mm)
}
