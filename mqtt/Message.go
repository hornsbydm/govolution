package mqtt

type Message struct {
	Time   string        `json:"time"`
	Action MQTTAction    `json:"action"`
	Epoch  string        `json:"epoch"`
	Nice   string        `json:"nice"`
	ID     string        `json:"id"`
	Value  string        `josn:"value"`
	Units  UnitofMeasure `json:"units"`
}
