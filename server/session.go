package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/tsingson/go-sessions"
	"github.com/tsingson/go-sessions/sessiondb/goredis"
	"github.com/valyala/fasthttp"
	"time"
)

var (
	db   *goredis.Database
	sess sessions.Sessions

	mySessionsConfig = sessions.Config{
		Cookie:                      "mysessioncookieid",
		Expires:                     time.Duration(2) * time.Hour,
		DisableSubdomainPersistence: false,
	}
	redisSessions sessions.Sessions
)

func init() {
	db = goredis.New(goredis.Config{})
	pong, err := db.Redis.Ping().Result()
	fmt.Println(pong, err)
	redisSessions = sessions.New(mySessionsConfig)
	redisSessions.UseDatabase(db)
	//	spew.Dump(redisSessions)
}

// set some values to the session
func setHandler(ctx *fasthttp.RequestCtx) {
	values := map[string]interface{}{
		"Name":   "go-sessions",
		"Days":   "1",
		"Secret": "dsads£2132215£%%Ssdsa",
	}

	sess := redisSessions.StartFasthttp(ctx) // init the session

	// sessions.StartFasthttp returns the, same, Session interface we saw before too
	//sess.UseDatabase(db)

	for k, v := range values {
		sess.Set(k, v) // fill session, set each of the key-value pair
	}
	ctx.WriteString("Session saved, go to /get to view the results")
}

// get the values from the session
func getHandler(reqCtx *fasthttp.RequestCtx) {
	sess := redisSessions.StartFasthttp(reqCtx) // init the session
	//sess.UseDatabase(db)

	sessValues := sess.GetAll() // get all values from this session
	spew.Dump(sessValues)
	reqCtx.WriteString(fmt.Sprintf("%#v", sessValues))
}

// clear all values from the session
func clearHandler(reqCtx *fasthttp.RequestCtx) {
	sess := redisSessions.StartFasthttp(reqCtx) // init the session
	//sess.UseDatabase(db)

	sess.Clear()
}

// destroys the session, clears the values and removes the server-side entry and client-side sessionid cookie
func destroyHandler(reqCtx *fasthttp.RequestCtx) {
	redisSessions.DestroyFasthttp(reqCtx)
}

func router() {
	fmt.Println("Open a browser tab and navigate to the localhost:8080/set")
	fasthttp.ListenAndServe(":8080", routerHandler)
}

func routerHandler(reqCtx *fasthttp.RequestCtx) {
	path := string(reqCtx.Path())

	if path == "/set" {
		setHandler(reqCtx)
	} else if path == "/get" {
		getHandler(reqCtx)
	} else if path == "/clear" {
		clearHandler(reqCtx)
	} else if path == "/destroy" {
		destroyHandler(reqCtx)
	} else {
		reqCtx.WriteString("Please navigate to /set or /get or /clear or /destroy")
	}
	return
}
