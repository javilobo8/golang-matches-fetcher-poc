package riot

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type RiotAPI struct {
	ApiKey string
}

func getRoutingHost(Region string) string {
	return RoutingHost[Region]
}

func getAccountRoutingHost(Region string) string {
	if Region == "sea" {
		return RoutingHost["asia"]
	}

	return RoutingHost[Region]
}

func getPlatformHost(Platform string) string {
	return PlatformHosts[Platform]
}

func (r *RiotAPI) createRequest(url string) *http.Response {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	request.Header.Add("X-Riot-Token", r.ApiKey)
	if err != nil {
		panic(err)
	}
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		panic(err)
	}
	if response.StatusCode != 200 {
		fmt.Printf("Request to %s returned status code %d\n", url, response.StatusCode)
		panic("Request failed")
	}
	return response
}

func (r *RiotAPI) AccountByName(Region string, GameName string, TagLine string) AccountDTO {
	url := fmt.Sprintf("%s/riot/account/v1/accounts/by-riot-id/%s/%s", getAccountRoutingHost(Region), GameName, TagLine)
	response := r.createRequest(url)
	var account AccountDTO
	json.NewDecoder(response.Body).Decode(&account)
	return account
}

func (r *RiotAPI) MatchList(Region string, PUUID string, Queue int, Count int) []string {
	url := fmt.Sprintf("%s/lol/match/v5/matches/by-puuid/%s/ids?queue=%s&start=0&count=%s", getRoutingHost(Region), PUUID, fmt.Sprint(Queue), fmt.Sprint(Count))
	response := r.createRequest(url)
	var matchList []string
	json.NewDecoder(response.Body).Decode(&matchList)
	return matchList
}

func (r *RiotAPI) Match(Region string, MatchId string) MatchDTO {
	url := fmt.Sprintf("%s/lol/match/v5/matches/%s", getRoutingHost(Region), MatchId)
	response := r.createRequest(url)
	var match MatchDTO
	json.NewDecoder(response.Body).Decode(&match)
	return match
}
