package slack

import (
	"os"

	"github.com/slack-go/slack"
	"github.com/tani-yu/dd-noti/datadog"
)

func PostMessageMonitor(mm []datadog.MutedMonitor) {
	// アクセストークンを使用してクライアントを生成する
	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	c := slack.New(token)

	attachment := slack.Attachment{
		Text:   "muteになっているMonitorをお知らせします",
		Fields: createAttachmentFields(mm),
	}

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage(channel, slack.MsgOptionText("datadog muted noti", true), slack.MsgOptionAttachments(attachment))
	if err != nil {
		panic(err)
	}
}

func PostMessageHost(mh []string) {
	// アクセストークンを使用してクライアントを生成する
	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")
	c := slack.New(token)

	attachment := slack.Attachment{
		Text:   "muteになっているHostお知らせします",
		Fields: createAttachmentFieldsFromString(mh),
	}

	// MsgOptionText() の第二引数に true を設定すると特殊文字をエスケープする
	_, _, err := c.PostMessage(channel, slack.MsgOptionText("datadog muted noti", true), slack.MsgOptionAttachments(attachment))
	if err != nil {
		panic(err)
	}
}

func createAttachmentFields(mm []datadog.MutedMonitor) []slack.AttachmentField {
	var af []slack.AttachmentField
	for _, data := range mm {
		af = append(af, slack.AttachmentField{
			Title: "MonitorName: \n" + data.GetMName(),
			Value: "Target: \n" + createStringFromMap(data.GetMInfo()),
			Short: false,
		})
	}
	return af
}

func createStringFromMap(mm map[string]int64) string {
	var rs string
	for k, _ := range mm {
		rs = rs + k + ", "
	}
	return rs
}

func createAttachmentFieldsFromString(mh []string) []slack.AttachmentField {
	var af []slack.AttachmentField
	for _, data := range mh {
		af = append(af, slack.AttachmentField{
			Title: "HostName:",
			Value: data,
			Short: false,
		})
	}
	return af
}
