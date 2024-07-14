package i18n

type Error struct {
	lang string
	res  string
	t    *Translator
}

func NewError(lang string, res string) *Error {
	return &Error{
		lang: lang,
		res:  res,
		t:    NewTranslator(),
	}
}

func (e *Error) translate() string {
	return e.t.Translate(e.lang, e.res)
}

func (e *Error) Error() string {
	return e.translate()
}
