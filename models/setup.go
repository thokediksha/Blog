package models

import (
    "fmt"
    "os"
    "strconv"

    "github.com/go-redis/redis"
    // "golang.org/x/crypto/bcrypt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB
var Rdb *redis.Client

type Settings struct {
    DB_HOST          string
    DB_NAME          string
    DB_USER          string
    DB_PASSWORD      string
    DB_PORT          string
}

func InitializeSettings() Settings {
    DB_HOST := os.Getenv("DB_HOST")
    DB_NAME := os.Getenv("DB_NAME")
    DB_USER := os.Getenv("DB_USER")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_PORT := os.Getenv("DB_PORT")

    var addr = os.Getenv("REDIS_PORT")
    var pass = os.Getenv("REDIS_PASSWORD")
    db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
    fmt.Println(addr, pass, db)
    Rdb = redis.NewClient(&redis.Options{
        Addr:     addr,
        Password: pass, // no password set
        DB:       db,   // use default DB
    })

    switch {
    case DB_HOST == "":
        fmt.Println("Environmet variable DB_HOST not set.")
        os.Exit(1)
    case DB_NAME == "":
        fmt.Println("Environmet variable DB_NAME not set.")
        os.Exit(1)
    case DB_USER == "":
        fmt.Println("Environmet variable DB_USER not set.")
        os.Exit(1)
    case DB_PASSWORD == "":
        fmt.Println("Environmet variable DB_PASSWORD not set.")
        os.Exit(1)
    }

    settings := Settings{
        DB_HOST:          DB_HOST,
        DB_NAME:          DB_NAME,
        DB_USER:          DB_USER,
        DB_PASSWORD:      DB_PASSWORD,
        DB_PORT:          DB_PORT,
    }

    return settings
}

func ConnectDataBase() {
    
    settings := InitializeSettings()
    

        url := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", settings.DB_USER, settings.DB_PASSWORD, settings.DB_HOST, settings.DB_PORT, settings.DB_NAME)
        fmt.Println(url)
        db, err := gorm.Open(mysql.Open(url), &gorm.Config{
            // NamingStrategy: schema.NamingStrategy{
            //  SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
            //  NoLowerCase: true, // skip the snake_casing of names
            //   },
        })
        if err != nil {
            panic("Failed to connect to database!")
        }
        db.AutoMigrate( &Article{}, &Comment{})
        DB = db
    
}