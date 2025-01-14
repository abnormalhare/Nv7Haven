package types

import (
	"text/template"
        "sync"
)

type translations struct {
	nameMayNotContain    string
	nameCannotBeEmpty    string
	successfullyUpdated  string
	playChannelReset     string
	youAreNotAuthorized  string
	successUpdateChannel string
}

type Language struct {
	name         string
	translations translations
}

type Variables struct {
	var1 interface{}
	var2 interface{}
	var3 interface{}
}

var LanguageTable = map[string]any{
        EnglishLock: sync.RWMutex,
	English: Language{
		name: "English",
		translations: translations{
			nameMayNotContain:    `A name may not contain '{{.var1}}'!`,
			nameCannotBeEmpty:    "Name cannot be empty!",
			successfullyUpdated:  `Successfully updated {{.var1}}!`,
			playChannelReset:     "**PLAY CHANNELS HAVE BEEN RESET**\nUpdate them below!",
			youAreNotAuthorized:  "You are not authorized to use this!",
			successUpdateChannel: "Successfully updated play channels!",
		},
	},
},}
