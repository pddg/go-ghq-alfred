package main

import (
	"encoding/json"
	"fmt"
	model "github.com/pddg/alfred-models"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func main() {
	app := cli.NewApp()
	app.Name = "ghq-alfred"
	app.Usage = "Search your local repos"
	app.Version = "0.1.1"
	app.Action = func(c *cli.Context) error {
		resp := model.NewResponse()
		query := strings.Trim(c.Args()[0], " \n")
		repos := c.Args()[1:c.NArg()]
		for index, repo := range repos {
			repo_path := strings.Split(repo, "/")
			if matchRepo(repo_path, query) {
				// Create normal item
				item := createNewItem(index, repo, repo_path)
				resp.Items = append(resp.Items, *item)
			}
		}
		if resp.Items == nil {
			// When any item is not found.
			item := createNoResultItem()
			resp.Items = append(resp.Items, *item)
		}
		j, err := json.Marshal(resp)
		if err != nil {
			// Json error
			fmt.Println("{'items': [{'title': 'Json object is invalid.', 'subtitle': 'Please contact to developper.', 'valid': false}]")
		} else {
			fmt.Println(string(j))
		}
		return nil
	}
	app.Run(os.Args)
}

func createNewItem(index int, repo string, repo_path []string) *model.Item {
	item := model.NewItem()
	item.Uid = string(index)
	item.Title = excludeDomain(repo_path, true)
	item.Subtitle = getDomainName(repo_path)
	item.Arg = repo
	item.Icon.Type = ""
	item.Icon.Path = getIconPath(repo_path)
	createModItems(repo_path, repo, &item.Mods)
	return item
}

func createNoResultItem() *model.Item {
	item := model.NewItem()
	item.Title = "No result found."
	item.Subtitle = "Please input again."
	item.Valid = false
	return item
}

func matchRepo(repo_path []string, query string) bool {
	if strings.Index(excludeDomain(repo_path, true), query) != -1 {
		return true
	}
	return false
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

func createModItems(repo []string, path string, mods *map[string]model.Mod) {
	for _, key := range model.Modifiers {
		var (
			arg string
			sub string
		)
		switch key {
		case model.Cmd:
			arg = path
			sub = "Open '" + path + "' in Finder."
		case model.Shift:
			arg = "https://" + excludeDomain(repo, false) + "/"
			sub = "Open '" + arg + "' in browser."
		case model.Ctrl:
			arg = path
			sub = "Open '" + path + "' in editor."
		case model.Fn:
			arg = path
			sub = "Open '" + path + "' in terminal app."
		case model.Alt:
			arg = excludeDomain(repo, true)
			sub = "Search '" + arg + "' with google."
		}
		mod := model.NewMod(arg, sub)
		(*mods)[key] = *mod
	}
}

func getIconPath(repo_path []string) string {
	var icon_path string
	domain := getDomainName(repo_path)
	prefix := "./resources"
	switch {
	case strings.Contains(domain, "github"):
		icon_path = prefix + "/github-logo.png"
	case strings.Contains(domain, "bitbucket"):
		icon_path = prefix + "/bitbucket-logo.png"
	default:
		icon_path = prefix + "/git-logo.png"
	}
	return icon_path
}
