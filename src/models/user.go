package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id bson.ObjectId `bson:"_id"`
	//Username  string        `bson:"username"` //用户名
	Email     string     `bson:"email"`    //Email
	Password  string     `bson:"password"` //密码
	Name      string     `bson:"name"`     //姓名
	Note      string     `bson:"note"`     //备考
	Address   string     `bson:"address"`  //住址
	Tel       string     `bson:"tel"`      //电话
	Fax       string     `bson:"fax"`      //传真
	Mobile    string     `bson:"mobile"`   //手机
	QQ        string     `bson:"qq"`       //QQ
	Wechat    string     `bson:"wechat"`   //微信
	Point     string     `bson:"point"`    //积分
	Shippings []Shipping `bson:"shipping"` //常用寄送地址
	BaseModel
}
type Shipping struct {
	Id      bson.ObjectId `bson:"_id"`
	Name    string        `bson:"name"`    //收货人
	Address string        `bson:"address"` //收货地址
	Tel     string        `bson:"tel"`     //收货人电话
	Mobile  string        `bson:"mobile"`  //收货人手机
}
