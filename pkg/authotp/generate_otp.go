package authotp

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateOtp() string {
	rand.Seed(time.Now().UnixNano())
	otp := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", otp)
}
