package utils

import "testing"

func TestGithubMainPage(t *testing.T) {
	parser, e := NewParser("https://github.com/ncw/swift")
	if e != nil {
		t.Error(e)
	}
	parser.TryGetLicenseFromGithubMainPage()
}
