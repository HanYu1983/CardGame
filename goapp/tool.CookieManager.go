package hello

import (
    "net/http"
    "time"
)

type SimpleCookieManager struct {
    HttpRequest *http.Request
    ResponseWriter *http.ResponseWriter
}

var defaultCookieName = "__SimpleCookieManager__"

func (mgr *SimpleCookieManager) GetValue() (bool, string) {
    cookie, err := mgr.HttpRequest.Cookie(defaultCookieName)
    if err == nil { return true, cookie.Value } else { return false, "" }
}

func (mgr *SimpleCookieManager) SetValue(value string) {
    expiration := time.Now()
    expiration = expiration.AddDate(0, 0, 1)
    cookie := http.Cookie{Name: defaultCookieName, Value: value, Expires: expiration}
    http.SetCookie(*mgr.ResponseWriter, &cookie)
}

func (mgr *SimpleCookieManager) Clear(){
    cookie := http.Cookie{Name: defaultCookieName, Path: "/", MaxAge: -1}
    http.SetCookie(*mgr.ResponseWriter, &cookie)
}