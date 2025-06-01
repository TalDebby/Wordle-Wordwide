package service

import (
	"errors"
	"server/models"
	"testing"
)

type MockLanguagesRepository struct {
	languages models.LanguagesList
	words     map[string]models.Words
}

func (m *MockLanguagesRepository) GetWords(language string) (*models.Words, error) {
	if words, ok := m.words[language]; ok {
		return &words, nil
	}
	return nil, errors.New("not found")
}

func (m *MockLanguagesRepository) GetLanguages() (*models.LanguagesList, error) {
	return &m.languages, nil
}

func TestGetDailyWord(t *testing.T) {
	type args struct {
		language string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{name: "test daily", args: args{language: "English"}, want: "apple", wantErr: false},
		{name: "test daily not found", args: args{language: "Spanish"}, want: "", wantErr: true},
		{name: "test daily empty", args: args{language: ""}, want: "", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewLanguagesService(&MockLanguagesRepository{words: map[string]models.Words{"English": {"apple"}}}).GetDailyWord(tt.args.language)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetDailyWord() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GetDailyWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
