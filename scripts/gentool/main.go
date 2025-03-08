package main

import "gorm.io/gen"


func main() {
  g := gen.NewGenerator(gen.Config{
    OutPath: "../query",s
    Mode: gen.WithoutContext|gen.WithDefaultQuery|gen.WithQueryInterface, // generate mode
  })

   gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/go-layout?charset=utf8mb4&parseTime=True&loc=Local"))
  g.UseDB(gormdb) // reuse your gorm db


  // Generate the code
  g.Execute()
}
