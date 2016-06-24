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

func (janitor *janitor) CleanDir(args []string, usedef bool) bool {
	if usedef {

		janitor.findFiles(janitor.list["janitor_defaultDir"].value)
	} else {
		janitor.findFiles(args[1])
	}
	fmt.Println("do you really want to clean up ", len(janitor.garbage_bag), " files ? [y/n]")

	if UserConfirm() {

		types := []string{}

		for i := 0; i < len(janitor.garbage_bag); i++ {

			if !sContains(types, p.Ext(janitor.garbage_bag[i])) {

				types = append(types, p.Ext(janitor.garbage_bag[i]))

			}

		}

		//makes the folders for the files by type.
		if usedef {
			for i := 0; i < len(types); i++ {

				fPath := strings.Join([]string{strings.Join([]string{janitor.list["janitor_defaultDir"].value, "/"}, ""), strings.Split(types[i], ".")[1]}, "")
				err := os.Mkdir(fPath, 0700)
				fmt.Println(fPath)
				if os.IsExist(err) {
					err = nil
				}
				if err != nil {
					//do nothing
				}

			}

		} else {
			for i := 0; i < len(types); i++ {

				fPath := strings.Join([]string{strings.Join([]string{args[1], "/"}, ""), strings.Split(types[i], ".")[1]}, "")
				err := os.Mkdir(fPath, 0700)
				fmt.Println(fPath)
				if os.IsExist(err) {
					err = nil
				}
				if err != nil {
					//do nothing
				}

			}
		}
		//actually move the files
		if usedef {
			for i := 0; i < len(janitor.garbage_bag); i++ {
				ext := strings.Split(p.Ext(janitor.garbage_bag[i]), ".")[1]
				sPath := strings.Join([]string{janitor.list["janitor_defaultDir"].value, "/"}, "")
				nPath := strings.Join([]string{sPath, ext}, "")
				err := os.Rename(strings.Join([]string{janitor.list["janitor_defaultDir"].value, janitor.garbage_bag[i]}, "/"), strings.Join([]string{nPath, janitor.garbage_bag[i]}, "/"))
				if err != nil {
					fmt.Println(err)
				}
			}

		} else {
			for i := 0; i < len(janitor.garbage_bag); i++ {
				ext := strings.Split(p.Ext(janitor.garbage_bag[i]), ".")[1]
				sPath := strings.Join([]string{args[1], "/"}, "")
				nPath := strings.Join([]string{sPath, ext}, "")
				err := os.Rename(strings.Join([]string{args[1], janitor.garbage_bag[i]}, "/"), strings.Join([]string{nPath, janitor.garbage_bag[i]}, "/"))
				if err != nil {
					fmt.Println(err)
				}
			}

		}
	}

	return true
}

func (janitor *janitor) findFiles(target string) {
	fs, err := ioutil.ReadDir(target)
	if err != nil {
		fmt.Println("invalid path : ", target)
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
