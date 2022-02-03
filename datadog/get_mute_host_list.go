package datadog

import (
	"context"
	"fmt"
	"os"

	datadog "github.com/DataDog/datadog-api-client-go/api/v1/datadog"
)

func GetMuteHostList() []string {
	var res []string
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
	for i, en := 0, 0; ; i++ {
		from := int64(789)            // int64 | Number of seconds from which you want to get total number of active hosts. (optional)
		filter := ""                  // string | String to filter search results. (optional)
		sortField := ""               // string | Sort hosts by this field. (optional)
		sortDir := ""                 // string | Direction of sort. Options include `asc` and `desc`. (optional)
		start := int64(en)            // int64 | Host result to start search from. (optional)
		count := int64(1000)          // int64 | Number of hosts to return. Max 1000. (optional)
		includeMutedHostsData := true // bool | Include information on the muted status of hosts and when the mute expires. (optional)
		includeHostsMetadata := false // bool | Include additional metadata about the hosts (agent_version, machine, platform, processor, etc.). (optional)
		optionalParams := datadog.ListHostsOptionalParameters{
			Filter:                &filter,
			SortField:             &sortField,
			SortDir:               &sortDir,
			Start:                 &start,
			Count:                 &count,
			From:                  &from,
			IncludeMutedHostsData: &includeMutedHostsData,
			IncludeHostsMetadata:  &includeHostsMetadata,
		}

		en += 1000
		configuration := datadog.NewConfiguration()

		apiClient := datadog.NewAPIClient(configuration)
		resp, r, err := apiClient.HostsApi.ListHosts(ctx, optionalParams)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error when calling `HostsApi.ListHosts`: %v\n", err)
			fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
		}

		hl := resp.GetHostList()
		for _, data := range hl {
			if data.GetIsMuted() {
				res = append(res, data.GetHostName())
			}
		}

		if resp.GetTotalMatching() != 1000 {
			break
		}
	}
	return res
}
