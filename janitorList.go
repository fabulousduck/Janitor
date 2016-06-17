package janitor

import (
	//"fmt"
	//reg "regexp"
	"strings"
)

func FormatList(list string) map[string]listItem {
	rMap := map[string]listItem{}
	IndivItems := strings.Split(list, "\n")

	for i := 0; i < len(IndivItems)-1; i++ {

		mapAddition := strings.Split(IndivItems[i], "=")
		var_name := mapAddition[0]
		var_value := mapAddition[1]

		if isList(var_value) {
			rMap[string(var_name)] = listItem{"", ul(var_value)}

		} else {
			//fmt.Println(mapAddition[1])
			rMap[string(var_name)] = listItem{var_value, []string{}}
		}
	}
	//fmt.Println(rMap)
	return rMap
}

func isList(uf string) bool {
	if string(uf[0]) != "{" || string(uf[len(uf)-1]) != "}" {
		return false
	}
	return true
}

func ul(l string) []string {
	f1 := strings.Replace(l, "{", "", -1)
	f2 := strings.Replace(f1, "}", "", -1)
	items := strings.Split(f2, ",")
	return items
}
