package media

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

// ENDPOINT is Twitter URL
const ENDPOINT = "https://twitter.com/i/profiles/show/${screen_name}/media_timeline"

// GetImageUrls Get images from screenName
func GetImageUrls(screenName string, size int, lastID uint64) (urls []string, err error) {
	endpoint := strings.Replace(ENDPOINT, "${screen_name}", screenName, -1)

	var urlList []string
	for len(urls) < size {
		urlList, lastID, err = request(endpoint, lastID)
		if err != nil {
			return
		}
		urls = append(urls, urlList...)
	}

	if len(urls) > size {
		urls = urls[:size]
	}

	log.Println("lastID:", lastID)

	return
}

func request(endpoint string, id uint64) (urls []string, lastID uint64, err error) {
	// API request
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		return
	}
	query := make(url.Values)
	query.Add("include_available_features", "1")
	query.Add("include_entities", "1")
	// query.Add("last_note_ts", "1445020184")
	if id > 0 {
		query.Add("max_position", strconv.FormatUint(id, 10))
	}
	query.Add("oldest_unread_id", "0")
	query.Add("reset_error_state", "false")
	req.URL.RawQuery = query.Encode()

	// NOTE: debug message
	// dump, _ := httputil.DumpRequestOut(req, true)
	// fmt.Printf("%s\n", dump)

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	// Decode JSON
	data := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&data)
	htmlStr := data["items_html"].(string)

	// Parse HTML
	doc, err := html.Parse(strings.NewReader(htmlStr))
	if err != nil {
		return
	}

	urls, lastID = parser(doc)

	return
}

// html parser wrapper
func parser(n *html.Node) (urls []string, lastID uint64) {
	urls = make([]string, 0, 10)

	if n.Type == html.ElementNode && n.Data == "div" {
		if src, ok := getAttr(n, "data-img-src"); ok {
			urls = append(urls, src)
		}
		if id, ok := getAttr(n, "data-tweet-id"); ok {
			_id, err := strconv.ParseUint(id, 10, 64)
			if err == nil {
				lastID = _id
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		_urls, _id := parser(c)
		if _id > 0 {
			lastID = _id
		}
		urls = append(urls, _urls...)
	}

	return
}

// get a attribute value
func getAttr(n *html.Node, key string) (val string, ok bool) {
	for _, e := range n.Attr {
		if e.Key == key {
			val = e.Val
			ok = true
			return
		}
	}

	return
}
