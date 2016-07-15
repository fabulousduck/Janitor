package janitor

import (
	"fmt"
	"io/ioutil"
	"os"
	p "path/filepath"
	"strings"
)

type listItem struct {
	value  string
	values []string
}

type janitor struct {
	current_task string
	busy         bool
	garbage_bag  []string
	list         map[string]listItem
	cmdArgs      []string
}

func NewJanitor() *janitor {
	if _, err := os.Stat("config.janitor"); err == nil {

	} else {
		fmt.Println("no config.janitor file found. making one for you.")
		file, _ := os.Create("config.janitor")
		defer file.Close()

	}

	list, err := ioutil.ReadFile("config.janitor")
	if err != nil {
		fmt.Println("Could not find config.janitor file")
		os.Exit(1)
	}
	mapped := FormatList(string(list))
	return &janitor{
		"none",
		false,
		[]string{},
		mapped,
		[]string{},
	}
}

func (janitor *janitor) findFiles(target string, args []string) {
	noignore := FlagGiven("-noignore", args)
	fs, err := ioutil.ReadDir(target)
	if err != nil {
		fmt.Println("invalid path : ", target)
	}

	for _, file := range fs {
		fp := strings.Join([]string{target, file.Name()}, "")
		fmt.Println(janitor.list)
		if noignore {
			if p.Ext(fp) != "" && file.Name()[0] != '.' {
				janitor.garbage_bag = append(janitor.garbage_bag, file.Name())
			}
		} else {
			if p.Ext(fp) != "" && !janitor.isIgnoring(p.Ext(fp)) && file.Name()[0] != '.' {
				janitor.garbage_bag = append(janitor.garbage_bag, file.Name())
			}
		}
	}
}

func (janitor *janitor) isIgnoring(ftype string) bool {
	if val, ok := janitor.list["janitor_ignore"]; ok {
		for i := 0; i < len(val.values); i++ {
			if val.values[i] == ftype {
				return true
			}
		}
		return false
	}
	return false
}

func sContains(slice []string, query string) bool {
	for i := 0; i < len(slice); i++ {
		if slice[i] == query {
			return true
		}
	}
	return false
}

func UserConfirm() bool {
	var resp int
	fmt.Scanf("%c", &resp)
	switch resp {
	case 'y':
		return true
	case 'n':
		return false
	default:
		return false
	}
}
