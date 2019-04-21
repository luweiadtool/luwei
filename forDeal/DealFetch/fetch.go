package DealFetch

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goDemo/Project/Luwei/config"
	"goDemo/Project/forDeal/model"
	"goDemo/Project/forDeal/tool"
	"myTool/common"
	"myTool/file"
	"net/http"
	"os"
	"path/filepath"
	"sort"
)

var pageLimit = 1000

var bufferSize = 1000
//http://m.fordeal.com/itemlist/70001061/1
//              http://m.fordeal.com/itemlist/70001061/1
var ItemList = "http://m.fordeal.com/itemlist/%v/%v"

var ItemDetail = "http://m.fordeal.com/detail/%v"

func Fetch(catID int64, filep string, con config.ForDealCon) {

	var result model.DownItems
	for i := 1; i <= pageLimit; i ++ {

		url := fmt.Sprintf(ItemList, catID, i)

		buf, err := tool.HttpGetRequest(url)

		if err != nil {
			continue
		}

		if len(buf) == 0 {
			break
		}

		var items model.ItemList
		err = json.Unmarshal(buf, &items)
		if err != nil {
			continue
		}
		for _, item := range items {

			imageName := fmt.Sprintf("%v-%v.%v", item.DisplayDiscountPrice, item.ID, "jpg")
			var dt = model.DownItem{}
			dt.Url = item.DisplayImage
			dt.LikeNum = item.LikeNum

			dt.Path = filepath.Join(filep, imageName)
			dt.Price = item.DisplayDiscountPrice
			result = append(result, &dt)
		}
	}
	if len(result) == 0 {
		fmt.Println("没有数据")
		return
	}
	//排序
	sort.Sort(model.ByDownItemLikeNum{result})

	fmt.Printf("此品类一共有图片 %v 张\n", len(result))
	if int64(len(result)) < con.Limit[0] {

		fmt.Printf("范围不够,区间为%v-%v,图片只有%v张",con.Limit[0],con.Limit[1], len(result))
	} else if int64(len(result)) >=  con.Limit[0] && int64(len(result)) <=  con.Limit[1]{
		result = result[con.Limit[0]:]
		fmt.Println("截取后图片数量：",len(result))
	} else {
		result = result[con.Limit[0]:con.Limit[1]]
		fmt.Println("截取后图片数量：",len(result))
	}

	//下载
	for _, item := range result {
		_ = common.Download(item.Url, item.Path)
	}

}
//URL	http://m.fordeal.com/detail/3988292?customer_trace=1.oms.1.2.3988292.0.w6Lu5OG3IQUqt.cat_70000055
func FetchDetails(Ids []int64) {
	path := config.DownLoadRootPath + "/" + "detail"
	if file.PathExist(path) == false {
		err := os.Mkdir(path, 0777)
		if err != nil {
			panic(err)
		}
	}

	for _, v := range Ids {
		url := fmt.Sprintf(ItemDetail, v)
		imgeurl := GetVideoUrlFromweb(url)
		fmt.Println(imgeurl)
		suf := fmt.Sprintf("/%v.jpg", v)
		_ = common.Download(imgeurl, filepath.Join(path, suf))
	}

}
func GetVideoUrlFromweb(weburl string)string  {

	res, err := http.Get(weburl)
	if err != nil {
		return ""
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return ""
	}

	// Load the HTML document
	doc, _ := goquery.NewDocumentFromReader(res.Body)
	// Find the review items

	url := ""
	doc.Find("meta").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//band := s.Find("").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band)

		op, _ := s.Attr("property")
		con, _ := s.Attr("content")
		if op == "og:image" {
			url = con
		}

	})
	//d-price col_ff0b0b
	doc.Find("span[class=\"d-price col_ff0b0b\"]").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		//band := s.Find("").Text()
		//fmt.Printf("Review %d: %s - %s\n", i, band)

		price := s.Text()
		fmt.Println("---",price)

	})


	return url

}