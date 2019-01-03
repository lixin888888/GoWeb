package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)
func RegisterHandler() *httprouter.Router{
   router :=httprouter.New()
   router.POST("/user",CreateUser)
   router.POST("/user/:user_name",Login)
   return router
}
 func main()  {
 	r:=RegisterHandler()
    http.ListenAndServe(":8080",r)
}
