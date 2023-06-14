package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego/validation"
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

	pageNum, ok := inputData["pageNum"].(int)
	if !ok {
		pageNum = 1
	}

	pageSize, ok := inputData["pageSize"].(int)
	if !ok {
		pageSize = 100
	}

	category, ok := inputData["category"].(string)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "category参数错误")
		c.ServeJSON()
		return
	}

	typeList, ok := inputData["typeList"].(string)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "typeList参数错误")
		c.ServeJSON()
		return
	}

	orderType, ok := inputData["orderType"].(int)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "orderType参数错误")
		c.ServeJSON()
		return
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
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "minWidth参数错误")
		c.ServeJSON()
		return
	}

	maxWidth, ok := inputData["maxWidth"].(int)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "maxWidth参数错误")
		c.ServeJSON()
		return
	}

	notLike, ok := inputData["notLike"].(int)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "notLike参数错误")
		c.ServeJSON()
		return
	}

	wordsType, ok := inputData["wordsType"].(int)
	if !ok {
		c.Data["json"] = c.Fail(c.Tr("参数错误"), "wordsType参数错误")
		c.ServeJSON()
		return
	}

	valid := validation.Validation{}
	// 2.验证获取到的数据
	valid.Numeric(minWidth, "minWidth")
	valid.Numeric(maxWidth, "maxWidth")
	valid.Max(orderType, 6, "orderType")

	// 3.判断有没有错误
	if valid.HasErrors() { // 说明有错误
		for _, err := range valid.Errors { // 循环打印错误
			c.Data["json"] = c.Fail(c.Tr("参数错误"), err.Key)
			c.ServeJSON()
			return
		}
	}

	session := c.O

	c.Data["json"] = c.Succ(c.Tr("查询成功"), outData)
	c.ServeJSON()
}
