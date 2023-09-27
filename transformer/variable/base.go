package variable

import "github.com/ysugimoto/vintage/transformer/value"

// VariablesImpl is underlying variable implementation.
// All get/set/unset variables will raise an error of NotImplementedError,
// intend to check all variables must be implemented.
// So it means that all variables implementation must be extended this struct
type VariablesImpl struct {
}

func (v *VariablesImpl) Get(name string) (*value.Value, error) {
	switch name {
	case LF:
		return value.NewValue(value.STRING, "\n"), nil
	case BACKEND_CONN_IS_TLS:
		return nil, ErrNotImplemented(name)
	case BACKEND_CONN_TLS_PROTOCOL:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_CONGESTION_ALGORITHM:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_CWND:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_ADVMSS:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_BYTES_ACKED:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_BYTES_RECEIVED:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_DATA_SEGS_IN:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_DATA_SEGS_OUT:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_DELIVERY_RATE:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_DELTA_RETRANS:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_LAST_DATA_SENT:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_MAX_PACING_RATE:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_MIN_RTT:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_NOTSENT_BYTES:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_PACING_RATE:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_PMTU:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_RCV_MSS:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_RCV_RTT:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_RCV_SPACE:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_RCV_SSTHRESH:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_REORDERING:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_RTT:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_RTTVAR:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_SEGS_IN:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_SEGS_OUT:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_SND_CWND:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_SND_MSS:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_SND_SSTHRESH:
		return nil, ErrNotImplemented(name)
	case BACKEND_SOCKET_TCPI_TOTAL_RETRANS:
		return nil, ErrNotImplemented(name)
	case BEREQ_BETWEEN_BYTES_TIMEOUT:
		return nil, ErrNotImplemented(name)
	case BEREQ_BODY_BYTES_WRITTEN:
		return nil, ErrNotImplemented(name)
	case BEREQ_BYTES_WRITTEN:
		return nil, ErrNotImplemented(name)
	case BEREQ_CONNECT_TIMEOUT:
		return nil, ErrNotImplemented(name)
	case BEREQ_FIRST_BYTE_TIMEOUT:
		return nil, ErrNotImplemented(name)
	case BEREQ_HEADER_BYTES_WRITTEN:
		return nil, ErrNotImplemented(name)
	case BEREQ_IS_CLUSTERING:
		return nil, ErrNotImplemented(name)
	case BEREQ_METHOD:
		return nil, ErrNotImplemented(name)
	case BEREQ_PROTO:
		return nil, ErrNotImplemented(name)
	case BEREQ_REQUEST:
		return nil, ErrNotImplemented(name)
	case BEREQ_URL:
		return nil, ErrNotImplemented(name)
	case BEREQ_URL_BASENAME:
		return nil, ErrNotImplemented(name)
	case BEREQ_URL_DIRNAME:
		return nil, ErrNotImplemented(name)
	case BEREQ_URL_EXT:
		return nil, ErrNotImplemented(name)
	case BEREQ_URL_PATH:
		return nil, ErrNotImplemented(name)
	case BEREQ_URL_QS:
		return nil, ErrNotImplemented(name)
	case BERESP_BACKEND_ALTERNATE_IPS:
		return nil, ErrNotImplemented(name)
	case BERESP_BACKEND_IP:
		return nil, ErrNotImplemented(name)
	case BERESP_BACKEND_NAME:
		return nil, ErrNotImplemented(name)
	case BERESP_BACKEND_PORT:
		return nil, ErrNotImplemented(name)
	case BERESP_BACKEND_REQUESTS:
		return nil, ErrNotImplemented(name)
	case BERESP_BACKEND_SRC_IP:
		return nil, ErrNotImplemented(name)
	case BERESP_BROTLI:
		return nil, ErrNotImplemented(name)
	case BERESP_CACHEABLE:
		return nil, ErrNotImplemented(name)
	case BERESP_DO_ESI:
		return nil, ErrNotImplemented(name)
	case BERESP_DO_STREAM:
		return nil, ErrNotImplemented(name)
	case BERESP_GRACE:
		return nil, ErrNotImplemented(name)
	case BERESP_GZIP:
		return nil, ErrNotImplemented(name)
	case BERESP_HANDSHAKE_TIME_TO_ORIGIN_MS:
		return nil, ErrNotImplemented(name)
	case BERESP_HIPAA:
		return nil, ErrNotImplemented(name)
	case BERESP_PCI:
		return nil, ErrNotImplemented(name)
	case BERESP_PROTO:
		return nil, ErrNotImplemented(name)
	case BERESP_RESPONSE:
		return nil, ErrNotImplemented(name)
	case BERESP_STALE_IF_ERROR:
		return nil, ErrNotImplemented(name)
	case BERESP_STALE_WHILE_REVALIDATE:
		return nil, ErrNotImplemented(name)
	case BERESP_STATUS:
		return nil, ErrNotImplemented(name)
	case BERESP_TTL:
		return nil, ErrNotImplemented(name)
	case BERESP_USED_ALTERNATE_PATH_TO_ORIGIN:
		return nil, ErrNotImplemented(name)
	case CLIENT_AS_NAME:
		return nil, ErrNotImplemented(name)
	case CLIENT_AS_NUMBER:
		return nil, ErrNotImplemented(name)
	case CLIENT_BOT_NAME:
		return nil, ErrNotImplemented(name)
	case CLIENT_BROWSER_NAME:
		return nil, ErrNotImplemented(name)
	case CLIENT_BROWSER_VERSION:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_BOT:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_BROWSER:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_CHECKER:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_DOWNLOADER:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_FEEDREADER:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_FILTER:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_MASQUERADING:
		return nil, ErrNotImplemented(name)
	case CLIENT_CLASS_SPAM:
		return nil, ErrNotImplemented(name)
	case CLIENT_DISPLAY_HEIGHT:
		return nil, ErrNotImplemented(name)
	case CLIENT_DISPLAY_PPI:
		return nil, ErrNotImplemented(name)
	case CLIENT_DISPLAY_TOUCHSCREEN:
		return nil, ErrNotImplemented(name)
	case CLIENT_DISPLAY_WIDTH:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_AREA_CODE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CITY:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CITY_ASCII:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CITY_LATIN1:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CITY_UTF8:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CONN_SPEED:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CONN_TYPE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_CONTINENT_CODE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_COUNTRY_CODE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_COUNTRY_CODE3:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_COUNTRY_NAME:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_COUNTRY_NAME_ASCII:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_COUNTRY_NAME_LATIN1:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_COUNTRY_NAME_UTF8:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_GMT_OFFSET:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_IP_OVERRIDE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_LATITUDE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_LONGITUDE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_METRO_CODE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_POSTAL_CODE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_PROXY_DESCRIPTION:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_PROXY_TYPE:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_REGION:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_REGION_ASCII:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_REGION_LATIN1:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_REGION_UTF8:
		return nil, ErrNotImplemented(name)
	case CLIENT_GEO_UTC_OFFSET:
		return nil, ErrNotImplemented(name)
	case CLIENT_IDENTIFIED:
		return nil, ErrNotImplemented(name)
	case CLIENT_IDENTITY:
		return nil, ErrNotImplemented(name)
	case CLIENT_IP:
		return nil, ErrNotImplemented(name)
	case CLIENT_OS_NAME:
		return nil, ErrNotImplemented(name)
	case CLIENT_OS_VERSION:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_EREADER:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_GAMECONSOLE:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_HWTYPE:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_MEDIAPLAYER:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_MOBILE:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_SMARTTV:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_TABLET:
		return nil, ErrNotImplemented(name)
	case CLIENT_PLATFORM_TVPLAYER:
		return nil, ErrNotImplemented(name)
	case CLIENT_PORT:
		return nil, ErrNotImplemented(name)
	case CLIENT_REQUESTS:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_CONGESTION_ALGORITHM:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_CWND:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_NEXTHOP:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_PACE:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_PLOSS:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCP_INFO:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_ADVMSS:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_BYTES_ACKED:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_BYTES_RECEIVED:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_DATA_SEGS_IN:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_DATA_SEGS_OUT:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_DELIVERY_RATE:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_DELTA_RETRANS:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_LAST_DATA_SENT:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_MAX_PACING_RATE:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_MIN_RTT:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_NOTSENT_BYTES:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_PACING_RATE:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_PMTU:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_RCV_MSS:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_RCV_RTT:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_RCV_SPACE:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_RCV_SSTHRESH:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_REORDERING:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_RTT:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_RTTVAR:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_SEGS_IN:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_SEGS_OUT:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_SND_CWND:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_SND_MSS:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_SND_SSTHRESH:
		return nil, ErrNotImplemented(name)
	case CLIENT_SOCKET_TCPI_TOTAL_RETRANS:
		return nil, ErrNotImplemented(name)
	case ESI_ALLOW_INSIDE_CDATA:
		return nil, ErrNotImplemented(name)
	case FASTLY_ERROR:
		return nil, ErrNotImplemented(name)
	case FASTLY_FF_VISITS_THIS_POP:
		return nil, ErrNotImplemented(name)
	case FASTLY_FF_VISITS_THIS_POP_THIS_SERVICE:
		return nil, ErrNotImplemented(name)
	case FASTLY_FF_VISITS_THIS_SERVICE:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_EDGE_IS_TLS:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_H2_FINGERPRINT:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_H2_IS_PUSH:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_H2_STREAM_ID:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_HOST_HEADER:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_IS_CLUSTER_EDGE:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_IS_CLUSTER_SHIELD:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_IS_H2:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_IS_H3:
		return nil, ErrNotImplemented(name)
	case FASTLY_INFO_STATE:
		return nil, ErrNotImplemented(name)
	case GEOIP_AREA_CODE:
		return nil, ErrNotImplemented(name)
	case GEOIP_CITY:
		return nil, ErrNotImplemented(name)
	case GEOIP_CITY_ASCII:
		return nil, ErrNotImplemented(name)
	case GEOIP_CITY_LATIN1:
		return nil, ErrNotImplemented(name)
	case GEOIP_CITY_UTF8:
		return nil, ErrNotImplemented(name)
	case GEOIP_CONTINENT_CODE:
		return nil, ErrNotImplemented(name)
	case GEOIP_COUNTRY_CODE:
		return nil, ErrNotImplemented(name)
	case GEOIP_COUNTRY_CODE3:
		return nil, ErrNotImplemented(name)
	case GEOIP_COUNTRY_NAME:
		return nil, ErrNotImplemented(name)
	case GEOIP_COUNTRY_NAME_ASCII:
		return nil, ErrNotImplemented(name)
	case GEOIP_COUNTRY_NAME_LATIN1:
		return nil, ErrNotImplemented(name)
	case GEOIP_COUNTRY_NAME_UTF8:
		return nil, ErrNotImplemented(name)
	case GEOIP_IP_OVERRIDE:
		return nil, ErrNotImplemented(name)
	case GEOIP_LATITUDE:
		return nil, ErrNotImplemented(name)
	case GEOIP_LONGITUDE:
		return nil, ErrNotImplemented(name)
	case GEOIP_METRO_CODE:
		return nil, ErrNotImplemented(name)
	case GEOIP_POSTAL_CODE:
		return nil, ErrNotImplemented(name)
	case GEOIP_REGION:
		return nil, ErrNotImplemented(name)
	case GEOIP_REGION_ASCII:
		return nil, ErrNotImplemented(name)
	case GEOIP_REGION_LATIN1:
		return nil, ErrNotImplemented(name)
	case GEOIP_REGION_UTF8:
		return nil, ErrNotImplemented(name)
	case GEOIP_USE_X_FORWARDED_FOR:
		return nil, ErrNotImplemented(name)
	case MATH_1_PI:
		return nil, ErrNotImplemented(name)
	case MATH_2PI:
		return nil, ErrNotImplemented(name)
	case MATH_2_PI:
		return nil, ErrNotImplemented(name)
	case MATH_2_SQRTPI:
		return nil, ErrNotImplemented(name)
	case MATH_E:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_DIG:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_EPSILON:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MANT_DIG:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MAX:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MAX_10_EXP:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MAX_EXP:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MIN:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MIN_10_EXP:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_MIN_EXP:
		return nil, ErrNotImplemented(name)
	case MATH_FLOAT_RADIX:
		return nil, ErrNotImplemented(name)
	case MATH_INTEGER_BIT:
		return nil, ErrNotImplemented(name)
	case MATH_INTEGER_MAX:
		return nil, ErrNotImplemented(name)
	case MATH_INTEGER_MIN:
		return nil, ErrNotImplemented(name)
	case MATH_LN10:
		return nil, ErrNotImplemented(name)
	case MATH_LN2:
		return nil, ErrNotImplemented(name)
	case MATH_LOG10E:
		return nil, ErrNotImplemented(name)
	case MATH_LOG2E:
		return nil, ErrNotImplemented(name)
	case MATH_NAN:
		return nil, ErrNotImplemented(name)
	case MATH_NEG_HUGE_VAL:
		return nil, ErrNotImplemented(name)
	case MATH_NEG_INFINITY:
		return nil, ErrNotImplemented(name)
	case MATH_PHI:
		return nil, ErrNotImplemented(name)
	case MATH_PI:
		return nil, ErrNotImplemented(name)
	case MATH_PI_2:
		return nil, ErrNotImplemented(name)
	case MATH_PI_4:
		return nil, ErrNotImplemented(name)
	case MATH_POS_HUGE_VAL:
		return nil, ErrNotImplemented(name)
	case MATH_POS_INFINITY:
		return nil, ErrNotImplemented(name)
	case MATH_SQRT1_2:
		return nil, ErrNotImplemented(name)
	case MATH_SQRT2:
		return nil, ErrNotImplemented(name)
	case MATH_TAU:
		return nil, ErrNotImplemented(name)
	case NOW:
		return nil, ErrNotImplemented(name)
	case NOW_SEC:
		return nil, ErrNotImplemented(name)
	case OBJ_AGE:
		return nil, ErrNotImplemented(name)
	case OBJ_CACHEABLE:
		return nil, ErrNotImplemented(name)
	case OBJ_ENTERED:
		return nil, ErrNotImplemented(name)
	case OBJ_GRACE:
		return nil, ErrNotImplemented(name)
	case OBJ_HITS:
		return nil, ErrNotImplemented(name)
	case OBJ_IS_PCI:
		return nil, ErrNotImplemented(name)
	case OBJ_LASTUSE:
		return nil, ErrNotImplemented(name)
	case OBJ_PROTO:
		return nil, ErrNotImplemented(name)
	case OBJ_RESPONSE:
		return nil, ErrNotImplemented(name)
	case OBJ_STALE_IF_ERROR:
		return nil, ErrNotImplemented(name)
	case OBJ_STALE_WHILE_REVALIDATE:
		return nil, ErrNotImplemented(name)
	case OBJ_STATUS:
		return nil, ErrNotImplemented(name)
	case OBJ_TTL:
		return nil, ErrNotImplemented(name)
	case QUIC_CC_CWND:
		return nil, ErrNotImplemented(name)
	case QUIC_CC_SSTHRESH:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_BYTES_RECEIVED:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_BYTES_SENT:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_PACKETS_ACK_RECEIVED:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_PACKETS_DECRYPTION_FAILED:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_PACKETS_LATE_ACKED:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_PACKETS_LOST:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_PACKETS_RECEIVED:
		return nil, ErrNotImplemented(name)
	case QUIC_NUM_PACKETS_SENT:
		return nil, ErrNotImplemented(name)
	case QUIC_RTT_LATEST:
		return nil, ErrNotImplemented(name)
	case QUIC_RTT_MINIMUM:
		return nil, ErrNotImplemented(name)
	case QUIC_RTT_SMOOTHED:
		return nil, ErrNotImplemented(name)
	case QUIC_RTT_VARIANCE:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_HEALTHY:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_IP:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_IS_CLUSTER:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_IS_ORIGIN:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_IS_SHIELD:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_NAME:
		return nil, ErrNotImplemented(name)
	case REQ_BACKEND_PORT:
		return nil, ErrNotImplemented(name)
	case REQ_BODY:
		return nil, ErrNotImplemented(name)
	case REQ_BODY_BASE64:
		return nil, ErrNotImplemented(name)
	case REQ_BODY_BYTES_READ:
		return nil, ErrNotImplemented(name)
	case REQ_BYTES_READ:
		return nil, ErrNotImplemented(name)
	case REQ_CUSTOMER_ID:
		return nil, ErrNotImplemented(name)
	case REQ_DIGEST:
		return nil, ErrNotImplemented(name)
	case REQ_DIGEST_RATIO:
		return nil, ErrNotImplemented(name)
	case REQ_ENABLE_RANGE_ON_PASS:
		return nil, ErrNotImplemented(name)
	case REQ_ENABLE_SEGMENTED_CACHING:
		return nil, ErrNotImplemented(name)
	case REQ_ESI:
		return nil, ErrNotImplemented(name)
	case REQ_ESI_LEVEL:
		return nil, ErrNotImplemented(name)
	case REQ_GRACE:
		return nil, ErrNotImplemented(name)
	case REQ_HASH:
		return nil, ErrNotImplemented(name)
	case REQ_HASH_ALWAYS_MISS:
		return nil, ErrNotImplemented(name)
	case REQ_HASH_IGNORE_BUSY:
		return nil, ErrNotImplemented(name)
	case REQ_HEADER_BYTES_READ:
		return nil, ErrNotImplemented(name)
	case REQ_IS_BACKGROUND_FETCH:
		return nil, ErrNotImplemented(name)
	case REQ_IS_CLUSTERING:
		return nil, ErrNotImplemented(name)
	case REQ_IS_ESI_SUBREQ:
		return nil, ErrNotImplemented(name)
	case REQ_IS_IPV6:
		return nil, ErrNotImplemented(name)
	case REQ_IS_PURGE:
		return nil, ErrNotImplemented(name)
	case REQ_IS_SSL:
		return nil, ErrNotImplemented(name)
	case REQ_MAX_STALE_IF_ERROR:
		return nil, ErrNotImplemented(name)
	case REQ_MAX_STALE_WHILE_REVALIDATE:
		return nil, ErrNotImplemented(name)
	case REQ_METHOD:
		return nil, ErrNotImplemented(name)
	case REQ_POSTBODY:
		return nil, ErrNotImplemented(name)
	case REQ_PROTO:
		return nil, ErrNotImplemented(name)
	case REQ_PROTOCOL:
		return nil, ErrNotImplemented(name)
	case REQ_REQUEST:
		return nil, ErrNotImplemented(name)
	case REQ_RESTARTS:
		return nil, ErrNotImplemented(name)
	case REQ_SERVICE_ID:
		return nil, ErrNotImplemented(name)
	case REQ_TOPURL:
		return nil, ErrNotImplemented(name)
	case REQ_URL:
		return nil, ErrNotImplemented(name)
	case REQ_URL_BASENAME:
		return nil, ErrNotImplemented(name)
	case REQ_URL_DIRNAME:
		return nil, ErrNotImplemented(name)
	case REQ_URL_EXT:
		return nil, ErrNotImplemented(name)
	case REQ_URL_PATH:
		return nil, ErrNotImplemented(name)
	case REQ_URL_QS:
		return nil, ErrNotImplemented(name)
	case REQ_VCL:
		return nil, ErrNotImplemented(name)
	case REQ_VCL_GENERATION:
		return nil, ErrNotImplemented(name)
	case REQ_VCL_MD5:
		return nil, ErrNotImplemented(name)
	case REQ_VCL_VERSION:
		return nil, ErrNotImplemented(name)
	case REQ_XID:
		return nil, ErrNotImplemented(name)
	case RESP_BODY_BYTES_WRITTEN:
		return nil, ErrNotImplemented(name)
	case RESP_BYTES_WRITTEN:
		return nil, ErrNotImplemented(name)
	case RESP_COMPLETED:
		return nil, ErrNotImplemented(name)
	case RESP_HEADER_BYTES_WRITTEN:
		return nil, ErrNotImplemented(name)
	case RESP_IS_LOCALLY_GENERATED:
		return nil, ErrNotImplemented(name)
	case RESP_PROTO:
		return nil, ErrNotImplemented(name)
	case RESP_RESPONSE:
		return nil, ErrNotImplemented(name)
	case RESP_STALE:
		return nil, ErrNotImplemented(name)
	case RESP_STALE_IS_ERROR:
		return nil, ErrNotImplemented(name)
	case RESP_STALE_IS_REVALIDATING:
		return nil, ErrNotImplemented(name)
	case RESP_STATUS:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_AUTOPURGED:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_BLOCK_NUMBER:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_BLOCK_SIZE:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_CANCELLED:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_CLIENT_REQ_IS_OPEN_ENDED:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_CLIENT_REQ_IS_RANGE:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_CLIENT_REQ_RANGE_HIGH:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_CLIENT_REQ_RANGE_LOW:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_COMPLETED:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_ERROR:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_FAILED:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_IS_INNER_REQ:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_IS_OUTER_REQ:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_OBJ_COMPLETE_LENGTH:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_ROUNDED_REQ_RANGE_HIGH:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_ROUNDED_REQ_RANGE_LOW:
		return nil, ErrNotImplemented(name)
	case SEGMENTED_CACHING_TOTAL_BLOCKS:
		return nil, ErrNotImplemented(name)
	case SERVER_BILLING_REGION:
		return nil, ErrNotImplemented(name)
	case SERVER_DATACENTER:
		return nil, ErrNotImplemented(name)
	case SERVER_HOSTNAME:
		return nil, ErrNotImplemented(name)
	case SERVER_IDENTITY:
		return nil, ErrNotImplemented(name)
	case SERVER_IP:
		return nil, ErrNotImplemented(name)
	case SERVER_POP:
		return nil, ErrNotImplemented(name)
	case SERVER_PORT:
		return nil, ErrNotImplemented(name)
	case SERVER_REGION:
		return nil, ErrNotImplemented(name)
	case STALE_EXISTS:
		return nil, ErrNotImplemented(name)
	case TIME_ELAPSED:
		return nil, ErrNotImplemented(name)
	case TIME_ELAPSED_MSEC:
		return nil, ErrNotImplemented(name)
	case TIME_ELAPSED_MSEC_FRAC:
		return nil, ErrNotImplemented(name)
	case TIME_ELAPSED_SEC:
		return nil, ErrNotImplemented(name)
	case TIME_ELAPSED_USEC:
		return nil, ErrNotImplemented(name)
	case TIME_ELAPSED_USEC_FRAC:
		return nil, ErrNotImplemented(name)
	case TIME_END:
		return nil, ErrNotImplemented(name)
	case TIME_END_MSEC:
		return nil, ErrNotImplemented(name)
	case TIME_END_MSEC_FRAC:
		return nil, ErrNotImplemented(name)
	case TIME_END_SEC:
		return nil, ErrNotImplemented(name)
	case TIME_END_USEC:
		return nil, ErrNotImplemented(name)
	case TIME_END_USEC_FRAC:
		return nil, ErrNotImplemented(name)
	case TIME_START:
		return nil, ErrNotImplemented(name)
	case TIME_START_MSEC:
		return nil, ErrNotImplemented(name)
	case TIME_START_MSEC_FRAC:
		return nil, ErrNotImplemented(name)
	case TIME_START_SEC:
		return nil, ErrNotImplemented(name)
	case TIME_START_USEC:
		return nil, ErrNotImplemented(name)
	case TIME_START_USEC_FRAC:
		return nil, ErrNotImplemented(name)
	case TIME_TO_FIRST_BYTE:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_DN:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_CERT_BAD:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_CERT_EXPIRED:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_CERT_MISSING:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_CERT_REVOKED:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_CERT_UNKNOWN:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_UNKNOWN_CA:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_IS_VERIFIED:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_ISSUER_DN:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_NOT_AFTER:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_NOT_BEFORE:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_RAW_CERTIFICATE_B64:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CERTIFICATE_SERIAL_NUMBER:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CIPHER:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CIPHERS_LIST:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CIPHERS_LIST_SHA:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CIPHERS_LIST_TXT:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_CIPHERS_SHA:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_HANDSHAKE_SENT_BYTES:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_IANA_CHOSEN_CIPHER_ID:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_JA3_MD5:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_PROTOCOL:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_SERVERNAME:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_TLSEXTS_LIST:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_TLSEXTS_LIST_SHA:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_TLSEXTS_LIST_TXT:
		return nil, ErrNotImplemented(name)
	case TLS_CLIENT_TLSEXTS_SHA:
		return nil, ErrNotImplemented(name)
	case TRANSPORT_BW_ESTIMATE:
		return nil, ErrNotImplemented(name)
	case TRANSPORT_TYPE:
		return nil, ErrNotImplemented(name)
	case WAF_ANOMALY_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_BLOCKED:
		return nil, ErrNotImplemented(name)
	case WAF_COUNTER:
		return nil, ErrNotImplemented(name)
	case WAF_EXECUTED:
		return nil, ErrNotImplemented(name)
	case WAF_FAILURES:
		return nil, ErrNotImplemented(name)
	case WAF_HTTP_VIOLATION_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_INBOUND_ANOMALY_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_LFI_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_LOGDATA:
		return nil, ErrNotImplemented(name)
	case WAF_LOGGED:
		return nil, ErrNotImplemented(name)
	case WAF_MESSAGE:
		return nil, ErrNotImplemented(name)
	case WAF_PASSED:
		return nil, ErrNotImplemented(name)
	case WAF_PHP_INJECTION_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_RCE_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_RFI_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_RULE_ID:
		return nil, ErrNotImplemented(name)
	case WAF_SESSION_FIXATION_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_SEVERITY:
		return nil, ErrNotImplemented(name)
	case WAF_SQL_INJECTION_SCORE:
		return nil, ErrNotImplemented(name)
	case WAF_XSS_SCORE:
		return nil, ErrNotImplemented(name)
	case WORKSPACE_BYTES_FREE:
		return nil, ErrNotImplemented(name)
	case WORKSPACE_BYTES_TOTAL:
		return nil, ErrNotImplemented(name)
	case WORKSPACE_OVERFLOWED:
		return nil, ErrNotImplemented(name)
	}

	return nil, ErrNotFound(name)
}

