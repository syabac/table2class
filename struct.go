package main

type GeneratorParam struct {
	DSN       string
	DbName    string
	TableName string
	ClassName string
	Language  string
}

type TableColumnInfo struct {
	TableType  string `gorm:"column:TABLE_TYPE"`
	TableName  string `gorm:"column:TABLE_NAME"`
	ColumnName string `gorm:"column:COLUMN_NAME"`
	TypeName   string `gorm:"column:TYPE_NAME"`
	Size       *int   `gorm:"column:SIZE"`
	Precision  *int   `gorm:"column:PRECISION"`
	Scale      *int   `gorm:"column:SCALE"`
	IsNullable bool   `gorm:"column:IS_NULLABLE"`
}
