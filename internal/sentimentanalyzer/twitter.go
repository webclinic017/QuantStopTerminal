package sentimentanalyzer

import (
	"context"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/pkg/vader"
	"github.com/quantstop/quantstopterminal/pkg/vader/sentitext"
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
	scraper.SetSearchMode(twitterScraper.SearchLatest)
	timer := time.NewTicker(time.Second * 30)

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

			var scores []sentitext.Sentiment

			for tweet := range scraper.SearchTweets(context.Background(),
				"$GME -filter:retweets", 50) {
				if tweet.Error != nil {
					log.Errorf(log.SentimentAnalyzer, "Twitter analyzer error. %v", tweet.Error)
				}
				//log.Debugln(log.SentimentAnalyzer, "Tweet | "+tweet.Text)
				scores = append(scores, vader.GetSentiment(tweet.Text))
				//log.Debugf(log.SentimentAnalyzer, "Twitter analyzer sentiment: %v", sentiment)

			}

			avgCompoundScore := 0.0
			avgNeutralScore := 0.0
			avgNegativeScore := 0.0
			avgPositiveScore := 0.0

			for _, sentiment := range scores {
				avgCompoundScore = avgCompoundScore + sentiment.Compound
				avgNeutralScore = avgNeutralScore + sentiment.Neutral
				avgNegativeScore = avgNegativeScore + sentiment.Negative
				avgPositiveScore = avgPositiveScore + sentiment.Positive
			}

			avgCompoundScore = avgCompoundScore / float64(len(scores))
			avgNeutralScore = avgNeutralScore / float64(len(scores))
			avgNegativeScore = avgNegativeScore / float64(len(scores))
			avgPositiveScore = avgPositiveScore / float64(len(scores))

			log.Debugf(
				log.SentimentAnalyzer,
				"Twitter analyzer sentiment: Compound: %v Neutral: %v Negative: %v Positive: %v",
				avgCompoundScore,
				avgNeutralScore,
				avgNegativeScore,
				avgPositiveScore,
			)

		}
	}

}
