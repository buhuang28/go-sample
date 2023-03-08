package main

import (
	"fmt"
	"time"
)

func SetVal(k string, v interface{}, exp int64) error {
	return rdb.Set(ctx, k, v, time.Duration(exp)).Err()
}

func GetVal(k string) string {
	result, err := rdb.Get(ctx, k).Result()
	if err != nil {
		fmt.Println(err)
	}
	return result
}

func PutList_R(k string, val interface{}) error {
	i := rdb.RPush(ctx, k, val)
	fmt.Println(i)
	fmt.Println(i.String())
	fmt.Println(i.Result())
	fmt.Println(i.Name())
	fmt.Println(i.FullName())
	return i.Err()
}

func RangeList_R(k string) (error, []string) {
	lLen := rdb.LLen(ctx, k)
	lRange := rdb.LRange(ctx, k, 0, lLen.Val())
	if lRange.Err() != nil {
		return lRange.Err(), nil
	}
	return nil, lRange.Val()
}
