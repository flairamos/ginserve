package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../srap-system/database/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	gormdb, _ := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/cwmdb"), &gorm.Config{})
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
	// Generate struct `TIotVideoDevice` based on table `t_iot_video_device`
	// g.GenerateModel("t_iot_video_device"),

	// g.GenerateModel("t_iot_video_capture_record"),

	// g.GenerateModel("t_iot_environment_device"),
	// g.GenerateModel("t_report_transport_apply"),
	// g.GenerateModel("t_report_transport_apply_trash"),
	// g.GenerateModel("t_iot_video_alarm_record"),

	// Generate struct `Employee` based on table `users`
	// g.GenerateModelAs("users", "Employee"),

	// Generate struct `User` based on table `users` and generating options
	// g.GenerateModel("users", gen.FieldIgnore("address"), gen.FieldType("id", "int64")),

	// Generate struct `Customer` based on table `customer` and generating options
	// customer table may have a tags column, it can be JSON type, gorm/gen tool can generate for your JSON data type
	// g.GenerateModel("customer", gen.FieldType("tags", "datatypes.JSON")),
	)
	//g.ApplyBasic(
	//	// Generate structs from all tables of current database
	//	g.GenerateAllTable()...,
	//)
	// Generate the code
	g.Execute()
}
