package datatypes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"github.com/spf13/viper"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Datapoint struct {
	Timestamp int `json:"timestamp"`
	Value float64 `json:"value"`
	Comment string `json:"comment"`
	Id string `json:"id"`
	UpdatedAt int `json:"updated_at"`
	RequestId string `json:"requestid"`
	Daystamp string `json:"daystamp"`
}

type DatapointResponse struct {
	Collection []Datapoint
}

func (f *DatapointResponse) UnmarshalJSON(bs []byte) error {
	return json.Unmarshal(bs, &f.Collection)
}

type Goal struct {
	Slug string `json:"slug"`
	UpdatedAt int `json:"updated_at"`
	Title string `json:"title"`
	FinePrint string `json:"fineprint"`
	YAxis string `json:"yaxis"`
	GoalDate int `json:"goaldate"`
	GoalVal int `json:"goalval"`
	Rate float64 `json:"rate"`
	SvgUrl string `json:"svg_url"`
	GraphUrl string `json:"graph_url"`
	ThumbUrl string `json:"thumb_url"`
	Autodata string `json:"autodata"`
	GoalType string `json:"goal_type"`
	LoseDate int `json:"losedate"`
	Queued bool `json:"queued"`
	Secret bool `json:"secret"`
	DataPublic bool `json:"datapublic"`
	Datapoints []Datapoint `json:"datapoints"`
	NumPts int `json:"numpts"`
	Pledge float64 `json:"pledge"`
	InitDay int `json:"initday"`
	InitVal float64 `json:"initval"`
	CurDay int `json:"curday"`
	CurVal float64 `json:"curval"`
	LastDay int `json:"lastday"`
	RUnits string `json:"runits"`
	Yaw int `json:"yaw"`
	Direction int `json:"dir"`
	Lane int `json:"lane"`
	MathIsHard [3]float64 `json:"mathishard"`
}

type User struct {
	Username string
	ApiKey string
	Goals []Goal
}