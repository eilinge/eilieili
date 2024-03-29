package comm

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"

	"eilieili/configs"
	"eilieili/models"

	"github.com/kataras/iris/sessions"
)

// MySessions ...
var MySessions = sessions.New(sessions.Config{
	Cookie: configs.CookieName,
	Encode: configs.SecureCookie.Encode,
	Decode: configs.SecureCookie.Decode,
})

// ClientIP get user host
func ClientIP(request *http.Request) string {
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

// Redirect ...
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}

// GetLoginUser ...
func GetLoginUser(request *http.Request) *models.ObjLoginuser {
	c, err := request.Cookie(configs.CookieName)
	if err != nil {
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil {
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}

	loginuser := &models.ObjLoginuser{
		Uid:      uid,
		Username: params.Get("username"),
		Now:      now,
		Ip:       ClientIP(request),
		Sign:     params.Get("sign"),
	}

	sign := createLoginuserSign(loginuser)
	if sign != loginuser.Sign {
		log.Println("func_web GetLoginuser createloginusersign not signed", sign, loginuser.Sign)
		return nil
	}
	return loginuser
}

// SetLoginuser set login user info to cookie
func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginuser) {
	if loginuser == nil || loginuser.Uid < 1 {
		c := &http.Cookie{
			Name:   configs.CookieName,
			Value:  "",
			Path:   "/",
			MaxAge: -1,
		}
		http.SetCookie(writer, c)
		return
	}
	if loginuser.Sign == "" {
		loginuser.Sign = createLoginuserSign(loginuser)
	}
	params := url.Values{}
	params.Add("uid", strconv.Itoa(loginuser.Uid))
	params.Add("username", loginuser.Username)
	params.Add("now", strconv.Itoa(loginuser.Now))
	params.Add("ip", loginuser.Ip)
	params.Add("sign", loginuser.Sign)

	c := &http.Cookie{
		Name:  configs.CookieName,
		Value: params.Encode(),
		Path:  "/",
	}
	http.SetCookie(writer, c)
}

func createLoginuserSign(loginuser *models.ObjLoginuser) string {
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d",
		loginuser.Uid, loginuser.Username, configs.SignSecret, loginuser.Now)

	sign := fmt.Sprintf("%x", md5.Sum([]byte(str)))

	return sign
}
