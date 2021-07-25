package main

// 导入包 驱动
import (
	"log"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // database/sql
)

/*
	标签名: orm  `orm:"pk"`
	指定主键: pk
	自动增长: auto
	列名:	column(name)
	字符串类型: 默认varchar(255)
					长度修改 size(lenght)
					指定类型 type(type)
	是否允许为null	默认不允许 null
	唯一值	unique
	索引	index1`
*/
// 用beego orm一定要设置显示设置主键: pk
type Account struct {
	// Id           string `orm:"pk"` // 非int类型 要制定主键
	Id           int64  `orm:"pk;auto;column(id);"`
	Name         string `orm:"size(64);unique"`
	Password     string
	Birthday     *time.Time
	Telephone    string
	Email        string `orm:"description(描述)"`
	Addr         string `orm:"default(中国,深圳)"`
	Status       int8   `orm:"default(1)"` // 默认数据
	RoleId       int64
	DepartmentId int64
	CreatedAt    *time.Time
	UpdatedAt    *time.Time
	DeletedAt    *time.Time `orm:"null"` // 允许为空
	Description  string
	Sex          bool
	Alinx        string `orm:"index;size(4096)"` // 索引
	C            string
	Height       float32 `orm:"digits(10);decimals(2)"` // digits长度  decimals小数位数
	Weihht       float64
}

// 设置表名称
func (account *Account) TableName() string {
	return "act"
}

// 索引设置 原理是反射的机制
func (acount *Account) TableIndex() [][]string {
	return [][]string{
		{"name", "password"},
		{"password"},
	}
}

func main() {
	// 注册驱动到orm 数据库类型
	orm.RegisterDriver("mysql", orm.DRMySQL)

	// 数据库连接串
	databaseName := "default"
	driverName := "mysql"
	dsn := "root:adl1314520@lx@tcp(106.12.217.9:3306)/beego?charset=utf8mb4&parseTime=true"
	// 数据库连接（连接池）,连接池名称 使用的驱动 数据库配置信息
	// beego 指定 defalut数据库名称
	err := orm.RegisterDataBase(databaseName, driverName, dsn)
	if err != nil {
		log.Fatal("error")
	}

	// 注册模型
	orm.RegisterModel(new(Account))

	// 使用

	// 表结构同步(库需要提前创建)
	/*
		参数含义
			databaseName	同步数据库
			force 			如果表存在则删除,再创建
			verbose 		打印出执行的sql语句
	*/
	orm.RunSyncdb(databaseName, true, true) // 只是更新表的列，并不会删除表

}
