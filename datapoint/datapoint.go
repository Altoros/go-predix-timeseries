// Package datapoint contains data point representation.
package datapoint

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Altoros/go-predix-timeseries/dataquality"
	"github.com/Altoros/go-predix-timeseries/measurement"
)

type Datapoint struct {
	Measure   measurement.Measurement
	Timestamp time.Time
	Quality   dataquality.Quality
}

// Custom JSON marshaling because Time Series expects a data point to be a 3-element array
func (p *Datapoint) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("[%d,%s,%d]", p.Timestamp.UnixNano()/int64(time.Millisecond), p.Measure.Value(), p.Quality)), nil
}

// Custom JSON unmarshalling because Time Series returns a datapoint as a 3-element array
func (p *Datapoint) UnmarshalJSON(bs []byte) error {
	s := string(bs)
	if !strings.HasPrefix(s, "[") && !strings.HasSuffix(s, "]") {
		return errors.New("Not a datapoint")
	}
	values := strings.Split(strings.Trim(s, "[]"), ",")

	timestamp, err := strconv.ParseInt(values[0], 10, 64)
	if err != nil {
		return nil
	}
	p.Timestamp = time.Unix(0, timestamp*int64(time.Millisecond))
	lastElement := len(values) - 1
	p.Measure, err = measurement.FromString(strings.Join(values[1:lastElement], ","))
	if err != nil {
		return err
	}
	q, err := strconv.Atoi(values[len(values)-1])
	if err != nil {
		return err
	}
	p.Quality = dataquality.Quality(q)
	return nil
}
