package quiz

import (
	"fmt"
	"math/rand"
	"time"

	"githab.com/oleglacto/translater/internal/pkg/cache"
	"githab.com/oleglacto/translater/internal/pkg/language"
)

type IrregularVerbsGetter interface {
	GetIrregularVerbs() ([]language.IrregularVerb, error)
}

type IrregularVerbs struct {
	irregularVerbsGetter IrregularVerbsGetter
	cache                *cache.Cache
}

func NewIrregularVerbs(irregularVerbsGetter IrregularVerbsGetter, cache *cache.Cache) *IrregularVerbs {
	return &IrregularVerbs{irregularVerbsGetter: irregularVerbsGetter, cache: cache}
}

type IrregularVerbsGameStruct struct {
	Verb               language.IrregularVerb
	Question           string
	CorrectAnswer      string
	CorrectAnswerIndex int
	AnswerOptions      []string
}

func (q IrregularVerbs) Start(chatID int) error {
	return q.cache.Set(fmt.Sprintf("quiz.irregular_verbs.chat_id.%d", chatID), IrregularVerbsGameStruct{})
}

func (q IrregularVerbs) GetQuestion(chatID int) (IrregularVerbsGameStruct, error) {
	rand.Seed(time.Now().UnixNano())

	verbs, err := q.irregularVerbsGetter.GetIrregularVerbs()
	if err != nil {
		return IrregularVerbsGameStruct{}, err
	}

	question := verbs[rand.Intn(len(verbs))-1]

	answerOptions := []string{
		verbs[rand.Intn(len(verbs))-1].Translate,
		verbs[rand.Intn(len(verbs))-1].Translate,
		verbs[rand.Intn(len(verbs))-1].Translate,
		question.Translate,
	}

	rand.Shuffle(len(answerOptions), func(i, j int) { answerOptions[i], answerOptions[j] = answerOptions[j], answerOptions[i] })

	index := 0
	for k, a := range answerOptions {
		if question.Translate == a {
			index = k
			break
		}
	}

	questionOptions := []string{question.Infinitive, question.V2, question.V3}

	gameStruct := IrregularVerbsGameStruct{
		Verb:               question,
		CorrectAnswer:      question.Translate,
		AnswerOptions:      answerOptions,
		Question:           questionOptions[rand.Intn(len(questionOptions)-1)],
		CorrectAnswerIndex: index,
	}

	err = q.cache.Set(fmt.Sprintf("quiz.irregular_verbs.chat_id.%d", chatID), gameStruct)
	if err != nil {
		return IrregularVerbsGameStruct{}, err
	}

	return gameStruct, nil
}

func (q IrregularVerbs) End(chatID int) {
	q.cache.Del(fmt.Sprintf("quiz.irregular_verbs.chat_id.%d", chatID))
}
