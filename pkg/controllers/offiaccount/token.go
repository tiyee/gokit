package offiaccount

import (
	context2 "context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/tiyee/gokit/pkg/component/log"
	"github.com/tiyee/gokit/pkg/component/redis"
	"github.com/tiyee/gokit/pkg/consts"
	"io"

	"net/http"
	"time"
)

const WxAccessTokenUrlFmt = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s"
const cacheKey = "access_token"

type AccessTokenValue struct {
	Token  string `json:"token"`
	Expire int64  `json:"expire"`
}

type OrigainalAccessToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

func GetAccessToken() (string, error) {

	if entry, err := redis.RedisClient.Get(context2.Background(), cacheKey).Bytes(); err == nil {
		var cacheValue AccessTokenValue
		if err := json.Unmarshal(entry, &cacheValue); err == nil {
			if cacheValue.Expire > time.Now().Unix() {
				log.Error("get access_token from cache")

				return cacheValue.Token, nil
			}
		}
	}
	return getAccessToken()
}
func getAccessToken() (string, error) {
	url := fmt.Sprintf(WxAccessTokenUrlFmt, consts.WxAppId, consts.WxAppSecret)
	res, err := http.Get(url)
	if err != nil {
		log.Error("A error occurred!")

		return "", err
	}
	defer res.Body.Close()
	if data, err := io.ReadAll(res.Body); err == nil {
		var accessToken OrigainalAccessToken
		if err := json.Unmarshal(data, &accessToken); err == nil {
			cacheValue := AccessTokenValue{
				accessToken.AccessToken,
				time.Now().Unix() + 6400,
			}
			if value, err := json.Marshal(cacheValue); err == nil {
				if err := redis.RedisClient.Set(context2.Background(), cacheKey, value, time.Second*7100).Err(); err != nil {
					log.Error(err.Error())
				}
			}
			return accessToken.AccessToken, nil
		}
		{
			return accessToken.AccessToken, nil
		}
	}
	return "", errors.New("get access_token error")
}
