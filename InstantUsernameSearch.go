package InstantUsernameSearch2

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"
)

type Resp struct {
	Service   string `json:"service"`
	URL       string `json:"url"`
	Available bool   `json:"available"`
}

type Resps struct {
	Resp *[]Resp
}

var (
	myClient = &http.Client{Timeout: 10 * time.Second}
	services = []string{"twitter", "facebook", "youtube", "medium", "reddit", "hackernews", "github", "quora", "9gag", "vk", "goodreads", "blogger", "patreon", "producthunt", "about.me", "academia.edu", "angellist", "aptoide", "askfm", "blip.fm", "bandcamp", "basecamp", "behance", "bitbucket", "bitcoinforum", "buzzfeed", "canva", "carbonmade", "cashme", "cloob", "codecademy", "codementor", "codepen", "coderwall", "colourlovers", "contently", "coroflot", "creativemarket", "crevado", "crunchyroll", "dev community", "dailymotion", "designspiration", "deviantart", "disqus", "dribbble", "ebay", "ello", "etsy", "eyeem", "flickr", "flipboard", "foursquare", "giphy", "gitlab", "gitee", "gravatar", "gumroad", "hackerone", "house-mixes.com", "houzz", "hubpages", "homescreen.me", "ifttt", "imageshack", "instructables", "investing.com", "issuu", "itch.io", "kaggle", "kanoworld", "keybase", "kik", "kongregate", "launchpad", "letterboxd", "livejournal", "mastodon", "mixcloud", "myanimelist", "namemc", "newgrounds", "pexels", "photobucket", "pinterest", "pixabay", "rajce.net", "repl.it", "roblox", "scribd", "signal", "slack", "slideshare", "soundcloud", "sourceforge", "spotify", "star citizen", "steam", "steamgroup", "telegram", "tradingview", "trakt", "trip", "tripadvisor", "twitch", "unsplash", "vsco", "venmo", "vimeo", "virustotal", "we heart it", "webnode", "fandom", "wikipedia", "wix", "wordpress", "youpic", "zhihu", "devrant", "last.fm", "makerlog"} // "imgsrc.ru" // "badoo", "500px", "jimdo", "reverbnation", "tinder", "meetme",  "taringa", "pastebin","imgur", "tiktok",
)

func Check(service, name string, res chan Resp) {
	url := "https://api.instantusername.com/check/" + service + "/" + name

	var resp Resp

	r, err := myClient.Get(url)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.Unmarshal(b, &resp)
	res <- resp
}

func CheckAll(username string) []Resp {
	res := make(chan Resp)
	arr := make([]Resp, 0)

	for _, service := range services {
		go Check(service, username, res)
	}

	for {
		select {
		case n := <-res:
			arr = append(arr, n)
		}
		if len(arr) == len(services) {
			break
		}
	}

	return arr
}

func ToString(arr []Resp) string {
	var s string = ""
	for _, v := range arr {
		if v.Service != "" {
			if !v.Available {
				s += v.Service + " " + v.URL + "\n"
			}
		}
	}

	return s
}
