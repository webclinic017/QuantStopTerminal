package sentimentanalyzer

import (
	"context"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/vader"
	"sync"
	"time"

	twitterScraper "github.com/n0madic/twitter-scraper"
)

type TwitterAnalyzer struct {
	Shutdown chan struct{}
	Wg       *sync.WaitGroup
}

func (r *TwitterAnalyzer) TestTwitterSentiment() {

	scraper := twitterScraper.New()

	timer := time.NewTimer(30 * time.Second)

	defer func() {
		timer.Stop()
		r.Wg.Done()
		log.Debugln(log.SentimentAnalyzer, "Twitter analyzer shutdown completed.")
	}()

	for {
		select {

		case <-r.Shutdown:
			log.Debugf(log.SentimentAnalyzer, "Twitter analyzer stopping ...\n")
			return
		case <-timer.C:
			for tweet := range scraper.SearchTweets(context.Background(),
				"$GME -filter:retweets", 50) {
				if tweet.Error != nil {
					log.Errorf(log.SentimentAnalyzer, "Twitter analyzer error. %v", tweet.Error)
				}
				log.Debugln(log.SentimentAnalyzer, "Tweet | "+tweet.Text)

				sentiment := vader.GetSentiment(tweet.Text)
				log.Debugf(log.SentimentAnalyzer, "Twitter analyzer sentiment: %v", sentiment)

			}
			return
		}
	}

}
