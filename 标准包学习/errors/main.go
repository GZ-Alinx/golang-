package main

import (
	"errors"
	"fmt"
)

type DbError struct {
}

func (e DbError) Error() string {
	return "dbError"
}

// 判断error是不是某种类型的
func main() {
	var err error = new(DbError)
	fmt.Println(err)
	fmt.Println(errors.Is(err, &DbError{}))
	var wrapError error = errors.Unwrap(err)
	fmt.Println(errors.Is(wrapError, &DbError{})) // 判断是否是DbError的类型
	var dbErr DbError
	fmt.Println(errors.As(dbErr, &DbError{}))

}
