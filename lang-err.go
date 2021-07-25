package validmor

const (
	ERR_ES = "Spanish"
	ERR_EN = "English"
)

var currentTypeErrLang = "English"

// change the language of the errors in Spanish or English depending on the project or what the user wants
func Errors(langErr string) {
	switch langErr {
	case ERR_EN:
		currentTypeErrLang = ERR_EN
	case ERR_ES:
		currentTypeErrLang = ERR_ES
	default:
		currentTypeErrLang = ERR_EN
	}
}
