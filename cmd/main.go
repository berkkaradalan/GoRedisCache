package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/berkkaradalan/GoRedisCache/api/route"
	"github.com/berkkaradalan/GoRedisCache/bootstrap"
	"github.com/berkkaradalan/GoRedisCache/utils"
	"github.com/gin-gonic/gin"
)

func main() {
	log.Println(`
	 _____      ______         _ _     _____            _          
	|  __ \     | ___ \       | (_)   /  __ \          | |         
	| |  \/ ___ | |_/ /___  __| |_ ___| /  \/ __ _  ___| |__   ___ 
	| | __ / _ \|    // _ \/ _' | / __| |    / _' |/ __| '_ \ / _ \
	| |_\ \ (_) | |\ \  __/ (_| | \__ \ \__/\ (_| | (__| | | |  __/
	 \____/\___/\_| \_\___|\__,_|_|___/\____/\__,_|\___|_| |_|\___|
	`)
    app := bootstrap.App()
    env := app.Env
    db := app.MongoDB
	redis := app.Redis
	utils.MigrateDB(db, env.MONGODB_DB_NAME)
	timeout := time.Duration(env.CONTEXT_TIMEOUT) * time.Second
	r:=gin.Default()
	route.Setup(env, timeout, db, redis,r)
    srv := &http.Server{
		Addr:         env.SERVER_ADDRESS,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
    go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("server started")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)
}