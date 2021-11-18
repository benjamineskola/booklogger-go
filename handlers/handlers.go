package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type App struct {
	router *mux.Router
	DB     *gorm.DB
}

func InitApp(router *mux.Router, db *gorm.DB) *App {
	app := App{router: router, DB: db}

	return &app
}

func (app *App) AddJSONRoute(route string, fun func(*Context) (int, []byte)) {
	ctx := app.NewContext()
	app.router.HandleFunc(route,
		func(resp http.ResponseWriter, req *http.Request) {
			ctx.Vars = mux.Vars(req)

			code, result := fun(ctx)
			if code != 0 {
				resp.WriteHeader(code)
				result, _ = json.Marshal(map[string]string{"error": string(result)})
			}
			_, err := resp.Write(result)
			if err != nil {
				panic(err)
			}
		})
}

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
	App      *App
	Vars     map[string]string
}

func (app *App) NewContext() *Context {
	ctx := Context{App: app}

	return &ctx
}
