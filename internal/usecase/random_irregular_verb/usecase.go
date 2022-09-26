package random_irregular_verb

import (
	"math/rand"
	"time"

	"githab.com/oleglacto/translater/internal/pkg/language"
)

type IrregularVerbsGetter interface {
	GetIrregularVerbs() ([]language.IrregularVerb, error)
}

type UseCase struct {
	irregularVerbsGetter IrregularVerbsGetter
}

func New(irregularVerbsGetter IrregularVerbsGetter) *UseCase {
	return &UseCase{irregularVerbsGetter: irregularVerbsGetter}
}

func (u UseCase) GetRandomVerb() (language.IrregularVerb, error) {
	verbs, err := u.irregularVerbsGetter.GetIrregularVerbs()
	if err != nil {
		return language.IrregularVerb{}, err
	}

	rand.Seed(time.Now().UnixNano())

	return verbs[rand.Intn(len(verbs))-1], nil
}
