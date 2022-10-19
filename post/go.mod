module forPractice/projectGolang/BlogWebsite/post

go 1.18

require (
	forPractice/projectGolang/BlogWebsite/authorization v0.0.0-00010101000000-000000000000
	forPractice/projectGolang/BlogWebsite/common v0.0.0-00010101000000-000000000000
	github.com/gofiber/fiber/v2 v2.38.1
	github.com/joho/godotenv v1.4.0
	gorm.io/driver/mysql v1.4.3
	gorm.io/gorm v1.24.0
)

require (
	github.com/andybalholm/brotli v1.0.4 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/klauspost/compress v1.15.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.40.0 // indirect
	github.com/valyala/tcplisten v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20221012134737-56aed061732a // indirect
	golang.org/x/sys v0.0.0-20220227234510-4e6760a101f9 // indirect
)

replace forPractice/projectGolang/BlogWebsite/common => ../common

replace forPractice/projectGolang/BlogWebsite/authorization => ../authorization
