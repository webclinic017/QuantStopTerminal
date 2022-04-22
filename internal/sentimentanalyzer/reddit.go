package sentimentanalyzer

import (
	"context"
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/vartanbeno/go-reddit/v2/reddit"
	"net/http"
	"sync"
	"time"
)

var ctx = context.Background()

type RedditAnalyzer struct {
	Shutdown chan struct{}
	Wg       *sync.WaitGroup
}

func (r *RedditAnalyzer) TestWSBSentiment() {

	/*posts, errs, stop := reddit.DefaultClient().Stream.Posts(
		"wallstreetbets",
		reddit.StreamInterval(time.Second*5),
		reddit.StreamDiscardInitial,
	)*/
	client, err := reddit.NewReadonlyClient()
	if err != nil {
		return
	}

	posts, res, err := client.Subreddit.HotPosts(ctx, "wallstreetbets", &reddit.ListOptions{Limit: 5})
	if err != nil {
		log.Errorf(log.SentimentAnalyzer, "Reddit analyzer error! %v\n", err)
		return
	}
	for _, post := range posts {
		log.Debugln(log.SentimentAnalyzer, post.Title)
	}
	logResponse(res.Request, res.Response)

	//client.OnRequestCompleted(logResponse)

	timer := time.NewTicker(time.Second * 30)

	defer func() {
		timer.Stop()
		r.Wg.Done()
		log.Debugln(log.SentimentAnalyzer, "Reddit analyzer shutdown completed.")
	}()

	for {
		select {
		/*case post, ok := <-posts:
			if !ok {
				return
			}
			log.Debugf(log.SentimentAnalyzer, "Received post: %s\n", post.Title)
		case err, ok := <-errs:
			if !ok {
				return
			}
			log.Errorf(log.SentimentAnalyzer, "Reddit analyzer error! %v\n", err)*/
		case <-r.Shutdown:
			log.Debugf(log.SentimentAnalyzer, "Reddit analyzer stopping ...\n")
			return
			/*case <-timer.C:
			_, _, err := client.Subreddit.HotPosts(ctx, "wallstreetbets", &reddit.ListOptions{Limit: 5})
			if err != nil {
				log.Errorf(log.SentimentAnalyzer, "Reddit analyzer error! %v\n", err)
				return
			}
			return*/
		}
	}

}

func logResponse(req *http.Request, res *http.Response) {
	log.Debugf(log.SentimentAnalyzer, "%s %s %s\n", req.Method, req.URL, res.Status)
}
