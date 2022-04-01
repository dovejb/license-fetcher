package utils

import (
	"io"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Parser struct {
	doc *goquery.Document
}

func NewParser(url string) (*Parser, error) {
	doc, e := goquery.NewDocument(url)
	if e != nil {
		return nil, e
	}
	return &Parser{
		doc: doc,
	}, nil
}

func NewParserFromReader(r io.Reader) (*Parser, error) {
	doc, e := goquery.NewDocumentFromReader(r)
	if e != nil {
		return nil, e
	}
	return &Parser{
		doc: doc,
	}, nil
}

func (p Parser) TryGetLicenseFromGithubMainPage() (found bool, license string, licenseUrl string) {
	p.doc.Find(".mt-2 a").Each(func(number int, s *goquery.Selection) {
		a := s.Nodes[0]
		isTarget := false
		for _, attr := range a.Attr {
			if attr.Key == "href" && strings.Contains(strings.ToUpper(attr.Val), "LICENSE") {
				isTarget = true
				licenseUrl = "https://github.com" + attr.Val
			}
		}
		if !isTarget {
			return
		}
		data := strings.Trim(a.LastChild.Data, "\n \t\r")
		if !strings.HasPrefix(strings.ToLower(data), "view") {
			license = data
			found = true
		}
	})
	return
}
