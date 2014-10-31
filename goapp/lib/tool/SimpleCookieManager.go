package tool

import (
    "net/http"
    "time"
)

type SimpleCookieManager struct {
}

var defaultCookieName = "__SimpleCookieManager__"

func (mgr *SimpleCookieManager) GetValue(sys ISystem) (bool, string) {
    cookie, err := sys.GetRequest().Cookie(defaultCookieName)
    if err == nil { return true, cookie.Value } else { return false, "" }
}

func (mgr *SimpleCookieManager) SetValue(sys ISystem, value string) {
    expiration := time.Now()
    expiration = expiration.AddDate(0, 0, 1)
    cookie := http.Cookie{Name: defaultCookieName, Value: value, Expires: expiration}
    http.SetCookie(sys.GetResponse(), &cookie)
}

func (mgr *SimpleCookieManager) Clear(sys ISystem){
    cookie := http.Cookie{Name: defaultCookieName, Path: "/", MaxAge: -1}
    http.SetCookie(sys.GetResponse(), &cookie)
}