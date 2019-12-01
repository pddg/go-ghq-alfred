package main

import (
	aw "github.com/deanishe/awgo"
	"strings"
	"testing"
)

const (
	full_repo_path      = "/hoge/fuga/github.com/user/repo"
	full_repo_bucket    = "/hoge/fuga/bitbucket.org/user/repo"
	full_repo_other     = "/hoge/fuga/other.git/user/repo"
	user_repo           = "user/repo"
	github_user_repo    = "github.com/user/repo"
	bitbucket_user_repo = "bitbucket.org/user/repo"
)

func TestExcludeDomain(t *testing.T) {
	testcases := []struct {
		path     string
		expected string
		domain   bool
	}{
		{
			full_repo_path,
			user_repo,
			true,
		}, {
			full_repo_path,
			github_user_repo,
			false,
		}, {
			full_repo_bucket,
			user_repo,
			true,
		}, {
			full_repo_bucket,
			bitbucket_user_repo,
			false,
		},
	}
	for _, tc := range testcases {
		repoPath := strings.Split(tc.path, "/")
		actual := excludeDomain(repoPath, tc.domain)
		if actual != tc.expected {
			t.Errorf("%s is expected, but actual %s\n", tc.expected, actual)
		}
	}
}

func TestGetDomainName(t *testing.T) {
	testcases := []struct {
		path   string
		domain string
	}{
		{
			full_repo_path,
			"github.com",
		}, {
			full_repo_bucket,
			"bitbucket.org",
		}, {
			full_repo_other,
			"other.git",
		},
	}
	for _, tc := range testcases {
		repoPath := strings.Split(tc.path, "/")
		if actual := getDomainName(repoPath); actual != tc.domain {
			t.Errorf("%s is expected, but actual %s\n", tc.domain, actual)
		}
	}
}

func TestGetIconPath(t *testing.T) {
	testcases := []struct {
		path string
		icon *aw.Icon
	}{
		{
			full_repo_path,
			gitHubIcon,
		}, {
			full_repo_bucket,
			bitBucketIcon,
		}, {
			full_repo_other,
			gitIcon,
		},
	}
	for _, tc := range testcases {
		repoPath := strings.Split(tc.path, "/")
		if icon := getIcon(repoPath); icon != tc.icon {
			t.Errorf("Expect: %s\nResult: %s\n", tc.icon, icon)
		}
	}
}
