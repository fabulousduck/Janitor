package janitor

import (
	"fmt"
	"os"
	p "path/filepath"
	"strings"
)

func (janitor *janitor) CleanDir(args []string) bool {
	usedef := FlagGiven("-defaultdir", args)
	if usedef {
		janitor.findFiles(janitor.list["janitor_defaultDir"].value, args)
	} else {
		janitor.findFiles(args[1], args)
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
				var fPath string
				if val, ok := janitor.list["defaultDump"]; ok {
					fPath = val.value
				} else {
					fPath = strings.Join([]string{strings.Join([]string{janitor.list["janitor_defaultDir"].value, "/"}, ""), strings.Split(types[i], ".")[1]}, "")
				}
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
