package output

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
)

func PrintError(value any) {
	intVal, ok := value.(int)
	if ok {
		color.Red("Код ошибки: %d", intVal)
		return
	}
	strVal, ok := value.(string)
	if ok {
		color.Red(strVal)
		return
	}
	errVal, ok := value.(error)
	if ok {
		color.Red(errVal.Error())
		return
	}
	color.Red("Неизвестный тип ошибки")
	// switch t := value.(type) {
	// case string:
	// 	color.Red(t)
	// case int:
	// 	color.Red("Код ошибки: %d", t)
	// case error:
	// 	color.Red(t.Error())
	// default:
	// 	color.Red("Неизвестный тип ошибки")
	// }
}

func PrintJson(data any) {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return
	}
	fmt.Println(string(bytes))
}
