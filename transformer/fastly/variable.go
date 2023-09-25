package fastly

import (
	"errors"
	"fmt"

	"github.com/ysugimoto/vintage"
	"github.com/ysugimoto/vintage/transformer/core"
)

type FastlyVariable struct {
	CoreVariable *core.CoreVariable
}

func NewFastlyVariable() *FastlyVariable {
	return &FastlyVariable{
		&core.CoreVariable{},
	}
}

func (v *FastlyVariable) Get(name string) (*core.ExpressionValue, error) {
	// Lookup variable for fastly specific field

	switch name {
	// User-Agent related variables
	case core.CLIENT_CLASS_BOT:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsBot()"), nil
	case core.CLIENT_CLASS_BROWSER:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsBrowser()"), nil
	case core.CLIENT_DISPLAY_TOUCHSCREEN:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsTouchScreen()"), nil
	case core.CLIENT_PLATFORM_EREADER:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsEReader()"), nil
	case core.CLIENT_PLATFORM_GAMECONSOLE:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsGameConsole()"), nil
	case core.CLIENT_PLATFORM_MOBILE:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsMobile()"), nil
	case core.CLIENT_PLATFORM_SMARTTV:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsSmartTV()"), nil
	case core.CLIENT_PLATFORM_TABLET:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsTablet()"), nil
	case core.CLIENT_PLATFORM_TVPLAYER:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.IsTvPlayer()"), nil
	case core.CLIENT_BOT_NAME:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.BotName()"), nil
	case core.CLIENT_BROWSER_NAME:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.BrowserName()"), nil
	case core.CLIENT_BROWSER_VERSION:
		return core.NewExpressionValue(vintage.BOOL, "ctx.UserAgent.BrowserVersion()"), nil
	case core.CLIENT_OS_NAME:
		return core.NewExpressionValue(vintage.STRING, "ctx.UserAgent.OSName()"), nil
	case core.CLIENT_OS_VERSION:
		return core.NewExpressionValue(vintage.STRING, "ctx.UserAgent.OSVersion()"), nil
	// Always empty string
	case core.CLIENT_PLATFORM_HWTYPE:
		return core.NewExpressionValue(vintage.STRING, ""), nil

	// Fastly Request related variables
	case core.FASTLY_INFO_IS_H2:
		return core.NewExpressionValue(vintage.BOOL, "ctx.Request.ProtoMajor == 2"), nil
	case core.FASTLY_INFO_IS_H3:
		return core.NewExpressionValue(vintage.BOOL, "ctx.Request.ProtoMajor == 3"), nil
	case core.FASTLY_INFO_HOST_HEADER:
		return core.NewExpressionValue(vintage.STRING, "ctx.Request.Host"), nil

	// Fastly GeoLocation related variables
	case core.CLIENT_GEO_LATITUDE:
		return core.NewExpressionValue(vintage.FLOAT, "ctx.Geo.Latitude"), nil
	case core.CLIENT_GEO_LONGITUDE:
		return core.NewExpressionValue(vintage.FLOAT, "ctx.Geo.Longitude"), nil
	case core.CLIENT_AS_NUMBER:
		return core.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.AsNumber)"), nil
	case core.CLIENT_AS_NAME:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.AsName"), nil
	case core.CLIENT_GEO_AREA_CODE:
		return core.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.AreaCode)"), nil
	case core.CLIENT_GEO_METRO_CODE:
		return core.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.MetroCode)"), nil
	case core.CLIENT_GEO_UTC_OFFSET:
		return core.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.UTCOffset)"), nil
	case core.CLIENT_GEO_GMT_OFFSET:
		return core.NewExpressionValue(vintage.INTEGER, "int64(ctx.Geo.UTCOffset)"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case core.CLIENT_GEO_CITY,
		core.CLIENT_GEO_CITY_ASCII,
		core.CLIENT_GEO_CITY_LATIN1,
		core.CLIENT_GEO_CITY_UTF8:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.City"), nil
	case core.CLIENT_GEO_CONN_SPEED:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.ConnSpeed"), nil
	case core.CLIENT_GEO_CONN_TYPE:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.ConnType"), nil
	case core.CLIENT_GEO_CONTINENT_CODE:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.ContinentCode"), nil
	case core.CLIENT_GEO_COUNTRY_CODE:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.CountryCode"), nil
	case core.CLIENT_GEO_COUNTRY_CODE3:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.CountryCode3"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case core.CLIENT_GEO_COUNTRY_NAME,
		core.CLIENT_GEO_COUNTRY_NAME_ASCII,
		core.CLIENT_GEO_COUNTRY_NAME_LATIN1,
		core.CLIENT_GEO_COUNTRY_NAME_UTF8:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.CountryName"), nil
	// @Tentative
	case core.CLIENT_GEO_IP_OVERRIDE:
		return core.NewExpressionValue(vintage.STRING, ""), nil
	case core.CLIENT_GEO_POSTAL_CODE:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.PostalCode"), nil
	case core.CLIENT_GEO_PROXY_DESCRIPTION:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.ProxyDescription"), nil
	case core.CLIENT_GEO_PROXY_TYPE:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.ProxyType"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case core.CLIENT_GEO_REGION,
		core.CLIENT_GEO_REGION_ASCII,
		core.CLIENT_GEO_REGION_LATIN1,
		core.CLIENT_GEO_REGION_UTF8:
		return core.NewExpressionValue(vintage.STRING, "ctx.Geo.Region"), nil

	// Request related values
	case core.REQ_HEADER_BYTES_READ:
		return core.NewExpressionValue(vintage.INTEGER, "ctx.RequestHeaderBytes"), nil
	case core.REQ_RESTARTS:
		return core.NewExpressionValue(vintage.INTEGER, "int64(ctx.Restarts)"), nil
	case core.CLIENT_IDENTITY:
		return core.NewExpressionValue(vintage.STRING, "ctx.ClientIdentity()"), nil
	case core.CLIENT_IP:
		return core.NewExpressionValue(vintage.IP, "ctx.ClientIP()"), nil
	case core.REQ_BODY:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.STRING,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBody()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.REQ_BODY_BASE64:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.STRING,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBodyBase64()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.REQ_DIGEST:
		return core.NewExpressionValue(vintage.STRING, "ctx.RequestDigest()"), nil
	case core.REQ_METHOD:
		return core.NewExpressionValue(vintage.STRING, "ctx.Request.Method"), nil
	case core.REQ_POSTBODY:
		return v.Get("req.body")
	case core.REQ_PROTO:
		return core.NewExpressionValue(vintage.STRING, "ctx.Request.Proto"), nil
	case core.REQ_REQUEST:
		return v.Get("req.method")

	// @Tentative
	case core.REQ_SERVICE_ID:
		return core.NewExpressionValue(vintage.STRING, ""), nil

	case core.REQ_TOPURL: // FIXME: what is the difference of req.url ?
		return v.Get("req.url")
	case core.REQ_URL:
		return core.NewExpressionValue(vintage.STRING, "ctx.RequestURL()"), nil
	case core.REQ_URL_BASENAME:
		return core.NewExpressionValue(
			vintage.STRING,
			"filepath.Base(ctx.Request.URL.Path)",
			core.Dependency("path/filepath", ""),
		), nil
	case core.REQ_URL_DIRNAME:
		return core.NewExpressionValue(
			vintage.STRING,
			"filepath.Dir(ctx.Request.URL.Path)",
			core.Dependency("path/filepath", ""),
		), nil
	case core.REQ_URL_EXT:
		return core.NewExpressionValue(
			vintage.STRING,
			`strings.TrimPrefix(filepath.Ext(ctx.Request.URL.Path), ".")`,
			core.Dependency("path/filepath", ""),
			core.Dependency("strings", ""),
		), nil
	case core.REQ_URL_PATH:
		return core.NewExpressionValue(vintage.STRING, "ctx.Request.URL.Path"), nil
	case core.REQ_URL_QS:
		return core.NewExpressionValue(vintage.STRING, "ctx.Request.URL.RawQuery"), nil
	// @Tentative
	case core.REQ_VCL:
		return core.NewExpressionValue(vintage.STRING, "vintage.vcl.transpile"), nil
	// Precalculated: md5("vintage.vcl.transpile")
	case core.REQ_VCL_MD5:
		return core.NewExpressionValue(vintage.STRING, "ce85569c98bce4df6334206466e65dc8"), nil
	case core.REQ_XID:
		return core.NewExpressionValue(vintage.STRING, "vintage.GenerateXid()"), nil

	// Backend request related variables
	case core.BEREQ_BODY_BYTES_WRITTEN:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.INTEGER,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBodyBytesWritten()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.BEREQ_BYTES_WRITTEN:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.INTEGER,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBytesWritten()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.BEREQ_HEADER_BYTES_WRITTEN:
		return core.NewExpressionValue(vintage.INTEGER, "ctx.BackendHeaderBytesWritten()"), nil

	// Cache object related variables
	// On Fastly runtime, obj.xxx will be treated as backend response
	case core.OBJ_AGE:
		return core.NewExpressionValue(vintage.RTIME, "ctx.ObjectAge()"), nil
	case core.OBJ_CACHEABLE:
		return core.NewExpressionValue(vintage.BOOL, "ctx.ObjectCacheable()"), nil
	case core.OBJ_HITS:
		return core.NewExpressionValue(vintage.INTEGER, "ctx.ObjectHits()"), nil
	// @Tentative
	case core.OBJ_ENTERED:
		return core.NewExpressionValue(vintage.RTIME, "time.Duration(0)"), nil
	case core.OBJ_GRACE:
		return core.NewExpressionValue(vintage.RTIME, "time.Duration(0)"), nil
	case core.OBJ_IS_PCI:
		return core.NewExpressionValue(vintage.BOOL, "false"), nil
	case core.OBJ_LASTUSE:
		return core.NewExpressionValue(vintage.RTIME, "time.Duration(0)"), nil
	case core.REQ_IS_IPV6:
		return core.NewExpressionValue(vintage.BOOL, "ctx.IsIpv6()"), nil
	case core.REQ_IS_PURGE:
		return core.NewExpressionValue(vintage.BOOL, `(ctx.Request.Method == "PURGE")`), nil
	// @Tentative
	case core.REQ_BACKEND_IP:
		return core.NewExpressionValue(vintage.IP, "net.IPv4(127, 0, 0, 1)", core.Dependency("net", "")), nil
	// @Tentative
	case core.REQ_BACKEND_IS_CLUSTER:
		return core.NewExpressionValue(vintage.BOOL, "false"), nil
	// @Tentative
	case core.REQ_BACKEND_PORT:
		return core.NewExpressionValue(vintage.INTEGER, "334"), nil

	case core.REQ_BODY_BYTES_READ:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.INTEGER,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBodyBytesRead()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.REQ_BYTES_READ:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.INTEGER,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBytesRead()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.RESP_PROTO:
		return core.NewExpressionValue(vintage.STRING, "ctx.Response.Request.Proto"), nil
	case core.RESP_RESPONSE:
		v := core.Temporary()
		return core.NewExpressionValue(
			vintage.STRING,
			v,
			core.Prepare(
				fmt.Sprintf("%s, err := ctx.ResponseBody()", v),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case core.RESP_STATUS:
		return core.NewExpressionValue(vintage.INTEGER, "ctx.Response.StatusCode"), nil
	}

	// Look up core variables
	return v.CoreVariable.Get(name)
}

func (v *FastlyVariable) Set(name string, value *core.ExpressionValue) error {
	return errors.New("Not Implemented")
}

func (v *FastlyVariable) Unset(name string) error {
	return errors.New("Not Implemented")
}

var _ core.Variable = (*FastlyVariable)(nil)
