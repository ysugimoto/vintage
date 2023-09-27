package fastly

import (
	"errors"
	"fmt"

	"github.com/ysugimoto/vintage/transformer/core"
	"github.com/ysugimoto/vintage/transformer/value"
	v "github.com/ysugimoto/vintage/transformer/variable"
)

type FastlyVariable struct {
	CoreVariable *core.CoreVariable
}

func NewFastlyVariable() *FastlyVariable {
	return &FastlyVariable{
		core.NewCoreVariables(),
	}
}

func (fv *FastlyVariable) Get(name string) (*value.Value, error) {
	// Lookup variable for fastly specific field

	switch name {
	case v.BEREQ_BETWEEN_BYTES_TIMEOUT:
		return value.NewValue(value.RTIME, "ctx.Backend.BetweenBytesTimeout"), nil
	// Backend request related variables
	case v.BEREQ_BODY_BYTES_WRITTEN:
		v := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			v,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBodyBytesWritten()", v),
				value.ErrorCheck,
			),
		), nil
	case v.BEREQ_BYTES_WRITTEN:
		v := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			v,
			value.Prepare(
				fmt.Sprintf("%s, err := ctx.BackendBytesWritten()", v),
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

	// @Tentative
	case v.BERESP_BACKEND_ALTERNATE_IPS:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil
	// @Tentative
	case v.BERESP_BACKEND_IP:
		return value.NewValue(value.IP, "nil", value.Comment(name)), nil

	case v.BERESP_BACKEND_NAME:
		return value.NewValue(value.STRING, "ctx.Backend.Backend()"), nil
	case v.BERESP_BACKEND_PORT:
		return value.NewValue(value.INTEGER, "ctx.Backend.Port"), nil

	// @Tentative
	case v.BERESP_BACKEND_REQUESTS:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil

	case v.BERESP_BROTLI:
		return value.NewValue(value.BOOL, "ctx.BackendResponseBrotli"), nil
	case v.BERESP_CACHEABLE:
		return value.NewValue(value.BOOL, "ctx.BackendResponseCacheable"), nil
	case v.BERESP_DO_ESI:
		return value.NewValue(value.BOOL, "ctx.BackendResponseDoESI"), nil
	case v.BERESP_DO_STREAM:
		return value.NewValue(value.BOOL, "ctx.BackendResponseDoStream"), nil
	case v.BERESP_GRACE:
		return value.NewValue(value.BOOL, "ctx.BackendResponseGrace"), nil
	case v.BERESP_GZIP:
		return value.NewValue(value.BOOL, "ctx.BackendResponseGzip"), nil

	// @Tentative: Edge runtime could not know handshake related info
	case v.BERESP_HANDSHAKE_TIME_TO_ORIGIN_MS:
		return value.NewValue(value.INTEGER, "100", value.Comment(name)), nil

	case v.BERESP_HIPAA:
		return value.NewValue(value.BOOL, "ctx.BackendResponseHipaa"), nil
	case v.BERESP_PCI:
		return value.NewValue(value.BOOL, "ctx.BackendResponsePCI"), nil
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
	case v.BERESP_TTL:
		return value.NewValue(value.RTIME, "ctx.BackendResponseTTL"), nil

	// @Tentative
	case v.BERESP_USED_ALTERNATE_PATH_TO_ORIGIN:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	case v.CLIENT_AS_NUMBER:
		return value.NewValue(value.INTEGER, "int64(ctx.Geo.AsNumber)"), nil
	case v.CLIENT_AS_NAME:
		return value.NewValue(value.STRING, "ctx.Geo.AsName"), nil
	case v.CLIENT_GEO_LATITUDE:
		return value.NewValue(value.FLOAT, "ctx.Geo.Latitude"), nil
	case v.CLIENT_GEO_LONGITUDE:
		return value.NewValue(value.FLOAT, "ctx.Geo.Longitude"), nil
	case v.CLIENT_GEO_AREA_CODE:
		return value.NewValue(value.INTEGER, "int64(ctx.Geo.AreaCode)"), nil
	case v.CLIENT_GEO_METRO_CODE:
		return value.NewValue(value.INTEGER, "int64(ctx.Geo.MetroCode)"), nil
	case v.CLIENT_GEO_UTC_OFFSET:
		return value.NewValue(value.INTEGER, "int64(ctx.Geo.UTCOffset)"), nil
	case v.CLIENT_GEO_GMT_OFFSET:
		return value.NewValue(value.INTEGER, "int64(ctx.Geo.UTCOffset)"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case v.CLIENT_GEO_CITY,
		v.CLIENT_GEO_CITY_ASCII,
		v.CLIENT_GEO_CITY_LATIN1,
		v.CLIENT_GEO_CITY_UTF8:
		return value.NewValue(value.STRING, "ctx.Geo.City"), nil
	case v.CLIENT_GEO_CONN_SPEED:
		return value.NewValue(value.STRING, "ctx.Geo.ConnSpeed"), nil
	case v.CLIENT_GEO_CONN_TYPE:
		return value.NewValue(value.STRING, "ctx.Geo.ConnType"), nil
	case v.CLIENT_GEO_CONTINENT_CODE:
		return value.NewValue(value.STRING, "ctx.Geo.ContinentCode"), nil
	case v.CLIENT_GEO_COUNTRY_CODE:
		return value.NewValue(value.STRING, "ctx.Geo.CountryCode"), nil
	case v.CLIENT_GEO_COUNTRY_CODE3:
		return value.NewValue(value.STRING, "ctx.Geo.CountryCode3"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case v.CLIENT_GEO_COUNTRY_NAME,
		v.CLIENT_GEO_COUNTRY_NAME_ASCII,
		v.CLIENT_GEO_COUNTRY_NAME_LATIN1,
		v.CLIENT_GEO_COUNTRY_NAME_UTF8:
		return value.NewValue(value.STRING, "ctx.Geo.CountryName"), nil

	// @Tentative
	case v.CLIENT_GEO_IP_OVERRIDE:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	case v.CLIENT_GEO_POSTAL_CODE:
		return value.NewValue(value.STRING, "ctx.Geo.PostalCode"), nil
	case v.CLIENT_GEO_PROXY_DESCRIPTION:
		return value.NewValue(value.STRING, "ctx.Geo.ProxyDescription"), nil
	case v.CLIENT_GEO_PROXY_TYPE:
		return value.NewValue(value.STRING, "ctx.Geo.ProxyType"), nil
	// ASCII, LATIN1 always return value as the same of UTF8
	case v.CLIENT_GEO_REGION,
		v.CLIENT_GEO_REGION_ASCII,
		v.CLIENT_GEO_REGION_LATIN1,
		v.CLIENT_GEO_REGION_UTF8:
		return value.NewValue(value.STRING, "ctx.Geo.Region"), nil
	case v.CLIENT_IDENTITY:
		return value.NewValue(value.STRING, "ctx.ClientIdentity()"), nil
	case v.CLIENT_IP:
		return value.NewValue(value.IP, "ctx.ClientIP()"), nil

	// @Tentative
	case v.CLIENT_PORT:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil

	case v.FASTLY_INFO_HOST_HEADER:
		return value.NewValue(value.STRING, "ctx.OriginalHost", value.Comment(name)), nil
	case v.FASTLY_INFO_IS_H2:
		return value.NewValue(value.BOOL, "ctx.Request.ProtoMajor == 2"), nil
	case v.FASTLY_INFO_IS_H3:
		return value.NewValue(value.BOOL, "ctx.Request.ProtoMajor == 3"), nil
	case v.GEOIP_AREA_CODE:
		return value.NewValue(value.INTEGER, "int64(ctx.Geo.AreaCode)", value.Deprecated()), nil
	case v.GEOIP_CITY,
		v.GEOIP_CITY_ASCII,
		v.GEOIP_CITY_LATIN1,
		v.GEOIP_CITY_UTF8:
		return value.NewValue(value.STRING, "ctx.Geo.City", value.Deprecated()), nil
	case v.GEOIP_CONTINENT_CODE:
		return value.NewValue(value.STRING, "ctx.Geo.ContinentCode", value.Deprecated()), nil
	case v.GEOIP_COUNTRY_CODE:
		return value.NewValue(value.STRING, "ctx.Geo.CountryCode", value.Deprecated()), nil
	case v.GEOIP_COUNTRY_CODE3:
		return value.NewValue(value.STRING, "ctx.Geo.CountryCode3", value.Deprecated()), nil
	case v.GEOIP_COUNTRY_NAME,
		v.GEOIP_COUNTRY_NAME_ASCII,
		v.GEOIP_COUNTRY_NAME_LATIN1,
		v.GEOIP_COUNTRY_NAME_UTF8:
		return value.NewValue(value.STRING, "ctx.Geo.CountryName", value.Deprecated()), nil

	// @Tentative
	case v.GEOIP_IP_OVERRIDE:
		return value.NewValue(value.STRING, "", value.Comment(name), value.Deprecated()), nil

	case v.GEOIP_LATITUDE:
		return value.NewValue(value.FLOAT, "ctx.Geo.Latitude", value.Deprecated()), nil
	case v.GEOIP_LONGITUDE:
		return value.NewValue(value.FLOAT, "ctx.Geo.Longitude", value.Deprecated()), nil
	case v.GEOIP_METRO_CODE:
		return value.NewValue(value.FLOAT, "ctx.Geo.MetroCode", value.Deprecated()), nil
	case v.GEOIP_POSTAL_CODE:
		return value.NewValue(value.FLOAT, "ctx.Geo.PostalCode", value.Deprecated()), nil
	case v.GEOIP_REGION,
		v.GEOIP_REGION_ASCII,
		v.GEOIP_REGION_LATIN1,
		v.GEOIP_REGION_UTF8:
		return value.NewValue(value.FLOAT, "ctx.Geo.Region", value.Deprecated()), nil
	case v.GEOIP_USE_X_FORWARDED_FOR:
		return value.NewValue(value.BOOL, "false", value.Comment(name), value.Deprecated()), nil

	// Cache object related variables
	// On Fastly runtime, obj.xxx will be treated as backend response
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
		return value.NewValue(value.STRING, "ctx.BackendResponse.Request.Proto", value.Comment(name)), nil
	case v.OBJ_RESPONSE:
		return value.NewValue(value.STRING, "ctx.ResponseBody(c.BackendResponse)", value.Comment(name)), nil
	case v.OBJ_STALE_IF_ERROR:
		return value.NewValue(value.RTIME, "ctx.ObjectStaleIfError"), nil
	case v.OBJ_STALE_WHILE_REVALIDATE:
		return value.NewValue(value.RTIME, "ctx.ObjectStaleWhileRevalidate"), nil
	case v.OBJ_STATUS:
		return value.NewValue(value.INTEGER, "int64(ctx.BackendResponse.StatusCode)", value.Comment(name)), nil
	case v.OBJ_TTL:
		return value.NewValue(value.RTIME, "ctx.ObjectTTL"), nil

	// @Tentative
	case v.REQ_BACKEND_IP:
		return value.NewValue(value.IP, "vintage.LocalHost", value.Comment(name)), nil
	// @Tentative
	case v.REQ_BACKEND_IS_CLUSTER:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	case v.REQ_BACKEND_IS_SHIELD:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

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
	case v.REQ_DIGEST:
		return value.NewValue(value.STRING, "ctx.RequestDigest()"), nil
	case v.REQ_ENABLE_RANGE_ON_PASS:
		return value.NewValue(value.BOOL, "ctx.EnableRangeOnPass"), nil
	case v.REQ_ENABLE_SEGMENTED_CACHING:
		return value.NewValue(value.BOOL, "ctx.EnableSegmentedCaching"), nil
	case v.REQ_HASH:
		return value.NewValue(value.STRING, "ctx.RequestHash"), nil
	case v.REQ_HASH_ALWAYS_MISS:
		return value.NewValue(value.BOOL, "ctx.HashAlwaysMiss"), nil
	case v.REQ_HASH_IGNORE_BUSY:
		return value.NewValue(value.BOOL, "ctx.HashIgnoreBusy"), nil
	case v.REQ_HEADER_BYTES_READ:
		return value.NewValue(value.INTEGER, "ctx.RequestHeaderBytes"), nil
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
		return value.NewValue(value.STRING, "ctx.RequestURL()"), nil
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

	// @Tentative
	case v.REQ_VCL:
		return value.NewValue(value.STRING, "vintage.vcl.transpiled", value.Comment(name)), nil
	// Precalculated: md5("vintage.vcl.transpiled")
	// @Tentative
	case v.REQ_VCL_MD5:
		return value.NewValue(value.STRING, "1061a3ce2c356a8a5e5423a824eba490", value.Comment(name)), nil

	case v.REQ_XID:
		return value.NewValue(value.STRING, "vintage.GenerateXid()"), nil
	case v.RESP_BODY_BYTES_WRITTEN:
		return value.NewValue(value.INTEGER, "ctx.ResponseBodyBytesWritten"), nil
	case v.RESP_BYTES_WRITTEN:
		return value.NewValue(value.INTEGER, "ctx.ResponseBytesWritten"), nil
	case v.RESP_HEADER_BYTES_WRITTEN:
		return value.NewValue(value.INTEGER, "ctx.ResponseHeaderBytesWritten"), nil

	// @Tentative
	case v.RESP_COMPLETED:
		return value.NewValue(value.BOOL, "true", value.Comment(name)), nil

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
	case v.TLS_CLIENT_CIPHER:
		return value.NewValue(value.STRING, "ctx.Request.TLSInfo.CipherOpenSSLName"), nil
	case v.TLS_CLIENT_PROTOCOL:
		return value.NewValue(value.STRING, "ctx.Request.TLSInfo.Protocol"), nil
	}

	// Look up core variables
	return fv.CoreVariable.Get(name)
}

func (v *FastlyVariable) Set(name string, value *value.Value) error {
	return errors.New("Not Implemented")
}

func (v *FastlyVariable) Unset(name string) error {
	return errors.New("Not Implemented")
}

var _ v.Variables = (*FastlyVariable)(nil)
