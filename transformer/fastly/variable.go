package fastly

import (
	"errors"
	"fmt"

	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/transformer"
	tf "github.com/ysugimoto/vintage/transformer"
)

type FastlyVariable struct {
	CoreVariable *tf.CoreVariable
}

func NewFastlyVariable() *FastlyVariable {
	return &FastlyVariable{
		&tf.CoreVariable{},
	}
}

func (v *FastlyVariable) Get(name string) (*tf.ExpressionValue, error) {
	// Lookup variable for fastly specific field

	switch name {
	// User-Agent related variables
	case tf.CLIENT_CLASS_BOT:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsBot()"), nil
	case tf.CLIENT_CLASS_BROWSER:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsBrowser()"), nil
	case tf.CLIENT_DISPLAY_TOUCHSCREEN:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsTouchScreen()"), nil
	case tf.CLIENT_PLATFORM_EREADER:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsEReader()"), nil
	case tf.CLIENT_PLATFORM_GAMECONSOLE:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsGameConsole()"), nil
	case tf.CLIENT_PLATFORM_MOBILE:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsMobile()"), nil
	case tf.CLIENT_PLATFORM_SMARTTV:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsSmartTV()"), nil
	case tf.CLIENT_PLATFORM_TABLET:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsTablet()"), nil
	case tf.CLIENT_PLATFORM_TVPLAYER:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsTvPlayer()"), nil
	case tf.CLIENT_BOT_NAME:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.BotName()"), nil
	case tf.CLIENT_BROWSER_NAME:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.BrowserName()"), nil
	case tf.CLIENT_BROWSER_VERSION:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.BrowserVersion()"), nil
	case tf.CLIENT_OS_NAME:
		return tf.NewExpressionValue(vintage.STRING, "ctx.UserAgent.OSName()"), nil
	case tf.CLIENT_OS_VERSION:
		return tf.NewExpressionValue(vintage.STRING, "ctx.UserAgent.OSVersion()"), nil
	// Always empty string
	case tf.CLIENT_PLATFORM_HWTYPE:
		return tf.NewExpressionValue(vintage.STRING, ""), nil

	// Fastly Request related variables
	case tf.FASTLY_INFO_IS_H2:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.Request.ProtoMajor == 2"), nil
	case tf.FASTLY_INFO_IS_H3:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.Request.ProtoMajor == 3"), nil
	case tf.FASTLY_INFO_HOST_HEADER:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Request.Host"), nil

	// Fastly GeoLocation related variables
	case tf.CLIENT_GEO_LATITUDE:
		return tf.NewExpressionValue(vintage.FLOAT, "ctx.Geo.Latitude"), nil
	case tf.CLIENT_GEO_LONGITUDE:
		return tf.NewExpressionValue(vintage.FLOAT, "ctx.Geo.Longitude"), nil
	case tf.CLIENT_AS_NUMBER:
		return tf.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.AsNumber)"), nil
	case tf.CLIENT_AS_NAME:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.AsName"), nil
	case tf.CLIENT_GEO_AREA_CODE:
		return tf.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.AreaCode)"), nil
	case tf.CLIENT_GEO_METRO_CODE:
		return tf.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.MetroCode)"), nil
	case tf.CLIENT_GEO_UTC_OFFSET:
		return tf.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.UTCOffset)"), nil
	case tf.CLIENT_GEO_GMT_OFFSET:
		return tf.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.UTCOffset)"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case tf.CLIENT_GEO_CITY,
		tf.CLIENT_GEO_CITY_ASCII,
		tf.CLIENT_GEO_CITY_LATIN1,
		tf.CLIENT_GEO_CITY_UTF8:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.City"), nil
	case tf.CLIENT_GEO_CONN_SPEED:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.ConnSpeed"), nil
	case tf.CLIENT_GEO_CONN_TYPE:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.ConnType"), nil
	case tf.CLIENT_GEO_CONTINENT_CODE:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.ContinentCode"), nil
	case tf.CLIENT_GEO_COUNTRY_CODE:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.CountryCode"), nil
	case tf.CLIENT_GEO_COUNTRY_CODE3:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.CountryCode3"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case tf.CLIENT_GEO_COUNTRY_NAME,
		tf.CLIENT_GEO_COUNTRY_NAME_ASCII,
		tf.CLIENT_GEO_COUNTRY_NAME_LATIN1,
		tf.CLIENT_GEO_COUNTRY_NAME_UTF8:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.CountryName"), nil
	// @Tentative
	case tf.CLIENT_GEO_IP_OVERRIDE:
		return tf.NewExpressionValue(vintage.STRING, ""), nil
	case tf.CLIENT_GEO_POSTAL_CODE:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.PostalCode"), nil
	case tf.CLIENT_GEO_PROXY_DESCRIPTION:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.ProxyDescription"), nil
	case tf.CLIENT_GEO_PROXY_TYPE:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.ProxyType"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case tf.CLIENT_GEO_REGION,
		tf.CLIENT_GEO_REGION_ASCII,
		tf.CLIENT_GEO_REGION_LATIN1,
		tf.CLIENT_GEO_REGION_UTF8:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Geo.Region"), nil

	// Request related values
	case tf.REQ_HEADER_BYTES_READ:
		return tf.NewExpressionValue(vintage.INTEGER, "ctx.RequestHeaderBytes"), nil
	case tf.REQ_RESTARTS:
		return tf.NewExpressionValue(vintage.INTEGER, "int64(ctx.Restarts)"), nil
	case tf.CLIENT_IDENTITY:
		return tf.NewExpressionValue(vintage.STRING, "ctx.ClientIdentity()"), nil
	case tf.CLIENT_IP:
		return tf.NewExpressionValue(vintage.IP, "ctx.ClientIP()"), nil
	case tf.REQ_BODY:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.STRING,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBody()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.REQ_BODY_BASE64:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.STRING,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBodyBase64()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.REQ_DIGEST:
		return tf.NewExpressionValue(vintage.STRING, "ctx.RequestDigest()"), nil
	case tf.REQ_METHOD:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Request.Method"), nil
	case tf.REQ_POSTBODY:
		return v.Get("req.body")
	case tf.REQ_PROTO:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Request.Proto"), nil
	case tf.REQ_REQUEST:
		return v.Get("req.method")

	// @Tentative
	case tf.REQ_SERVICE_ID:
		return tf.NewExpressionValue(vintage.STRING, ""), nil

	case tf.REQ_TOPURL: // FIXME: what is the difference of req.url ?
		return v.Get("req.url")
	case tf.REQ_URL:
		return tf.NewExpressionValue(vintage.STRING, "ctx.RequestURL()"), nil
	case tf.REQ_URL_BASENAME:
		return tf.NewExpressionValue(
			vintage.STRING,
			"filepath.Base(ctx.Request.URL.Path)",
			tf.Dependency("path/filepath", ""),
		), nil
	case tf.REQ_URL_DIRNAME:
		return tf.NewExpressionValue(
			vintage.STRING,
			"filepath.Dir(ctx.Request.URL.Path)",
			tf.Dependency("path/filepath", ""),
		), nil
	case tf.REQ_URL_EXT:
		return tf.NewExpressionValue(
			vintage.STRING,
			`strings.TrimPrefix(filepath.Ext(ctx.Request.URL.Path), ".")`,
			tf.Dependency("path/filepath", ""),
			tf.Dependency("strings", ""),
		), nil
	case tf.REQ_URL_PATH:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Request.URL.Path"), nil
	case tf.REQ_URL_QS:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Request.URL.RawQuery"), nil
	// @Tentative
	case tf.REQ_VCL:
		return tf.NewExpressionValue(vintage.STRING, "vintage.vcl.transpile"), nil
	// Precalculated: md5("vintage.vcl.transpile")
	case tf.REQ_VCL_MD5:
		return tf.NewExpressionValue(vintage.STRING, "ce85569c98bce4df6334206466e65dc8"), nil
	case tf.REQ_XID:
		return tf.NewExpressionValue(vintage.STRING, "vintage.GenerateXid()"), nil

	// Backend request related variables
	case tf.BEREQ_BODY_BYTES_WRITTEN:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.INTEGER,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBodyBytesWritten()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.BEREQ_BYTES_WRITTEN:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.INTEGER,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBytesWritten()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.BEREQ_HEADER_BYTES_WRITTEN:
		return tf.NewExpressionValue(vintage.INTEGER, "ctx.BackendHeaderBytesWritten()"), nil

	// Cache object related variables
	// On Fastly runtime, obj.xxx will be treated as backend response
	case tf.OBJ_AGE:
		return tf.NewExpressionValue(vintage.RTIME, "ctx.ObjectAge()"), nil
	case tf.OBJ_CACHEABLE:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.ObjectCacheable()"), nil
	case tf.OBJ_HITS:
		return tf.NewExpressionValue(vintage.INTEGER, "ctx.ObjectHits()"), nil
	// @Tentative
	case tf.OBJ_ENTERED:
		return tf.NewExpressionValue(vintage.RTIME, "time.Duration(0)"), nil
	case tf.OBJ_GRACE:
		return tf.NewExpressionValue(vintage.RTIME, "time.Duration(0)"), nil
	case tf.OBJ_IS_PCI:
		return tf.NewExpressionValue(vintage.BOOL, "false"), nil
	case tf.OBJ_LASTUSE:
		return tf.NewExpressionValue(vintage.RTIME, "time.Duration(0)"), nil
	case tf.REQ_IS_IPV6:
		return tf.NewExpressionValue(vintage.BOOL, "ctx.IsIpv6()"), nil
	case tf.REQ_IS_PURGE:
		return tf.NewExpressionValue(vintage.BOOL, `(ctx.Request.Method == "PURGE")`), nil
	// @Tentative
	case tf.REQ_BACKEND_IP:
		return tf.NewExpressionValue(vintage.IP, "net.IPv4(127, 0, 0, 1)", tf.Dependency("net", "")), nil
	// @Tentative
	case tf.REQ_BACKEND_IS_CLUSTER:
		return tf.NewExpressionValue(vintage.BOOL, "false"), nil
	// @Tentative
	case tf.REQ_BACKEND_PORT:
		return tf.NewExpressionValue(vintage.INTEGER, "334"), nil

	case tf.REQ_BODY_BYTES_READ:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.INTEGER,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBodyBytesRead()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.REQ_BYTES_READ:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.INTEGER,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBytesRead()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.RESP_PROTO:
		return tf.NewExpressionValue(vintage.STRING, "ctx.Response.Request.Proto"), nil
	case tf.RESP_RESPONSE:
		v := tf.Temporary()
		return tf.NewExpressionValue(
			vintage.STRING,
			v,
			tf.Prepare(
				fmt.Sprintf("%s, err := ctx.ResponseBody()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case tf.RESP_STATUS:
		return tf.NewExpressionValue(vintage.INTEGER, "ctx.Response.StatusCode"), nil
	}

	// Look up core variables
	return v.CoreVariable.Get(name)
}

func (v *FastlyVariable) Set(name string, value *transformer.ExpressionValue) error {
	return errors.New("Not Implemented")
}

func (v *FastlyVariable) Unet(name string) error {
	return errors.New("Not Implemented")
}
