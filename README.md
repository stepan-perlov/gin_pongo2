# gin_pongo2
pongo2 middleware for Gin framework.

##Example:
```Go
package main

import (
    "os"

    "github.com/gin-gonic/gin"
    "github.com/stepan-perlov/gin_pongo2"
    "github.com/flosch/pongo2"
)

func main() {
    switch os.Getenv("MODE") {
    case "RELEASE":
        gin.SetMode(gin.ReleaseMode)

    case "DEBUG":
        gin.SetMode(gin.DebugMode)

    case "TEST":
        gin.SetMode(gin.TestMode)

    default:
        gin.SetMode(gin.ReleaseMode)
    }

    engine := gin.New()
    engine.Use(gin.Recovery())

    if gin.IsDebugging() {
        engine.HTMLRender = gin_pongo2.NewDebug("resources")
    } else {
        engine.HTMLRender = gin_pongo2.NewProduction("resources")
    }

    engine.Static("/static", "resources/static")
    engine.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.tpl", pongo2.Context{"title": "Gin-pongo2!"})
    })
    engine.GET("/other", func(c *gin.Context) {
        # if this key not exists in third parameter
        # gin_pongo2.MakeContext will copy key1, key2
        # from gin.Context to pongo2.Context using gin.Get function
        c.HTML(http.StatusOK, "other.tpl", gin_pongo2.MakeContext(c, []string{'key1', key2}, map[string]interface{}{}))
    })

    engine.Run(":3000")
}
```
