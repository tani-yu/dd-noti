package slack

import (
	"fmt"
	"os"
	"strconv"

	"github.com/slack-go/slack"
	"github.com/tani-yu/dd-noti/datadog"
)

func PostMessageMonitor(mm []datadog.MutedMonitor) {
	var trigger = true
	var err error
	// アクセストークンを使用してクライアントを生成する
	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")

	if os.Getenv("NOTIFY_MOMIT") != "" {
		trigger, err = strconv.ParseBool(os.Getenv("NOTIFY_MOMIT"))
		if err != nil {
			trigger = true
			fmt.Fprintf(os.Stderr, "Error Failed to cast NOTIFY_MOMIT: %v\n", err)
		}
	}

	if trigger {
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
}

func PostMessageHost(mh []string) {
	var trigger = true
	var err error
	// アクセストークンを使用してクライアントを生成する
	token := os.Getenv("SLACK_API_TOKEN")
	channel := os.Getenv("SLACK_CHANNEL")

	if os.Getenv("NOTIFY_HOST") != "" {
		trigger, err = strconv.ParseBool(os.Getenv("NOTIFY_HOST"))
		if err != nil {
			trigger = true
			fmt.Fprintf(os.Stderr, "Error Failed to cast NOTIFY_HOST: %v\n", err)
		}
	}

	if trigger {
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
