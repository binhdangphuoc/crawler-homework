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

/*
//Find number of page
func totalPage(url string) int {
	response, err := h.HttpClient.GetRequestWithRetries(url)
	h.CheckError(err)
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	h.CheckError(err)

	lastPageLink, _ := doc.Find("ul.pagination li:last-child a").Attr("href") // Đọc dữ liệu từ thẻ a của ul.pagination
	fmt.Println("STRINGFOUND: ",lastPageLink)
	split := strings.Split(lastPageLink, "/")[5]
	totalPages, _ := strconv.Atoi(split)
	fmt.Println("totalPage->", totalPages)
	return totalPages
}
*/

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

/*
func AllPage(url string) {
	fileName := "info_web_deface.csv"
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Could not create %s", fileName)
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	_=writer.Write([]string{"Attacker", "Country", "Web Url", "Ip", "Date"})

	sem := semaphore.NewWeighted(int64(runtime.NumCPU())) // Tạo ra số lượng group goroutines bằng 8 lần số luồng CPU, cùng đồng thời đi thu thập thông tin
	group, ctx := errgroup.WithContext(context.Background())
	var totalResults int = 0

	for page := 1; page <= 1; page ++ { // Lặp qua từng trang đã được phân trang
		pathURL := fmt.Sprintf("https://mirror-h.org/search/country/VN/pages/%d", page) // Tìm ra url của từng trang bằng cách nối chuỗi với số trang
		err := sem.Acquire(ctx, 1)
		if err != nil {
			fmt.Printf("Acquire err = %+v\n", err)
			continue
		}
		group.Go(func() error {
			defer sem.Release(1)

			// do work
			infoList, err := GetInfoPage(pathURL) // Thu thập thông tin web qua url của page
			if err != nil {
				log.Println(err)
			}
			totalResults += len(infoList)
			for _, info := range infoList {
				_ = writer.Write([]string{info.Attacker, info.Country, info.WebUrl, info.Ip, info.Date})
			}
			return nil
		})
	}
	if err := group.Wait(); err != nil { // Error Group chờ đợi các group goroutines done, nếu có lỗi thì trả về
		fmt.Printf("g.Wait() err = %+v\n", err)
	}
	fmt.Println("crawler done!")
	fmt.Println("total results:", totalResults)
}

*/
