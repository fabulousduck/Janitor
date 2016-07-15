#Janitor [![Build Status](https://travis-ci.org/fabulousduck/Janitor.svg?branch=master)](https://travis-ci.org/fabulousduck/Janitor)

A Cleaning tool for your directories.

## Description 
  
  Janitor is a Cleaning tool for directories.
  It sorts files into individual folders that are named after the type of file in them.
  
  It also makes use of a config.janitor file in which you can specify certain variables that Janitor can look at.
  
  You can also use flags with the commands now. the ones currently availible are listed below
  
## Available flags/commands

### Clean
  
| Flag | function |Currently usable |
| :--: | :------: | :-------------: |
| -defaultdir | uses the user predefined default directory to clean | yes |
| -noIgnore   | sort the user predefined ignored file types         | yes |
| revert      | reverts the last cleanup                            | no  |

### install

| Flag | function |Currently usable |
| :--: | :------: | :-------------: |
| -preconfig | install janitor with  base config.janitor variables | no |
| -noconfig  | install without the config.janitor files (not recommended) | no |

## Available config.janitor variables
  
  You can store any varibles you'd like in the config.Janitor file, but there are only a few that actually do something at the moment.
  
  The ones that do do something at the moment are listed below, along with examples of how to use them.
  
  General syntax : ` janitor_Myvariable=MyValue `
  
### variables
  
| Name | Example | Available| 
| :--: | :-----: | :------: |
| janitor\_ignore | `janitor_ignore={.txt,.png}` | yes |
| janitor\_defaultDir | `janitor_desktopPath=/Users/ryanvlaming/Desktop` | yes |
| janitor\_defaultDump | `janitor_defaultDump=/Users/ryanvlaming/Desktop/dump` | no |




## future todo's

- [x] Implement basic flags
- [x] Implement a readline
- [ ] Implement folder sorting
- [ ] Allow for clean/sort by file name
- [ ] Error logging for config.janitor file
- [x] Automatically generate config.janitor file if it doesnt exist and on install
