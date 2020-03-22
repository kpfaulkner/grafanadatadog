package helpers

import (
	"github.com/kpfaulkner/grafanadatadog/pkg/models"
	"sort"
	"time"
)

// ConvertDDResponseToGrafanaResponse converts response into something Grafana can handle
// For now it will convert to number of errors per minute.
func ConvertDDResponseToGrafanaResponse(ddResponse models.DatadogQueryResponse) ([]models.TimeSeriesResponse,error) {

	resp := models.TimeSeriesResponse{}

	data := getNumberOfEntriesPerMinute(ddResponse)
	resp.Datapoints = data
	resp.Target = "test1"
	return []models.TimeSeriesResponse{resp}, nil
}

func getNumberOfEntriesPerMinuteUnsorted(ddResponse models.DatadogQueryResponse ) [][]float64 {

	timeMap := make(map[time.Time]float64)

	// get number of entries within a given minute.
	for _,l := range ddResponse.Logs {
		//key := l.Content.Timestamp.Format("2006-01-02T15:04")
		key := l.Content.Timestamp.Truncate(time.Minute)
		// if key doesn't exist, it will return 0 anyway.
    entry := timeMap[key]
    entry++
    timeMap[key] = entry
	}

	// get key list
	keyList := make([]time.Time, len(timeMap), len(timeMap))
	for k,_ := range timeMap {
		keyList = append(keyList, k)
	}

	// sort it.
	sort.Slice(keyList, func(i int, j int) bool{
		return keyList[i].Before(keyList[j])
	})

	l := make([][]float64, len(keyList), len(keyList))
	for i,k := range keyList {
		v := timeMap[k]
		timeCountArray := []float64{v , float64(k.Unix()*1000)}
		l[i] = timeCountArray
	}

	return l
}

// assumption that DD returns data in chronological order.
func getNumberOfEntriesPerMinute(ddResponse models.DatadogQueryResponse ) [][]float64 {
	finalList := [][]float64{}
	currentTime := ddResponse.Logs[0].Content.Timestamp.Truncate(time.Minute)
	count := float64(1)
	for _,l := range ddResponse.Logs[1:] {
		if !l.Content.Timestamp.Truncate(time.Minute).Equal( currentTime) {
			finalList = append(finalList, []float64{count, float64(currentTime.Unix()*1000)})
			currentTime = l.Content.Timestamp.Truncate(time.Minute)
			count=1
		} else {
			count++
		}
	}
	finalList = append(finalList, []float64{count, float64(currentTime.Unix()*1000)})
  return finalList
}
