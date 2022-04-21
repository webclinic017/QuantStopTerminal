package vader

import (
	"github.com/quantstop/quantstopterminal/pkg/vader/lexicon"
	"github.com/quantstop/quantstopterminal/pkg/vader/sentitext"
)

//DoCalcuateSentiment of a sentence with a specific lexicon
func DoCalcuateSentiment(text string, lexicon lexicon.Lexicon) sentitext.Sentiment {
	senti := sentitext.Parse(text, lexicon)
	return sentitext.PolarityScore(senti)
}

//GetSentiment of a sentence with the default lexicon
func GetSentiment(text string) sentitext.Sentiment {
	senti := sentitext.Parse(text, lexicon.DefaultLexicon)
	return sentitext.PolarityScore(senti)
}
