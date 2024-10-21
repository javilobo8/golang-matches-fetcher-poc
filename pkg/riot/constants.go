package riot

var RoutingHost = map[string]string{
	"americas": "https://americas.api.riotgames.com",
	"asia":     "https://asia.api.riotgames.com",
	"europe":   "https://europe.api.riotgames.com",
	"sea":      "https://sea.api.riotgames.com",
}

var PlatformHosts = map[string]string{
	"na":   "https://na1.api.riotgames.com",
	"euw":  "https://euw1.api.riotgames.com",
	"eune": "https://eun1.api.riotgames.com",
	"kr":   "https://kr.api.riotgames.com",
	"br":   "https://br1.api.riotgames.com",
	"lan":  "https://la1.api.riotgames.com",
	"las":  "https://la2.api.riotgames.com",
	"oce":  "https://oc1.api.riotgames.com",
	"tr":   "https://tr1.api.riotgames.com",
	"ru":   "https://ru.api.riotgames.com",
	"jp":   "https://jp1.api.riotgames.com",
	"vn":   "https://vn1.api.riotgames.com",
	"ph":   "https://ph1.api.riotgames.com",
}

var QueueIds = map[string]int{
	"RANKED_SOLO_5x5": 420,
	"RANKED_FLEX_SR":  440,
	"ARAM":            450,
}
