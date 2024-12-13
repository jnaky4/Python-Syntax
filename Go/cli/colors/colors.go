package colors

import (
	tc "Go/const/terminalColors"
	"fmt"
)

func SetColor(value interface{}, color tc.TerminalColor) string {
	return fmt.Sprint(color, value, tc.Reset)
}
