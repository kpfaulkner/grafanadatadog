package pkg

import "github.com/kpfaulkner/grafanadatadog/pkg/models"

type Datadog struct {

}

func NewDatadog() *Datadog {
	d := Datadog{}
	return &d
}

// queryDatadog does the query... for now, just return fake data.
func (d *Datadog) queryDatadog(query string) (*models.DatadogQueryResponse,error) {


	return nil, nil
}
