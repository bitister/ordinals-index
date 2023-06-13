package models

type DoMain struct {
	Id            int64  `orm:"pk;description(主键id)"  form:"id" json:"id" `
	Name          string `orm:"size(255);description(名称)" form:"name" json:"name"`
	InscriptionId string `orm:"size(66);description(序列id)" form:"inscription_id" json:"inscription_id"`
	Value         uint64 `orm:"description(铭文余额)" form:"value" json:"value"`
	ContentLength uint64 `orm:"description(字符长度)" form:"content_length" json:"content_length"`
	Type          string `orm:"size(10);description(类型)" form:"type" json:"type"`
	Owner         string `orm:"size(62);description(所有者地址)" form:"owner" json:"owner"`
	Ctime         uint64 `orm:"description(铭刻时间)" form:"ctime" json:"ctime"`
}

func (a *DoMain) TableName() string {
	return DoMainTBName()
}

// 多字段索引
func (u *DoMain) TableIndex() [][]string {
	return [][]string{
		[]string{"id"},
		[]string{"name"},
		[]string{"inscription_id"},
		[]string{"type"},
		[]string{"owner"},
	}
}

// 多字段唯一键
func (u *DoMain) TableUnique() [][]string {
	return [][]string{
		[]string{"id"},
		[]string{"name"},
	}
}
