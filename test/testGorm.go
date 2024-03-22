package test

import (
	"Go_Code/GinChat/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//打印logger
	// newLogger := logger.New(
	// 	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
	// 	logger.Config{
	// 		SlowThreshold: time.Second, // Slow SQL threshold
	// 		LogLevel:      logger.Info, // Log level
	// 		Colorful:      false,       // Disable color
	// 	},
	// )

	db, err := gorm.Open(mysql.Open("root:ztztzwd4250@tcp(127.0.0.1:3306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		//Logger: newLogger,
	})
	if err != nil {
		panic("failed to connect database")
	}

	//迁移schema，生成表
	err = db.AutoMigrate(&models.Community{})
	//err = db.AutoMigrate(&models.Message{})
	if err != nil {
		panic(err)
	}

	// message := &models.Message{}
	// message.Content = "hello"
	// db.Create(message)

	// //创建
	// user := &models.UserBasic{}
	// user.Name = "宵宫"
	// db.Create(user)

	// //读取
	// //user = &models.UserBasic{} //没有这一行的话，会变成Select * from user_basics where id = 3 and id = 10(最新创建的哪一行)
	// data := db.First(user, 1) // 根据整型主键查找,找id==3
	// fmt.Println(data)
	// //db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// // Update - 将 product 的 price 更新为 200
	// db.Model(user).Update("Password", 1234)
	// Update - 更新多个字段
	// db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // 仅更新非零值字段
	// db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// // Delete - 删除 product
	// db.Delete(&product, 1)
}
