package main

import (
	"flag"
	"fmt"
	"lobobot/matches-fetcher/pkg/config"
	"lobobot/matches-fetcher/pkg/riot"
	"log/slog"
	"strings"
	"time"
)

type Arguments struct {
	Region   string
	GameName string
	TagLine  string
}

func parseArguments() Arguments {
	regionParam := flag.String("region", "", "Region")
	summonerParam := flag.String("name", "", "Summoner name and tag")
	flag.Parse()

	if *regionParam == "" {
		panic("Region is required")
	}

	if *summonerParam == "" {
		panic("Summoner name is required")
	}

	region := *regionParam

	// Split summoner name and tag
	summonerParts := strings.Split(*summonerParam, "#")
	if len(summonerParts) != 2 {
		slog.Error("Summoner name and tag must be in the format SummonerName#Tag")
	}

	GameName := summonerParts[0]
	TagLine := summonerParts[1]

	return Arguments{
		Region:   region,
		GameName: GameName,
		TagLine:  TagLine,
	}
}

func fetchAllMatches(riotapi *riot.RiotAPI, Region string, matchList []string, totalMatches *[]riot.MatchDTO) {
	matchResponsesChannel := make(chan riot.MatchDTO, len(matchList))
	for _, matchId := range matchList {
		go func(matchId string) {
			matchResponsesChannel <- riotapi.Match(Region, matchId)
		}(matchId)
	}
	matches := make([]riot.MatchDTO, 0, len(matchList))
	for i := 0; i < len(matchList); i++ {
		match := <-matchResponsesChannel
		matches = append(matches, match)
	}

	// Add matches to totalMatches
	*totalMatches = append(*totalMatches, matches...)
}

func main() {
	slog.Info("Init!")
	start := time.Now()
	config := config.ReadConfig()

	args := parseArguments()

	riotapi := riot.RiotAPI{
		ApiKey: config.Riot.ApiKey,
	}

	account := riotapi.AccountByName(args.Region, args.GameName, args.TagLine)
	fmt.Println(account.PUUID)

	matchList := riotapi.MatchList(args.Region, account.PUUID, 450, 100)
	startfetch := time.Now()

	matches := make([]riot.MatchDTO, 0, len(matchList))
	fetchAllMatches(&riotapi, args.Region, matchList, &matches)
	elapsedfetch := time.Since(startfetch)
	fmt.Println("Elapsed time for fetching:", elapsedfetch)

	startcalc := time.Now()
	var totalKills = 0
	var totalAssists = 0
	var totalDeaths = 0
	for _, match := range matches {
		for _, participant := range match.Info.Participants {
			if participant.PUUID == account.PUUID {
				totalKills += participant.Kills
				totalDeaths += participant.Deaths
				totalAssists += participant.Assists
			}
		}
	}
	meanKills := float64(totalKills) / float64(len(matches))
	meanDeaths := float64(totalDeaths) / float64(len(matches))
	meanAssists := float64(totalAssists) / float64(len(matches))
	elapsedcalc := time.Since(startcalc)
	fmt.Println("Elapsed time for calculations:", elapsedcalc)

	elapsed := time.Since(start)

	fmt.Println("Elapsed time:", elapsed)
	fmt.Println("Fetched matches:", len(matches))
	fmt.Println("Total kills:", totalKills)
	fmt.Println("Total deaths:", totalDeaths)
	fmt.Println("Total assists:", totalAssists)
	fmt.Printf("KDA: %.2f / %.2f / %.2f\n", meanKills, meanDeaths, meanAssists)
}
