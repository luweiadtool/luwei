package forDeal

import (
	"fmt"
	"goDemo/Project/Luwei/config"
	"goDemo/Project/Luwei/forDeal/DealFetch"
	"myTool/file"
	"os"
	"strconv"
)

func ForDealRun(con config.ForDealCon) {
	fmt.Println(con)

	if con.CatId > 0 {
		path := con.RootPath + "/" + strconv.Itoa(int(con.CatId))

		if file.PathExist(path) == false {
			err := os.Mkdir(path,0777)
				if err != nil {
				panic(err)
			}
		}
		
		if con.CatId > 0 {
			DealFetch.Fetch(con.CatId, path, con)

			puzDir := fmt.Sprintf("%v/%v/%v_puzzle",con.RootPath, con.CatId, con.CatId)
			DealFetch.Puzzle(path, con, puzDir, true, true)

		}
	}

	if len(con.DetailIds) > 0 {
		DealFetch.FetchDetails(con.DetailIds)
	}

}
