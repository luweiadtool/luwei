package config

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"myTool/file"
	"myTool/img"
	"path/filepath"
	"strings"
)


var Con config
var DownLoadRootPath string

type config struct {
	ForDeal ForDealCon `toml:"forDeal"`
}

func LoadConfig()config  {

	cPath := configPath()
	if file.PathExist(cPath) == false {
		fmt.Println("配置文件不存在", cPath)
		return config{}
	}

	if _, err := toml.DecodeFile(cPath, &Con); err != nil {
		fmt.Println("配置文件填写出错", err)
	}

	DownLoadRootPath = filepath.Join(file.GetHomeDir(),"Downloads")

	Con.ForDeal.RootPath = filepath.Join(DownLoadRootPath,"forDeal")
	logoPath :=filepath.Join(GetSourceDir(), Con.ForDeal.Puzzle.Logo)

	var err error
	Con.ForDeal.Puzzle.LogoImage,err = img.GetImageObj(logoPath)
	if err != nil {
		fmt.Println("logo 读取 出错")
		panic(err)
	}

	fmt.Println("配置文件信息：", Con)
	return Con

}


func configPath() string {

	conPath := filepath.Join(GetConfigDir(),"config.toml")
	if file.PathExist(conPath) == true {
		return conPath
	}
	cmdPath := file.GetCurrentDirectory()
	proPath := strings.TrimSuffix(cmdPath,"/cmd")
	conPath = filepath.Join(proPath,"/config/config.toml")
	return conPath

}

//字体文件
func GetFontSource()string  {

	return filepath.Join(GetSourceDir(),"arial.ttf")

}

//资源文件
func GetSourceDir()string  {
	return filepath.Join(GetProDir(),"source")
}

// 配置文件
func GetConfigDir()string  {
	return filepath.Join(GetProDir(),"config")
}

func GetProDir()string  {
	proPath := file.GetCurrentDirectory()
	if strings.HasSuffix(proPath,"Luwei") {
		//fmt.Println("1", proPath)
		return proPath
	}
	home := file.GetHomeDir()
	proPath = filepath.Join(home, "Desktop/Luwei")
	if file.PathExist(proPath) {
		//fmt.Println("2", proPath)
		return proPath
	}

	proPath = filepath.Join(home, "Desktop/qianjianeng")
	if file.PathExist(proPath) {
		//fmt.Println("3", proPath)
		return proPath
	}

	return "/Users/qianjianeng/go/src/goDemo/Project/Luwei"

}