package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"server/models"
)

type LanguagesRepository interface {
	GetWords(language string) (*models.Words, error)
	GetLanguages() (*models.LanguagesList, error)
}

type JsonLanguagesRepository struct {
	languagesFileLocation string
	wordsFileLocation     string
}

func NewJsonWordsRepository() *JsonLanguagesRepository {
	return &JsonLanguagesRepository{wordsFileLocation: "./data/words.json", languagesFileLocation: "./data/languages.json"}
}

func (r *JsonLanguagesRepository) GetWords(language string) (*models.Words, error) {
	jsonFile, err := os.Open(r.wordsFileLocation)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("fetch error")
	}

	fmt.Println("successfuly Opened words.json")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var languages models.Languages
	if err := json.Unmarshal(byteValue, &languages); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("fetch error")
	}

	for index := range languages.Languages {
		if languages.Languages[index].Language == language {
			return &languages.Languages[index].Words, nil
		}
	}

	return nil, errors.New("not found")
}

func (r *JsonLanguagesRepository) GetLanguages() (*models.LanguagesList, error) {
	jsonFile, err := os.Open(r.languagesFileLocation)

	if err != nil {
		fmt.Println(err)
		return nil, errors.New("fetch error")
	}

	fmt.Println("successfuly Opened languages.json")

	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var languages models.LanguagesList
	if err := json.Unmarshal(byteValue, &languages); err != nil {
		fmt.Println(err.Error())
		return nil, errors.New("fetch error")
	}

	return &languages, nil
}
