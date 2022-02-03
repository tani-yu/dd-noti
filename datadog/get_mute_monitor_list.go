package datadog

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

type MutedMonitor struct {
	mName string
	mInfo map[string]int64
}

func (mm MutedMonitor) GetMName() string {
	return mm.mName
}

func (mm MutedMonitor) GetMInfo() map[string]int64 {
	return mm.mInfo
}

func GetMuteMonitorList() []MutedMonitor {
	ctx := context.WithValue(
		context.Background(),
		datadog.ContextAPIKeys,
		map[string]datadog.APIKey{
			"apiKeyAuth": {
				Key: os.Getenv("DD_CLIENT_API_KEY"),
			},
			"appKeyAuth": {
				Key: os.Getenv("DD_CLIENT_APP_KEY"),
			},
		},
	)
	var mm []MutedMonitor
	groupStates := ""     // string | When specified, shows additional information about the group states. Choose one or more from `all`, `alert`, `warn`, and `no data`. (optional)
	name := ""            // string | A string to filter monitors by name. (optional)
	tags := ""            // string | A comma separated list indicating what tags, if any, should be used to filter the list of monitors by scope. For example, `host:host0`. (optional)
	monitorTags := ""     // string | A comma separated list indicating what service and/or custom tags, if any, should be used to filter the list of monitors. Tags created in the Datadog UI automatically have the service key prepended. For example, `service:my-app`. (optional)
	withDowntimes := true // bool | If this argument is set to true, then the returned data includes all current active downtimes for each monitor. (optional)
	optionalParams := datadog.ListMonitorsOptionalParameters{
		GroupStates:   &groupStates,
		Name:          &name,
		Tags:          &tags,
		MonitorTags:   &monitorTags,
		WithDowntimes: &withDowntimes,
	}

	configuration := datadog.NewConfiguration()

	apiClient := datadog.NewAPIClient(configuration)
	resp, r, err := apiClient.MonitorsApi.ListMonitors(ctx, optionalParams)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MonitorsApi.ListMonitors`: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	for _, data := range resp {
		monit := data.GetOptions()
		if len(monit.GetSilenced()) != 0 {
			mm = append(mm, MutedMonitor{mName: data.GetName(), mInfo: data.Options.GetSilenced()})
		}
	}
	return mm
}
