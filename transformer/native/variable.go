package native

import (
	"fmt"

	"github.com/ysugimoto/vintage/transformer/core"
	"github.com/ysugimoto/vintage/transformer/value"
	v "github.com/ysugimoto/vintage/transformer/variable"
)

type NativeVariable struct {
	CoreVariable *core.CoreVariable
}

func NewNativeVariable() *NativeVariable {
	return &NativeVariable{
		core.NewCoreVariables(),
	}
}

// nolint:funlen,gocyclo
func (fv *NativeVariable) Get(name string) (*value.Value, error) {
	// Lookup variable for fastly specific field

	switch name {
	case v.BEREQ_BETWEEN_BYTES_TIMEOUT:
		return value.NewValue(value.RTIME, "ctx.Backend.BetweenBytesTimeout"), nil
	// Backend request related variables
	case v.BEREQ_BODY_BYTES_WRITTEN:
		tmp := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBodyBytesWritten()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.BEREQ_BYTES_WRITTEN:
		tmp := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBytesWritten()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.BEREQ_HEADER_BYTES_WRITTEN:
		return value.NewValue(value.INTEGER, "ctx.BackendHeaderBytesWritten()"), nil
	case v.BEREQ_CONNECT_TIMEOUT:
		return value.NewValue(value.RTIME, "ctx.Backend.ConnectTimeout"), nil
	case v.BEREQ_FIRST_BYTE_TIMEOUT:
		return value.NewValue(value.RTIME, "ctx.Backend.FirstByteTimeout"), nil
	case v.BEREQ_METHOD,
		v.BEREQ_REQUEST:
		return value.NewValue(value.STRING, "ctx.BackendRequest.Method"), nil
	case v.BEREQ_PROTO:
		return value.NewValue(value.STRING, "ctx.BackendRequest.Proto"), nil
	case v.BEREQ_URL:
		return value.NewValue(value.STRING, "ctx.RequestURL(ctx.BackendRequest)"), nil
	case v.BEREQ_URL_BASENAME:
		return value.NewValue(
			value.STRING,
			"filepath.Base(ctx.BackendRequest.URL.Path)",
			value.Dependency("path/filepath", ""),
		), nil
	case v.BEREQ_URL_DIRNAME:
		return value.NewValue(
			value.STRING,
			"filepath.Dir(ctx.BackendRequest.URL.Path)",
			value.Dependency("path/filepath", ""),
		), nil
	case v.BEREQ_URL_EXT:
		return value.NewValue(
			value.STRING,
			`strings.TrimPrefix(filepath.Ext(ctx.BackendRequest.URL.Path), ".")`,
			value.Dependency("path/filepath", ""),
			value.Dependency("strings", ""),
		), nil
	case v.BEREQ_URL_PATH:
		return value.NewValue(value.STRING, "ctx.BackendRequest.URL.Path"), nil
	case v.BEREQ_URL_QS:
		return value.NewValue(value.STRING, "ctx.BackendRequest.URL.RawQuery"), nil

	case v.BERESP_BACKEND_NAME:
		return value.NewValue(value.STRING, "ctx.Backend.Backend()"), nil
	case v.BERESP_BACKEND_PORT:
		return value.NewValue(value.INTEGER, "ctx.Backend.Port"), nil

	case v.BERESP_PROTO:
		return value.NewValue(value.BOOL, "ctx.BackendResponse.Request.Proto"), nil
	case v.BERESP_RESPONSE:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.ResponseBody(ctx.BackendResponse)", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.BERESP_STALE_IF_ERROR:
		return value.NewValue(value.RTIME, "ctx.BackendResponseStaleIfError"), nil
	case v.BERESP_STALE_WHILE_REVALIDATE:
		return value.NewValue(value.RTIME, "ctx.BackendResponseStaleWhileRevalidate"), nil
	case v.BERESP_STATUS:
		return value.NewValue(value.INTEGER, "int64(ctx.BackendResponse.StatusCode)"), nil

	// AS Number always indicates "Reserved" defined by RFC7300
	// see: https://datatracker.ietf.org/doc/html/rfc7300
	// @Tentative
	case v.CLIENT_AS_NUMBER:
		return value.NewValue(value.INTEGER, "4294967294", value.Comment(name)), nil

	// AS Name always indicates "Reserved"
	// @Tentative
	case v.CLIENT_AS_NAME:
		return value.NewValue(value.STRING, "Reserved", value.Comment(name)), nil

	// Geo Longitude and Latitude always indicates Fastly.inc
	// @Tentative
	case v.CLIENT_GEO_LATITUDE:
		return value.NewValue(value.FLOAT, "37.7786941", value.Comment(name)), nil
	// @Tentative
	case v.CLIENT_GEO_LONGITUDE:
		return value.NewValue(value.FLOAT, "-122.3981452", value.Comment(name)), nil

	// Other geo related integer indicates 0
	// @Tentative
	case v.CLIENT_GEO_AREA_CODE,
		v.CLIENT_GEO_METRO_CODE,
		v.CLIENT_GEO_UTC_OFFSET,
		v.CLIENT_GEO_GMT_OFFSET:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil

	// Other geo related string indicates "unknown"
	// @Tentative
	case v.CLIENT_GEO_CITY,
		v.CLIENT_GEO_CITY_ASCII,
		v.CLIENT_GEO_CITY_LATIN1,
		v.CLIENT_GEO_CITY_UTF8,
		v.CLIENT_GEO_CONN_SPEED,
		v.CLIENT_GEO_CONN_TYPE,
		v.CLIENT_GEO_CONTINENT_CODE,
		v.CLIENT_GEO_COUNTRY_CODE,
		v.CLIENT_GEO_COUNTRY_CODE3,
		v.CLIENT_GEO_COUNTRY_NAME,
		v.CLIENT_GEO_COUNTRY_NAME_ASCII,
		v.CLIENT_GEO_COUNTRY_NAME_LATIN1,
		v.CLIENT_GEO_COUNTRY_NAME_UTF8,
		v.CLIENT_GEO_IP_OVERRIDE,
		v.CLIENT_GEO_POSTAL_CODE,
		v.CLIENT_GEO_PROXY_DESCRIPTION,
		v.CLIENT_GEO_PROXY_TYPE,
		v.CLIENT_GEO_REGION,
		v.CLIENT_GEO_REGION_ASCII,
		v.CLIENT_GEO_REGION_LATIN1,
		v.CLIENT_GEO_REGION_UTF8:
		return value.NewValue(value.STRING, "unknown", value.Comment(name)), nil

	// @Tentative
	case v.CLIENT_PORT:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil

	case v.FASTLY_INFO_HOST_HEADER:
		return value.NewValue(value.STRING, "ctx.OriginalHost", value.Comment(name)), nil
	case v.FASTLY_INFO_IS_H2:
		return value.NewValue(value.BOOL, "ctx.Request.ProtoMajor == 2"), nil
	case v.FASTLY_INFO_IS_H3:
		return value.NewValue(value.BOOL, "ctx.Request.ProtoMajor == 3"), nil

	// @Tentative
	case v.GEOIP_AREA_CODE:
		return value.NewValue(value.INTEGER, "0", value.Comment(name), value.Deprecated()), nil

	// Other deprecated geo related string indicates "unknown"
	// @Tentative
	case v.GEOIP_CITY,
		v.GEOIP_CITY_ASCII,
		v.GEOIP_CITY_LATIN1,
		v.GEOIP_CITY_UTF8,
		v.GEOIP_CONTINENT_CODE,
		v.GEOIP_COUNTRY_CODE,
		v.GEOIP_COUNTRY_CODE3,
		v.GEOIP_COUNTRY_NAME,
		v.GEOIP_COUNTRY_NAME_ASCII,
		v.GEOIP_COUNTRY_NAME_LATIN1,
		v.GEOIP_COUNTRY_NAME_UTF8:
		return value.NewValue(value.STRING, "unknown", value.Comment(name), value.Deprecated()), nil

	// @Tentative
	case v.GEOIP_IP_OVERRIDE:
		return value.NewValue(value.STRING, "", value.Comment(name), value.Deprecated()), nil

	// @Tentative
	case v.GEOIP_LATITUDE:
		return value.NewValue(value.FLOAT, "37.7786941", value.Comment(name), value.Deprecated()), nil
	case v.GEOIP_LONGITUDE:
		return value.NewValue(value.FLOAT, "-122.3981452", value.Comment(name), value.Deprecated()), nil

	// @Tentative
	case v.GEOIP_METRO_CODE,
		v.GEOIP_POSTAL_CODE,
		v.GEOIP_REGION,
		v.GEOIP_REGION_ASCII,
		v.GEOIP_REGION_LATIN1,
		v.GEOIP_REGION_UTF8:
		return value.NewValue(value.STRING, "unknown", value.Comment(name), value.Deprecated()), nil

	// @Tentative
	case v.GEOIP_USE_X_FORWARDED_FOR:
		return value.NewValue(value.BOOL, "false", value.Comment(name), value.Deprecated()), nil

	// Cache object related variables
	// On Native runtime, obj.xxx will be treated as backend response
	case v.OBJ_AGE:
		return value.NewValue(value.RTIME, "ctx.ObjectAge()"), nil
	case v.OBJ_CACHEABLE:
		return value.NewValue(value.BOOL, "ctx.ObjectCacheable()"), nil
	case v.OBJ_HITS:
		return value.NewValue(value.INTEGER, "ctx.ObjectHits()"), nil

	// @Tentative
	case v.OBJ_ENTERED:
		return value.NewValue(value.RTIME, "time.Duration(0)", value.Comment(name)), nil
	// @Tentative
	case v.OBJ_GRACE:
		return value.NewValue(value.RTIME, "time.Duration(0)", value.Comment(name)), nil
	// @Tentative
	case v.OBJ_IS_PCI:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	case v.OBJ_LASTUSE:
		return value.NewValue(value.RTIME, "time.Duration(0)", value.Comment(name)), nil
	case v.OBJ_PROTO:
		return value.NewValue(value.STRING, "ctx.BackendResponse.Request.Proto"), nil
	case v.OBJ_RESPONSE:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.ResponseBody(ctx.BackendResponse)", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.OBJ_STALE_IF_ERROR:
		return value.NewValue(value.RTIME, "ctx.ObjectStaleIfError"), nil
	case v.OBJ_STALE_WHILE_REVALIDATE:
		return value.NewValue(value.RTIME, "ctx.ObjectStaleWhileRevalidate"), nil
	case v.OBJ_STATUS:
		return value.NewValue(value.INTEGER, "int64(ctx.BackendResponse.StatusCode)"), nil
	case v.OBJ_TTL:
		return value.NewValue(value.RTIME, "ctx.ObjectTTL"), nil

	case v.REQ_BACKEND_PORT:
		return value.NewValue(value.INTEGER, "ctx.Backend.Port"), nil
	case v.REQ_BODY:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBody()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.REQ_BODY_BASE64:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBodyBase64()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.REQ_BODY_BYTES_READ:
		tmp := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBodyBytesRead()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.REQ_BYTES_READ:
		tmp := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.RequestBytesRead()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.REQ_IS_IPV6:
		return value.NewValue(value.BOOL, "ctx.IsIpv6()"), nil
	case v.REQ_IS_PURGE:
		return value.NewValue(value.BOOL, `(ctx.Request.Method == "PURGE")`), nil
	case v.REQ_METHOD:
		return value.NewValue(value.STRING, "ctx.Request.Method"), nil
	case v.REQ_POSTBODY:
		return fv.Get("req.body") // alias of req.body
	case v.REQ_PROTO:
		return value.NewValue(value.STRING, "ctx.Request.Proto"), nil
	case v.REQ_REQUEST:
		return fv.Get("req.method") // alias of req.method

	// @Tentative
	case v.REQ_SERVICE_ID:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	case v.REQ_TOPURL:
		return fv.Get("req.url")
	case v.REQ_URL:
		return value.NewValue(value.STRING, "ctx.RequestURL(ctx.Request)"), nil
	case v.REQ_URL_BASENAME:
		return value.NewValue(
			value.STRING,
			"filepath.Base(ctx.Request.URL.Path)",
			value.Dependency("path/filepath", ""),
		), nil
	case v.REQ_URL_DIRNAME:
		return value.NewValue(
			value.STRING,
			"filepath.Dir(ctx.Request.URL.Path)",
			value.Dependency("path/filepath", ""),
		), nil
	case v.REQ_URL_EXT:
		return value.NewValue(
			value.STRING,
			`strings.TrimPrefix(filepath.Ext(ctx.Request.URL.Path), ".")`,
			value.Dependency("path/filepath", ""),
			value.Dependency("strings", ""),
		), nil
	case v.REQ_URL_PATH:
		return value.NewValue(value.STRING, "ctx.Request.URL.Path"), nil
	case v.REQ_URL_QS:
		return value.NewValue(value.STRING, "ctx.Request.URL.RawQuery"), nil

	case v.RESP_PROTO:
		return value.NewValue(value.STRING, "ctx.Response.Request.Proto"), nil
	case v.RESP_RESPONSE:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			tmp,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.ResponseBody()", tmp),
				value.ErrorCheck,
			),
		), nil
	case v.RESP_STATUS:
		return value.NewValue(value.INTEGER, "ctx.Response.StatusCode"), nil

	// @Tentative
	case v.TLS_CLIENT_CIPHER:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil
	// @Tentative
	case v.TLS_CLIENT_PROTOCOL:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil
	}

	// Look up core variables
	return fv.CoreVariable.Get(name)
}

