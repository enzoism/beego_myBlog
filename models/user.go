package models


import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
)
type User struct {
	Id          int
	Phone       string
	UserProfile *UserProfile `orm:"rel(one)"`
	Password    string
	Status      int
	Created     int64
	Changed     int64
}
type UserProfile struct {
	Id       int
	Realname string
	Sex      int
	Birth    string
	Email    string
	Phone    string
	Address  string
	Hobby    string
	Intro    string
	User     *User `orm:"reverse(one)"`
}

func (this *User) TableName() string {
	return "user"
}
func init() {
	orm.RegisterModel(new(User), new(UserProfile)) //
}

func LoginUser(phone string, password string) (err error, user []User) {
	fmt.Println("-------------------models/user--LoginUser")
	o := orm.NewOrm()
	qs := o.QueryTable("user")

	cond := orm.NewCondition()
	cond = cond.And("phone", phone)
	cond = cond.And("password", com.Md5(password))
	cond = cond.And("status", 1)

	qs = qs.SetCond(cond)
	err = qs.Limit(1).One(&user)
	return err, user
}

func GetUser(id int) (User, error) {
	var user User
	var err error
	o := orm.NewOrm()

	user = User{Id: id}
	err = o.Read(&user)

	if err == orm.ErrNoRows {
		return user, nil
	}
	return user, err
}