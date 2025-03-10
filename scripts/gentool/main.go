package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./internal/data/gen/query",
		ModelPkgPath:      "./internal/data/gen/model",
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable:     true,
		FieldCoverable:    false,
		FieldWithIndexTag: false,
		FieldWithTypeTag:  false,
	})

	dsn := "root:4b21ac7296fdf64a2e72d1b77fe5866c@tcp(127.0.0.1:3306)/go-layout?charset=utf8mb4&parseTime=True&loc=Local"
	gormdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	g.UseDB(gormdb)

	autoUpdateTimeField := gen.FieldGORMTag("updated_at", func(tag field.GormTag) field.GormTag {
		return tag.Append("autoUpdateTime")
	})

	autoCreateTimeField := gen.FieldGORMTag("created_at", func(tag field.GormTag) field.GormTag {
		return tag.Append("autoCreateTime")
	})

	deleteField := gen.FieldType("deleted_at", "gorm.DeletedAt")

	ignoreField := gen.FieldIgnore("extra_info")

	fieldOpts := []gen.ModelOpt{
		autoCreateTimeField,
		autoUpdateTimeField,
		deleteField,
		ignoreField,
	}

	allModel := g.GenerateAllTable(fieldOpts...)
	g.ApplyBasic(allModel...)
	g.Execute()
}
