package main

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"gotest/sql/Operation"
	"gotest/sql/mysql"
	"strconv"
	"time"
)

var (
	testfile = "/Users/ludongdong/zaqizaba/test.xlsx"
)

func main() {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("szcitest")
	if err != nil {
		fmt.Println(err.Error())
	}

	ffrow := sheet.AddRow()
	fphcell := ffrow.AddCell()
	fphcell.Value = "手机号"

	fovcell := ffrow.AddCell()
	fovcell.Value = "下单价格"

	fsvcell := ffrow.AddCell()
	fsvcell.Value = "结算价格"

	fsycell := ffrow.AddCell()
	fsycell.Value = "标的物"

	fhtcell := ffrow.AddCell()
	fhtcell.Value = "下单时间"

	fincell := ffrow.AddCell()
	fincell.Value = "下单金额"

	foutcell := ffrow.AddCell()
	foutcell.Value = "结算金额"

	foddscell := ffrow.AddCell()
	foddscell.Value = "赔率"

	frscell := ffrow.AddCell()
	frscell.Value = "输赢"



	cond := map[string]interface{}{"Handletime__gte":1567995000,"Handletime__lt":1568013000}
	rs := []mysql.Realtrade{}
	mysql.GetAllRecord("Realtrade",cond,&rs)

	for _,v := range rs {
		frow := sheet.AddRow()
		phcell := frow.AddCell()
		phcell.Value = v.Uid

		ovcell := frow.AddCell()
		ovcell.Value = Operation.HPstring(v.Ordervalue)

		svcell := frow.AddCell()
		svcell.Value = Operation.HPstring(v.Settlevalue)

		sycell := frow.AddCell()
		sycell.Value = v.Symbol

		htcell := frow.AddCell()
		htcell.Value = time.Unix(v.Handletime,0).Format("2006-01-02 15:04:05")

		incell := frow.AddCell()
		incell.Value = Operation.HPstring(v.Inputamount)

		outcell := frow.AddCell()
		outcell.Value = Operation.HPstring(v.Outputamount)

		oddscell := frow.AddCell()
		oddscell.Value = Operation.HPstring(v.Odds)

		rscell := frow.AddCell()
		rscell.Value = strconv.FormatInt(int64(v.Orderresult),10)
	}

	err = file.Save(testfile)
	if err != nil {
		fmt.Println("写入失败")
	}

}
