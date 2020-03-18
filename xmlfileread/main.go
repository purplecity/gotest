package main


import (
	"bufio"
	"encoding/xml"
	"fmt"
	"log"
	"os"
	"strings"
)

type row struct {
	DataType string `xml:"dataType,attr"`
	BillNo string `xml:"billNo,attr"`
	PlayerName string `xml:"playerName,attr"`
	AgentCode string `xml:"agentCode,attr"`
	GameCode string `xml:"gameCode,attr"`
	NetAmount string `xml:"netAmount,attr"`
	BetTime string `xml:"betTime,attr"`
	GameType string `xml:"gameType,attr"`
	BetAmount string `xml:"betAmount,attr"`
	ValidBetAmount string `xml:"validBetAmount,attr"`
	Flag string `xml:"flag,attr"`
	PlayType string `xml:"playType,attr"`
	Currency string `xml:"currency,attr"`
	TableCode string `xml:"tableCode,attr"`
	LoginIP string `xml:"loginIP,attr"`
	RecalcuTime string `xml:"recalcuTime,attr"`
	PlatformType string `xml:"platformType,attr"`
	Remark string `xml:"remark,attr"`
	Round string `xml:"round,attr"`
	Result string `xml:"result,attr"`
	BeforeCredit string `xml:"beforeCredit,attr"`
	DeviceType string `xml:"deviceType,attr"`
}
func main() {
	//filenameinfoList := []string{"202003140007.xml","202003140009.xml","202003140011.xml","202003140013.xml"}
	xmlpath := "/Users/ludongdong/Downloads/testreadxml/202003140011.xml"
	readnum := int64(0)
	file, err := os.Open(xmlpath)
	if err != nil {
		log.Fatalf("open file failed :%+v\n",err)
	}

	filescanner := bufio.NewScanner(file)
	linecount := int64(0)
	beginread := false
	for filescanner.Scan() {
		//从记录的上一次读取的行开始读
		if readnum  == linecount || beginread {
			filetext := filescanner.Text()
			beginread = true
			//是否包含 那个字符串
			if strings.Contains(filetext,"dataType=\"BR\"")  {
				//解析 查是否存在  存在更新否则插入
				data := row{}
				xml.Unmarshal([]byte(filetext),&data)
				log.Printf("%+v\n",data)
			}
		}
		linecount++
	}
	file.Close()
	fmt.Println(linecount)
}
