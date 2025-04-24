package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	dsn := "postgres://myapp:myapp@localhost:5432/myapp?sslmode=disable" // ←調整
	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/dao",       // 生成コード出力先
		ModelPkgPath: "internal/dao/model", // モデルのみ分離可
		Mode:         gen.WithoutContext | gen.WithQueryInterface,
	})

	g.UseDB(db)

	// DB → struct 自動生成（products テーブルだけ）
	g.GenerateModel("products")

	// 型安全 DAO/DSL 生成
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
