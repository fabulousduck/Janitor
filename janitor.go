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
}

func NewJanitor() *janitor {
	file, er := os.Create("config.janitor")
	if os.IsExist(er) {
		er = nil
	}
	if er != nil {
		//do nothing
	}

	defer file.Close()

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
	}
}

func (janitor *janitor) CleanDir(dir string) bool {

	janitor.findFiles(dir)
	fmt.Println("do you really want to clean up ", len(janitor.garbage_bag), " files ? [y/n]")

	if userConfirm() {

		types := []string{}

		for i := 0; i < len(janitor.garbage_bag); i++ {

			if !sContains(types, p.Ext(janitor.garbage_bag[i])) {

				types = append(types, p.Ext(janitor.garbage_bag[i]))

			}

		}

		//makes the folders for the files by type.

		for i := 0; i < len(types); i++ {

			fPath := strings.Join([]string{strings.Join([]string{dir, "/"}, ""), strings.Split(types[i], ".")[1]}, "")
			err := os.Mkdir(fPath, 0700)

			if os.IsExist(err) {
				err = nil
			}
			if err != nil {
				//do nothing
			}

		}

		//actually move the files

		for i := 0; i < len(janitor.garbage_bag); i++ {
			ext := strings.Split(p.Ext(janitor.garbage_bag[i]), ".")[1]
			sPath := strings.Join([]string{dir, "/"}, "")
			nPath := strings.Join([]string{sPath, ext}, "")
			err := os.Rename(strings.Join([]string{dir, janitor.garbage_bag[i]}, "/"), strings.Join([]string{nPath, janitor.garbage_bag[i]}, "/"))
			if err != nil {
				fmt.Println(err)
			}
		}

	}

	return true
}

func (janitor *janitor) findFiles(target string) {
	fs, err := ioutil.ReadDir(target)
	if err != nil {
		panic(err)
	}
	for _, file := range fs {
		fp := strings.Join([]string{target, file.Name()}, "")
		if p.Ext(fp) != "" && !janitor.isIgnoring(p.Ext(fp)) && file.Name()[0] != '.' {
			janitor.garbage_bag = append(janitor.garbage_bag, file.Name())
		}
	}
}

func (janitor *janitor) isIgnoring(ftype string) bool {
	if val, ok := janitor.list["janitor_ignore"]; ok {
		for i := 0; i < len(val.values)-1; i++ {
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

func userConfirm() bool {
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
