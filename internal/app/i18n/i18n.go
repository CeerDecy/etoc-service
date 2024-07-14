package i18n

import (
	"os"

	"gopkg.in/yaml.v3"
)

const configPath = "/conf/i18n/i18n.yaml"

const LangEnglish = "en"
const LangChinese = "cn"

type Translator struct {
	def string
	en  map[string]string
	cn  map[string]string
}

func NewTranslator() *Translator {
	dir, _ := os.Getwd()
	file, _ := os.Open(dir + configPath)
	decoder := yaml.NewDecoder(file)

	dict := make(map[string]map[string]string)
	_ = decoder.Decode(&dict)

	return &Translator{
		def: LangEnglish,
		en:  dict[LangEnglish],
		cn:  dict[LangChinese],
	}
}

func (t *Translator) Translate(lang string, res string) string {
	switch lang {
	case "zh-CN":
		return t.cn[res]
	case "en":
		return t.en[res]
	default:
		return t.en[res]
	}
}
