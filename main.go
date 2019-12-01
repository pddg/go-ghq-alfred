package main

import (
	"os"
	"path"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/urfave/cli"
)

const (
	appName    = "ghq-alfred"
	appDesc    = "Search your local repos"
	appVersion = "0.3.1"
)

var (
	wf            *aw.Workflow
	gitHubIcon    = &aw.Icon{Value: path.Join("github-logo.png")}
	bitBucketIcon = &aw.Icon{Value: path.Join("bitbucket-logo.png")}
	gitIcon       = &aw.Icon{Value: path.Join("git-logo.png")}
	modKeys       = []aw.ModKey{
		aw.ModCmd,
		aw.ModOpt,
		aw.ModFn,
		aw.ModCtrl,
		aw.ModShift,
	}
)

func init() {
	wf = aw.New()
}

func run() {
	app := cli.NewApp()
	app.Name = appName
	app.Usage = appDesc
	app.Version = appVersion
	app.Action = func(c *cli.Context) error {
		query := strings.Trim(c.Args()[0], " \n")
		repos := c.Args()[1:c.NArg()]
		for _, repo := range repos {
			addNewItem(repo)
		}
		if len(query) > 0 {
			wf.Filter(query)
		}
		wf.WarnEmpty("No matching repository", "Try different query?")
		wf.SendFeedback()
		return nil
	}
	app.Run(os.Args)
}

func main() {
	wf.Run(run)
}

func addNewItem(repo string) {
	repoPath := strings.Split(repo, "/")
	it := wf.NewItem(repo).
		Title(excludeDomain(repoPath, true)).
		UID(repo).
		Arg(repo).
		Subtitle(getDomainName(repoPath)).
		Icon(getIcon(repoPath)).
		Valid(true)
	for _, modKey := range modKeys {
		mod := createModItem(repoPath, repo, modKey)
		it.SetModifier(mod)
	}
}

func excludeDomain(repo []string, domain bool) string {
	// full_repo_path: strings.Split("/path/to/github.com/user/full_repo_path", "/")
	var i int
	if domain {
		// return user/full_repo_path
		i = 2
	} else {
		// return github.com/user/full_repo_path
		i = 3
	}
	length := len(repo)
	return strings.Join(repo[length-i:length], "/")
}

func getDomainName(repo_path []string) string {
	// return github.com
	return repo_path[len(repo_path)-3]
}

func createModItem(repo []string, path string, modKey aw.ModKey) *aw.Modifier {
	var (
		arg string
		sub string
	)
	switch modKey {
	case aw.ModCmd:
		arg = path
		sub = "Open in Finder."
	case aw.ModShift:
		arg = "https://" + excludeDomain(repo, false) + "/"
		sub = "Open '" + arg + "' in browser."
	case aw.ModCtrl:
		arg = path
		sub = "Open in editor."
	case aw.ModFn:
		arg = path
		sub = "Open in terminal app."
	case aw.ModOpt:
		arg = excludeDomain(repo, true)
		sub = "Search '" + arg + "' with google."
	}
	mod := &aw.Modifier{Key: modKey}
	return mod.
		Arg(arg).
		Subtitle(sub).
		Valid(true)
}

func getIcon(repoPath []string) *aw.Icon {
	domain := getDomainName(repoPath)
	switch {
	case strings.Contains(domain, "github"):
		return gitHubIcon
	case strings.Contains(domain, "bitbucket"):
		return bitBucketIcon
	default:
		return gitIcon
	}
}
