zenv
====
Powerful environments

##Description
`zenv` activates or deactivates your environment settings when you do `cd`.

##Features
You Can ...
* Activate or Deactivate Automatically
* Make Global or Local Environment
* Make Recursive Environment
* [WIP] Merge or Combine Environments Made By `zenv`
* [WIP] Run Commands When Activating or Deactivating

##Demo
Changing the version of python `2.7.10` to `3.4.3` in `~/tmp/example` directory.
![demo](https://cloud.githubusercontent.com/assets/8780513/9543051/946a5990-4db0-11e5-9c29-bdf4cb82333e.png)

##Install
First, install [Go](https://golang.org/doc/install).  
And, run the following commands.(replace `.bashrc` if needed.)  
    $ go get -u github.com/ryo33/zenv  
    $ go install github.com/ryo33/zenv  
    $ echo source \$GOPATH/src/github.com/ryo33/zenv/scripts/init.sh >> ~/.bashrc  
Finally, restart your shell.

##Usage
1. Initialize your local environment.  
`$ zenv local`  
2. Edit your environment settings.  
`$ zenv COMMAND ARGUMENT...`
3. Activate it!  
`$ cd YOUR_DIRECTORY`  

##Commands
When you want to edit global environments.  
`$ zenv --global NAME COMMAND ARGUMENT...`  
or  
`$ zenv -g NAME COMMAND ARGUMENT...`  
###local
Initialize local environment.  
`$ zenv local`  
Options:  
`--not-recursive`: not to activate this environment in child directories.  
`--exclusive`: not to activate environments of parent directories.  
[WIP]`--merge NAME` or `-m NAME`: see `merge`  
[WIP]`--load NAME` or `-l NAME`: see `load`  
###global
Initialize global environment.  
`$ zenv global NAME`  
Options:  
Same as `local`'s  
###link
Add executable files.  
`$ zenv link NAME EXECUTABLE`  
Remove them.  
`$ zenv link --remove NAME...`  
###[WIP]task
Add tasks you want to run when activating of deactivating.  
`$ zenv task SHELL`  
Option:  
`--both` or `-b`: when activating or deactivating **DEFAULT**  
`--activate` or `-a`: when activating  
`--deactivate` or `-d`: when deactivating  
###[WIP]git
**checkout**  
`$ zenv git checkout BRANCH`  
**config**  
`$ zenv git config CONFIG`  
###[WIP]merge
Merge environments into current environment.  
`$ zenv merge NAME...`  
Option of NAME:  
`--global` or `-g`: global environment **DEFAULT**  
`--local` or `-l`: local environment  
###[WIP]load
Activate of Deactivate environments at the same time  
`$ zenv load NAME...`  
Options:  
Same as `merge`'s  
