package translate

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

// Translator
type Translator struct {
	IsInit    bool // check if translator initialized or not
	bundle    *i18n.Bundle
	localizer *i18n.Localizer
}

// func (t *Translator) InitTranslator(initLang string) {
//
// 	t.bundle = i18n.NewBundle(language.English)
// 	t.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)
// 	t.bundle.MustLoadMessageFile("./lang/en.json")
// 	t.bundle.MustLoadMessageFile("./lang/zh.json")
// 	t.bundle.MustLoadMessageFile("./lang/ar.json")
// 	t.localizer = i18n.NewLocalizer(t.bundle, initLang) // should add additional check here
// 	t.IsInit = true
// 	// end of adding
// }

// InitTranslator added to support multi-language
func (t *Translator) InitTranslator(langPath string) {

	langFiles, err := os.ReadDir(langPath)
	if err != nil {
		fmt.Printf("InitTranslator() Error: %v", err)
	}

	t.bundle = i18n.NewBundle(language.English)
	t.bundle.RegisterUnmarshalFunc("json", json.Unmarshal)

	for _, langFile := range langFiles {
		if !langFile.IsDir() {
			langFilePath := langPath + langFile.Name()
			fmt.Printf("Language file: %s loaded.\r\n", langFilePath)
			t.bundle.MustLoadMessageFile(langFilePath)
		}
	}

	t.localizer = i18n.NewLocalizer(t.bundle, language.English.String(), language.Arabic.String()) // should add additional check here
	t.IsInit = true
}

// SetTranslatorLanguage set translator language
func (t *Translator) SetTranslatorLanguage(lang string, accept string) {
	t.localizer = i18n.NewLocalizer(t.bundle, lang, accept)
}

// TranslateMessage translate msg
func (t *Translator) TranslateMessage(messageID string) string {
	// translatedMsg, err := t.localizer.LocalizeMessage(&i18n.Message{ID: messageID})
	// if err != nil {
	// 	return messageID
	// }
	// // return t.localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
	// return translatedMsg

	localizeConfig := i18n.LocalizeConfig{
		MessageID: messageID,
	}

	localization, err := t.localizer.Localize(&localizeConfig)
	if err != nil {
		log.Println("error in localization: ", err)
	}

	return localization
}
