# go-playdb

golang 向 数据库进行转换的工具

# 函数说明

## GenerateCreateDatabaseSQL

函数名：GenerateCreateDatabaseSQL(dbName string) string
参数：dbName (string) 数据库名
返回值: 生成的创建数据库 SQL 语句

## GenerateCreateTableSQL

函数名：GenerateCreateTableSQL(data any) (string, error)  
参数：data （any） golang 结构体。其中字段应标记 tag 值为`axis` 或 `axis_y`
返回值: 生成的创建表的 SQL 语句
