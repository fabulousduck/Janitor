#Janitor [![Build Status](https://travis-ci.org/fabulousduck/Janitor.svg?branch=master)](https://travis-ci.org/fabulousduck/Janitor)

A Cleaning tool for your directories.

## Description 
  
  Janitor is a Cleaning tool for directories.
  It sorts files into individual folders that are named after the type of file in them.
  
  It also makes use of a config.janitor file in which you can specify certain variables that Janitor can look at.
  
  You can also use flags with the commands now. the ones currently availible are listed below
  
## Availible flags

### Clean
  
| Flag | function |Currently usable |
| :--: | :------: | :-------------: |
| -defaultDir | uses the user predefined default directory to clean | yes |
| -noIgnore   | sort the user predefined ignored file types         | no  |

## Available config.janitor variables
  
  You can store any varibles you'd like in the config.Janitor file, but there are only a few that actually do something at the moment.
  
  The ones that do do something at the moment are listed below, along with examples of how to use them.
  
  General syntax : ` janitor_Myvariable=MyValue `
  
### variables
  
| Name | Example |
| :--: | :-----: |
| janitor\_ignore | `janitor_ignore={.txt,.png}` |
| janitor\_defaultDir | `janitor_desktopPath=/Users/ryanvlaming/Desktop`|


## functions

  At the moment Janitor only has a cleaning function for putting stray files into folders.
  
  The path to the directory must be specified fully from root. ( `/Users/ryanvlaming/whereever` )
### example 

```go
func main(){
  var myJanitor = janitor.NewJanitor();
    
   myJanitor.CleanDir("DirectoryToBeCleaned");
    
  }
```

## future todo's

- [x] Implement basic flags
- [ ] Implement a readline
- [ ] Implement folder sorting
- [ ] Allow for clean/sort by file name
- [ ] Error logging for config.janitor file
- [x] Automatically generate config.janitor file if it doesnt exist and on install
