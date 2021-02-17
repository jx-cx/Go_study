package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库连接信息
const (
	USERNAME = "root"

	PASSWORD = ""

	NETWORK = "tcp"

	SERVER = "localhost"

	PORT = 3306

	DATABASE = "test"
)

var (
	db *sql.DB
)

type user struct {
	id       int
	username string
	pwssword int
}

func iniDb() (err error) {
	conn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", USERNAME, PASSWORD, NETWORK, SERVER, PORT, DATABASE)

	db, err = sql.Open("mysql", conn)
	if err != nil {
		fmt.Println("connection to mysql failed:", err)
		return
	}
	err = db.Ping() // 尝试连接数据库
	if err != nil {
		return
	}
	// 设置数据库连接池最大连接数
	db.SetMaxOpenConns(10)
	//设置连接池中的最大闲置连接数
	db.SetMaxIdleConns(5)
	return
}

// 查询单条
func queryOne(id int) {
	var u1 user
	//单条记录的sql语句
	sqlStr := `select id,username,password from users where id=?`
	//执行
	db.QueryRow(sqlStr, id).Scan(&u1.id, &u1.username, &u1.pwssword)
	fmt.Printf("u1:%#v\n", u1)
}

// 查询多条
func queryMore(id int) {
	sqlStr := "select id,username,password from users where id>?"
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据库链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u1 user
		err := rows.Scan(&u1.id, &u1.username, &u1.pwssword)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("u1:%#v\n", u1)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into users(username, password) values (?,?)"
	ret, err := db.Exec(sqlStr, "test3", 123456)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	theID, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert ID failed, err:%v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d.\n", theID)
}

// 更新数据
func updateRowDemo(id int) {
	sqlStr := "update users set password=? where id = ?"
	ret, err := db.Exec(sqlStr, 111111, id)
	if err != nil {
		fmt.Printf("update failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
func deleteRowDemo(id int) {
	sqlStr := "delete from users where id = ?"
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err:%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err:%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

// 预处理查询示例
func prepareQueryDemo(id int) {
	sqlStr := "select id, username, password from users where id > ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Printf("query failed, err:%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.username, &u.pwssword)
		if err != nil {
			fmt.Printf("scan failed, err:%v\n", err)
			return
		}
		fmt.Printf("id:%d name:%s password:%d\n", u.id, u.username, u.pwssword)
	}
}

func prepareInsertDemo() {
	sqlStr := "insert into users(username, password) values (?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err:%v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec("test5", 11)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	_, err = stmt.Exec("test6", 22)
	if err != nil {
		fmt.Printf("insert failed, err:%v\n", err)
		return
	}
	var m = map[string]int{
		"test7": 33,
		"test8": 44,
	}
	for s, i := range m {
		stmt.Exec(s, i)
	}
	fmt.Println("insert success.")
}

/*
事务相关方法
Go语言中使用以下三个方法实现MySQL中的事务操作。 开始事务

func (db *DB) Begin() (*Tx, error)
提交事务

func (tx *Tx) Commit() error
回滚事务

func (tx *Tx) Rollback() error
*/
// 事务操作示例
func transactionDemo() {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err:%v\n", err)
		return
	}
	sqlStr1 := "update users set password=66 where id = ?"
	ret1, err := tx.Exec(sqlStr1, 2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql1 failed, err:%v\n", err)
		return
	}
	affRow1, err := ret1.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	sqlStr2 := "update users set password=88 where id = ?"
	ret2, err := tx.Exec(sqlStr2, 3)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err:%v\n", err)
		return
	}
	affRow2, err := ret2.RowsAffected()
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec ret1.RowsAffected() failed, err:%v\n", err)
		return
	}

	fmt.Println(affRow1, affRow2)
	if affRow1 == 1 && affRow2 == 1 {
		fmt.Println("事务提交啦...")
		tx.Commit() // 提交事务
	} else {
		tx.Rollback()
		fmt.Println("事务回滚啦...")
	}

	fmt.Println("exec trans success!")
}
func main() {
	err := iniDb()
	if err != nil {
		fmt.Printf("ini DB failed,err:%v\n", err)
	}
	fmt.Println("连接数据库成功")
	//queryOne(1)
	//queryMore(0)
	//updateRowDemo(3)
	//deleteRowDemo(6)
	//prepareQueryDemo(0)
	//prepareInsertDemo()
	transactionDemo()
	// 拿到结果
	//rowObj.Scan(&u1.id, &u1.username, &u1.pwssword) //scan 释放数据库连接
	//打印结果

}
