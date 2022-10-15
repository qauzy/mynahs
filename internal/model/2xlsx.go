package model

import (
	"github.com/tealeg/xlsx"
	"strconv"
)

func Goods2xlsx() {

	goodsodel := NewGoods()
	var p = &Page{
		Page:     0,
		PageSize: 10000,
	}
	list, err := goodsodel.GetGoodsList(p)
	if err != nil {

	}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("商品信息汇总")
	if err != nil {
		panic(err.Error())
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "ID"

	cell = row.AddCell()
	cell.Value = "条形码"

	cell = row.AddCell()
	cell.Value = "商品名"

	cell = row.AddCell()
	cell.Value = "价格"

	cell = row.AddCell()
	cell.Value = "规格"

	cell = row.AddCell()
	cell.Value = "生产厂家"

	for k, v := range list {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = strconv.Itoa(k)

		cell = row.AddCell()
		cell.Value = v.Barcode

		cell = row.AddCell()
		cell.Value = v.GoodsName

		cell = row.AddCell()
		cell.Value = v.Price

		cell = row.AddCell()
		cell.Value = v.Spec

		cell = row.AddCell()
		cell.Value = v.Supplier

	}

	err = file.Save("goods.xlsx")
	if err != nil {
		panic(err.Error())
	}
}
