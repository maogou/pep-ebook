package downloader

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewRequestViaCURL(t *testing.T) {
	curl := `
curl 'https://book.pep.com.cn/1212001301244/files/mobile/1.jpg?240828122428' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'Referer: https://book.pep.com.cn/1212001301244/mobile/index.html?u_atoken=tokenxyz' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36' \
  -H 'sec-ch-ua: "Chromium";v="134", "Not:A-Brand";v="24", "Google Chrome";v="134"' \
  -H 'sec-ch-ua-mobile: ?0'`
	data, err := ParseCURL(curl)
	assert.NoError(t, err)
	req, err := NewRequestViaCurl(data, "123456789", 1)
	assert.NoError(t, err)
	assert.Equal(t, req.Header.Get("referer"), "https://book.pep.com.cn/123456789/mobile/index.html?u_atoken=tokenxyz")

	req, err = NewRequestViaCurl(data, "00123456789", 1)
	assert.NoError(t, err)
	assert.Equal(t, req.Header.Get("referer"), "https://book.pep.com.cn/00123456789/mobile/index.html?u_atoken=tokenxyz")
}

func TestParseCURL(t *testing.T) {
	cases := []struct {
		name   string
		curl   string
		want   CurlData
		errMsg string
	}{
		{
			name: "valid_curl 1",
			curl: `curl 'https://book.pep.com.cn/1212001301244/files/mobile/1.jpg?240828122428' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'Referer: https://book.pep.com.cn/1212001301244/mobile/index.html?u_atoken=token' \
  -H 'User-Agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36' \
  -H 'sec-ch-ua: "Chromium";v="134", "Not:A-Brand";v="24", "Google Chrome";v="134"' \
  -H 'sec-ch-ua-mobile: ?0'`,
			want: CurlData{
				PageURLFormat: "https://book.pep.com.cn/%s/files/mobile/%d.jpg?240828122428",
				Headers: func() http.Header {
					h := make(http.Header)
					m := map[string]string{
						"Referer":            "https://book.pep.com.cn/%s/mobile/index.html?u_atoken=token",
						"User-Agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
						"sec-ch-ua":          "Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Google Chrome\";v=\"134",
						"sec-ch-ua-mobile":   "?0",
						"sec-ch-ua-platform": "macOS",
					}
					for k, v := range m {
						h.Add(k, v)
					}
					return h
				}(),
			},
		},
		{
			name: "valid_curl 2",
			curl: `curl 'https://book.pep.com.cn/1212001301245/files/mobile/11.jpg?240828111803' \
  -H 'accept: image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8' \
  -H 'accept-language: zh-CN,zh;q=0.9,en;q=0.8' \
  -b 'HMACCOUNT=account; cdn_sec_tc=xyz' \
  -H 'priority: i' \
  -H 'referer: https://book.pep.com.cn/1212001301245/mobile/index.html' \
  -H 'sec-ch-ua: "Chromium";v="134", "Not:A-Brand";v="24", "Google Chrome";v="134"' \
  -H 'sec-ch-ua-mobile: ?0' \
  -H 'sec-ch-ua-platform: "macOS"' \
  -H 'sec-fetch-dest: image' \
  -H 'sec-fetch-mode: no-cors' \
  -H 'sec-fetch-site: same-origin' \
  -H 'user-agent: Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36'`,
			want: CurlData{
				PageURLFormat: "https://book.pep.com.cn/%s/files/mobile/%d.jpg?240828111803",
				Headers: func() http.Header {
					h := make(http.Header)
					m := map[string]string{
						"accept":             "image/avif,image/webp,image/apng,image/svg+xml,image/*,*/*;q=0.8",
						"accept-language":    "zh-CN,zh;q=0.9,en;q=0.8",
						"priority":           "i",
						"referer":            "https://book.pep.com.cn/%s/mobile/index.html",
						"sec-ch-ua":          "Chromium\";v=\"134\", \"Not:A-Brand\";v=\"24\", \"Google Chrome\";v=\"134",
						"sec-ch-ua-mobile":   "?0",
						"sec-ch-ua-platform": "macOS",
						"sec-fetch-dest":     "image",
						"sec-fetch-mode":     "no-cors",
						"sec-fetch-site":     "same-origin",
						"user-agent":         "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/134.0.0.0 Safari/537.36",
						"Cookie":             "HMACCOUNT=account; cdn_sec_tc=xyz",
					}
					for k, v := range m {
						h.Add(k, v)
					}
					return h
				}(),
			},
		},
		{
			name:   "invalid_curl_format",
			curl:   "invalid curl command",
			errMsg: "invalid URL format",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseCURL(tt.curl)

			if tt.errMsg != "" {
				require.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want.PageURLFormat, got.PageURLFormat)
			assert.Len(t, got.Headers, len(tt.want.Headers))
			for k, v := range tt.want.Headers {
				assert.Equal(t, v, got.Headers[k])
			}
			assert.Equal(t, tt.want.Headers, got.Headers)
		})
	}
}