func (v *VariablesImpl) Set(name string, value *value.Value) error {
	switch name {
	case BEREQ_BETWEEN_BYTES_TIMEOUT:
		return ErrNotImplemented(name)
	case BEREQ_CONNECT_TIMEOUT:
		return ErrNotImplemented(name)
	case BEREQ_FIRST_BYTE_TIMEOUT:
		return ErrNotImplemented(name)
	case BEREQ_METHOD:
		return ErrNotImplemented(name)
	case BEREQ_REQUEST:
		return ErrNotImplemented(name)
	case BEREQ_URL:
		return ErrNotImplemented(name)
	case BERESP_BROTLI:
		return ErrNotImplemented(name)
	case BERESP_CACHEABLE:
		return ErrNotImplemented(name)
	case BERESP_DO_ESI:
		return ErrNotImplemented(name)
	case BERESP_DO_STREAM:
		return ErrNotImplemented(name)
	case BERESP_GRACE:
		return ErrNotImplemented(name)
	case BERESP_GZIP:
		return ErrNotImplemented(name)
	case BERESP_HIPAA:
		return ErrNotImplemented(name)
	case BERESP_PCI:
		return ErrNotImplemented(name)
	case BERESP_RESPONSE:
		return ErrNotImplemented(name)
	case BERESP_SAINTMODE:
		return ErrNotImplemented(name)
	case BERESP_STALE_IF_ERROR:
		return ErrNotImplemented(name)
	case BERESP_STALE_WHILE_REVALIDATE:
		return ErrNotImplemented(name)
	case BERESP_STATUS:
		return ErrNotImplemented(name)
	case BERESP_TTL:
		return ErrNotImplemented(name)
	case CLIENT_GEO_IP_OVERRIDE:
		return ErrNotImplemented(name)
	case CLIENT_IDENTITY:
		return ErrNotImplemented(name)
	case CLIENT_SOCKET_CONGESTION_ALGORITHM:
		return ErrNotImplemented(name)
	case CLIENT_SOCKET_CWND:
		return ErrNotImplemented(name)
	case CLIENT_SOCKET_PACE:
		return ErrNotImplemented(name)
	case ESI_ALLOW_INSIDE_CDATA:
		return ErrNotImplemented(name)
	case GEOIP_IP_OVERRIDE:
		return ErrNotImplemented(name)
	case GEOIP_USE_X_FORWARDED_FOR:
		return ErrNotImplemented(name)
	case OBJ_GRACE:
		return ErrNotImplemented(name)
	case OBJ_RESPONSE:
		return ErrNotImplemented(name)
	case OBJ_STATUS:
		return ErrNotImplemented(name)
	case OBJ_TTL:
		return ErrNotImplemented(name)
	case REQ_BACKEND:
		return ErrNotImplemented(name)
	case REQ_ENABLE_RANGE_ON_PASS:
		return ErrNotImplemented(name)
	case REQ_ENABLE_SEGMENTED_CACHING:
		return ErrNotImplemented(name)
	case REQ_ESI:
		return ErrNotImplemented(name)
	case REQ_GRACE:
		return ErrNotImplemented(name)
	case REQ_HASH:
		return ErrNotImplemented(name)
	case REQ_HASH_ALWAYS_MISS:
		return ErrNotImplemented(name)
	case REQ_HASH_IGNORE_BUSY:
		return ErrNotImplemented(name)
	case REQ_MAX_STALE_IF_ERROR:
		return ErrNotImplemented(name)
	case REQ_MAX_STALE_WHILE_REVALIDATE:
		return ErrNotImplemented(name)
	case REQ_METHOD:
		return ErrNotImplemented(name)
	case REQ_REQUEST:
		return ErrNotImplemented(name)
	case REQ_URL:
		return ErrNotImplemented(name)
	case RESP_RESPONSE:
		return ErrNotImplemented(name)
	case RESP_STALE:
		return ErrNotImplemented(name)
	case RESP_STALE_IS_ERROR:
		return ErrNotImplemented(name)
	case RESP_STALE_IS_REVALIDATING:
		return ErrNotImplemented(name)
	case RESP_STATUS:
		return ErrNotImplemented(name)
	case SEGMENTED_CACHING_BLOCK_SIZE:
		return ErrNotImplemented(name)
	case WAF_ANOMALY_SCORE:
		return ErrNotImplemented(name)
	case WAF_BLOCKED:
		return ErrNotImplemented(name)
	case WAF_COUNTER:
		return ErrNotImplemented(name)
	case WAF_EXECUTED:
		return ErrNotImplemented(name)
	case WAF_HTTP_VIOLATION_SCORE:
		return ErrNotImplemented(name)
	case WAF_INBOUND_ANOMALY_SCORE:
		return ErrNotImplemented(name)
	case WAF_LFI_SCORE:
		return ErrNotImplemented(name)
	case WAF_LOGDATA:
		return ErrNotImplemented(name)
	case WAF_LOGGED:
		return ErrNotImplemented(name)
	case WAF_MESSAGE:
		return ErrNotImplemented(name)
	case WAF_PASSED:
		return ErrNotImplemented(name)
	case WAF_RFI_SCORE:
		return ErrNotImplemented(name)
	case WAF_RULE_ID:
		return ErrNotImplemented(name)
	case WAF_SESSION_FIXATION_SCORE:
		return ErrNotImplemented(name)
	case WAF_SEVERITY:
		return ErrNotImplemented(name)
	case WAF_XSS_SCORE:
		return ErrNotImplemented(name)
	}

	return ErrCannotSet(name)
}

func (v *VariablesImpl) Unset(name string) error {
	switch name {
	case FASTLY_ERROR:
		return ErrNotImplemented(name)
	}

	return ErrCannotUnset(name)
}
