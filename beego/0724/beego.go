

用户登录模块

1.会话session
	session服务器端存储
		内存
		本地磁盘文件
		数据库
	序列化
	sid => session id
	set-cookie session id
	cookie


	session
		配置 => 配置文件中配置 conf/app.conf
				ini格式配置文件
		    是否开启: SessionOn=true/false
		    存储类型: SessionProvider=memory/file/mysql/redis
		    存储的位置:  SessionProviderConfig=
		    失效时间: SessionGCMaxLifetime=3600 (失效时间一小时)
		    		SessionCookieLifeTime=3600
		    Sid名称(可以手动设置)： SessionName=BeegoSessionId

		操作
			Controller 对象做以下操作
				key=>value
				    获取
				    	GetSession(key) interface{} 检查权限
		    		更新
		    			SetSession(key,value) 登录认证成功
		    		删除
		    			DelSession(key)
		    		销毁
		    			DestorySession() 退出登录
			默认gob编码

		Sesssion存储数据:
			1.尽量少
			2.尽量使用基本类型
			3.自动逸类型一定要gob注册


	用户登录
		1.用户名,密码
		2.密码加密方式
			验证
			修改密码
			添加用户

			hash运算: md5 彩虹表查询(不建议使用)
				sha1,sha256,sha512
				text => md5 => hash
				123456 => md5 => xxxxxxx
				md5 32位16进制字符串组成集合有限

				验证: text => md5 => hash
					input(用户输入数据) => md5 => input(用户输入数据)hash
					hash input(用户输入数据)hash => text nput

				加盐MD5
					text => 随机字符串(6) salt(随机字符串) => salt(随机字符串)+"|"+text => md5 => hash
					存储slat(随机字符串) hash => 
					验证: 
						text => salt(随机字符串) => connect(salt(随机字符串)+text) => hash
						input => 获取对应的salt(随机字符串) => connect(salt(随机字符串), input(用户输入数据)) => input hash

						hash inputhash =  text input   对比 如果相同即相同

				慢速加盐算法
					bcrypt
						genate     	 生成hash密码
						check/verify 检查/验证

				项目中使用sha256或sha512,样本集更大
			golang.org/x/crypto/bcrypt


		代码
			package main

			import (
				"fmt"

				"golang.org/x/crypto/bcrypt"
			)

			func main() {
				password := "123@456"

				// 通过hash加密密码
				hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10) // 计算重复次数10 所以比较慢 默认10
				fmt.Println(string(hashed), err)
				if err != nil {
					fmt.Println("计算错误")
				}
				hashed1 := "$2a$16$6wvzhj7HlQZD2PuFDZ/OD.i56VoTV3e2uhqCzpcXcL6QxYwIruuMK"
				hashed2 := "$2a$16$CXZ7HCV3fTFTqAglofNf/OSj5yvpB5/tCpoIXvtdvLYnohMsw0qPK"

				// 运算对比 返回err类型  为nil则
				// 参数1:  hash密码  参数2: 原密码(text)
				fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed1), []byte(password)))
				fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed2), []byte(password)))
				fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)))

				password1 := "123123123"
				fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed1), []byte(password1)))
				fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed2), []byte(password1)))
				fmt.Println(bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password1)))

			}



		3.登录逻辑
			a.打开登录页面
			b.用户输入用户名,密码，点击登录
				验证： name => password hash input
				成功 => 用户列表
				失败 => 返回原有页面(登录页面) 提示用户登录失败


		beego 运行模式
		dev: 开发模式
		prod: 生产模式(默认模式)
		app.conf 配置文件中进行修改 尽量使用dev模式开发项目
		RunMode=dev  项目: conf/app.conf


		session 
			存 ID =>　ID　　登录成功后
			取 ID => nil 未登录/ID存在 已登录
			销毁 退出登录

			开启session
			SessionOn=true
			# SessionOn=true/false
			# SessionProvider="redis/file/mysql/memory"
			# SessionProviderCOnfig=""根据存储类型配置参数
			# SessionGCMaxLifetime=3600 session失效时间
			# SessionName=alinx SID名称设置


			注意: 开启session后如果不对session进行存储,那么又将会重定向回到登录页面
			c.SetSession("user", user.ID) // 存储session id

			读取时取的字段和 存储的字段要匹配否则无法验证


			所有登录以后访问的操作都加上验证session的步骤

				user := c.GetSession("user")
				if user == nil {
					// session等于nil 那就属于未登录,然后重定向到登录页面
					c.Redirect("/auth/login", 302)
					return // 保证后面的代码不执行
				}


			项目路径: D:\开发集合\goproject\user




			Init
			Prepare		路由执行前检查函数 可以将用户验证的逻辑放进去 在资源操作中添加即可
			架构图地址: https://beego.me/docs/mvc/

			通过封装beego.Controllers 进行登录验证

			定义

				package controllers

				import (
					"fmt"

					"github.com/astaxie/beego"
				)

				type RequiredAuthController struct {
					beego.Controller
				}

				// 控制器添加功能
				func (c *RequiredAuthController) Prepare() {
					user := c.GetSession("user")
					if user == nil {
						// 为空则未找到session数据库，状态为未登录
						fmt.Println("用户未登录,重定向到登录页面")
						c.Redirect("/auth/login", 302)
					}
				}


			使用
				package controllers

				import (
					base "user/base/controllers"
				)

				type HomeController struct {
					base.RequiredAuthController // 通过封装的 Controller控制器
				}

				func (c *HomeController) Index() {
					c.Ctx.WriteString("home")
				}

			只要使用了RequiredAuthController控制器，那后续的每一次操作都会检查sessionID是否存在后才进行资源重定向




