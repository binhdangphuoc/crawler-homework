package crawler

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	//"project1/model"
	h "project1/myHttp"
)

func GetProducts(pathURL string) {
	fmt.Println("begin Get products..............")
	response, err := h.HttpClient.GetRequestWithRetries(pathURL)
	h.CheckError(err)
	defer response.Body.Close()
	doc, err := goquery.NewDocumentFromReader(response.Body)
	h.CheckError(err)
	//infoList := make(map[string]model.Product)
	//dem :=1
	fmt.Println("Get products..............")
	doc.Find(".main-content .view #collection .container.collection-detail").Each(func(index int, html *goquery.Selection) {
		data := html.Text()
		fmt.Println("data: ", data)
	})
}
