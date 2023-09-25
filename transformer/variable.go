package transformer

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/ysugimoto/vintage"
)

type CoreVariable struct{}

func (v *CoreVariable) Get(name string) (*ExpressionValue, error) {
	switch name {

	// All scoped variables
	case BEREQ_IS_CLUSTERING:
		return NewExpressionValue(vintage.BOOL, "false"), nil

	// Following values are always false in Edge Runtime
	case CLIENT_CLASS_CHECKER,
		CLIENT_CLASS_DOWNLOADER,
		CLIENT_CLASS_FEEDREADER,
		CLIENT_CLASS_FILTER,
		CLIENT_CLASS_MASQUERADING,
		CLIENT_CLASS_SPAM,
		CLIENT_PLATFORM_MEDIAPLAYER,
		REQ_IS_BACKGROUND_FETCH,
		REQ_IS_CLUSTERING,
		REQ_IS_ESI_SUBREQ,
		RESP_STALE,
		RESP_STALE_IS_ERROR,
		RESP_STALE_IS_REVALIDATING,
		WORKSPACE_OVERFLOWED:
		return NewExpressionValue(vintage.BOOL, "false"), nil

	// Edge always works on the TLS
	case FASTLY_INFO_EDGE_IS_TLS:
		return NewExpressionValue(vintage.BOOL, "true"), nil

	// Edge does not check backend is healthy but value should be true
	case REQ_BACKEND_HEALTHY:
		return NewExpressionValue(vintage.BOOL, "true"), nil

	// Edge always works on the TLS
	case REQ_IS_SSL:
		return NewExpressionValue(vintage.BOOL, "true"), nil
	case REQ_PROTOCOL:
		return NewExpressionValue(vintage.STRING, "https"), nil

	// Error is always handles on-the-fly via golang way
	case FASTLY_ERROR:
		return NewExpressionValue(vintage.STRING, ""), nil

	// Math related values
	case MATH_1_PI:
		return NewExpressionValue(vintage.FLOAT, "(1/math.Pi)", Dependency("math", "")), nil
	case MATH_2_PI:
		return NewExpressionValue(vintage.FLOAT, "(2/math.Pi)", Dependency("math", "")), nil
	case MATH_2_SQRTPI:
		return NewExpressionValue(vintage.FLOAT, "(2/math.SqrtPi)", Dependency("math", "")), nil
	case MATH_2PI:
		return NewExpressionValue(vintage.FLOAT, "(2*math.Pi)", Dependency("math", "")), nil
	case MATH_E:
		return NewExpressionValue(vintage.FLOAT, "math.E", Dependency("math", "")), nil
	case MATH_FLOAT_EPSILON:
		return NewExpressionValue(vintage.FLOAT, "math.Pow(2, -52)", Dependency("math", "")), nil
	case MATH_FLOAT_MAX:
		return NewExpressionValue(vintage.FLOAT, "math.MaxFloat64", Dependency("math", "")), nil
	case MATH_FLOAT_MIN:
		return NewExpressionValue(vintage.FLOAT, "math.SmallestNonzeroFloat64", Dependency("math", "")), nil
	case MATH_LN10:
		return NewExpressionValue(vintage.FLOAT, "math.Ln10", Dependency("math", "")), nil
	case MATH_LN2:
		return NewExpressionValue(vintage.FLOAT, "math.Ln2", Dependency("math", "")), nil
	case MATH_LOG10E:
		return NewExpressionValue(vintage.FLOAT, "math.Log10E", Dependency("math", "")), nil
	case MATH_LOG2E:
		return NewExpressionValue(vintage.FLOAT, "math.Log2E", Dependency("math", "")), nil
	case MATH_NAN:
		return NewExpressionValue(vintage.FLOAT, "math.NaN", Dependency("math", "")), nil
	case MATH_NEG_HUGE_VAL:
		return NewExpressionValue(vintage.FLOAT, "math.Inf(-1)", Dependency("math", "")), nil
	case MATH_NEG_INFINITY:
		return NewExpressionValue(vintage.FLOAT, "math.Inf(-1)", Dependency("math", "")), nil
	case MATH_PHI:
		return NewExpressionValue(vintage.FLOAT, "math.Phi", Dependency("math", "")), nil
	case MATH_PI:
		return NewExpressionValue(vintage.FLOAT, "math.Pi", Dependency("math", "")), nil
	case MATH_PI_2:
		return NewExpressionValue(vintage.FLOAT, "(math.Pi/2)", Dependency("math", "")), nil
	case MATH_PI_4:
		return NewExpressionValue(vintage.FLOAT, "(math.Pi/4)", Dependency("math", "")), nil
	case MATH_POS_HUGE_VAL:
		return NewExpressionValue(vintage.FLOAT, "math.Inf(1)", Dependency("math", "")), nil
	case MATH_POS_INFINITY:
		return NewExpressionValue(vintage.FLOAT, "math.Inf(1)", Dependency("math", "")), nil
	case MATH_SQRT1_2:
		return NewExpressionValue(vintage.FLOAT, "(1/math.Sqrt2)", Dependency("math", "")), nil
	case MATH_SQRT2:
		return NewExpressionValue(vintage.FLOAT, "math.Sqrt2", Dependency("math", "")), nil
	case MATH_TAU:
		return NewExpressionValue(vintage.FLOAT, "(2*math.Pi)", Dependency("math", "")), nil

	// Client display infos are unknown. Always returns -1
	// @Tentative
	case CLIENT_DISPLAY_HEIGHT,
		CLIENT_DISPLAY_PPI,
		CLIENT_DISPLAY_WIDTH:
		return NewExpressionValue(vintage.INTEGER, "-1"), nil

	// Client could not fully identified so returns false
	// @Tentative
	case CLIENT_IDENTIFIED:
		return NewExpressionValue(vintage.BOOL, "false"), nil

	// Client requests always returns 1, means new connection is coming
	// @Tentative
	case CLIENT_REQUESTS:
		return NewExpressionValue(vintage.INTEGER, "1"), nil

	// Edge runtime does not have POP info, returns 1
	// @Tentative
	case FASTLY_FF_VISITS_THIS_POP:
		return NewExpressionValue(vintage.INTEGER, "1"), nil

	// Returns common value -- do not consider of clustering
	// see: https://developer.fastly.com/reference/vcl/variables/miscellaneous/fastly-ff-visits-this-service/
	// Edge runtime does not have VCL service info, always returns 1
	// @Tentative
	case FASTLY_FF_VISITS_THIS_SERVICE:
		return NewExpressionValue(vintage.INTEGER, "1"), nil

	// Returns tentative value -- you may know your customer_id in the contraction :-)
	case REQ_CUSTOMER_ID:
		return NewExpressionValue(vintage.STRING, ""), nil

	// Returns fixed value which is presented on Fastly fiddle
	case MATH_FLOAT_DIG:
		return NewExpressionValue(vintage.INTEGER, "15"), nil
	case MATH_FLOAT_MANT_DIG:
		return NewExpressionValue(vintage.INTEGER, "53"), nil
	case MATH_FLOAT_MAX_10_EXP:
		return NewExpressionValue(vintage.INTEGER, "308"), nil
	case MATH_FLOAT_MAX_EXP:
		return NewExpressionValue(vintage.INTEGER, "1024"), nil
	case MATH_FLOAT_MIN_10_EXP:
		return NewExpressionValue(vintage.INTEGER, "-307"), nil
	case MATH_FLOAT_MIN_EXP:
		return NewExpressionValue(vintage.INTEGER, "-1021"), nil
	case MATH_FLOAT_RADIX:
		return NewExpressionValue(vintage.INTEGER, "2"), nil
	case MATH_INTEGER_BIT:
		return NewExpressionValue(vintage.INTEGER, "64"), nil
	case MATH_INTEGER_MAX:
		return NewExpressionValue(vintage.INTEGER, "9223372036854775807"), nil
	case MATH_INTEGER_MIN:
		return NewExpressionValue(vintage.INTEGER, "-9223372036854775808"), nil

	case REQ_RESTARTS:
		return NewExpressionValue(vintage.INTEGER, "int64(ctx.Restarts)"), nil

	// Returns always 1 because VCL is generated locally
	// @Tentative
	case REQ_VCL_GENERATION:
		return NewExpressionValue(vintage.INTEGER, "1"), nil
	case REQ_VCL_VERSION:
		return NewExpressionValue(vintage.INTEGER, "1"), nil

	// Edge runtime could not know server info, respects Fastly fiddle one
	// @Tentative
	case SERVER_BILLING_REGION:
		return NewExpressionValue(vintage.STRING, "North America"), nil
	case SERVER_PORT:
		return NewExpressionValue(vintage.INTEGER, "0"), nil
	case SERVER_POP:
		return NewExpressionValue(vintage.STRING, "CHI"), nil // Chicago POP

	// workspace related values respects Fastly fiddle one
	// @Tentative
	case WORKSPACE_BYTES_FREE:
		return NewExpressionValue(vintage.INTEGER, "125008"), nil
	case WORKSPACE_BYTES_TOTAL:
		return NewExpressionValue(vintage.INTEGER, "139392"), nil

	// Edge Runtime does not know backend and server IP info
	// @Tentative
	case BERESP_BACKEND_SRC_IP:
		return NewExpressionValue(vintage.IP, "net.IPv4(127, 0, 0, 1)", Dependency("net", "")), nil
	case SERVER_IP:
		return NewExpressionValue(vintage.IP, "net.IPv4(127, 0, 0, 1)", Dependency("net", "")), nil

	// Core request related values
	case REQ_BACKEND:
		return NewExpressionValue(vintage.BACKEND, "ctx.Backend"), nil
	case REQ_GRACE: // Alias of req.max_stale_if_error
		return NewExpressionValue(vintage.RTIME, "ctx.MaxStaleIfError"), nil
	case REQ_MAX_STALE_IF_ERROR:
		return NewExpressionValue(vintage.RTIME, "ctx.MaxStaleIfError"), nil
	case REQ_MAX_STALE_WHILE_REVALIDATE:
		return NewExpressionValue(vintage.RTIME, "ctx.MaxStaleWhileRevalidate"), nil

	// Edge runtime does not consider the state, returns empty string
	case FASTLY_INFO_STATE:
		return NewExpressionValue(vintage.STRING, ""), nil

	// Return empty string to stale.exists
	case STALE_EXISTS:
		return NewExpressionValue(vintage.STRING, ""), nil

	case LF:
		return NewExpressionValue(vintage.STRING, "\n"), nil
	case NOW_SEC:
		return NewExpressionValue(vintage.STRING, "ctx.NowSec()"), nil

	// Fixed values because Edge Runtime does not know what DC is chosen
	case SERVER_DATACENTER:
		return NewExpressionValue(vintage.STRING, "Vintage"), nil
	case SERVER_HOSTNAME:
		return NewExpressionValue(vintage.STRING, "Vintage.Runtime"), nil
	case SERVER_IDENTITY:
		return NewExpressionValue(vintage.STRING, "Vintage.Runtime"), nil
	case SERVER_REGION:
		return NewExpressionValue(vintage.STRING, "US"), nil

	// Time related variables
	case TIME_ELAPSED:
		return NewExpressionValue(vintage.RTIME, "time.Since(ctx.RequestStartTime)", Dependency("time", "")), nil
	case TIME_ELAPSED_MSEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(time.Since(ctx.RequestStartTime).Milliseconds())",
			Dependency("time", ""),
			Dependency("fmt", ""),
		), nil
	case TIME_ELAPSED_MSEC_FRAC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(time.Since(ctx.RequestStartTime).Milliseconds() % 1000))",
			Dependency("time", ""),
			Dependency("fmt", ""),
		), nil
	case TIME_ELAPSED_SEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(time.Since(ctx.RequestStartTime).Seconds()))",
			Dependency("time", ""),
			Dependency("fmt", ""),
		), nil
	case TIME_ELAPSED_USEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(time.Since(ctx.RequestStartTime).Microseconds())",
			Dependency("time", ""),
			Dependency("fmt", ""),
		), nil
	case TIME_ELAPSED_USEC_FRAC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(time.Since(ctx.RequestStartTime).Microseconds() % 1000000))",
			Dependency("time", ""),
			Dependency("fmt", ""),
		), nil
	case TIME_START_MSEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(ctx.RequestStartTime.UnixMilli())",
			Dependency("fmt", ""),
		), nil
	case TIME_START_MSEC_FRAC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(ctx.RequestStartTime.UnixMilli() % 1000))",
			Dependency("fmt", ""),
		), nil
	case TIME_START_SEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(ctx.RequestStartTime.Unix())",
			Dependency("fmt", ""),
		), nil
	case TIME_START_USEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(ctx.RequestStartTime.UnixMicro())",
			Dependency("fmt", ""),
		), nil
	case TIME_START_USEC_FRAC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(ctx.RequestStartTime.UnixMicro() % 1000000))",
			Dependency("fmt", ""),
		), nil
	case NOW:
		return NewExpressionValue(vintage.TIME, "time.Now()", Dependency("time", "")), nil
	case TIME_START:
		return NewExpressionValue(vintage.TIME, "ctx.RequestStartTime"), nil

	// Deliver scope variables
	case CLIENT_SOCKET_CONGESTION_ALGORITHM:
		return NewExpressionValue(vintage.STRING, "ctx.ClientSocketCongestionAlgorithm"), nil
	// @Tentative
	case CLIENT_SOCKET_CWND:
		return NewExpressionValue(vintage.INTEGER, "60"), nil
	case CLIENT_SOCKET_NEXTHOP:
		return NewExpressionValue(vintage.IP, "net.IPv4(127, 0, 0, 1)"), nil
	case CLIENT_SOCKET_PACE:
		// Minimum value: 128KiB
		// @see: https://developer.fastly.com/reference/vcl/variables/client-connection/client-socket-pace/
		return NewExpressionValue(vintage.INTEGER, "131072"), nil
	case CLIENT_SOCKET_PLOSS:
		return NewExpressionValue(vintage.FLOAT, "0"), nil
	case ESI_ALLOW_INSIDE_CDATA:
		return NewExpressionValue(vintage.BOOL, "ctx.EsiAllowInsideCData"), nil

	// Always false due to edge runtime does not use clustering
	case FASTLY_INFO_IS_CLUSTER_EDGE:
		return NewExpressionValue(vintage.BOOL, "false"), nil

	case REQ_ESI:
		return NewExpressionValue(vintage.BOOL, "ctx.EnableESI"), nil
	case REQ_ESI_LEVEL:
		return NewExpressionValue(vintage.BOOL, "ctx.ESILevel"), nil
	case REQ_BACKEND_NAME:
		return NewExpressionValue(vintage.STRING, "ctx.Backend.Backend()"), nil

	case RESP_IS_LOCALLY_GENERATED:
		return NewExpressionValue(vintage.BOOL, "ctx.IsLocallyGenerated"), nil

	case TIME_TO_FIRST_BYTE:
		return NewExpressionValue(vintage.RTIME, "ctx.TimeToFirstByte"), nil
	case TIME_END:
		return NewExpressionValue(vintage.TIME, "ctx.RequestEndTime"), nil
	case TIME_END_MSEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(ctx.RequestEndTime.UnixMilli())",
			Dependency("fmt", ""),
		), nil
	case TIME_END_MSEC_FRAC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(ctx.RequestEndTime.UnixMilli() % 1000)",
			Dependency("fmt", ""),
		), nil
	case TIME_END_SEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(ctx.RequestEndTime.Unix())",
			Dependency("fmt", ""),
		), nil
	case TIME_END_USEC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(ctx.RequestEndTime.UnixMicro())",
			Dependency("fmt", ""),
		), nil
	case TIME_END_USEC_FRAC:
		return NewExpressionValue(
			vintage.STRING,
			"fmt.Sprint(int64(ctx.RequestEndTime.UnixMicro() % 1000000))",
			Dependency("fmt", ""),
		), nil

	// @Tentative
	case REQ_DIGEST_RATIO:
		return NewExpressionValue(vintage.FLOAT, "0,4"), nil
	}

	// Finished Deliver

	return nil, errors.WithStack(
		fmt.Errorf("Undefined Variable %s", name),
	)
}

func (v *CoreVariable) Set(name string, value *ExpressionValue) error {
	return fmt.Errorf("Unimplemented")
}

func (v *CoreVariable) Unet(name string) error {
	return fmt.Errorf("Unimplemented")
}
