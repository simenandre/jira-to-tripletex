# Copy Jira issues to Tripletex CLI

This utility is written in Go and features functionality to copy Jira issues to Tripletex, a Norwegian accounting software.

## Install

```shell
brew install cobraz/tools/jira-to-tripletex
```

**Notes**: The library is not tested on Linux or Windows. There are [executables available](https://github.com/cobraz/jira-to-tripletex/releases/latest) at every release >1.0.3

### Using Snap?

```shell
sudo snap install jira-to-tripletex
```

## Help

```shell
> jira-to-tripletex --help
NAME:
   Jira-To-Tripletex - A new cli application

USAGE:
   main [global options] command [command options] [arguments...]

DESCRIPTION:
   Transport Jira issues to Tripletex

COMMANDS:
   get:config, gc   
   set:config, cnf  
   list:activities  
   copy, c          
   help, h          Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h  show help (default: false)
```

## Contribute

Please, oh pretty please do contribute! If you feel this helps you out, but you want to increase the quality of this software, please submit pull requests. 🎉
