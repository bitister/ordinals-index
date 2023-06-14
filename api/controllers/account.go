package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"models"
)

type Domain struct {
	BaseController
}

/*
*
  - showdoc
  - @catalog API接口/域名查询
  - @title 域名查询
  - @description 域名查询
  - @method post
  - @url http://54.250.244.153:8080/domains/query
  - @param name 选填 string 搜索框的输入
  - @param pageNum 选填 int 页码(默认1)
  - @param pageSize 选填 int 每页数量(默认100)
  - @param category 选填  string 分类(10K,)
  - @param typeList 选填  string数组 后缀类型([
    "sats",
    "btc",
    "txt",
    "ord",
    "unisat",
    "x"
    ])
  - @param orderType 选填  int 排序方式(0:铭文序号倒序 1：铭文需要升序 2：字母升序 3：字母降序 4：铭文余额升序 5: 铭文余额降序 6: 最短字符)
  - @param startWith 选填  string 开头
  - @param endWith 选填  string 结尾
  - @param minWidth 选填  int 字符最小长度
  - @param maxWidth 选填  int 字符最大长度
  - @param notLike 选填  string 不包含
  - @param wordsType 选填  int 字符类型(0:仅含数字 1:仅含字母 2:仅含Emoji)
  - @return {"code":1001,"status":true,"message":"导入成功","data":nil}
  - @return_param  code	string	状态码
  - @return_param status	string	状态
  - @return_param message	string	信息
  - @return_param data	nil	返回数据
  - @number 99
*/
/**
* showdoc
* @catalog API接口/域名查询
* @title 域名查询
* @description 域名查询的接口
* @method post
* @url http://54.250.244.153:8080/domains/query
* @header token 可选 string 设备token
* @param name 选填 string 搜索框的输入
* @param pageNum 选填 int 页码(默认1)
* @param pageSize 选填 int 每页数量(默认100)
* @param typeList 选填  string数组 后缀类型([
  "sats",
  "btc",
  "txt",
  "ord",
  "unisat",
  "x"
  ])
* @param orderType 选填  int 排序方式(0:铭文序号倒序 1：铭文需要升序 2：字母升序 3：字母降序 4：铭文余额升序 5: 铭文余额降序 6: 最短字符)
* @param startWith 选填  string 开头
* @param endWith 选填  string 结尾
* @param minWidth 选填  int 字符最小长度
* @param maxWidth 选填  int 字符最大长度
* @param notLike 选填  string 不包含
* @param wordsType 选填  int 字符类型(1:仅含数字 2:仅含字母 3:仅含Emoji)
* @return {"error_code":0,"data":{"uid":"1","username":"12154545","name":"吴系挂","groupid":2,"reg_time":"1436864169","last_login_time":"0"}}
* @return_param groupid int 用户组id
* @return_param name string 用户昵称
* @remark 这里是备注信息
* @number 99
*/
func (c *Domain) Query() {
	inputData := make(map[string]interface{}, 0)
	data := c.Ctx.Input.RequestBody
	err := json.Unmarshal(data, &inputData)
	if err != nil {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "解析参数错误")
		c.ServeJSON()
		return
	}

	name, ok := inputData["name"].(string)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "name参数错误")
		c.ServeJSON()
		return
	}

	page, ok := inputData["pageNum"].(float64)
	if !ok {
		beego.Info("==pageNum==", pageNum)
		page = 1
	}

	pageNum := int(page)

	size, ok := inputData["pageSize"].(float64)
	if !ok {
		beego.Info("==pageSize==", pageSize)
		size = 100
	}
	beego.Info("pageSize:", pageSize)

	pageSize := int(size)

	//category, ok := inputData["category"].(string)
	//if !ok {
	//	c.Data["json"] = c.Fail(c.Tr("参数错误"), "category参数错误")
	//	c.ServeJSON()
	//	return
	//}

	typeList, ok := inputData["typeList"].([]interface{})
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "typeList参数错误")
		c.ServeJSON()
		return
	}

	orderType, ok := inputData["orderType"].(int)
	if !ok {
		orderType = 0
	}

	startWith, ok := inputData["startWith"].(string)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "startWith参数错误")
		c.ServeJSON()
		return
	}

	endWith, ok := inputData["endWith"].(string)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "endWith参数错误")
		c.ServeJSON()
		return
	}

	minWidth, ok := inputData["minWidth"].(int)
	if !ok {
		minWidth = 0
	}

	maxWidth, ok := inputData["maxWidth"].(int)
	if !ok {
		maxWidth = 0
	}

	notLike, ok := inputData["notLike"].(string)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "notLike参数错误")
		c.ServeJSON()
		return
	}

	//wordsType, ok := inputData["wordsType"].(int)
	//if !ok {
	//	c.Data["json"] = c.Fail(c.Tr("参数错误"), "wordsType参数错误")
	//	c.ServeJSON()
	//	return
	//}

	//valid := validation.Validation{}
	//// 2.验证获取到的数据
	//valid.Numeric(minWidth, "minWidth")
	//valid.Numeric(maxWidth, "maxWidth")
	//valid.Max(orderType, 6, "orderType")
	//
	//// 3.判断有没有错误
	//if valid.HasErrors() { // 说明有错误
	//	for _, err := range valid.Errors { // 循环打印错误
	//		c.Data["json"] = c.Fail(c.Tr("参数错误"), err.Key)
	//		c.ServeJSON()
	//		return
	//	}
	//}

	session := c.O

	var typeLists []string
	for _, tl := range typeList {
		typeLists = append(typeLists, tl.(string))
	}

	beego.Info("typeLists:", typeLists)
	qs := session.QueryTable(models.DoMainTBName())
	if name != "" {
		qs = qs.Filter("content__icontains", name)
	}

	if typeLists != nil {
		qs = qs.Filter("type__in", typeLists)
	}

	if startWith != "" {
		qs = qs.Filter("content__istartswith", startWith)
	}

	if endWith != "" {
		qs = qs.Filter("content__iendswith", endWith)
	}

	if minWidth != 0 {
		qs = qs.Filter("content__len__gte", minWidth)
	}

	if maxWidth != 0 {
		qs = qs.Filter("content__len__lte", maxWidth)
	}

	if notLike != "" {
		qs = qs.Filter("content__icontains", notLike)
	}

	totalCount, err := qs.Count()
	if err != nil {
		beego.Error(err)
		c.Data["json"] = c.Fail(c.Tr("服务异常请重试"), nil)
		c.ServeJSON()
	}

	//orderType 选填  int 排序方式(0:铭文序号倒序 1：铭文需要升序 2：字母升序 3：字母降序 4：铭文余额升序 5: 铭文余额降序 6: 最短字符)
	if orderType == 1 {
		qs = qs.OrderBy("id")
	} else if orderType == 2 {
		qs = qs.OrderBy("content")
	} else if orderType == 3 {
		qs = qs.OrderBy("-content")
	} else if orderType == 4 {
		qs = qs.OrderBy("value")
	} else if orderType == 5 {
		qs = qs.OrderBy("-value")
	} else if orderType == 6 {
		qs = qs.OrderBy("length(content)")
	} else {
		qs = qs.OrderBy("-id")
	}

	domains := make([]models.DoMain, 0)
	if _, err := qs.Limit(pageSize, (pageNum-1)*pageSize).All(&domains); err != nil {
		beego.Error(err)
		c.Data["json"] = c.Fail(c.Tr("服务异常请重试"), nil)
		c.ServeJSON()
	}

	outData := struct {
		TotalCount    int64
		ProfitDetails []models.DoMain
	}{
		totalCount,
		domains,
	}

	c.Data["json"] = c.Succ(c.Tr("查询成功"), outData)
	c.ServeJSON()
}
