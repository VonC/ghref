# ghref
SHA1 for branch or tag of a GitHub repo

## Context

[Stack Overflow question 30815555](http://stackoverflow.com/q/30815555/6309):
"**Get the hash of the current HEAD of a repo using only HTTP**"

## Implementation

`ghref` uses only https calls to the GitHub API 
through the [google/go-github](https://github.com/google/go-github) Go library
which only calls URLs from the [GitHub V3 API](https://developer.github.com/v3/).

It is a simple executable, compiled from sources in Go: no dll or dynamic library dependency.  
Copy the compiled executable anywhere you want and use it.

## Usage

````
ghref -hdv <owner>/<repo> [<branchname>|<tagname>|heads]",
      => refs from a github repo (owner/rename)
````

`ghref` displays, for a given GitHub repo, the SHA1 for a branch or a tag,
or it find the default branch (remote HEAD) when no parameter is given.

By default, the output is only the SHA1 or branch name, 
or nothing if the repo (or branch or tag) does not exist.  
That way, the `ghref` output can be passed to another command (batch usage)

For example, to find the SHA1 of the *default* branch of a repo:
````
     ghref git/git    | xargs  ghref git/git
    (default branch)          (sha1 of that branch)
````
Note: it is also copied to the clipboard. (`CTRL`/`Command`+`V` to paste the result)

### Parameters:

- `<owner>/<repo>`: owner and repo name on github
- `<branchname>`: name of the branch, to look for its SHA1
- `<tagname>`: name of the tag, to look for its SHA1

### Flags (posix)

- `--help`, `-h`: display this usage and exit
- `--verbose`, `-v`: if set, prints a detailed output (not for batch usage)
- `--debug`, `-d`: if set, add debug information

## Installation

- get go: [download go](https://golang.org/dl/), unzip it and set `GOROOT` to its path
- add `$GOROOT/bin` to your `$PATH`
- then:
````
git clone https://github.com/VonC/ghref
cd ghref
git submodule update --init
./gb.sh
````

- test it with:

        bin/ghref git/git | xargs bin/ghref git/git
        
        

