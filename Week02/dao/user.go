package dao

import (
	"Week02/code"
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	Id int
	Name string
}

func BatchGetUser()	([]User, error) {
	// db
	// 暂时不知道放哪里
	//db, err := sql.Open("mysql", "user:password@tcp(127.0.0.1:3306)/hello")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	// 查询
	//sql := "SELECT * FROM users WHERE is_vip = ?"
	//rows, err := db.Query(sql, 1)
	//if err != nil {
	//	// ...
	//}

	// 获取 error
	//err = rows.Err()
	err := sql.ErrNoRows
	if err != nil {
		// sql.ErrNoRows
		// 重点：内部吞掉，用一个默认的错误返回
		// 好处：上层看不到底层实现，看到的是一个业务错误码（DDD 思想）
		// wrap 包含 stack stace 堆栈信息
		return nil, errors.Wrapf(code.NotFound, "sql error: %v", err)
	}

	return []User{}, nil
}
