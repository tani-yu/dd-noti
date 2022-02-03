package main

import (
	dd "github.com/tani-yu/dd-noti/datadog"
	"github.com/tani-yu/dd-noti/slack"
)

func main() {
	ss := dd.GetMuteHostList()
	mm := dd.GetMuteMonitorList()
	//for _, data := range mm {
	//	fmt.Println(data.GetHostName())
	//	fmt.Println(data.GetMInfo())
	//}
	slack.PostMessageHost(ss)
	slack.PostMessageMonitor(mm)
}
