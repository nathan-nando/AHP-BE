package app

func Run(api *Api) {
	api.routes.Logger.Fatal(api.routes.Start(":" + api.cfg.Server.Port))
}
