package service

import (
	"errors"
	"fmt"
	"server/crud"
	"server/models"
	"time"
)

type LanguagesService struct {
	LanguagesRepository crud.LanguagesRepository
}

func NewLanguagesService(wordsRepository crud.LanguagesRepository) *LanguagesService {
	return &LanguagesService{LanguagesRepository: wordsRepository}
}

func (languagesService *LanguagesService) GetLanguages() (*models.LanguagesList, error) {
	return languagesService.LanguagesRepository.GetLanguages()
}

func (languagesService *LanguagesService) GetWords(language string) (*models.Words, error) {
	return languagesService.LanguagesRepository.GetWords(language)
}

func (languagesService *LanguagesService) GetDailyWord(language string) (string, error) {
	words, err := languagesService.LanguagesRepository.GetWords(language)
	if err != nil {
		fmt.Println("couldn't get words")
		return "", err
	}

	if len(*words) == 0 {
		return "", errors.New("not found")
	}

	currentDate := time.Now()

	magicNumber := currentDate.YearDay() % len(*words)

	return (*words)[magicNumber], nil
}
