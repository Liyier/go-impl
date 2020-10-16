package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	eng := gin.New()
	eng.GET("/set", func(context *gin.Context) {
		c := http.Cookie{
			Name:       "key-cookie",
			Value:      "value-cookie",
			Path:       "/",
			Expires:    time.Now().Add(time.Hour),
			RawExpires: "",
			MaxAge:     1000,
			Secure:     true,
			HttpOnly:   true,
			SameSite:   http.SameSiteNoneMode,
			Raw:        "",
			Unparsed:   nil,
		}
		http.SetCookie(context.Writer, &c)
		context.String(http.StatusOK, "done")
	})

	eng.GET("/delete", func(ctx *gin.Context) {
		c := http.Cookie{
			Name:       "Pycharm-d25e2260",
			Path:       "/",
			RawExpires: "",
			MaxAge:     -1,
			Secure:     true,
			HttpOnly:   true,
			SameSite:   http.SameSiteNoneMode,
		}
		http.SetCookie(ctx.Writer, &c)
	})
	eng.Run(":8080")
}
