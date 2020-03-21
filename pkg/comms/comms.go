package comms

type Comms struct {
	apiKey string
	appKey string

}

func NewComms( apiKey string, appKey string) Comms {
	c := Comms{}
	c.apiKey = apiKey
	c.appKey = appKey
	return c
}

func (c Comms) DoPost( query string ) {

}
