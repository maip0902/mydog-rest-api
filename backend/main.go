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
    api.Use(&rest.CorsMiddleware{
            RejectNonCorsRequests: false,
            OriginValidator: func(origin string, request *rest.Request) bool {
                return origin == "http://localhost"
            },
            AllowedMethods: []string{"GET", "POST", "PUT"},
            AllowedHeaders: []string{
                "Accept", "Content-Type", "X-Custom-Header", "Origin"},
            AccessControlAllowCredentials: true,
            AccessControlMaxAge:           3600,
        })
    router, err := rest.MakeRouter(
        rest.Get("/api/code/:code", models.GetImageByCode),
        rest.Get("/api/codeImage/:id", models.GetImageById),
        rest.Get("/api/codeImage/image/:code", models.GetStatusImage),
        rest.Post("/api/codeImage/:id", models.UpdateImage),
        rest.Get("/api/code", models.GetAll),
        rest.Post("/api/signUp", auth.SignUp),
        rest.Post("/api/signIn", auth.SignIn),
        rest.Get("/api/email", auth.VerifyEmail),
        rest.Post("/api/authUser", auth.GetAuthenticatedUser),
    )
    if err != nil {
        log.Fatal(err)
    }
    api.SetApp(router)
   	log.Fatal(http.ListenAndServe(":" + os.Getenv("DEFAULT_PORT"), api.MakeHandler()))
//     r := mux.NewRouter()
//     r.Handle("/private", jwt.JwtMiddleware.Handler(private))
//     r.Handle("/public", jwt.GetTokenHandler)
//     //サーバー起動
//         if err := http.ListenAndServe(":3001", r); err != nil {
//             log.Fatal("ListenAndServe:", nil)
//         }
}