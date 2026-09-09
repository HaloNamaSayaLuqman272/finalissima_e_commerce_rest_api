package middlewares

import (
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

// file "logger.go" akan digunakan sebagai middleware pencatat aktivitas setiap request
// yg masuk aplikasi ini
type LoggerConfig struct {
	Config middleware.RequestLoggerConfig
	// dalam type struct "LoggerConfig" hanya berisi field "Config" bertipe data
	// "middleware.RequestLoggerConfig" yg memiliki fungsi untuk menentukan apa saja
	// yg akan dicatat dan bagaimana bentuk formatnya
}

func (c *LoggerConfig) Init() echo.MiddlewareFunc {
	// fungsi "Init()" digunakan untuk membungkus konfigurasi menjadi middleware
	return middleware.RequestLoggerWithConfig(c.Config)
	// ".RequestLoggerWithConfig" fungsi bawaan "echo" yg menghasilkan middleware
	// logger berdasarkan konfigurasi yg diberikan
}
