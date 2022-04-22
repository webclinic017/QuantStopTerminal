package edgar

import (
	"encoding/xml"
	"fmt"
	"net/http"
)

var Address = []string{"https://www.sec.gov/Archives/edgar/usgaap.rss.xml"}

type Enclosure struct {
	Url    string `xml:"url,attr"`
	Length int64  `xml:"length,attr"`
	Type   string `xml:"type,attr"`
}

type Item struct {
	Title     string    `xml:"title"`
	Link      string    `xml:"link"`
	Desc      string    `xml:"description"`
	City      string    `xml:"city"`
	Company   string    `xml:"company"`
	Logo      string    `xml:"logo"`
	JobType   string    `xml:"jobtype"`
	Category  string    `xml:"category"`
	PubDate   string    `xml:"date"`
	Enclosure Enclosure `xml:"enclosure"`
}

type Channel struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Desc  string `xml:"description"`
	Items []Item `xml:"item"`
}

type Rss struct {
	Channel Channel `xml:"channel"`
}

// Crawler will store the account by given request body
func Crawler(address string) {
	resp, err := http.Get(address)
	if err != nil {
		fmt.Printf("Error GET: %v\n", err)
		return
	}
	defer resp.Body.Close()

	rss := Rss{}

	decoder := xml.NewDecoder(resp.Body)
	err = decoder.Decode(&rss)
	if err != nil {
		fmt.Printf("Error Decode: %v\n", err)
		return
	}
	var data Item
	for _, item := range rss.Channel.Items {
		data = item
		fmt.Println(data)
		//Store Data Here(Establish Prior Connection With DB)
		//Store.Insert(data)
	}
}

func WebCrawler() {
	for _, value := range Address {
		Crawler(value)
	}
}

func main() {
	/*cr := cron.New()
	cr.AddFunc("@midnight", WebCrawler)
	cr.AddFunc("@midnight", func() { fmt.Println("Every day") })
	cr.Start()*/
}
