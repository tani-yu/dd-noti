package main

import (
	dd "github.com/tani-yu/dd-noti/datadog"
	"github.com/tani-yu/dd-noti/slack"
)

func main() {
	mh := dd.GetMuteHostList()
	mm := dd.GetMuteMonitorList()
	//for _, data := range mm {
	//	fmt.Println(data.GetHostName())
	//	fmt.Println(data.GetMInfo())
	//}
	slack.PostMessageHost(mh)
	slack.PostMessageMonitor(mm)
}
