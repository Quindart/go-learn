package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// Khóa bí mật dùng để ký token khi đăng nhập và kiểm tra chữ ký khi nhận request.
var secretKey = "superkey"

// Claims là dữ liệu nằm trong token. Struct này chứa các trường cần đọc khi xác thực.
type tokenClaims struct {
	UserID int64 `json:"userId"`
	// Các trường JWT chuẩn, trong đó có thời điểm hết hạn (exp).
	jwt.RegisteredClaims
}

// GenerateToken tạo JWT chứa email, ID người dùng và có hiệu lực trong 2 giờ.
func GenerateToken(email string, userID int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	// Ký bằng khóa bí mật để phát hiện token bị chỉnh sửa; dữ liệu không được mã hóa.
	return token.SignedString([]byte(secretKey))
}

// VerifyToken kiểm tra JWT và trả về ID người dùng nếu token hợp lệ.
// Tham số token chỉ chứa chuỗi JWT, không gồm tiền tố "Bearer ".
func VerifyToken(token string) (int64, error) {
	// Truyền con trỏ để thư viện ghi dữ liệu đọc từ token vào claims.
	claims := &tokenClaims{}

	// Đọc dữ liệu, kiểm tra chữ ký và kiểm tra các điều kiện của token.
	parsedToken, err := jwt.ParseWithClaims(
		token,
		claims,
		func(_ *jwt.Token) (any, error) {
			// Thư viện gọi hàm này để lấy khóa kiểm tra chữ ký.
			// Dấu _ nghĩa là không sử dụng tham số *jwt.Token được truyền vào.
			return []byte(secretKey), nil
		},
		// Chỉ chấp nhận HS256, đúng với thuật toán dùng trong GenerateToken.
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		// Bắt buộc có exp; thư viện cũng kiểm tra token đã hết hạn hay chưa.
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		// Token sai định dạng, sai chữ ký hoặc hết hạn sẽ bị từ chối ở đây.
		return 0, err
	}
	// Chỉ chấp nhận token hợp lệ và ID người dùng là số dương.
	if !parsedToken.Valid || claims.UserID <= 0 {
		return 0, errors.New("invalid token")
	}

	// Handler dùng ID này để xác định người đang gửi request.
	return claims.UserID, nil
}
