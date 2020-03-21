package helpers

import "github.com/kpfaulkner/grafanadatadog/pkg/models"

// ConvertDDResponseToGrafanaResponse converts raw datadog responses to something grafana can handle
// with the simple json plugin.
func ConvertDDResponseToGrafanaResponse(ddResponse models.DatadogQueryResponse) ([]models.TimeSeriesResponse,error) {

	return nil, nil
}
