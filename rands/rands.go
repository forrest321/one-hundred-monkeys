package rands
//shamelessly copied/modified from http://pastie.org/1156941#35,37
//after reading discussion here: https://groups.google.com/forum/#!topic/golang-nuts/f0Y02OR06WM
import (
	"math/rand"
	"time"
	"fmt"
)

const chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890 abcdefghijklmnopqrstuvwxyz" +
	"~!@#$%^&*()-_+={}[]\\|<,>.?/\"';:`"

const MaxLen = 100

func init() {
	rand.Seed(int64(time.Now().Nanosecond()))
}

func RandStrings(N int) string {
	r := make([]string, N)
	ri := 0
	buf := make([]byte, MaxLen)
	known := map[string]bool{}

	for i := 0; i < N; i++ {
	retry:
		l := rand.Intn(MaxLen)
		for j := 0; j < l; j++ {
			buf[j] = chars[rand.Intn(len(chars))]
		}
		s := string(buf[0:l])
		if known[s] {
			goto retry
		}
		known[s] = true
		r[ri] = s
		ri++
	}
	return fmt.Sprintf("%s",r)
}