package language

import (
	"encoding/csv"
	"os"

	"github.com/pkg/errors"
)

type IrregularVerb struct {
	Infinitive string
	V2         string
	V3         string
	Translate  string
}

type IrregularVerbs struct{}

func NewIrregularVerbsRepository() *IrregularVerbs {
	return &IrregularVerbs{}
}

func (i IrregularVerbs) GetIrregularVerbs() ([]IrregularVerb, error) {
	fileName := "data/irregular_verbs.csv"
	f, err := os.Open(fileName)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read input file "+fileName)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse file as CSV for "+fileName)
	}

	irregularVerbs := make([]IrregularVerb, 0, len(records))
	for _, line := range records {
		irregularVerbs = append(irregularVerbs, IrregularVerb{
			Infinitive: line[0],
			V2:         line[1],
			V3:         line[2],
			Translate:  line[3],
		})
	}

	return irregularVerbs, nil
}
