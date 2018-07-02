package main

func initializeRoutes() {
	router.GET("/", showIndexPage)

	articleRoutes := router.Group("/article")
	{
		articleRoutes.GET("/view/:article_id", getArticle)
	}
}
