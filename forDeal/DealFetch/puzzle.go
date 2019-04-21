package DealFetch

import (
	"fmt"
	"goDemo/Project/Luwei/config"
	"image"
	"myTool/file"
	"myTool/img"
	"os"
	"path/filepath"
	"strings"
)

func Puzzle(dir string, con config.ForDealCon, exportDir string, water, puzzle bool) {


	files, err := file.GetCurrentFiles(dir)
	if err != nil {
		fmt.Println("puzzle 读取图片失败")
		return
	}

	// 给每张图片添加文字水印
	if water {  //wanter
		waterPath := waterMelen(files,con,exportDir)
		files, err = file.GetAllFiles(waterPath)
		if err != nil {
			fmt.Println("puzzle 读取图片失败")
			return
		}
	}
	// 拼图 + 左上角加上logo
	if puzzle {
		puzzleAction(files, con, exportDir)
	}


}

func waterMelen(files []string,con config.ForDealCon, exportDir string) string  {
	// 生成临时目录
	waterDir := filepath.Join(exportDir,"temp")
	_ = os.MkdirAll(waterDir, 0777)

	//添加文字水印
	for i := 0; i < len(files); i ++ {

		im, err := img.GetImageObj(files[i])
		if err != nil || im == nil{
			continue
		}

		title := files[i]
		fileName := filepath.Base(title) //105.26-1294953.jpg
		price := strings.Split(fileName,"-")[0]
		price += " sar"
		export := fmt.Sprintf("%v/%v", waterDir,fileName)
		so := config.GetFontSource()
		img.MarkFontWithImage(im,price,export,float64(con.Puzzle.Fontsize),image.Point{85, 295},so)
	}
	return waterDir
}

func puzzleAction(files []string,con config.ForDealCon, exportDir string)  {

	images := img.GetImages(files)

	if con.Puzzle.Count == 0 {
		con.Puzzle.Count = 6
	}
	if len(images) / con.Puzzle.Count < 1 {
		fmt.Println("图片数量太少")
		return
	}

	yu := len(images) % con.Puzzle.Count
	images = images[:len(images) - yu]


	// 导出目录
	if file.PathExist(exportDir) == false {
		_ = os.MkdirAll(exportDir, 0777)
	}

	waterPath := filepath.Join(exportDir,"water")
	if file.PathExist(waterPath) == false {
		_ = os.MkdirAll(waterPath, 0777)
	}

	for i := 0; i < len(images); i += con.Puzzle.Count {
		exportPath := fmt.Sprintf("%v/%d.jpg", exportDir, i+1)
		imgs := images[i:i+con.Puzzle.Count]
		img.PuzzleImagesThumbnail(imgs,con.Puzzle.Style,exportPath)

		// 添加水印
		AddLogo(exportPath, con, waterPath,i+1)
	}
	// 删除临时文件目录
	_ = os.RemoveAll(filepath.Join(exportDir,"temp"))
}

func AddLogo(bjImgPath string,con config.ForDealCon, waterDir string, index int)  {
	bjImg ,_ := img.GetImageObj(bjImgPath)
	exportPath := fmt.Sprintf("%v/%v.png", waterDir,index)
	img.MarkLogoImage(bjImg,con.Puzzle.LogoImage,exportPath,image.Point{0,0})

}