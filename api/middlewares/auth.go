package middlewares

import (
	"context"
	"errors"
	"finalissima_e_commerce_rest_api/database/models"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
)

// file "auth.go" package "middlewares" adalah berisi proses klaim token terjadi
type JWTCustomsClaims struct {
	ID   int         `json:"id"`
	Role models.Role `json:"role"`
	jwt.RegisteredClaims
}

type JWTConfig struct {
	SecretKey      string
	ExpireDuration int
}

type ContextKey string

const userContextKey = ContextKey("user")

func (jwtConfig *JWTConfig) Init() echojwt.Config {
	// kita mendeklarasikan sebuah fungsi dan memberi nama "Init" tanpa ada parameter penerima
	// kita membutuhkan keberadaan variabel penerima "jwtConfig" bertipe data "*JWTConfig"
	// ini adalah cara mengaitkan sebuah fungsi ke dalam struct "JWTConfig" tujuannya
	// untuk mengubah nilai asli dari properti JWTConfig, bukan nilai salinan
	// fungsi "Init" ini rencananya akan mengembalikan "echojwt.Config" yg bermaksud
	// "Init" akan menghasilkan sebuah objek konfigurasi yg digunakan untuk
	// mengatur cara kerja middleware JWT
	return echojwt.Config{
		// mengembalikan "echojwt.Config" sebuah struct konfigurasi bagaimana middleware
		// JWT harus bekerja
		NewClaimsFunc: func(c *echo.Context) jwt.Claims {
			// "func(c *echo.Context)" fungsi yg menerima "*echo.Context" dan mengembalikan
			// sesuatu bertipe jwt.Claims
			return new(JWTCustomsClaims)
			// dilakukan pengembalian pointer kosongan ke struct "JWTCustomsClaims"
			// field "NewClaimsFunc" digunakan untuk memproses token yg masuk
		},
		SigningKey: []byte(jwtConfig.SecretKey),
		// "SigningKey" berisi kunci rahasia untuk verifikasi keaslian token JWT
		// "[]byte(jwtConfig.SecretKey)" mengakses field "SecretKet" melalui
		// variabel penerima "jwtConfig" kemudian melakukan konversi tipe "SecretKey"
		// yg bertipe "string" menjadi "[]byte", sesuai dengan aturan dari "SigningKey"
		// yg mewajibkan untuk bertipe "[]byte"
	}
}

func (jwtConfig *JWTConfig) GenerateToken(userID int, role models.Role) (string, error) {
	// kali ini kita mendeklarasikan sebuah fungsi untuk melakukan generate token dan
	// kita beri nama fungsi ini "GenerateToken" memiliki parameter penerima "userID" dan
	// "role", kemudian fungsi ini rencananya akan mengembalikan data bertipe "string"
	// dan "error" apabila terjadi masalah
	expire := jwt.NewNumericDate(time.Now().Local().Add(time.Minute * time.Duration(int64(jwtConfig.ExpireDuration))))
	// "int64(jwtConfig.ExpireDuration)" mengubah tipe data "int" dari "ExpireDuration" menjadi "int64"
	// "time.Duration(...)" mengubah bentuk tipe data "int64" menjadi "time.Duration" yg masih satuan "nanosecond"
	// "time.Minute * ..." kita mengubah "time.Duration" yg masih "nanosecond" menjadi "time.Minute"
	// "time.Now()" adalah fungsi package "time" yg mengembalikan waktu saat ini
	// ".Local()" adalah method yg meng-konversi waktu ke zona waktu lokal sistem
	// ".Add(...)" adalah method dari "time.Time" untuk menambahkan durasi ke suatu waktu
	// maka "time.Now().Local()" yg merepresentsikan waktu sekarang
	// dan ".Add()" yg merepresentsikan durasi waktu yg telah dihitung, menghasilkan titik waktu
	// yg akan merepresentasikan waktu token akan kadaluwarsa
	// "jwt.NewNumericDate" adalah fungsi dari package "JWT" untuk konversi "time.Time" menjadi
	// "*jwt.NumericDate" yg merupakan format waktu standar spesifikasi JWT untuk klaim, seperti expiration time
	// "expire" variabel untuk menampung proses penghitungan kapan token JWT akan kadaluwarsa

	claims := &JWTCustomsClaims{
		ID: userID,
		// klaim ini milik siapa
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: expire,
			// kadaluwarsa token dimulai
		},
		Role: role,
	}

	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// token mulai dibentuk
	token, err := rawToken.SignedString([]byte(jwtConfig.SecretKey))
	if err != nil {
		return "", err
	}

	return token, nil
}

func GetUser(ctx context.Context) (*JWTCustomsClaims, error) {
	user, ok := ctx.Value(userContextKey).(*jwt.Token)
	if !ok || user == nil {
		return nil, errors.New("invalid token")
	}

	claims, ok := user.Claims.(*JWTCustomsClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	return claims, nil
	// fungsi "GetUser" ini untuk mengeluarkan data user dari token JWT yg berhasil divalidasi
}

func GetUserID(ctx context.Context) (int, error) {
	claims, err := GetUser(ctx)
	if err != nil {
		return 0, errors.New("invalid token")
	}

	return claims.ID, nil
}

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		if user == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		ctx := context.WithValue(c.Request().Context(), userContextKey, user)
		c.SetRequest(c.Request().WithContext(ctx))

		userData, err := GetUser(ctx)
		if userData == nil || err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		c.Set("userData", userData)
		return next(c)
	}
}

func VerifyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		user, err := GetUser(c.Request().Context())
		if err != nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}

		if user.Role != models.Admin {
			return c.JSON(http.StatusForbidden, map[string]string{
				"message": "access denied",
			})
		}

		return next(c)
	}

}
