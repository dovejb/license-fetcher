package utils

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func FetchForGo(pkg string) (license string, repo string, e error) {
	u := "https://pkg.go.dev/" + pkg
	doc, e := goquery.NewDocument(u)
	if e != nil {
		return
	}
	doc.Find(".go-Main-headerDetailItem a").Each(func(number int, s *goquery.Selection) {
		n := s.Nodes[0]
		ok := false
		for _, attr := range n.Attr {
			if attr.Key == "aria-label" && attr.Val == "Go to Licenses" {
				ok = true
				break
			}
		}
		if !ok {
			return
		}
		license = n.FirstChild.Data
	})
	doc.Find(".UnitMeta-repo a").Each(func(number int, s *goquery.Selection) {
		n := s.Nodes[0]
		repo = strings.Trim(n.FirstChild.Data, "\n\r \t")
	})
	return
}
