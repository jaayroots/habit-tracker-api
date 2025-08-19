package enums

type ExceptionType string

const (
	Debug   ExceptionType = "Debug"
	Info    ExceptionType = "Info"    // Logger openSearch
	Warning ExceptionType = "Warning" // Logger openSearch
	Error   ExceptionType = "Error"   // Logger openSearch and break
	Fatal   ExceptionType = "Fatal"   // Logger openSearch and break
)
