package core

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage/transformer/value"
	v "github.com/ysugimoto/vintage/transformer/variable"
)

type CoreVariable struct {
	*v.VariablesImpl
}

func NewCoreVariables() *CoreVariable {
	return &CoreVariable{
		VariablesImpl: &v.VariablesImpl{},
	}
}

// nolint:funlen,gocyclo
func (cv *CoreVariable) Get(name string) (*value.Value, error) {
	switch name {
	case v.LF:
		return value.NewValue(value.STRING, "\n"), nil

	case v.CLIENT_IDENTITY:
		return value.NewValue(value.STRING, "ctx.GetClientIdentity()"), nil
	case v.CLIENT_IP:
		return value.NewValue(value.IP, "ctx.ClientIp"), nil

	// @Tentative
	case v.BEREQ_IS_CLUSTERING:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	case v.CLIENT_BOT_NAME:
		return value.NewValue(value.STRING, "ctx.UserAgent.BotName()"), nil
	case v.CLIENT_BROWSER_NAME:
		return value.NewValue(value.STRING, "ctx.UserAgent.BrowserName()"), nil
	case v.CLIENT_BROWSER_VERSION:
		return value.NewValue(value.STRING, "ctx.UserAgent.BrowserVersion()"), nil
	case v.CLIENT_CLASS_BOT:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsBot()"), nil
	case v.CLIENT_CLASS_BROWSER:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsBrowser()"), nil
	case v.CLIENT_DISPLAY_TOUCHSCREEN:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsTouchScreen()"), nil
	case v.CLIENT_OS_NAME:
		return value.NewValue(value.STRING, "ctx.UserAgent.OSName()"), nil
	case v.CLIENT_OS_VERSION:
		return value.NewValue(value.STRING, "ctx.UserAgent.OSVersion()"), nil
	case v.CLIENT_PLATFORM_EREADER:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsEReader()"), nil
	case v.CLIENT_PLATFORM_GAMECONSOLE:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsGameConsole()"), nil

	// @Tentative
	case v.CLIENT_PLATFORM_HWTYPE:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	case v.CLIENT_PLATFORM_MOBILE:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsMobile()"), nil
	case v.CLIENT_PLATFORM_SMARTTV:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsSmartTV()"), nil
	case v.CLIENT_PLATFORM_TABLET:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsTablet()"), nil
	case v.CLIENT_PLATFORM_TVPLAYER:
		return value.NewValue(value.BOOL, "ctx.UserAgent.IsTvPlayer()"), nil

	// Following values are always false in Edge Runtime
	case v.CLIENT_CLASS_CHECKER,
		v.CLIENT_CLASS_DOWNLOADER,
		v.CLIENT_CLASS_FEEDREADER,
		v.CLIENT_CLASS_FILTER,
		v.CLIENT_CLASS_MASQUERADING,
		v.CLIENT_CLASS_SPAM,
		v.CLIENT_PLATFORM_MEDIAPLAYER,
		v.REQ_IS_BACKGROUND_FETCH,
		v.REQ_IS_CLUSTERING,
		v.REQ_IS_ESI_SUBREQ:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	case v.RESP_STALE:
		return value.NewValue(value.BOOL, "ctx.ResponseStale"), nil
	case v.RESP_STALE_IS_ERROR:
		return value.NewValue(value.BOOL, "ctx.ResponseStaleIsError"), nil
	case v.RESP_STALE_IS_REVALIDATING:
		return value.NewValue(value.BOOL, "ctx.ResponseStaleIsRevalidating"), nil

	// Edge does not check backend is healthy but value should be true
	// @Tentative
	case v.REQ_BACKEND_HEALTHY:
		return value.NewValue(value.BOOL, "true", value.Comment(name)), nil

	// Edge always works on the TLS
	// @Tentative
	case v.REQ_IS_SSL:
		return value.NewValue(value.BOOL, "true", value.Comment(name)), nil
	// @Tentative
	case v.REQ_PROTOCOL:
		return value.NewValue(value.STRING, "https", value.Comment(name)), nil

	// Client display infos are unknown. Always returns -1
	// @Tentative
	case v.CLIENT_DISPLAY_HEIGHT,
		v.CLIENT_DISPLAY_PPI,
		v.CLIENT_DISPLAY_WIDTH:
		return value.NewValue(value.INTEGER, "-1", value.Comment(name)), nil

	// Client could not fully identified so returns false
	// @Tentative
	case v.CLIENT_IDENTIFIED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	// Client requests always returns 1, means new connection is coming
	// @Tentative
	case v.CLIENT_REQUESTS:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil

	// Returns tentative value -- you may know your customer_id in the contraction :-)
	// @Tentative
	case v.REQ_CUSTOMER_ID:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	// Returns fixed value which is presented on Fastly fiddle
	case v.REQ_RESTARTS:
		return value.NewValue(value.INTEGER, "int64(ctx.Restarts)"), nil

	// Returns always 1 because VCL is generated locally
	// @Tentative
	case v.REQ_VCL_GENERATION:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil
	// @Tentative
	case v.REQ_VCL_VERSION:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil
	// @Tentative
	case v.REQ_BACKEND_IP:
		return value.NewValue(value.IP, "vintage.LocalHost", value.Comment(name)), nil
	// @Tentative
	case v.REQ_BACKEND_IS_CLUSTER:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.REQ_BACKEND_IS_SHIELD:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	case v.REQ_HEADER_BYTES_READ:
		return value.NewValue(value.INTEGER, "ctx.RequestHeaderBytes"), nil
	case v.REQ_DIGEST:
		return value.NewValue(value.STRING, "ctx.RequestDigest()"), nil

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

	case v.RESP_COMPLETED:
		return value.NewValue(value.BOOL, "ctx.ResponseCompleted"), nil
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

	// Edge Runtime does not know backend and server IP info
	// @Tentative
	case v.BERESP_BACKEND_SRC_IP:
		return value.NewValue(value.IP, "vintage.LocalHost)", value.Comment(name)), nil
	// @Tentative
	case v.BERESP_BACKEND_ALTERNATE_IPS:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil
	// @Tentative
	case v.BERESP_BACKEND_IP:
		return value.NewValue(value.IP, "nil", value.Comment(name)), nil

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
	case v.BERESP_TTL:
		return value.NewValue(value.RTIME, "ctx.BackendResponseTTL"), nil

	// @Tentative
	case v.BERESP_USED_ALTERNATE_PATH_TO_ORIGIN:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	// @Tentative: Edge runtime could not know handshake related info
	case v.BERESP_HANDSHAKE_TIME_TO_ORIGIN_MS:
		return value.NewValue(value.INTEGER, "100", value.Comment(name)), nil

	case v.BERESP_HIPAA:
		return value.NewValue(value.BOOL, "ctx.BackendResponseHipaa"), nil
	case v.BERESP_PCI:
		return value.NewValue(value.BOOL, "ctx.BackendResponsePCI"), nil

	// Core request related values
	case v.REQ_BACKEND:
		return value.NewValue(value.BACKEND, "ctx.Backend"), nil
	case v.REQ_GRACE: // Alias of req.max_stale_if_error
		return value.NewValue(value.RTIME, "ctx.MaxStaleIfError", value.Comment(name)), nil
	case v.REQ_MAX_STALE_IF_ERROR:
		return value.NewValue(value.RTIME, "ctx.MaxStaleIfError"), nil
	case v.REQ_MAX_STALE_WHILE_REVALIDATE:
		return value.NewValue(value.RTIME, "ctx.MaxStaleWhileRevalidate"), nil
	case v.REQ_ESI:
		return value.NewValue(value.BOOL, "ctx.EnableESI"), nil
	case v.REQ_ESI_LEVEL:
		return value.NewValue(value.BOOL, "ctx.ESILevel"), nil
	case v.REQ_BACKEND_NAME:
		return value.NewValue(value.STRING, "ctx.Backend.Backend()"), nil

	case v.RESP_IS_LOCALLY_GENERATED:
		return value.NewValue(value.BOOL, "ctx.IsLocallyGenerated"), nil

	// Always true because edge could not have origin-shielding
	case v.REQ_BACKEND_IS_ORIGIN:
		return value.NewValue(value.BOOL, "true", value.Comment(name)), nil

	case v.BACKEND_CONN_IS_TLS:
		return value.NewValue(value.BOOL, "ctx.Backend.SSL"), nil
	// @Tentative
	case v.BACKEND_CONN_TLS_PROTOCOL:
		return value.NewValue(value.STRING, "TLSv1.2", value.Comment(name)), nil
	// @Tentative
	case v.BACKEND_SOCKET_CONGESTION_ALGORITHM:
		return value.NewValue(value.STRING, "cubic", value.Comment(name)), nil
	// @Tentative
	case v.BACKEND_SOCKET_CWND:
		return value.NewValue(value.INTEGER, "60", value.Comment(name)), nil

	// Edge runtime could not know backend socket information.
	// @Tentative
	case v.BACKEND_SOCKET_TCPI_ADVMSS,
		v.BACKEND_SOCKET_TCPI_BYTES_ACKED,
		v.BACKEND_SOCKET_TCPI_BYTES_RECEIVED,
		v.BACKEND_SOCKET_TCPI_DATA_SEGS_IN,
		v.BACKEND_SOCKET_TCPI_DATA_SEGS_OUT,
		v.BACKEND_SOCKET_TCPI_DELIVERY_RATE,
		v.BACKEND_SOCKET_TCPI_DELTA_RETRANS,
		v.BACKEND_SOCKET_TCPI_LAST_DATA_SENT,
		v.BACKEND_SOCKET_TCPI_MAX_PACING_RATE,
		v.BACKEND_SOCKET_TCPI_MIN_RTT,
		v.BACKEND_SOCKET_TCPI_NOTSENT_BYTES,
		v.BACKEND_SOCKET_TCPI_PACING_RATE,
		v.BACKEND_SOCKET_TCPI_PMTU,
		v.BACKEND_SOCKET_TCPI_RCV_MSS,
		v.BACKEND_SOCKET_TCPI_RCV_RTT,
		v.BACKEND_SOCKET_TCPI_RCV_SPACE,
		v.BACKEND_SOCKET_TCPI_RCV_SSTHRESH,
		v.BACKEND_SOCKET_TCPI_REORDERING,
		v.BACKEND_SOCKET_TCPI_RTT,
		v.BACKEND_SOCKET_TCPI_RTTVAR,
		v.BACKEND_SOCKET_TCPI_SEGS_IN,
		v.BACKEND_SOCKET_TCPI_SEGS_OUT,
		v.BACKEND_SOCKET_TCPI_SND_CWND,
		v.BACKEND_SOCKET_TCPI_SND_MSS,
		v.BACKEND_SOCKET_TCPI_SND_SSTHRESH,
		v.BACKEND_SOCKET_TCPI_TOTAL_RETRANS:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	case v.CLIENT_SOCKET_CONGESTION_ALGORITHM:
		return value.NewValue(value.STRING, "ctx.ClientSocketCongestionAlgorithm"), nil
	case v.CLIENT_SOCKET_CWND:
		return value.NewValue(value.INTEGER, "ctx.ClientSocketCWND"), nil
	// @Tentative
	case v.CLIENT_SOCKET_NEXTHOP:
		return value.NewValue(value.IP, "net.IPv4(127, 0, 0, 1)", value.Comment(name)), nil
	// @Tentative
	case v.CLIENT_SOCKET_PACE:
		// Minimum value: 128KiB
		// @see: https://developer.fastly.com/reference/vcl/variables/client-connection/client-socket-pace/
		return value.NewValue(value.INTEGER, "131072", value.Comment(name)), nil
	// @Tentative
	case v.CLIENT_SOCKET_PLOSS:
		return value.NewValue(value.FLOAT, "0", value.Comment(name)), nil

	// Edge runtime does not know client TCP info
	// @Tentativce
	case v.CLIENT_SOCKET_TCP_INFO,
		v.CLIENT_SOCKET_TCPI_ADVMSS,
		v.CLIENT_SOCKET_TCPI_BYTES_ACKED,
		v.CLIENT_SOCKET_TCPI_BYTES_RECEIVED,
		v.CLIENT_SOCKET_TCPI_DATA_SEGS_IN,
		v.CLIENT_SOCKET_TCPI_DATA_SEGS_OUT,
		v.CLIENT_SOCKET_TCPI_DELIVERY_RATE,
		v.CLIENT_SOCKET_TCPI_DELTA_RETRANS,
		v.CLIENT_SOCKET_TCPI_LAST_DATA_SENT,
		v.CLIENT_SOCKET_TCPI_MAX_PACING_RATE,
		v.CLIENT_SOCKET_TCPI_MIN_RTT,
		v.CLIENT_SOCKET_TCPI_NOTSENT_BYTES,
		v.CLIENT_SOCKET_TCPI_PACING_RATE,
		v.CLIENT_SOCKET_TCPI_PMTU,
		v.CLIENT_SOCKET_TCPI_RCV_MSS,
		v.CLIENT_SOCKET_TCPI_RCV_RTT,
		v.CLIENT_SOCKET_TCPI_RCV_SPACE,
		v.CLIENT_SOCKET_TCPI_RCV_SSTHRESH,
		v.CLIENT_SOCKET_TCPI_REORDERING,
		v.CLIENT_SOCKET_TCPI_RTT,
		v.CLIENT_SOCKET_TCPI_RTTVAR,
		v.CLIENT_SOCKET_TCPI_SEGS_IN,
		v.CLIENT_SOCKET_TCPI_SEGS_OUT,
		v.CLIENT_SOCKET_TCPI_SND_CWND,
		v.CLIENT_SOCKET_TCPI_SND_MSS,
		v.CLIENT_SOCKET_TCPI_SND_SSTHRESH,
		v.CLIENT_SOCKET_TCPI_TOTAL_RETRANS:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	case v.ESI_ALLOW_INSIDE_CDATA:
		return value.NewValue(value.BOOL, "ctx.EsiAllowInsideCData"), nil

	// Error is always handles on-the-fly via golang way
	// @Tentative
	case v.FASTLY_ERROR:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil
	// Edge runtime does not have POP info, returns 1
	// @Tentative
	case v.FASTLY_FF_VISITS_THIS_POP:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil

	// Returns common value -- do not consider of clustering
	// see: https://developer.fastly.com/reference/vcl/variables/miscellaneous/fastly-ff-visits-this-service/
	// Edge runtime does not have VCL service info, always returns 1
	// @Tentative
	case v.FASTLY_FF_VISITS_THIS_SERVICE,
		v.FASTLY_FF_VISITS_THIS_POP_THIS_SERVICE:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil

	// Edge always works on the TLS
	case v.FASTLY_INFO_EDGE_IS_TLS:
		return value.NewValue(value.BOOL, "true", value.Comment(name)), nil

	// Undocumented this spec in Fastly
	// @Tentative
	case v.FASTLY_INFO_H2_FINGERPRINT:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	// Edge does not accept h2.push, maybe
	// @Tentative
	case v.FASTLY_INFO_H2_IS_PUSH:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.FASTLY_INFO_H2_STREAM_ID:
		return value.NewValue(value.STRING, "0", value.Comment(name)), nil

	// @Tentative
	case v.FASTLY_INFO_IS_CLUSTER_EDGE:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.FASTLY_INFO_IS_CLUSTER_SHIELD:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	case v.FASTLY_INFO_STATE:
		return value.NewValue(value.STRING, "ctx.State"), nil
	case v.MATH_1_PI:
		return value.NewValue(value.FLOAT, "(1/math.Pi)", value.Dependency("math", "")), nil
	case v.MATH_2_PI:
		return value.NewValue(value.FLOAT, "(2/math.Pi)", value.Dependency("math", "")), nil
	case v.MATH_2_SQRTPI:
		return value.NewValue(value.FLOAT, "(2/math.SqrtPi)", value.Dependency("math", "")), nil
	case v.MATH_2PI:
		return value.NewValue(value.FLOAT, "(2*math.Pi)", value.Dependency("math", "")), nil
	case v.MATH_E:
		return value.NewValue(value.FLOAT, "math.E", value.Dependency("math", "")), nil
	case v.MATH_FLOAT_EPSILON:
		return value.NewValue(value.FLOAT, "math.Pow(2, -52)", value.Dependency("math", "")), nil
	case v.MATH_FLOAT_MAX:
		return value.NewValue(value.FLOAT, "math.MaxFloat64", value.Dependency("math", "")), nil
	case v.MATH_FLOAT_MIN:
		return value.NewValue(value.FLOAT, "math.SmallestNonzeroFloat64", value.Dependency("math", "")), nil
	case v.MATH_LN10:
		return value.NewValue(value.FLOAT, "math.Ln10", value.Dependency("math", "")), nil
	case v.MATH_LN2:
		return value.NewValue(value.FLOAT, "math.Ln2", value.Dependency("math", "")), nil
	case v.MATH_LOG10E:
		return value.NewValue(value.FLOAT, "math.Log10E", value.Dependency("math", "")), nil
	case v.MATH_LOG2E:
		return value.NewValue(value.FLOAT, "math.Log2E", value.Dependency("math", "")), nil
	case v.MATH_NAN:
		return value.NewValue(value.FLOAT, "math.NaN", value.Dependency("math", "")), nil
	case v.MATH_NEG_HUGE_VAL:
		return value.NewValue(value.FLOAT, "math.Inf(-1)", value.Dependency("math", "")), nil
	case v.MATH_NEG_INFINITY:
		return value.NewValue(value.FLOAT, "math.Inf(-1)", value.Dependency("math", "")), nil
	case v.MATH_PHI:
		return value.NewValue(value.FLOAT, "math.Phi", value.Dependency("math", "")), nil
	case v.MATH_PI:
		return value.NewValue(value.FLOAT, "math.Pi", value.Dependency("math", "")), nil
	case v.MATH_PI_2:
		return value.NewValue(value.FLOAT, "(math.Pi/2)", value.Dependency("math", "")), nil
	case v.MATH_PI_4:
		return value.NewValue(value.FLOAT, "(math.Pi/4)", value.Dependency("math", "")), nil
	case v.MATH_POS_HUGE_VAL:
		return value.NewValue(value.FLOAT, "math.Inf(1)", value.Dependency("math", "")), nil
	case v.MATH_POS_INFINITY:
		return value.NewValue(value.FLOAT, "math.Inf(1)", value.Dependency("math", "")), nil
	case v.MATH_SQRT1_2:
		return value.NewValue(value.FLOAT, "(1/math.Sqrt2)", value.Dependency("math", "")), nil
	case v.MATH_SQRT2:
		return value.NewValue(value.FLOAT, "math.Sqrt2", value.Dependency("math", "")), nil
	case v.MATH_TAU:
		return value.NewValue(value.FLOAT, "(2*math.Pi)", value.Dependency("math", "")), nil
	case v.MATH_FLOAT_DIG:
		return value.NewValue(value.INTEGER, "15", value.Comment(name)), nil
	case v.MATH_FLOAT_MANT_DIG:
		return value.NewValue(value.INTEGER, "53", value.Comment(name)), nil
	case v.MATH_FLOAT_MAX_10_EXP:
		return value.NewValue(value.INTEGER, "308", value.Comment(name)), nil
	case v.MATH_FLOAT_MAX_EXP:
		return value.NewValue(value.INTEGER, "1024", value.Comment(name)), nil
	case v.MATH_FLOAT_MIN_10_EXP:
		return value.NewValue(value.INTEGER, "-307", value.Comment(name)), nil
	case v.MATH_FLOAT_MIN_EXP:
		return value.NewValue(value.INTEGER, "-1021", value.Comment(name)), nil
	case v.MATH_FLOAT_RADIX:
		return value.NewValue(value.INTEGER, "2", value.Comment(name)), nil
	case v.MATH_INTEGER_BIT:
		return value.NewValue(value.INTEGER, "64", value.Comment(name)), nil
	case v.MATH_INTEGER_MAX:
		return value.NewValue(value.INTEGER, "9223372036854775807", value.Comment(name)), nil
	case v.MATH_INTEGER_MIN:
		return value.NewValue(value.INTEGER, "-9223372036854775808", value.Comment(name)), nil

	case v.NOW:
		return value.NewValue(value.TIME, "time.Now()", value.Dependency("time", "")), nil
	case v.NOW_SEC:
		return value.NewValue(value.STRING, "ctx.NowSec()"), nil

	// Edge runtime could now know quic info
	// @Tentative
	case v.QUIC_CC_CWND,
		v.QUIC_CC_SSTHRESH,
		v.QUIC_NUM_BYTES_RECEIVED,
		v.QUIC_NUM_BYTES_SENT,
		v.QUIC_NUM_PACKETS_ACK_RECEIVED,
		v.QUIC_NUM_PACKETS_DECRYPTION_FAILED,
		v.QUIC_NUM_PACKETS_LATE_ACKED,
		v.QUIC_NUM_PACKETS_LOST,
		v.QUIC_NUM_PACKETS_RECEIVED,
		v.QUIC_NUM_PACKETS_SENT,
		v.QUIC_RTT_LATEST,
		v.QUIC_RTT_MINIMUM,
		v.QUIC_RTT_SMOOTHED,
		v.QUIC_RTT_VARIANCE:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil

	// @Tentative
	case v.REQ_DIGEST_RATIO:
		return value.NewValue(value.FLOAT, "0,4", value.Comment(name)), nil

	// @Tentative
	case v.SEGMENTED_CACHING_AUTOPURGED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_BLOCK_NUMBER:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_BLOCK_SIZE:
		return value.NewValue(value.INTEGER, "1", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_CANCELLED: // nolint:misspell
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_CLIENT_REQ_IS_OPEN_ENDED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	case v.SEGMENTED_CACHING_CLIENT_REQ_IS_RANGE:
		return value.NewValue(value.BOOL, `(ctx.RequestHeader.Get("Range") != "")`, value.Comment(name)), nil
	case v.SEGMENTED_CACHING_CLIENT_REQ_RANGE_HIGH:
		tmp := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			tmp,
			value.Prepare(
				fmt.Sprintf(`%s, err := ctx.RequestRangeHeader("high")`, tmp),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil
	case v.SEGMENTED_CACHING_CLIENT_REQ_RANGE_LOW:
		tmp := value.Temporary()
		return value.NewValue(
			value.INTEGER,
			tmp,
			value.Prepare(
				fmt.Sprintf(`%s, err := ctx.RequestRangeHeader("low")`, tmp),
				"if err != nil {",
				"return vintage.NONE, err",
				"}",
			),
		), nil

	// @Tentative
	case v.SEGMENTED_CACHING_COMPLETED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_ERROR:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_FAILED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_IS_INNER_REQ:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_IS_OUTER_REQ:
		return value.NewValue(value.BOOL, "true", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_OBJ_COMPLETE_LENGTH:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_ROUNDED_REQ_RANGE_HIGH:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_ROUNDED_REQ_RANGE_LOW:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.SEGMENTED_CACHING_TOTAL_BLOCKS:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil

	// Edge runtime could not know server info, respects Fastly fiddle one
	// @Tentative
	case v.SERVER_BILLING_REGION:
		return value.NewValue(value.STRING, "North America", value.Comment(name)), nil
	// @Tentative
	case v.SERVER_PORT:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.SERVER_POP:
		return value.NewValue(value.STRING, "CHI", value.Comment(name)), nil // Chicago POP
	// @Tentative
	case v.SERVER_DATACENTER:
		return value.NewValue(value.STRING, "Vintage", value.Comment(name)), nil
	// @Tentative
	case v.SERVER_HOSTNAME:
		return value.NewValue(value.STRING, "Vintage.Runtime", value.Comment(name)), nil
	// @Tentative
	case v.SERVER_IDENTITY:
		return value.NewValue(value.STRING, "Vintage.Runtime", value.Comment(name)), nil
	// @Tentative
	case v.SERVER_REGION:
		return value.NewValue(value.STRING, "US", value.Comment(name)), nil
	// @Tentative
	case v.SERVER_IP:
		return value.NewValue(value.IP, "vintage.LocalHost", value.Comment(name)), nil

	// Return empty string to stale.exists
	// @Tentative
	case v.STALE_EXISTS:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	// Time related variables
	case v.TIME_ELAPSED:
		return value.NewValue(value.RTIME, "time.Since(ctx.RequestStartTime)", value.Dependency("time", "")), nil
	case v.TIME_ELAPSED_MSEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(time.Since(ctx.RequestStartTime).Milliseconds())",
			value.Dependency("time", ""),
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_ELAPSED_MSEC_FRAC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(time.Since(ctx.RequestStartTime).Milliseconds() % 1000))",
			value.Dependency("time", ""),
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_ELAPSED_SEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(time.Since(ctx.RequestStartTime).Seconds()))",
			value.Dependency("time", ""),
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_ELAPSED_USEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(time.Since(ctx.RequestStartTime).Microseconds())",
			value.Dependency("time", ""),
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_ELAPSED_USEC_FRAC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(time.Since(ctx.RequestStartTime).Microseconds() % 1000000))",
			value.Dependency("time", ""),
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_START_MSEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(ctx.RequestStartTime.UnixMilli())",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_START_MSEC_FRAC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(ctx.RequestStartTime.UnixMilli() % 1000))",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_START_SEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(ctx.RequestStartTime.Unix())",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_START_USEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(ctx.RequestStartTime.UnixMicro())",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_START_USEC_FRAC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(ctx.RequestStartTime.UnixMicro() % 1000000))",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_START:
		return value.NewValue(value.TIME, "ctx.RequestStartTime"), nil
	case v.TIME_TO_FIRST_BYTE:
		return value.NewValue(value.RTIME, "ctx.TimeToFirstByte"), nil
	case v.TIME_END:
		return value.NewValue(value.TIME, "ctx.RequestEndTime"), nil
	case v.TIME_END_MSEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(ctx.RequestEndTime.UnixMilli())",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_END_MSEC_FRAC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(ctx.RequestEndTime.UnixMilli() % 1000)",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_END_SEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(ctx.RequestEndTime.Unix())",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_END_USEC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(ctx.RequestEndTime.UnixMicro())",
			value.Dependency("fmt", ""),
		), nil
	case v.TIME_END_USEC_FRAC:
		return value.NewValue(
			value.STRING,
			"fmt.Sprint(int64(ctx.RequestEndTime.UnixMicro() % 1000000))",
			value.Dependency("fmt", ""),
		), nil

	// TLS related values could not find in SDK
	// @Tentative
	case v.TLS_CLIENT_CERTIFICATE_DN,
		v.TLS_CLIENT_CERTIFICATE_ISSUER_DN,
		v.TLS_CLIENT_CERTIFICATE_RAW_CERTIFICATE_B64,
		v.TLS_CLIENT_CERTIFICATE_SERIAL_NUMBER:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	// @Tentative
	case v.TLS_CLIENT_CERTIFICATE_IS_CERT_BAD,
		v.TLS_CLIENT_CERTIFICATE_IS_CERT_EXPIRED,
		v.TLS_CLIENT_CERTIFICATE_IS_CERT_MISSING,
		v.TLS_CLIENT_CERTIFICATE_IS_CERT_REVOKED,
		v.TLS_CLIENT_CERTIFICATE_IS_CERT_UNKNOWN,
		v.TLS_CLIENT_CERTIFICATE_IS_UNKNOWN_CA,
		v.TLS_CLIENT_CERTIFICATE_IS_VERIFIED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil

	// @Tentative
	case v.TLS_CLIENT_CERTIFICATE_NOT_AFTER:
		return value.NewValue(
			value.TIME,
			"time.Now().Add(-24 * time.Hour).Add(8760 * time.Hour)",
			value.Dependency("time", ""),
			value.Comment(name),
		), nil
	// @Tentative
	case v.TLS_CLIENT_CERTIFICATE_NOT_BEFORE:
		return value.NewValue(
			value.TIME,
			"time.Now().Add(-24 * time.Hour)",
			value.Dependency("time", ""),
			value.Comment(name),
		), nil

	// @Tentative
	case v.TLS_CLIENT_CIPHERS_LIST:
		return value.NewValue(
			value.STRING,
			"130213031301C02FC02BC030C02C009EC0270067C028006B00A3009FCCA9CCA8CCAAC0AFC0ADC0A3C09FC05DC061C057C05300A2C0AEC0ACC0A2C09EC05CC060C056C052C024006AC0230040C00AC01400390038C009C01300330032009DC0A1C09DC051009CC0A0C09CC050003D003C0035002F00FF", // nolint:lll
			value.Comment(name),
		), nil
	// @Tentative
	case v.TLS_CLIENT_CIPHERS_LIST_SHA:
		return value.NewValue(
			value.STRING, "JZtiTn8H/ntxORk+XXvU2EvNoz8=", value.Comment(name),
		), nil
	// @Tentative
	case v.TLS_CLIENT_CIPHERS_LIST_TXT:
		return value.NewValue(
			value.STRING,
			"TLS_AES_256_GCM_SHA384:TLS_CHACHA20_POLY1305_SHA256:TLS_AES_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384:TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:TLS_DHE_RSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256:TLS_DHE_RSA_WITH_AES_128_CBC_SHA256:TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384:TLS_DHE_RSA_WITH_AES_256_CBC_SHA256:TLS_DHE_DSS_WITH_AES_256_GCM_SHA384:TLS_DHE_RSA_WITH_AES_256_GCM_SHA384:TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256:TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:TLS_DHE_RSA_WITH_CHACHA20_POLY1305_SHA256:TLS_ECDHE_ECDSA_WITH_AES_256_CCM_8:TLS_ECDHE_ECDSA_WITH_AES_256_CCM:TLS_DHE_RSA_WITH_AES_256_CCM_8:TLS_DHE_RSA_WITH_AES_256_CCM:TLS_ECDHE_ECDSA_WITH_ARIA_256_GCM_SHA384:TLS_ECDHE_RSA_WITH_ARIA_256_GCM_SHA384:TLS_DHE_DSS_WITH_ARIA_256_GCM_SHA384:TLS_DHE_RSA_WITH_ARIA_256_GCM_SHA384:TLS_DHE_DSS_WITH_AES_128_GCM_SHA256:TLS_ECDHE_ECDSA_WITH_AES_128_CCM_8:TLS_ECDHE_ECDSA_WITH_AES_128_CCM:TLS_DHE_RSA_WITH_AES_128_CCM_8:TLS_DHE_RSA_WITH_AES_128_CCM:TLS_ECDHE_ECDSA_WITH_ARIA_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_ARIA_128_GCM_SHA256:TLS_DHE_DSS_WITH_ARIA_128_GCM_SHA256:TLS_DHE_RSA_WITH_ARIA_128_GCM_SHA256:TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA384:TLS_DHE_DSS_WITH_AES_256_CBC_SHA256:TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA256:TLS_DHE_DSS_WITH_AES_128_CBC_SHA256:TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA:TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA:TLS_DHE_RSA_WITH_AES_256_CBC_SHA:TLS_DHE_DSS_WITH_AES_256_CBC_SHA:TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA:TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA:TLS_DHE_RSA_WITH_AES_128_CBC_SHA:TLS_DHE_DSS_WITH_AES_128_CBC_SHA:TLS_RSA_WITH_AES_256_GCM_SHA384:TLS_RSA_WITH_AES_256_CCM_8:TLS_RSA_WITH_AES_256_CCM:TLS_RSA_WITH_ARIA_256_GCM_SHA384:TLS_RSA_WITH_AES_128_GCM_SHA256:TLS_RSA_WITH_AES_128_CCM_8:TLS_RSA_WITH_AES_128_CCM:TLS_RSA_WITH_ARIA_128_GCM_SHA256:TLS_RSA_WITH_AES_256_CBC_SHA256:TLS_RSA_WITH_AES_128_CBC_SHA256:TLS_RSA_WITH_AES_256_CBC_SHA:TLS_RSA_WITH_AES_128_CBC_SHA:TLS_EMPTY_RENEGOTIATION_INFO_SCSV", // nolint:lll
			value.Comment(name),
		), nil
	// @Tentative
	case v.TLS_CLIENT_CIPHERS_SHA:
		return value.NewValue(
			value.STRING, "+7dB1w3Ov9S4Ct3HG3Qed68pSko=", value.Comment(name),
		), nil

	// @Tentative
	case v.TLS_CLIENT_HANDSHAKE_SENT_BYTES:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.TLS_CLIENT_IANA_CHOSEN_CIPHER_ID:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.TLS_CLIENT_JA3_MD5:
		return value.NewValue(
			value.STRING, "582a3b42ab84f78a5b376b1e29d6d367", value.Comment(name),
		), nil

	// case v.TLS_CLIENT_PROTOCOL:
	// 	return nil, ErrNotImplemented(name)
	case v.TLS_CLIENT_SERVERNAME,
		v.TLS_CLIENT_TLSEXTS_LIST,
		v.TLS_CLIENT_TLSEXTS_LIST_SHA,
		v.TLS_CLIENT_TLSEXTS_LIST_TXT,
		v.TLS_CLIENT_TLSEXTS_SHA:
		return value.NewValue(value.STRING, "", value.Comment(name)), nil

	// @Tentative
	case v.TRANSPORT_BW_ESTIMATE:
		return value.NewValue(value.INTEGER, "0", value.Comment(name)), nil
	// @Tentative
	case v.TRANSPORT_TYPE:
		// TODO: will be "quic" if we have support quic protocol
		return value.NewValue(value.STRING, "tcp", value.Comment(name)), nil

	// Waf related variables are all tentative value
	case v.WAF_ANOMALY_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.AnomalyScore"), nil
	case v.WAF_BLOCKED:
		return value.NewValue(value.BOOL, "ctx.Waf.Blocked"), nil
	case v.WAF_COUNTER:
		return value.NewValue(value.INTEGER, "ctx.Waf.Counter"), nil
	case v.WAF_EXECUTED:
		return value.NewValue(value.BOOL, "ctx.Waf.Executed"), nil
	case v.WAF_FAILURES:
		return value.NewValue(value.INTEGER, "ctx.Waf.Failures"), nil
	case v.WAF_HTTP_VIOLATION_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.HttpViolationScore"), nil
	case v.WAF_INBOUND_ANOMALY_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.InboundAnomalyScore"), nil
	case v.WAF_LFI_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.LFIScore"), nil
	case v.WAF_LOGDATA:
		return value.NewValue(value.STRING, "ctx.Waf.Logdata"), nil
	case v.WAF_LOGGED:
		return value.NewValue(value.BOOL, "ctx.Waf.Logged"), nil
	case v.WAF_MESSAGE:
		return value.NewValue(value.STRING, "ctx.Waf.Message"), nil
	case v.WAF_PASSED:
		return value.NewValue(value.BOOL, "ctx.Waf.Passed"), nil
	case v.WAF_PHP_INJECTION_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.PHPInjectionScore"), nil
	case v.WAF_RCE_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.RCEScore"), nil
	case v.WAF_RFI_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.RFIScore"), nil
	case v.WAF_RULE_ID:
		return value.NewValue(value.INTEGER, "ctx.Waf.RuleId"), nil
	case v.WAF_SESSION_FIXATION_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.SessionFixationScore"), nil
	case v.WAF_SEVERITY:
		return value.NewValue(value.INTEGER, "ctx.Waf.Severity"), nil
	case v.WAF_SQL_INJECTION_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.SQLInjectionScore"), nil
	case v.WAF_XSS_SCORE:
		return value.NewValue(value.INTEGER, "ctx.Waf.XSSScore"), nil

	// Workspace related values are tentative
	// @Tentative
	case v.WORKSPACE_BYTES_FREE:
		return value.NewValue(value.INTEGER, "125008", value.Comment(name)), nil
	case v.WORKSPACE_BYTES_TOTAL:
		return value.NewValue(value.INTEGER, "139392", value.Comment(name)), nil
	case v.WORKSPACE_OVERFLOWED:
		return value.NewValue(value.BOOL, "false", value.Comment(name)), nil
	}

	// Lookup dynamic variables (e.g HTTP Headers)
	if match := v.RequestHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.RequestHeader.Get("%s")`, match[1]),
		), nil
	}
	if match := v.BackendRequestHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendRequestHeader.Get("%s")`, match[1]),
		), nil
	}
	if match := v.BackendResponseHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendResponseHeader.Get("%s")`, match[1]),
		), nil
	}
	// object header is treated the same as backend response
	if match := v.ObjectHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendResponseHeader.Get("%s")`, match[1]),
		), nil
	}
	if match := v.ResponseHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.ResponseHeader.Get("%s")`, match[1]),
		), nil
	}

	return cv.VariablesImpl.Get(name)
}

// nolint:funlen,gocyclo
func (cv *CoreVariable) Set(name string, val *value.Value) (*value.Value, error) {
	switch name {
	case v.BEREQ_BETWEEN_BYTES_TIMEOUT:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Backend.BetweenBytesTimeout = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.BEREQ_CONNECT_TIMEOUT:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Backend.ConnectTimeout = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.BEREQ_FIRST_BYTE_TIMEOUT:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Backend.FirstByteTimeout = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_BROTLI:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseBrotli = %s\nctx.BackendResponseGzip = false", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_CACHEABLE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseCacheable = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_DO_ESI:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseDoESI = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_DO_STREAM:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseDoStream = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_GRACE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseGrace = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_GZIP:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseGzip = %s\nctx.BackendResponseBrotli = false", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_HIPAA:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseHipaa = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_PCI:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponsePCI = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_SAINTMODE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.SaintMode = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_STALE_IF_ERROR:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseStaleIfError = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_STALE_WHILE_REVALIDATE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseStaleWhileRevalidate = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.BERESP_TTL:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.BackendResponseTTL = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil

	case v.CLIENT_GEO_IP_OVERRIDE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.GeoIpOverride = %s", val.Conversion(value.IP).String()),
			value.FromValue(val),
		), nil

	case v.CLIENT_IDENTITY:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ClientIdentity = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil

	case v.CLIENT_SOCKET_CONGESTION_ALGORITHM:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ClientSocketCongestionAlgorithm = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.CLIENT_SOCKET_CWND:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ClientSocketCWND = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.CLIENT_SOCKET_PACE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ClientSocketPace = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.ESI_ALLOW_INSIDE_CDATA:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.EsiAllowInsideCData = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.GEOIP_IP_OVERRIDE:
		vv, err := cv.Set(v.CLIENT_GEO_IP_OVERRIDE, val)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		vv.Deprecated = true
		return vv, nil
	case v.GEOIP_USE_X_FORWARDED_FOR:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.GeoIpUseXForwardedFor = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
			value.Deprecated(),
		), nil

	case v.OBJ_GRACE:
		return cv.Set(v.BERESP_STALE_IF_ERROR, val) // Alias of beresp.stale_if_error
	case v.OBJ_TTL:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ObjectTTL = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.REQ_BACKEND:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Backend = %s", val.Conversion(value.BACKEND).String()),
			value.FromValue(val),
		), nil

	case v.REQ_ENABLE_RANGE_ON_PASS:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.EnableRangeOnPass = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.REQ_ENABLE_SEGMENTED_CACHING:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.EnableSegmentedCaching = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.REQ_ESI:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.EnableESI = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.REQ_GRACE:
		return cv.Set(v.REQ_MAX_STALE_IF_ERROR, val)
	case v.REQ_HASH:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.RequestHash = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.REQ_HASH_ALWAYS_MISS:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.HashAlwaysMiss = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.REQ_HASH_IGNORE_BUSY:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.HashIgnoreBusy = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.REQ_MAX_STALE_IF_ERROR:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.MaxStaleIfError = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.REQ_MAX_STALE_WHILE_REVALIDATE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.MaxStaleWhileRevalidate = %s", val.Conversion(value.RTIME).String()),
			value.FromValue(val),
		), nil
	case v.RESP_STALE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ResponseStale = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.RESP_STALE_IS_ERROR:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ResponseStaleIsError = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.RESP_STALE_IS_REVALIDATING:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.ResponseStaleIsRevalidating = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.SEGMENTED_CACHING_BLOCK_SIZE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.SegmentedCachingBlockSize = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil

	case v.WAF_ANOMALY_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.AnomalyScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_BLOCKED:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Blocked = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.WAF_COUNTER:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Counter = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_EXECUTED:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Executed = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.WAF_HTTP_VIOLATION_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.HttpViolationScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_INBOUND_ANOMALY_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.InboundAnomalyScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_LFI_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.LFIScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_LOGDATA:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Logdata = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.WAF_LOGGED:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Logged = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.WAF_MESSAGE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Message = %s", val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	case v.WAF_PASSED:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Passed = %s", val.Conversion(value.BOOL).String()),
			value.FromValue(val),
		), nil
	case v.WAF_RFI_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.RFIScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_RULE_ID:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.RuleId = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_SESSION_FIXATION_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.SessionFixationScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_SEVERITY:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.Severity = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	case v.WAF_XSS_SCORE:
		return value.NewValue(
			value.STRING,
			fmt.Sprintf("ctx.Waf.XSSScore = %s", val.Conversion(value.INTEGER).String()),
			value.FromValue(val),
		), nil
	}

	// Set dynamic variables (e.g HTTP Headers)
	if match := v.RequestHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.RequestHeader.Set("%s", %s)`, match[1], val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	}
	if match := v.BackendRequestHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendRequestHeader.Set("%s", %s)`, match[1], val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	}
	if match := v.BackendResponseHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendResponseHeader.Set("%s", %s)`, match[1], val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	}
	// object header is treated the same as backend response
	if match := v.ObjectHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendResponseHeader.Set("%s", %s)`, match[1], val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	}
	if match := v.ResponseHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.ResponseHeader.Set("%s", %s)`, match[1], val.Conversion(value.STRING).String()),
			value.FromValue(val),
		), nil
	}

	return cv.VariablesImpl.Set(name, val)
}

func (cv *CoreVariable) Unset(name string) (*value.Value, error) {
	if name == v.FASTLY_ERROR {
		// Nothing to do
		return value.NewValue(value.NULL, ""), nil
	}

	// Unset dynamic variables (e.g HTTP Headers)
	if match := v.RequestHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.RequestHeader.Unset("%s")`, match[1]),
		), nil
	}
	if match := v.BackendRequestHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendRequestHeader.Unset("%s")`, match[1]),
		), nil
	}
	if match := v.BackendResponseHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendResponseHeader.Unset("%s")`, match[1]),
		), nil
	}
	// object header is treated the same as backend response
	if match := v.ObjectHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.BackendResponseHeader.Unset("%s")`, match[1]),
		), nil
	}
	if match := v.ResponseHttpHeaderRegex.FindStringSubmatch(name); match != nil {
		return value.NewValue(
			value.STRING,
			fmt.Sprintf(`ctx.ResponseHeader.Unset("%s")`, match[1]),
		), nil
	}
	return cv.VariablesImpl.Unset(name)
}
