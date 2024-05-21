package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)



type NBA struct {
	Events[] Event `json:"events"`
}



type Event struct {
	ID            string          `json:"id"`
	UID           string          `json:"uid"`
	Date          string          `json:"date"`
	Name          string          `json:"name"`
	ShortName     string          `json:"shortName"`
	Season        Season          `json:"season"`
	Competitions  []Competition   `json:"competitions"`
	Status        Status          `json:"status"`
	Notes         []Note          `json:"notes"`
	Broadcasts    []Broadcast     `json:"broadcasts"`
	Format        Format          `json:"format"`
	Tickets       []Ticket        `json:"tickets"`
	StartDate     string          `json:"startDate"`
	Series        Series          `json:"series"`
	GeoBroadcasts []GeoBroadcast  `json:"geoBroadcasts"`
	Odds          []Odd           `json:"odds"`
}

type Season struct {
	Year int    `json:"year"`
	Type int    `json:"type"`
	Slug string `json:"slug"`
}

type Competition struct {
	ID                   string        `json:"id"`
	UID                  string        `json:"uid"`
	Date                 string        `json:"date"`
	Attendance           int           `json:"attendance"`
	Type                 Type          `json:"type"`
	TimeValid            bool          `json:"timeValid"`
	NeutralSite          bool          `json:"neutralSite"`
	ConferenceCompetition bool        `json:"conferenceCompetition"`
	PlayByPlayAvailable  bool          `json:"playByPlayAvailable"`
	Recent               bool          `json:"recent"`
	Venue                Venue         `json:"venue"`
	Competitors          []Competitor  `json:"competitors"`
}

type Type struct {
	ID           string `json:"id"`
	Abbreviation string `json:"abbreviation"`

}

type Venue struct {
	ID       string  `json:"id"`
	FullName string  `json:"fullName"`
	Address  Address `json:"address"`
	Indoor   bool    `json:"indoor"`
}

type Address struct {
	City  string `json:"city"`
	State string `json:"state"`
}

type Competitor struct {
	ID         string        `json:"id"`
	UID        string        `json:"uid"`
	Type       string        `json:"type"`
	Order      int           `json:"order"`
	HomeAway   string        `json:"homeAway"`
	Team       Team          `json:"team"`
	Score      string        `json:"score"`
	Statistics []Statistic   `json:"statistics"`
	Records    []Record      `json:"records"`
	Leaders    []LeaderGroup `json:"leaders"`
}

type Team struct {
	ID                string  `json:"id"`
	UID               string  `json:"uid"`
	Location          string  `json:"location"`
	Name              string  `json:"name"`
	Abbreviation      string  `json:"abbreviation"`
	DisplayName       string  `json:"displayName"`
	ShortDisplayName  string  `json:"shortDisplayName"`
	Color             string  `json:"color"`
	AlternateColor    string  `json:"alternateColor"`
	IsActive          bool    `json:"isActive"`
	Venue             Venue   `json:"venue"`
	Links             []Link  `json:"links"`
	Logo              string  `json:"logo"`
}

type Link struct {
	Rel        []string `json:"rel"`
	Href       string   `json:"href"`
	Text       string   `json:"text"`
	IsExternal bool     `json:"isExternal"`
	IsPremium  bool     `json:"isPremium"`
}

type Statistic struct {
	Name           string  `json:"name"`
	Abbreviation   string  `json:"abbreviation"`
	DisplayValue   string  `json:"displayValue"`
	RankDisplayValue string `json:"rankDisplayValue"`
}

type Record struct {
	Name     string `json:"name"`
	Type     string `json:"type"`
	Summary  string `json:"summary"`
}

type LeaderGroup struct {
	Name            string   `json:"name"`
	DisplayName     string   `json:"displayName"`
	ShortDisplayName string  `json:"shortDisplayName"`
	Abbreviation    string   `json:"abbreviation"`
	Leaders         []Leader `json:"leaders"`
}

type Leader struct {
	DisplayValue string  `json:"displayValue"`
	Value        float64 `json:"value"`
	Athlete      Athlete `json:"athlete"`
	Team         Team    `json:"team"`
}

type Athlete struct {
	ID          string  `json:"id"`
	FullName    string  `json:"fullName"`
	DisplayName string  `json:"displayName"`
	ShortName   string  `json:"shortName"`
	Links       []Link  `json:"links"`
	Headshot    string  `json:"headshot"`
	Jersey      string  `json:"jersey"`
	Position    Position `json:"position"`
	Team        Team    `json:"team"`
	Active      bool    `json:"active"`
}

type Position struct {
	Abbreviation string `json:"abbreviation"`
}

type Note struct {
	Type     string `json:"type"`
	Headline string `json:"headline"`
}

type Status struct {
	Clock        float64 `json:"clock"`
	DisplayClock string  `json:"displayClock"`
	Period       int     `json:"period"`
	Type         StatusType    `json:"type"`
}

type StatusType struct{
	ID  string `json:"id"`
	Name string `json: "name"`
	State string `json: "state"`
	Completed bool `json: "completed"`
	Description string `json: "description"`
	Detail string `json: "Detail"`
	ShortDetail string `json: "shortDetail"`
	


}

type Broadcast struct {
	Market Market `json:"market"`
	Names  []string `json:"names"`
}

type Market struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type Format struct {
	Regulation Regulation `json:"regulation"`
}

type Regulation struct {
	Periods int `json:"periods"`
}

type Ticket struct {
	Summary        string `json:"summary"`
	NumberAvailable int   `json:"numberAvailable"`
	Links          []Link `json:"links"`
}

type Series struct {
	Type              string       `json:"type"`
	Title             string       `json:"title"`
	Summary           string       `json:"summary"`
	Completed         bool         `json:"completed"`
	TotalCompetitions int          `json:"totalCompetitions"`
	Competitors       []Competitor `json:"competitors"`
}

type GeoBroadcast struct {
	Type    Type   `json:"type"`
	Market  Market `json:"market"`
	Media   Media  `json:"media"`
	Lang    string `json:"lang"`
	Region  string `json:"region"`
}

type Media struct {
	ShortName string `json:"shortName"`
}

type Odd struct {
	Provider Provider `json:"provider"`
	Details  string   `json:"details"`
}

type Provider struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Priority int    `json:"priority"`
}
func main(){
	res, err:= http.Get("https://site.api.espn.com/apis/site/v2/sports/basketball/nba/scoreboard")
	if err!= nil{
		panic(err)
	}

	defer res.Body.Close()

	if res.StatusCode !=  http.StatusOK{
		panic("Espn Api not available")
	}

	body, err := io.ReadAll(res.Body)

	if err!= nil{
		panic(err)
	}


	

	var nba NBA
	json.Unmarshal(body, &nba)
	events := nba.Events

	for _, event := range events{
		fmt.Println(event.Name, event.Status.Type.ShortDetail)
	}


}