2. orm

	ORM: object relaction mapping

	数据表/数据表操作 <=> 面向对象
		类，对象，方法
		结构体 结构体对象 结构体实例 结构体方法

		数据库和结构体对应关系
			表(通过列定义数据模型) <=>   结构体(通过属性定义)
			数据 				 <=>   实例
			对数据的操作 			 <=>   方法

			go中操作表
				通过定义结构体 自定生成表
				通过对结构体实例的方法调用，对表进行操作

	ORM框架
		对数据库性能要求比较高的时候，直接使用sql语句去操作即可
		beego/orm   : 底层就是调用sql去进行执行，中间会有转换的过程

		
		数据库 => 配置类信息
		操作 => orm包内提供的接口 (函数 方法)


		表定义(表对应结构体)

	
	使用
		导入orm包
		注册数据库驱动，mysql已经自动注册
		注册数据库(数据库配置信息对接)
		注册模型(通过结构体完成)

		数据库操作，同步数据库模型到数据库中的表

	数据库准备
		mysql> create database beego default charset=utf8mb4;
		mysql> grant all privileges on beego.* to '账号'@'%';
		mysql> flush privileges;
		mysql> show grant for 账号;
		mysql> drop table beego_user;


	代码
		package main

		import (
			"log"
			"time"

			"github.com/beego/beego/v2/client/orm"
			_ "github.com/go-sql-driver/mysql"
		)

		// 结构体定义
		/*
			表名称为结构体名称 除了第一个大写字母外其余的大写字母将会转换为_小写字母
			可以通过函数定义表名称


		*/
		type User struct { // Id 自动设置为主键 如果不指定的话
			// orm标签设置属性
			/*
				column(uid) 字段名称
				pk  主键设置
				auto 自动增长
				default(123) 在not null时设置默认值
				type(text) 类型设置
				size(64) 长度设置
				unique   唯一索引
				index 创建所以，联合索引
				null  允许为null，默认是not null
				auto_now_add 如果时创建会更新当前时间
				auto_now 当更新或者插入时更新时间
				digits(12) 整数长度
				decimals(3) 小数长度

			*/
			Id        int64      `orm:"column(uid);pk;auto"`
			Name      string     `orm:"size(64);unique"`
			Password  string     `orm:"size(128)"`
			Tel       string     `orm:"null"`
			Sex       bool       `orm:"null"`
			Height    float64    `orm:"digits(12);decimals(3);default(1.78)"`
			Addr      string     `orm:"type(text)"`
			CreatedAt *time.Time `orm:"auto_now_add"`
			UpdatedAt *time.Time `orm:"auto_now"`
			DeletedAt *time.Time `orm:"null"`
		}

		func (u *User) TableName() string {
			return "beego_user" // 自定义表名称
		}

		// 联合索引设置
		func (u *User) TableIndex() [][]string {
			return [][]string{
				{"Name"},
				{"Password"},
				{"Tel", "Addr"},
			}
		}

		func main() {
			// 数据库连接信息
			dsn := "root:123456@tcp(127.0.0.1:3306)/beego?parseTime=true&loc=Local&charset=utf8mb4"

			// 注册数据库驱动 pgsql/mysql 已经自动注册
			orm.RegisterDriver("mysql", orm.DRMySQL)

			// 注册数据库
			if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
				log.Fatal(err) // 记录日志并退出
			}

			// 注册模型 必须是指针类型的实例
			orm.RegisterModel(&User{})

			// 操作

			// DDL 同你数据库
			/*
				结构体对应表是否存在
				表不存在 创建对应的表
				若表存在 属性列是否在表中存在
				属性不存在 添加列
				索引是否存在 索引不存在 添加索引


				三个参数
					1 更改的数据
					2 是否先删除后再创建(危险操作)
					3 是否显示详细信息
			*/
			orm.RunSyncdb("default", true, true) // 同步数据建
		}




	DML DQL
	创建可操作的orm对象，然后对数据库进行CRUD操作


	项目对接到orm中
	导入驱
	导入orm包
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	orm.RegisterDriver("mysql", orm.DRMySQL) // 注册mysql驱动,默认已注册
	if err := orm.RegisterDataBase("default", "mysql", dsn); err != nil {
		log.Fatal(err) // 记录错误并退出
	}
	// default 数据库链接别名
	orm.RunSyncdb("default", false, true) // 同步数据库 
	/*
	三个参数:
		1. 数据库名称
		2. 是否删除重建
		3. 是否显示详细执行过程
	*/



	作业
		添加密码字段
		添加修改密码







