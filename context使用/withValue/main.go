package main

import (
	"context"
	"fmt"
	"net/http"
)

/*
case1:
首先，我们创建一个空上下文并将其分配给ctx变量。
	1.ctx使用其键和值创建 3 个上下文作为父值。
	2.然后我们创建另一个ctx1作为父级的上下文并给它一个键和值。
	3.我们将尝试从中提取价值ctx1，ctx2，ctx3用正确的密钥。它将根据键返回给我们值。
	4.如果我们尝试从具有错误键的上下文中提取值，它将返回nil值。

case2:
	在http上下文中插入固值(中间件的方式)，如 userID，Token等。
*/

func main() {
	//case1()
	http.Handle("/", middleware1(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// After Insert Claim into Context

		fmt.Printf("%+v\n", CtxClaim(r.Context()))
		fmt.Fprintf(w, "%+v\n", CtxClaim(r.Context()))
	})))

	http.ListenAndServe(":8080", nil)
}

type Claims struct {
	ID int `json:"id"`
}

var contextKey = "ctx-key"

func middleware1(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Example: Get UserID From Token
		// ....
		// ....

		fmt.Println("Insert Claim into Context")
		newCtx := context.WithValue(r.Context(), contextKey, &Claims{
			ID: 1, // example User ID: 1
		})

		r = r.WithContext(newCtx)

		next.ServeHTTP(w, r)
	})
}

func CtxClaim(ctx context.Context) *Claims {
	raw, _ := ctx.Value(contextKey).(*Claims)
	return raw
}

func case1() {
	ctx := context.Background() // Empty Context

	ctx1 := context.WithValue(ctx, "key1", "value1") // parent: ctx
	ctx2 := context.WithValue(ctx, "key2", "value2") // parent: ctx
	ctx3 := context.WithValue(ctx, "key3", "value3") // parent: ctx

	ctx4 := context.WithValue(ctx1, "key4", "value4") // parent: ctx1

	fmt.Println(ctx1.Value("key1")) // value1
	fmt.Println(ctx2.Value("key2")) // value2
	fmt.Println(ctx3.Value("key3")) // value3

	fmt.Println(ctx4.Value("key4")) // value4
	fmt.Println(ctx4.Value("key1")) // value1

	fmt.Println(ctx3.Value("key1")) // nil
}
