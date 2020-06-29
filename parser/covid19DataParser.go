package parser

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/ambigudus/covid19-china-Go/model"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

type Covid19DataParser interface {
	DownloadAndParse(url string) (model.Covid19StatMap, error)
	ParseFromReader(r io.Reader) (model.Covid19StatMap, error)
}

func NewCovid19DataParser() Covid19DataParser {
	return &covid19DataParser{
		data:              model.Covid19StatMap{},
		currentParsedData: model.Covid19Stat{},
	}
}

type covid19DataParser struct {
	data              model.Covid19StatMap
	currentParsedData model.Covid19Stat
}

func (c *covid19DataParser) DownloadAndParse(url string) (model.Covid19StatMap, error) {
	resp, err := http.Get(url)
	if err != nil {
		return c.data, err
	}

	return c.ParseFromReader(resp.Body)
}

func (c *covid19DataParser) ParseFromReader(r io.Reader) (model.Covid19StatMap, error) {
	html, err := goquery.NewDocumentFromReader(r)
	tableHtml := html.Find("#statisticContainer .statisticChart > table > tbody")

	tr := tableHtml.Find("tr")
	trLen := tr.Length()

	fmt.Println("Total tr: ", trLen)
	tr = tr.Next()
	for i := 1; i < trLen; i++ {
		//
		//if c.currentParsedData.StateName == "Total" {
		//	tr = tr.Next()
		//	continue
		//}
		c.processTRSelection(i, tr)
		tr = tr.Next()
		//I know it's BAD. but I've no other workaround. :| ...
		// 最后一个省之后跟着 待确认

	}

	return c.data, err
}

func (c *covid19DataParser) processTRSelection(index int, selection *goquery.Selection) model.Covid19Stat {
	c.currentParsedData = model.Covid19Stat{}
	selection.Find("td").Each(c.processTDSelection)

	c.data[c.currentParsedData.StateName] = c.currentParsedData

	return c.currentParsedData

}

func (c *covid19DataParser) processTDSelection(index int, selection *goquery.Selection) {

	re := regexp.MustCompile(`[*#]`)
	text := re.ReplaceAllString(selection.Text(), "")
	text = strings.ReplaceAll(text, ",", "")

	switch index {
	case 0:
		c.currentParsedData.StateName = strings.ReplaceAll(text, "\n", "")
		break
	case 1:
		break
	case 2:
		i, _ := strconv.Atoi(text)
		c.currentParsedData.ConfirmedCase = i
		break
	case 3:
		i, _ := strconv.Atoi(text)
		c.currentParsedData.Death = i
		break
	case 4:
		i, _ := strconv.Atoi(text)
		c.currentParsedData.Cured = i
		break
	}
}
