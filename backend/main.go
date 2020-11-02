package main

import (
    "github.com/ant0ine/go-json-rest/rest"
    "github.com/maip0902/mydog-rest-api/models"
    "github.com/maip0902/mydog-rest-api/auth"
    "net/http"
    "log"
    "github.com/joho/godotenv"
    "os"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    api := rest.NewApi()
    api.Use(rest.DefaultDevStack...)
//     api.Use(&rest.CorsMiddleware{
//             RejectNonCorsRequests: false,
//             OriginValidator: func(origin string, request *rest.Request) bool {
//                 return origin == "http://localhost:8080"
//             },
//             AllowedMethods: []string{"GET", "POST", "PUT", "OPTIONS"},
//             AllowedHeaders: []string{
//                 "Accept", "Content-Type", "X-Custom-Header", "Origin"},
//             AccessControlAllowCredentials: true,
//             AccessControlMaxAge:           3600,
//         })
    router, err := rest.MakeRouter(
        rest.Get("/code/:code", models.GetImageByCode),
        rest.Get("/code", models.GetAll),
        rest.Options("/signUp", auth.SignUp),
        rest.Post("/signUp", auth.SignUp),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
   	log.Fatal(http.ListenAndServe(":" + os.Getenv("DEFAULT_PORT"), api.MakeHandler()))
}