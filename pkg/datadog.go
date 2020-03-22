package pkg

import (
	"encoding/json"
	ddComms "github.com/kpfaulkner/grafanadatadog/pkg/comms"
	"github.com/kpfaulkner/grafanadatadog/pkg/models"
	"time"
)

type Datadog struct {
  comms ddComms.DatadogComms
}

func NewDatadog(apiKey string, appKey string) *Datadog {
	d := Datadog{}
	d.comms = ddComms.NewDatadogComms(apiKey, appKey)
	return &d
}

// queryDatadog does the query... for now, just return fake data.
func (d *Datadog) queryDatadog(query string, from time.Time, to time.Time) (*models.DatadogQueryResponse,error) {

	ddQuery := models.GenerateDatadogQuery(query, from, to)
	queryBytes,err := json.Marshal(ddQuery)
	if err != nil {
		return nil, err
	}

  resp, err := d.comms.DoPost(queryBytes)
  if err != nil {
  	return nil, err
  }

  var ddResp models.DatadogQueryResponse
  err = json.Unmarshal(resp, &ddResp)
  if err != nil {
  	return nil, err
  }

	return &ddResp, err
}
