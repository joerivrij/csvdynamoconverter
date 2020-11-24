package impl

import "encoding/json"

type Configurations []Configuration

func UnmarshalConfigurations(data []byte) (Configurations, error) {
	var r Configurations
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Configurations) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type Configuration struct {
	Configuration string `json:"configuration"`
	Accountid     int64  `json:"accountid"`
	Budget        int64  `json:"budget"`
	Email         string `json:"email"`
}
