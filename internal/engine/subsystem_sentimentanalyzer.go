package engine

import (
	"github.com/quantstop/quantstopterminal/internal/log"
	"github.com/quantstop/quantstopterminal/internal/sentimentanalyzer"
	"sync"
)

type SentimentAnalyzerSubsystem struct {
	Subsystem
	wg *sync.WaitGroup
}

func (s *SentimentAnalyzerSubsystem) init(bot *Engine, name string) error {

	if err := s.Subsystem.init(bot, name); err != nil {
		return err
	}
	/*s.enabled = bot.Config*/

	s.enabled = true
	s.initialized = true
	log.Debugln(log.SentimentAnalyzer, s.name+MsgSubsystemInitialized)
	return nil
}

func (s *SentimentAnalyzerSubsystem) start(wg *sync.WaitGroup) (err error) {
	if err = s.Subsystem.start(wg); err != nil {
		return err
	}

	s.started = true

	/*redditAnalyzer := sentimentanalyzer.RedditAnalyzer{
		Shutdown: s.shutdown,
		Wg: wg,
	}

	redditAnalyzer.Wg.Add(1)
	go redditAnalyzer.TestWSBSentiment()*/

	twitterAnalyzer := sentimentanalyzer.TwitterAnalyzer{
		Shutdown: s.shutdown,
		Wg:       wg,
	}

	wg.Add(1)
	go twitterAnalyzer.TestTwitterSentiment()

	log.Debugln(log.SentimentAnalyzer, s.name+MsgSubsystemStarted)
	return nil
}

func (s *SentimentAnalyzerSubsystem) stop() error {

	if err := s.Subsystem.stop(); err != nil {
		return err
	}
	close(s.shutdown)
	s.started = false
	log.Debugln(log.SentimentAnalyzer, s.name+MsgSubsystemShutdown)
	return nil
}
