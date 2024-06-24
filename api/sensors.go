package api

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// Sensor is a struct that represents a sensor from the BMC
type sensorJson struct {
	ID                             int             `json:"id"`
	SensorNumber                   int             `json:"sensor_number"`
	Name                           string          `json:"name"`
	OwnerID                        int             `json:"owner_id"`
	OwnerLUN                       int             `json:"owner_lun"`
	RawReading                     float64         `json:"raw_reading"`
	Reading                        json.RawMessage `json:"reading"`
	Unit                           string          `json:"unit"`
	Type                           string          `json:"type"`
	TypeNumber                     int             `json:"type_number"`
	SensorState                    int             `json:"sensor_state"`
	DiscreteState                  int             `json:"discrete_state"`
	LowerNonRecoverableThreshold   json.RawMessage `json:"lower_non_recoverable_threshold"`
	LowerCriticalThreshold         json.RawMessage `json:"lower_critical_threshold"`
	LowerNonCriticalThreshold      json.RawMessage `json:"lower_non_critical_threshold"`
	HigherNonCriticalThreshold     json.RawMessage `json:"higher_non_critical_threshold"`
	HigherCriticalThreshold        json.RawMessage `json:"higher_critical_threshold"`
	HigherNonRecoverableThreshHold json.RawMessage `json:"higher_non_recoverable_threshold"`
	Accessible                     int             `json:"accessible"`
}

type Sensor struct {
	ID                           int
	SensorNumber                 int
	Name                         string
	Reading                      string
	Unit                         string
	Type                         string
	State                        string
	Accessible                   string
	LowerNonRecoverableThreshold string
	LowerCriticalThreshold       string
	LowerNonCriticalThreshold    string
	HigherNonCriticalThreshold   string
	HigherCriticalThreshold      string
	Alert                        string
}

// GetSensorsList returns a list of sensors from the BMC
func (a *Api) GetSensorsList(host string) ([]*Sensor, error) {
	resp, err := a.getRequest(host, "sensors")
	if err != nil {
		return nil, fmt.Errorf("failed to get sensors list: %v", err)
	}

	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return nil, &BMCAPIError{StatusCode: resp.StatusCode, Message: "failed to get sensors list"}
	}

	var sensors []*sensorJson
	err = json.NewDecoder(resp.Body).Decode(&sensors)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	var sensorsList []*Sensor
	for _, sensor := range sensors {
		s := &Sensor{
			ID:           sensor.ID,
			SensorNumber: sensor.SensorNumber,
			Name:         sensor.Name,
			Type:         sensor.Type,
			Accessible:   toAccessible(sensor.Accessible),
			State:        toActive(sensor.SensorState),
			Unit:         sensor.Unit,
		}

		s.Reading = fromRawMessage(sensor.Reading)
		s.LowerNonRecoverableThreshold = fromRawMessage(sensor.LowerNonRecoverableThreshold)
		s.Alert = checkAlert(s.Reading, s.LowerNonRecoverableThreshold, "<", "Lower non-recoverable threshold exceeded")
		s.LowerCriticalThreshold = fromRawMessage(sensor.LowerCriticalThreshold)
		if s.Alert == "" {
			s.Alert = checkAlert(s.Reading, s.LowerCriticalThreshold, "<", "Lower critical threshold exceeded")
		}
		s.LowerNonCriticalThreshold = fromRawMessage(sensor.LowerNonCriticalThreshold)
		if s.Alert == "" {
			s.Alert = checkAlert(s.Reading, s.LowerNonCriticalThreshold, "<", "Lower non-critical threshold exceeded")
		}
		s.HigherNonCriticalThreshold = fromRawMessage(sensor.HigherNonCriticalThreshold)
		if s.Alert == "" {
			s.Alert = checkAlert(s.Reading, s.HigherNonCriticalThreshold, ">", "Higher non-critical threshold exceeded")
		}
		s.HigherCriticalThreshold = fromRawMessage(sensor.HigherCriticalThreshold)
		if s.Alert == "" {
			s.Alert = checkAlert(s.Reading, s.HigherCriticalThreshold, ">", "Higher critical threshold exceeded")
		}
		sensorsList = append(sensorsList, s)
	}

	return sensorsList, nil
}

func toActive(val int) string {
	if val == 1 {
		return "active"
	}
	return "inactive"
}

func toAccessible(val int) string {
	if val == 0 {
		return "accessible"
	}
	return "inaccessible"
}

func fromRawMessage(data json.RawMessage) string {
	if data != nil {
		data, err := data.MarshalJSON()
		if err != nil {
			return "N/A"
		} else {
			reading := string(data)
			floatValue, err := strconv.ParseFloat(reading, 64)
			if err == nil {
				return strconv.FormatFloat(floatValue, 'f', -1, 64)
			} else {
				return reading
			}
		}
	} else {
		return "N/A"
	}
}

func checkAlert(value string, threshold string, operation string, alertMessage string) string {
	if value != "N/A" && threshold != "N/A" {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return ""
		}
		floatThreshold, err := strconv.ParseFloat(threshold, 64)
		if err != nil {
			return ""
		}
		switch operation {
		case ">":
			if floatValue > floatThreshold {
				return alertMessage
			}
		case "<":
			if floatValue < floatThreshold {
				return alertMessage
			}
		}
	}
	return ""
}
