package main

import (
	"strings"
	"testing"
)

const (
	full_repo_path   = "/hoge/fuga/github.com/user/repo"
	full_repo_bucket = "/hoge/fuga/bitbucket.org/user/repo"
	full_repo_other  = "/hoge/fuga/other.git/user/repo"
	user_repo        = "user/repo"
	domain_user_repo = "github.com/user/repo"
	valid_query      = "re"
	invalid_query    = "aaa"
	githublogo       = "./resources/github-logo.png"
	bucketlogo       = "./resources/bitbucket-logo.png"
	gitlogo          = "./resources/git-logo.png"
)

func TestExcludeDomain(t *testing.T) {
	repo_path := strings.Split(full_repo_path, "/")
	ur := excludeDomain(repo_path, true)
	if ur != user_repo {
		t.Errorf("Did:\texcludeDoamin(%s, '/')\nExpect:\t%s\nResult:\t%s\n", full_repo_path, user_repo, ur)
	}
	dur := excludeDomain(repo_path, false)
	if dur != domain_user_repo {
		t.Errorf("Did:\texcludeDoamin(%s, '/')\nExpect:\t%s\nResult:\t%s\n", full_repo_path, domain_user_repo, dur)
	}
}

func TestGetDomainName(t *testing.T) {
	repo_path := strings.Split(full_repo_path, "/")
	if !matchRepo(repo_path, valid_query) {
		t.Errorf("Could not match '%s' in '%s'. 'Match' is expected, but there is no match.\n", valid_query, full_repo_path)
	}
	if matchRepo(repo_path, invalid_query) {
		t.Errorf("Invalid match with query ('%s') in '%s'. ", invalid_query, full_repo_path)
	}
}

func TestGetIconPath(t *testing.T) {
	repo_path := strings.Split(full_repo_path, "/")
	if icon := getIconPath(repo_path); icon != githublogo {
		t.Errorf("Domain:\tgithub.com\nExpect:\t%s\nResult:\t%s\n", githublogo, icon)
	}
	bucket_path := strings.Split(full_repo_bucket, "/")
	if icon := getIconPath(bucket_path); icon != bucketlogo {
		t.Errorf("Domain:\tbitbucket.org\nExpect:\t%s\nResult:\t%s\n", bucketlogo, icon)
	}
	other_path := strings.Split(full_repo_other, "/")
	if icon := getIconPath(other_path); icon != gitlogo {
		t.Errorf("Domain:\tother.git\nExpect:\t%s\nResult:\t%s\n", gitlogo, icon)
	}
}