func (fv *NativeVariable) Set(name string, val *value.Value) (*value.Value, error) {
	switch name {
	case v.BEREQ_METHOD:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendRequest.Method = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.BEREQ_REQUEST:
		return fv.Set(v.BEREQ_METHOD, val)
	case v.BEREQ_URL:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.SetURL(ctx.BackendRequest, %s)", tmp),
			value.Prepare(
				fmt.Sprintf("%s, err := url.Parse(%s)", tmp, val.Conversion(value.STRING).String()),
				value.ErrorCheck,
			),
			value.FromValue(val),
			value.Dependency("net/url", ""),
		), nil
	case v.BERESP_RESPONSE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.SetResponseBody(ctx.BackendResponse, %s)", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_STATUS:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponse.StatusCode = int(%s)", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.OBJ_RESPONSE:
		return fv.Set(v.BERESP_RESPONSE, val)
	case v.OBJ_STATUS:
		return fv.Set(v.BERESP_STATUS, val)
	case v.REQ_METHOD:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Request.Method = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.REQ_REQUEST:
		return fv.Set(v.REQ_METHOD, val)
	case v.REQ_URL:
		tmp := value.Temporary()
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.SetURL(ctx.Request, %s)", tmp),
			value.Prepare(
				fmt.Sprintf("%s, err := url.Parse(%s)", tmp, val.Conversion(value.STRING).String()),
				value.ErrorCheck,
			),
			value.FromValue(val),
			value.Dependency("net/url", ""),
		), nil
	case v.RESP_RESPONSE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.SetResponseBody(ctx.Response, %s)", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.RESP_STATUS:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Response.StatusCode = int(%s)", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	}

	return fv.CoreVariable.Set(name, val)
}

func (fv *NativeVariable) Unset(name string) (*value.Value, error) {
	return fv.CoreVariable.Unset(name)
}

var _ v.Variables = (*NativeVariable)(nil)
