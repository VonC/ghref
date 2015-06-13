package main

import (
	"fmt"
	// "os"
	// "regexp"
	"strings"
	// "time"

	"github.com/VonC/godbg"
	"github.com/VonC/godbg/exit"
	"github.com/atotto/clipboard"
	"github.com/google/go-github/github"
	flag "github.com/spf13/pflag"
)

var client *github.Client
var ex *exit.Exit
var pdbg *godbg.Pdbg

var debug bool
var verbose bool
var help bool

func init() {
	ex = exit.Default()
	client = github.NewClient(nil)
	flag.BoolVarP(&help, "help", "h", false, "ghref usage")
	flag.BoolVarP(&verbose, "verbose", "v", false, `instead of just the SHA1, display a verbose output
		not suited for batch usage`)
	flag.BoolVarP(&debug, "debug", "d", false, "output debug informations (not for batch usage)")
}

func main() {
	flag.Parse()
	if help {
		usage()
		ex.Exit(0)
	}
	nargs := len(flag.Args())
	if nargs > 2 {
		fmt.Println("!! Too many arguments (only repo and ref names are expected) !!\n")
		usage()
		ex.Exit(1)
	}
	if nargs == 0 {
		fmt.Println("!! Too few arguments (at least the repo name is expected) !!\n")
		usage()
		ex.Exit(1)
	}

	ownerrepo := flag.Args()[0]
	or := strings.Split(ownerrepo, "/")
	if len(or) != 2 {
		fmt.Println("!! expects the first argument to be <owner>/<repo> !!\n")
		usage()
		ex.Exit(1)
	}
	owner := or[0]
	reponame := or[1]

	if !debug {
		pdbg = godbg.NewPdbg(godbg.OptExcludes([]string{"/ghref.go"}))
	} else {
		pdbg = godbg.NewPdbg()
		displayRateLimit()
	}

	repo, _, err := client.Repositories.Get(owner, reponame)
	if err != nil {
		if verbose {
			fmt.Printf("Unable to find repo '%s'\n", owner, reponame)
		}
		if debug {
			fmt.Println(err.Error())
		}
		ex.Exit(1)
	}

	res := ""
	if nargs == 2 {
		res = displayRef(repo, flag.Args()[1])
	} else {
		res = displayDefaultBranch(repo)
	}

	fmt.Println(res)
	if res != "" {
		clipboard.WriteAll(res)
	}
	if debug {
		fmt.Println("(Copied to the clipboard)")
		displayRateLimit()
	}
}

func usage() {
	fmt.Println(`ghref -hdv <owner>/<repo> [<branchname>|<tagname>|heads]",
      => refs from a github repo (owner/rename)

ghref displays, for a given GitHub repo, the SHA1 for a branch or a tag,
or it find the default branch (remote HEAD) when no parameter is given.

By default, the output is only the SHA1 or branch name, 
or nothing if the repo (or branch or tag) does not exist.
That way, the ghref output can be pass to another command (batch usage)

For example, to find the SHA1 of the *default* branch of a repo:

     ghref git/git    | xargs  ghref git/git
    (default branch)          (sha1 of that branch)

Note: it is also copied to the clipboard. (CTRL/Command+V to paste the result)

Parameters:

<owner>/<repo>: owner and repo name on github
<branchname>: name of the branch, to look for its SHA1
<tagname>: name of the tag, to look for its SHA1
heads: means that all branches SH1 will be displayed

Flags (posix)

--help, -h: display this usage and exit
--verbose, -v: if set, prints a detailed output (not for batch usage)
--debug, -d: if set, add debug information
`)
}

func displayRef(repo *github.Repository, refname string) string {
	ref, _, err := client.Git.GetRef(*repo.Owner.Login, *repo.Name, "heads/"+refname)
	if err != nil {
		if strings.Contains(err.Error(), "404 Not Found") {
			ref, _, err = client.Git.GetRef(*repo.Owner.Login, *repo.Name, "tags/"+refname)
		}
		if err != nil {
			if verbose {
				fmt.Printf("'%s' not found in heads or tags for repo '%s/%s' \n",
					refname, *repo.Owner.Login, *repo.Name)
			}
			if debug {
				fmt.Println(err.Error())
			}
			ex.Exit(1)
		}
	}
	res := *ref.Object.SHA
	if verbose {
		res = fmt.Sprintf("%s '%s' has for SHA1 %s", *ref.Object.Type, refname, res)
	}
	return res
}

func displayDefaultBranch(repo *github.Repository) string {
	res := *repo.DefaultBranch
	if verbose {
		res = fmt.Sprintf("The name of the default branch (remote HEAD) for %s/%s is %s",
			*repo.Owner.Login, *repo.Name, res)
	}
	return res
}

func displayRateLimit() {
	rate, _, err := client.RateLimits()
	if err != nil {
		fmt.Printf("Error fetching rate limit: %#v\n\n", err)
	} else {
		const layout = "15:04pm (MST)"
		tc := rate.Core.Reset.Time
		tcs := fmt.Sprintf("%s", tc.Format(layout))
		ts := rate.Search.Reset.Time
		tss := fmt.Sprintf("%s", ts.Format(layout))
		fmt.Printf("\nAPI Rate Core Limit: %d/%d (reset at %s) - Search Limit: %d/%d (reset at %s)\n",
			rate.Core.Remaining, rate.Core.Limit, tcs,
			rate.Search.Remaining, rate.Search.Limit, tss)
	}
}
