package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"regexp"
	"strconv"
	"strings"

	//"context"
	//"encoding/csv"
	//"fmt"
	//
	//"golang.org/x/sync/errgroup"
	//"golang.org/x/sync/semaphore"
	//"log"
	//"os"
	//"runtime"
	"project1/model"
	h "project1/myHttp"
)

func GetInfoPage(pathURL string) (map[string]model.Film, error) {
	response, err := h.HttpClient.GetRequestWithRetries(pathURL)
	h.CheckError(err)
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	h.CheckError(err)
	infoList := make(map[string]model.Film)
	dem := 1
	doc.Find("table tbody").Each(func(index int, tableHtml *goquery.Selection) {

		tableHtml.Find("tr").Each(func(indexTr int, rowHtml *goquery.Selection) {

			row := make([]string, 0)
			rowHtml.Find("td").Each(func(ndexTd int, tableCell *goquery.Selection) {
				row = append(row, tableCell.Text())
			})
			var info model.Film

			info.Id = strconv.Itoa(dem)
			dot := strings.Index(row[1], ".")
			//strings.ReplaceAll(row[1],"/n","")

			//info.Name = row[1][dot+2:]
			rak := row[1][:dot]
			name := row[1][dot+1:]
			re := regexp.MustCompile(`\r?\n`)
			name = re.ReplaceAllString(name, " ")
			space := regexp.MustCompile(`\s+`)
			name = space.ReplaceAllString(name, " ")
			name = strings.Trim(name, " ")
			rak = re.ReplaceAllString(rak, " ")
			rak = space.ReplaceAllString(rak, "")
			fmt.Println("-----------")

			info.Rank = rak
			info.Name = name
			row[2] = re.ReplaceAllString(row[2], "")
			row[2] = space.ReplaceAllString(row[2], "")
			info.Rating = row[2]
			fmt.Println(row[2])

			infoList[info.Id] = info
			dem++
			//fmt.Println(info.Id," => ",infoList[info.Id])
		})
	})
	return infoList, nil
}
