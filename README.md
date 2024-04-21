# go-playdb

golang 向 数据库进行转换的工具

# 函数说明

## GenerateCreateDatabaseSQL

功能：生成创建数据库的 SQL 语句
函数名：GenerateCreateDatabaseSQL(dbName string) string
参数：dbName (string) 数据库名
返回值: 生成的创建数据库 SQL 语句

## GenerateCreateTableSQL

功能：生成创建表的 SQL 语句。被认为需要转化成数据库字段的结构字段应使用默认 Tag 名称 `axis` 或 `axis_y`
函数名：GenerateCreateTableSQL(data any) (string, error)  
参数：data （any） golang 结构体。其中字段应标记 tag 值为`axis` 或 `axis_y`
返回值: 生成的创建表的 SQL 语句

## GenerateCreateTableSQLAllField

功能：生成创建表的 SQL 语句。自定义结构体中的所有字段全部转化成数据库字段。
函数名：GenerateCreateTableSQLAllField(data any) (string, error)  
参数：data （any） golang 结构体。
返回值: 生成的创建表的 SQL 语句

## GenerateCreateTableSQLCustomTag

功能：生成创建表的 SQL 语句。自定义结构体中的含有被指定的 Tag 名称的字段转化成数据库字段。
函数名：GenerateCreateTableSQLCustomTag(data any, fieldTag string) (string, error)  
参数：data （any） golang 结构体。
fieldTag (string) 自定义 tag 名称
返回值: 生成的创建表的 SQL 语句
