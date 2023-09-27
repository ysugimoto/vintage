package variable

const (
	LF                                         = "LF"
	BACKEND_CONN_IS_TLS                        = "backend.conn.is_tls"
	BACKEND_CONN_TLS_PROTOCOL                  = "backend.conn.tls_protocol"
	BACKEND_SOCKET_CONGESTION_ALGORITHM        = "backend.socket.congestion_algorithm"
	BACKEND_SOCKET_CWND                        = "backend.socket.cwnd"
	BACKEND_SOCKET_TCPI_ADVMSS                 = "backend.socket.tcpi_advmss"
	BACKEND_SOCKET_TCPI_BYTES_ACKED            = "backend.socket.tcpi_bytes_acked"
	BACKEND_SOCKET_TCPI_BYTES_RECEIVED         = "backend.socket.tcpi_bytes_received"
	BACKEND_SOCKET_TCPI_DATA_SEGS_IN           = "backend.socket.tcpi_data_segs_in"
	BACKEND_SOCKET_TCPI_DATA_SEGS_OUT          = "backend.socket.tcpi_data_segs_out"
	BACKEND_SOCKET_TCPI_DELIVERY_RATE          = "backend.socket.tcpi_delivery_rate"
	BACKEND_SOCKET_TCPI_DELTA_RETRANS          = "backend.socket.tcpi_delta_retrans"
	BACKEND_SOCKET_TCPI_LAST_DATA_SENT         = "backend.socket.tcpi_last_data_sent"
	BACKEND_SOCKET_TCPI_MAX_PACING_RATE        = "backend.socket.tcpi_max_pacing_rate"
	BACKEND_SOCKET_TCPI_MIN_RTT                = "backend.socket.tcpi_min_rtt"
	BACKEND_SOCKET_TCPI_NOTSENT_BYTES          = "backend.socket.tcpi_notsent_bytes"
	BACKEND_SOCKET_TCPI_PACING_RATE            = "backend.socket.tcpi_pacing_rate"
	BACKEND_SOCKET_TCPI_PMTU                   = "backend.socket.tcpi_pmtu"
	BACKEND_SOCKET_TCPI_RCV_MSS                = "backend.socket.tcpi_rcv_mss"
	BACKEND_SOCKET_TCPI_RCV_RTT                = "backend.socket.tcpi_rcv_rtt"
	BACKEND_SOCKET_TCPI_RCV_SPACE              = "backend.socket.tcpi_rcv_space"
	BACKEND_SOCKET_TCPI_RCV_SSTHRESH           = "backend.socket.tcpi_rcv_ssthresh"
	BACKEND_SOCKET_TCPI_REORDERING             = "backend.socket.tcpi_reordering"
	BACKEND_SOCKET_TCPI_RTT                    = "backend.socket.tcpi_rtt"
	BACKEND_SOCKET_TCPI_RTTVAR                 = "backend.socket.tcpi_rttvar"
	BACKEND_SOCKET_TCPI_SEGS_IN                = "backend.socket.tcpi_segs_in"
	BACKEND_SOCKET_TCPI_SEGS_OUT               = "backend.socket.tcpi_segs_out"
	BACKEND_SOCKET_TCPI_SND_CWND               = "backend.socket.tcpi_snd_cwnd"
	BACKEND_SOCKET_TCPI_SND_MSS                = "backend.socket.tcpi_snd_mss"
	BACKEND_SOCKET_TCPI_SND_SSTHRESH           = "backend.socket.tcpi_snd_ssthresh"
	BACKEND_SOCKET_TCPI_TOTAL_RETRANS          = "backend.socket.tcpi_total_retrans"
	BEREQ_BETWEEN_BYTES_TIMEOUT                = "bereq.between_bytes_timeout"
	BEREQ_BODY_BYTES_WRITTEN                   = "bereq.body_bytes_written"
	BEREQ_BYTES_WRITTEN                        = "bereq.bytes_written"
	BEREQ_CONNECT_TIMEOUT                      = "bereq.connect_timeout"
	BEREQ_FIRST_BYTE_TIMEOUT                   = "bereq.first_byte_timeout"
	BEREQ_HEADER_BYTES_WRITTEN                 = "bereq.header_bytes_written"
	BEREQ_IS_CLUSTERING                        = "bereq.is_clustering"
	BEREQ_METHOD                               = "bereq.method"
	BEREQ_PROTO                                = "bereq.proto"
	BEREQ_REQUEST                              = "bereq.request"
	BEREQ_URL                                  = "bereq.url"
	BEREQ_URL_BASENAME                         = "bereq.url.basename"
	BEREQ_URL_DIRNAME                          = "bereq.url.dirname"
	BEREQ_URL_EXT                              = "bereq.url.ext"
	BEREQ_URL_PATH                             = "bereq.url.path"
	BEREQ_URL_QS                               = "bereq.url.qs"
	BERESP_BACKEND_ALTERNATE_IPS               = "beresp.backend.alternate_ips"
	BERESP_BACKEND_IP                          = "beresp.backend.ip"
	BERESP_BACKEND_NAME                        = "beresp.backend.name"
	BERESP_BACKEND_PORT                        = "beresp.backend.port"
	BERESP_BACKEND_REQUESTS                    = "beresp.backend.requests"
	BERESP_BACKEND_SRC_IP                      = "beresp.backend.src_ip"
	BERESP_BROTLI                              = "beresp.brotli"
	BERESP_CACHEABLE                           = "beresp.cacheable"
	BERESP_DO_ESI                              = "beresp.do_esi"
	BERESP_DO_STREAM                           = "beresp.do_stream"
	BERESP_GRACE                               = "beresp.grace"
	BERESP_GZIP                                = "beresp.gzip"
	BERESP_HANDSHAKE_TIME_TO_ORIGIN_MS         = "beresp.handshake_time_to_origin_ms"
	BERESP_HIPAA                               = "beresp.hipaa"
	BERESP_PCI                                 = "beresp.pci"
	BERESP_PROTO                               = "beresp.proto"
	BERESP_RESPONSE                            = "beresp.response"
	BERESP_SAINTMODE                           = "beresp.saintmode"
	BERESP_STALE_IF_ERROR                      = "beresp.stale_if_error"
	BERESP_STALE_WHILE_REVALIDATE              = "beresp.stale_while_revalidate"
	BERESP_STATUS                              = "beresp.status"
	BERESP_TTL                                 = "beresp.ttl"
	BERESP_USED_ALTERNATE_PATH_TO_ORIGIN       = "beresp.used_alternate_path_to_origin"
	CLIENT_AS_NAME                             = "client.as.name"
	CLIENT_AS_NUMBER                           = "client.as.number"
	CLIENT_BOT_NAME                            = "client.bot.name"
	CLIENT_BROWSER_NAME                        = "client.browser.name"
	CLIENT_BROWSER_VERSION                     = "client.browser.version"
	CLIENT_CLASS_BOT                           = "client.class.bot"
	CLIENT_CLASS_BROWSER                       = "client.class.browser"
	CLIENT_CLASS_CHECKER                       = "client.class.checker"
	CLIENT_CLASS_DOWNLOADER                    = "client.class.downloader"
	CLIENT_CLASS_FEEDREADER                    = "client.class.feedreader"
	CLIENT_CLASS_FILTER                        = "client.class.filter"
	CLIENT_CLASS_MASQUERADING                  = "client.class.masquerading"
	CLIENT_CLASS_SPAM                          = "client.class.spam"
	CLIENT_DISPLAY_HEIGHT                      = "client.display.height"
	CLIENT_DISPLAY_PPI                         = "client.display.ppi"
	CLIENT_DISPLAY_TOUCHSCREEN                 = "client.display.touchscreen"
	CLIENT_DISPLAY_WIDTH                       = "client.display.width"
	CLIENT_GEO_AREA_CODE                       = "client.geo.area_code"
	CLIENT_GEO_CITY                            = "client.geo.city"
	CLIENT_GEO_CITY_ASCII                      = "client.geo.city.ascii"
	CLIENT_GEO_CITY_LATIN1                     = "client.geo.city.latin1"
	CLIENT_GEO_CITY_UTF8                       = "client.geo.city.utf8"
	CLIENT_GEO_CONN_SPEED                      = "client.geo.conn_speed"
	CLIENT_GEO_CONN_TYPE                       = "client.geo.conn_type"
	CLIENT_GEO_CONTINENT_CODE                  = "client.geo.continent_code"
	CLIENT_GEO_COUNTRY_CODE                    = "client.geo.country_code"
	CLIENT_GEO_COUNTRY_CODE3                   = "client.geo.country_code3"
	CLIENT_GEO_COUNTRY_NAME                    = "client.geo.country_name"
	CLIENT_GEO_COUNTRY_NAME_ASCII              = "client.geo.country_name.ascii"
	CLIENT_GEO_COUNTRY_NAME_LATIN1             = "client.geo.country_name.latin1"
	CLIENT_GEO_COUNTRY_NAME_UTF8               = "client.geo.country_name.utf8"
	CLIENT_GEO_GMT_OFFSET                      = "client.geo.gmt_offset"
	CLIENT_GEO_IP_OVERRIDE                     = "client.geo.ip_override"
	CLIENT_GEO_LATITUDE                        = "client.geo.latitude"
	CLIENT_GEO_LONGITUDE                       = "client.geo.longitude"
	CLIENT_GEO_METRO_CODE                      = "client.geo.metro_code"
	CLIENT_GEO_POSTAL_CODE                     = "client.geo.postal_code"
	CLIENT_GEO_PROXY_DESCRIPTION               = "client.geo.proxy_description"
	CLIENT_GEO_PROXY_TYPE                      = "client.geo.proxy_type"
	CLIENT_GEO_REGION                          = "client.geo.region"
	CLIENT_GEO_REGION_ASCII                    = "client.geo.region.ascii"
	CLIENT_GEO_REGION_LATIN1                   = "client.geo.region.latin1"
	CLIENT_GEO_REGION_UTF8                     = "client.geo.region.utf8"
	CLIENT_GEO_UTC_OFFSET                      = "client.geo.utc_offset"
	CLIENT_IDENTIFIED                          = "client.identified"
	CLIENT_IDENTITY                            = "client.identity"
	CLIENT_IP                                  = "client.ip"
	CLIENT_OS_NAME                             = "client.os.name"
	CLIENT_OS_VERSION                          = "client.os.version"
	CLIENT_PLATFORM_EREADER                    = "client.platform.ereader"
	CLIENT_PLATFORM_GAMECONSOLE                = "client.platform.gameconsole"
	CLIENT_PLATFORM_HWTYPE                     = "client.platform.hwtype"
	CLIENT_PLATFORM_MEDIAPLAYER                = "client.platform.mediaplayer"
	CLIENT_PLATFORM_MOBILE                     = "client.platform.mobile"
	CLIENT_PLATFORM_SMARTTV                    = "client.platform.smarttv"
	CLIENT_PLATFORM_TABLET                     = "client.platform.tablet"
	CLIENT_PLATFORM_TVPLAYER                   = "client.platform.tvplayer"
	CLIENT_PORT                                = "client.port"
	CLIENT_REQUESTS                            = "client.requests"
	CLIENT_SOCKET_CONGESTION_ALGORITHM         = "client.socket.congestion_algorithm"
	CLIENT_SOCKET_CWND                         = "client.socket.cwnd"
	CLIENT_SOCKET_NEXTHOP                      = "client.socket.nexthop"
	CLIENT_SOCKET_PACE                         = "client.socket.pace"
	CLIENT_SOCKET_PLOSS                        = "client.socket.ploss"
	CLIENT_SOCKET_TCP_INFO                     = "client.socket.tcp_info"
	CLIENT_SOCKET_TCPI_ADVMSS                  = "client.socket.tcpi_advmss"
	CLIENT_SOCKET_TCPI_BYTES_ACKED             = "client.socket.tcpi_bytes_acked"
	CLIENT_SOCKET_TCPI_BYTES_RECEIVED          = "client.socket.tcpi_bytes_received"
	CLIENT_SOCKET_TCPI_DATA_SEGS_IN            = "client.socket.tcpi_data_segs_in"
	CLIENT_SOCKET_TCPI_DATA_SEGS_OUT           = "client.socket.tcpi_data_segs_out"
	CLIENT_SOCKET_TCPI_DELIVERY_RATE           = "client.socket.tcpi_delivery_rate"
	CLIENT_SOCKET_TCPI_DELTA_RETRANS           = "client.socket.tcpi_delta_retrans"
	CLIENT_SOCKET_TCPI_LAST_DATA_SENT          = "client.socket.tcpi_last_data_sent"
	CLIENT_SOCKET_TCPI_MAX_PACING_RATE         = "client.socket.tcpi_max_pacing_rate"
	CLIENT_SOCKET_TCPI_MIN_RTT                 = "client.socket.tcpi_min_rtt"
	CLIENT_SOCKET_TCPI_NOTSENT_BYTES           = "client.socket.tcpi_notsent_bytes"
	CLIENT_SOCKET_TCPI_PACING_RATE             = "client.socket.tcpi_pacing_rate"
	CLIENT_SOCKET_TCPI_PMTU                    = "client.socket.tcpi_pmtu"
	CLIENT_SOCKET_TCPI_RCV_MSS                 = "client.socket.tcpi_rcv_mss"
	CLIENT_SOCKET_TCPI_RCV_RTT                 = "client.socket.tcpi_rcv_rtt"
	CLIENT_SOCKET_TCPI_RCV_SPACE               = "client.socket.tcpi_rcv_space"
	CLIENT_SOCKET_TCPI_RCV_SSTHRESH            = "client.socket.tcpi_rcv_ssthresh"
	CLIENT_SOCKET_TCPI_REORDERING              = "client.socket.tcpi_reordering"
	CLIENT_SOCKET_TCPI_RTT                     = "client.socket.tcpi_rtt"
	CLIENT_SOCKET_TCPI_RTTVAR                  = "client.socket.tcpi_rttvar"
	CLIENT_SOCKET_TCPI_SEGS_IN                 = "client.socket.tcpi_segs_in"
	CLIENT_SOCKET_TCPI_SEGS_OUT                = "client.socket.tcpi_segs_out"
	CLIENT_SOCKET_TCPI_SND_CWND                = "client.socket.tcpi_snd_cwnd"
	CLIENT_SOCKET_TCPI_SND_MSS                 = "client.socket.tcpi_snd_mss"
	CLIENT_SOCKET_TCPI_SND_SSTHRESH            = "client.socket.tcpi_snd_ssthresh"
	CLIENT_SOCKET_TCPI_TOTAL_RETRANS           = "client.socket.tcpi_total_retrans"
	ESI_ALLOW_INSIDE_CDATA                     = "esi.allow_inside_cdata"
	FASTLY_ERROR                               = "fastly.error"
	FASTLY_FF_VISITS_THIS_POP                  = "fastly.ff.visits_this_pop"
	FASTLY_FF_VISITS_THIS_POP_THIS_SERVICE     = "fastly.ff.visits_this_pop_this_service"
	FASTLY_FF_VISITS_THIS_SERVICE              = "fastly.ff.visits_this_service"
	FASTLY_INFO_EDGE_IS_TLS                    = "fastly_info.edge.is_tls"
	FASTLY_INFO_H2_FINGERPRINT                 = "fastly_info.h2.fingerprint"
	FASTLY_INFO_H2_IS_PUSH                     = "fastly_info.h2.is_push"
	FASTLY_INFO_H2_STREAM_ID                   = "fastly_info.h2.stream_id"
	FASTLY_INFO_HOST_HEADER                    = "fastly_info.host_header"
	FASTLY_INFO_IS_CLUSTER_EDGE                = "fastly_info.is_cluster_edge"
	FASTLY_INFO_IS_CLUSTER_SHIELD              = "fastly_info.is_cluster_shield"
	FASTLY_INFO_IS_H2                          = "fastly_info.is_h2"
	FASTLY_INFO_IS_H3                          = "fastly_info.is_h3"
	FASTLY_INFO_STATE                          = "fastly_info.state"
	GEOIP_AREA_CODE                            = "geoip.area_code"
	GEOIP_CITY                                 = "geoip.city"
	GEOIP_CITY_ASCII                           = "geoip.city.ascii"
	GEOIP_CITY_LATIN1                          = "geoip.city.latin1"
	GEOIP_CITY_UTF8                            = "geoip.city.utf8"
	GEOIP_CONTINENT_CODE                       = "geoip.continent_code"
	GEOIP_COUNTRY_CODE                         = "geoip.country_code"
	GEOIP_COUNTRY_CODE3                        = "geoip.country_code3"
	GEOIP_COUNTRY_NAME                         = "geoip.country_name"
	GEOIP_COUNTRY_NAME_ASCII                   = "geoip.country_name.ascii"
	GEOIP_COUNTRY_NAME_LATIN1                  = "geoip.country_name.latin1"
	GEOIP_COUNTRY_NAME_UTF8                    = "geoip.country_name.utf8"
	GEOIP_IP_OVERRIDE                          = "geoip.ip_override"
	GEOIP_LATITUDE                             = "geoip.latitude"
	GEOIP_LONGITUDE                            = "geoip.longitude"
	GEOIP_METRO_CODE                           = "geoip.metro_code"
	GEOIP_POSTAL_CODE                          = "geoip.postal_code"
	GEOIP_REGION                               = "geoip.region"
	GEOIP_REGION_ASCII                         = "geoip.region.ascii"
	GEOIP_REGION_LATIN1                        = "geoip.region.latin1"
	GEOIP_REGION_UTF8                          = "geoip.region.utf8"
	GEOIP_USE_X_FORWARDED_FOR                  = "geoip.use_x_forwarded_for"
	MATH_1_PI                                  = "math.1_PI"
	MATH_2PI                                   = "math.2PI"
	MATH_2_PI                                  = "math.2_PI"
	MATH_2_SQRTPI                              = "math.2_SQRTPI"
	MATH_E                                     = "math.E"
	MATH_FLOAT_DIG                             = "math.FLOAT_DIG"
	MATH_FLOAT_EPSILON                         = "math.FLOAT_EPSILON"
	MATH_FLOAT_MANT_DIG                        = "math.FLOAT_MANT_DIG"
	MATH_FLOAT_MAX                             = "math.FLOAT_MAX"
	MATH_FLOAT_MAX_10_EXP                      = "math.FLOAT_MAX_10_EXP"
	MATH_FLOAT_MAX_EXP                         = "math.FLOAT_MAX_EXP"
	MATH_FLOAT_MIN                             = "math.FLOAT_MIN"
	MATH_FLOAT_MIN_10_EXP                      = "math.FLOAT_MIN_10_EXP"
	MATH_FLOAT_MIN_EXP                         = "math.FLOAT_MIN_EXP"
	MATH_FLOAT_RADIX                           = "math.FLOAT_RADIX"
	MATH_INTEGER_BIT                           = "math.INTEGER_BIT"
	MATH_INTEGER_MAX                           = "math.INTEGER_MAX"
	MATH_INTEGER_MIN                           = "math.INTEGER_MIN"
	MATH_LN10                                  = "math.LN10"
	MATH_LN2                                   = "math.LN2"
	MATH_LOG10E                                = "math.LOG10E"
	MATH_LOG2E                                 = "math.LOG2E"
	MATH_NAN                                   = "math.NAN"
	MATH_NEG_HUGE_VAL                          = "math.NEG_HUGE_VAL"
	MATH_NEG_INFINITY                          = "math.NEG_INFINITY"
	MATH_PHI                                   = "math.PHI"
	MATH_PI                                    = "math.PI"
	MATH_PI_2                                  = "math.PI_2"
	MATH_PI_4                                  = "math.PI_4"
	MATH_POS_HUGE_VAL                          = "math.POS_HUGE_VAL"
	MATH_POS_INFINITY                          = "math.POS_INFINITY"
	MATH_SQRT1_2                               = "math.SQRT1_2"
	MATH_SQRT2                                 = "math.SQRT2"
	MATH_TAU                                   = "math.TAU"
	NOW                                        = "now"
	NOW_SEC                                    = "now.sec"
	OBJ_AGE                                    = "obj.age"
	OBJ_CACHEABLE                              = "obj.cacheable"
	OBJ_ENTERED                                = "obj.entered"
	OBJ_GRACE                                  = "obj.grace"
	OBJ_HITS                                   = "obj.hits"
	OBJ_IS_PCI                                 = "obj.is_pci"
	OBJ_LASTUSE                                = "obj.lastuse"
	OBJ_PROTO                                  = "obj.proto"
	OBJ_RESPONSE                               = "obj.response"
	OBJ_STALE_IF_ERROR                         = "obj.stale_if_error"
	OBJ_STALE_WHILE_REVALIDATE                 = "obj.stale_while_revalidate"
	OBJ_STATUS                                 = "obj.status"
	OBJ_TTL                                    = "obj.ttl"
	QUIC_CC_CWND                               = "quic.cc.cwnd"
	QUIC_CC_SSTHRESH                           = "quic.cc.ssthresh"
	QUIC_NUM_BYTES_RECEIVED                    = "quic.num_bytes.received"
	QUIC_NUM_BYTES_SENT                        = "quic.num_bytes.sent"
	QUIC_NUM_PACKETS_ACK_RECEIVED              = "quic.num_packets.ack_received"
	QUIC_NUM_PACKETS_DECRYPTION_FAILED         = "quic.num_packets.decryption_failed"
	QUIC_NUM_PACKETS_LATE_ACKED                = "quic.num_packets.late_acked"
	QUIC_NUM_PACKETS_LOST                      = "quic.num_packets.lost"
	QUIC_NUM_PACKETS_RECEIVED                  = "quic.num_packets.received"
	QUIC_NUM_PACKETS_SENT                      = "quic.num_packets.sent"
	QUIC_RTT_LATEST                            = "quic.rtt.latest"
	QUIC_RTT_MINIMUM                           = "quic.rtt.minimum"
	QUIC_RTT_SMOOTHED                          = "quic.rtt.smoothed"
	QUIC_RTT_VARIANCE                          = "quic.rtt.variance"
	REQ_BACKEND                                = "req.backend"
	REQ_BACKEND_HEALTHY                        = "req.backend.healthy"
	REQ_BACKEND_IP                             = "req.backend.ip"
	REQ_BACKEND_IS_CLUSTER                     = "req.backend.is_cluster"
	REQ_BACKEND_IS_ORIGIN                      = "req.backend.is_origin"
	REQ_BACKEND_IS_SHIELD                      = "req.backend.is_shield"
	REQ_BACKEND_NAME                           = "req.backend.name"
	REQ_BACKEND_PORT                           = "req.backend.port"
	REQ_BODY                                   = "req.body"
	REQ_BODY_BASE64                            = "req.body.base64"
	REQ_BODY_BYTES_READ                        = "req.body_bytes_read"
	REQ_BYTES_READ                             = "req.bytes_read"
	REQ_CUSTOMER_ID                            = "req.customer_id"
	REQ_DIGEST                                 = "req.digest"
	REQ_DIGEST_RATIO                           = "req.digest.ratio"
	REQ_ENABLE_RANGE_ON_PASS                   = "req.enable_range_on_pass"
	REQ_ENABLE_SEGMENTED_CACHING               = "req.enable_segmented_caching"
	REQ_ESI                                    = "req.esi"
	REQ_ESI_LEVEL                              = "req.esi_level"
	REQ_GRACE                                  = "req.grace"
	REQ_HASH                                   = "req.hash"
	REQ_HASH_ALWAYS_MISS                       = "req.hash_always_miss"
	REQ_HASH_IGNORE_BUSY                       = "req.hash_ignore_busy"
	REQ_HEADER_BYTES_READ                      = "req.header_bytes_read"
	REQ_IS_BACKGROUND_FETCH                    = "req.is_background_fetch"
	REQ_IS_CLUSTERING                          = "req.is_clustering"
	REQ_IS_ESI_SUBREQ                          = "req.is_esi_subreq"
	REQ_IS_IPV6                                = "req.is_ipv6"
	REQ_IS_PURGE                               = "req.is_purge"
	REQ_IS_SSL                                 = "req.is_ssl"
	REQ_MAX_STALE_IF_ERROR                     = "req.max_stale_if_error"
	REQ_MAX_STALE_WHILE_REVALIDATE             = "req.max_stale_while_revalidate"
	REQ_METHOD                                 = "req.method"
	REQ_POSTBODY                               = "req.postbody"
	REQ_PROTO                                  = "req.proto"
	REQ_PROTOCOL                               = "req.protocol"
	REQ_REQUEST                                = "req.request"
	REQ_RESTARTS                               = "req.restarts"
	REQ_SERVICE_ID                             = "req.service_id"
	REQ_TOPURL                                 = "req.topurl"
	REQ_URL                                    = "req.url"
	REQ_URL_BASENAME                           = "req.url.basename"
	REQ_URL_DIRNAME                            = "req.url.dirname"
	REQ_URL_EXT                                = "req.url.ext"
	REQ_URL_PATH                               = "req.url.path"
	REQ_URL_QS                                 = "req.url.qs"
	REQ_VCL                                    = "req.vcl"
	REQ_VCL_GENERATION                         = "req.vcl.generation"
	REQ_VCL_MD5                                = "req.vcl.md5"
	REQ_VCL_VERSION                            = "req.vcl.version"
	REQ_XID                                    = "req.xid"
	RESP_BODY_BYTES_WRITTEN                    = "resp.body_bytes_written"
	RESP_BYTES_WRITTEN                         = "resp.bytes_written"
	RESP_COMPLETED                             = "resp.completed"
	RESP_HEADER_BYTES_WRITTEN                  = "resp.header_bytes_written"
	RESP_IS_LOCALLY_GENERATED                  = "resp.is_locally_generated"
	RESP_PROTO                                 = "resp.proto"
	RESP_RESPONSE                              = "resp.response"
	RESP_STALE                                 = "resp.stale"
	RESP_STALE_IS_ERROR                        = "resp.stale.is_error"
	RESP_STALE_IS_REVALIDATING                 = "resp.stale.is_revalidating"
	RESP_STATUS                                = "resp.status"
	SEGMENTED_CACHING_AUTOPURGED               = "segmented_caching.autopurged"
	SEGMENTED_CACHING_BLOCK_NUMBER             = "segmented_caching.block_number"
	SEGMENTED_CACHING_BLOCK_SIZE               = "segmented_caching.block_size"
	SEGMENTED_CACHING_CANCELLED                = "segmented_caching.cancelled"
	SEGMENTED_CACHING_CLIENT_REQ_IS_OPEN_ENDED = "segmented_caching.client_req.is_open_ended"
	SEGMENTED_CACHING_CLIENT_REQ_IS_RANGE      = "segmented_caching.client_req.is_range"
	SEGMENTED_CACHING_CLIENT_REQ_RANGE_HIGH    = "segmented_caching.client_req.range_high"
	SEGMENTED_CACHING_CLIENT_REQ_RANGE_LOW     = "segmented_caching.client_req.range_low"
	SEGMENTED_CACHING_COMPLETED                = "segmented_caching.completed"
	SEGMENTED_CACHING_ERROR                    = "segmented_caching.error"
	SEGMENTED_CACHING_FAILED                   = "segmented_caching.failed"
	SEGMENTED_CACHING_IS_INNER_REQ             = "segmented_caching.is_inner_req"
	SEGMENTED_CACHING_IS_OUTER_REQ             = "segmented_caching.is_outer_req"
	SEGMENTED_CACHING_OBJ_COMPLETE_LENGTH      = "segmented_caching.obj.complete_length"
	SEGMENTED_CACHING_ROUNDED_REQ_RANGE_HIGH   = "segmented_caching.rounded_req.range_high"
	SEGMENTED_CACHING_ROUNDED_REQ_RANGE_LOW    = "segmented_caching.rounded_req.range_low"
	SEGMENTED_CACHING_TOTAL_BLOCKS             = "segmented_caching.total_blocks"
	SERVER_BILLING_REGION                      = "server.billing_region"
	SERVER_DATACENTER                          = "server.datacenter"
	SERVER_HOSTNAME                            = "server.hostname"
	SERVER_IDENTITY                            = "server.identity"
	SERVER_IP                                  = "server.ip"
	SERVER_POP                                 = "server.pop"
	SERVER_PORT                                = "server.port"
	SERVER_REGION                              = "server.region"
	STALE_EXISTS                               = "stale.exists"
	TIME_ELAPSED                               = "time.elapsed"
	TIME_ELAPSED_MSEC                          = "time.elapsed.msec"
	TIME_ELAPSED_MSEC_FRAC                     = "time.elapsed.msec_frac"
	TIME_ELAPSED_SEC                           = "time.elapsed.sec"
	TIME_ELAPSED_USEC                          = "time.elapsed.usec"
	TIME_ELAPSED_USEC_FRAC                     = "time.elapsed.usec_frac"
	TIME_END                                   = "time.end"
	TIME_END_MSEC                              = "time.end.msec"
	TIME_END_MSEC_FRAC                         = "time.end.msec_frac"
	TIME_END_SEC                               = "time.end.sec"
	TIME_END_USEC                              = "time.end.usec"
	TIME_END_USEC_FRAC                         = "time.end.usec_frac"
	TIME_START                                 = "time.start"
	TIME_START_MSEC                            = "time.start.msec"
	TIME_START_MSEC_FRAC                       = "time.start.msec_frac"
	TIME_START_SEC                             = "time.start.sec"
	TIME_START_USEC                            = "time.start.usec"
	TIME_START_USEC_FRAC                       = "time.start.usec_frac"
	TIME_TO_FIRST_BYTE                         = "time.to_first_byte"
	TLS_CLIENT_CERTIFICATE_DN                  = "tls.client.certificate.dn"
	TLS_CLIENT_CERTIFICATE_IS_CERT_BAD         = "tls.client.certificate.is_cert_bad"
	TLS_CLIENT_CERTIFICATE_IS_CERT_EXPIRED     = "tls.client.certificate.is_cert_expired"
	TLS_CLIENT_CERTIFICATE_IS_CERT_MISSING     = "tls.client.certificate.is_cert_missing"
	TLS_CLIENT_CERTIFICATE_IS_CERT_REVOKED     = "tls.client.certificate.is_cert_revoked"
	TLS_CLIENT_CERTIFICATE_IS_CERT_UNKNOWN     = "tls.client.certificate.is_cert_unknown"
	TLS_CLIENT_CERTIFICATE_IS_UNKNOWN_CA       = "tls.client.certificate.is_unknown_ca"
	TLS_CLIENT_CERTIFICATE_IS_VERIFIED         = "tls.client.certificate.is_verified"
	TLS_CLIENT_CERTIFICATE_ISSUER_DN           = "tls.client.certificate.issuer_dn"
	TLS_CLIENT_CERTIFICATE_NOT_AFTER           = "tls.client.certificate.not_after"
	TLS_CLIENT_CERTIFICATE_NOT_BEFORE          = "tls.client.certificate.not_before"
	TLS_CLIENT_CERTIFICATE_RAW_CERTIFICATE_B64 = "tls.client.certificate.raw_certificate_b64"
	TLS_CLIENT_CERTIFICATE_SERIAL_NUMBER       = "tls.client.certificate.serial_number"
	TLS_CLIENT_CIPHER                          = "tls.client.cipher"
	TLS_CLIENT_CIPHERS_LIST                    = "tls.client.ciphers_list"
	TLS_CLIENT_CIPHERS_LIST_SHA                = "tls.client.ciphers_list_sha"
	TLS_CLIENT_CIPHERS_LIST_TXT                = "tls.client.ciphers_list_txt"
	TLS_CLIENT_CIPHERS_SHA                     = "tls.client.ciphers_sha"
	TLS_CLIENT_HANDSHAKE_SENT_BYTES            = "tls.client.handshake_sent_bytes"
	TLS_CLIENT_IANA_CHOSEN_CIPHER_ID           = "tls.client.iana_chosen_cipher_id"
	TLS_CLIENT_JA3_MD5                         = "tls.client.ja3_md5"
	TLS_CLIENT_PROTOCOL                        = "tls.client.protocol"
	TLS_CLIENT_SERVERNAME                      = "tls.client.servername"
	TLS_CLIENT_TLSEXTS_LIST                    = "tls.client.tlsexts_list"
	TLS_CLIENT_TLSEXTS_LIST_SHA                = "tls.client.tlsexts_list_sha"
	TLS_CLIENT_TLSEXTS_LIST_TXT                = "tls.client.tlsexts_list_txt"
	TLS_CLIENT_TLSEXTS_SHA                     = "tls.client.tlsexts_sha"
	TRANSPORT_BW_ESTIMATE                      = "transport.bw_estimate"
	TRANSPORT_TYPE                             = "transport.type"
	WAF_ANOMALY_SCORE                          = "waf.anomaly_score"
	WAF_BLOCKED                                = "waf.blocked"
	WAF_COUNTER                                = "waf.counter"
	WAF_EXECUTED                               = "waf.executed"
	WAF_FAILURES                               = "waf.failures"
	WAF_HTTP_VIOLATION_SCORE                   = "waf.http_violation_score"
	WAF_INBOUND_ANOMALY_SCORE                  = "waf.inbound_anomaly_score"
	WAF_LFI_SCORE                              = "waf.lfi_score"
	WAF_LOGDATA                                = "waf.logdata"
	WAF_LOGGED                                 = "waf.logged"
	WAF_MESSAGE                                = "waf.message"
	WAF_PASSED                                 = "waf.passed"
	WAF_PHP_INJECTION_SCORE                    = "waf.php_injection_score"
	WAF_RCE_SCORE                              = "waf.rce_score"
	WAF_RFI_SCORE                              = "waf.rfi_score"
	WAF_RULE_ID                                = "waf.rule_id"
	WAF_SESSION_FIXATION_SCORE                 = "waf.session_fixation_score"
	WAF_SEVERITY                               = "waf.severity"
	WAF_SQL_INJECTION_SCORE                    = "waf.sql_injection_score"
	WAF_XSS_SCORE                              = "waf.xss_score"
	WORKSPACE_BYTES_FREE                       = "workspace.bytes_free"
	WORKSPACE_BYTES_TOTAL                      = "workspace.bytes_total"
	WORKSPACE_OVERFLOWED                       = "workspace.overflowed"
)

var Getable = []string{
	"LF",
	"backend.conn.is_tls",
	"backend.conn.tls_protocol",
	"backend.socket.congestion_algorithm",
	"backend.socket.cwnd",
	"backend.socket.tcpi_advmss",
	"backend.socket.tcpi_bytes_acked",
	"backend.socket.tcpi_bytes_received",
	"backend.socket.tcpi_data_segs_in",
	"backend.socket.tcpi_data_segs_out",
	"backend.socket.tcpi_delivery_rate",
	"backend.socket.tcpi_delta_retrans",
	"backend.socket.tcpi_last_data_sent",
	"backend.socket.tcpi_max_pacing_rate",
	"backend.socket.tcpi_min_rtt",
	"backend.socket.tcpi_notsent_bytes",
	"backend.socket.tcpi_pacing_rate",
	"backend.socket.tcpi_pmtu",
	"backend.socket.tcpi_rcv_mss",
	"backend.socket.tcpi_rcv_rtt",
	"backend.socket.tcpi_rcv_space",
	"backend.socket.tcpi_rcv_ssthresh",
	"backend.socket.tcpi_reordering",
	"backend.socket.tcpi_rtt",
	"backend.socket.tcpi_rttvar",
	"backend.socket.tcpi_segs_in",
	"backend.socket.tcpi_segs_out",
	"backend.socket.tcpi_snd_cwnd",
	"backend.socket.tcpi_snd_mss",
	"backend.socket.tcpi_snd_ssthresh",
	"backend.socket.tcpi_total_retrans",
	"bereq.between_bytes_timeout",
	"bereq.body_bytes_written",
	"bereq.bytes_written",
	"bereq.connect_timeout",
	"bereq.first_byte_timeout",
	"bereq.header_bytes_written",
	"bereq.is_clustering",
	"bereq.method",
	"bereq.proto",
	"bereq.request",
	"bereq.url",
	"bereq.url.basename",
	"bereq.url.dirname",
	"bereq.url.ext",
	"bereq.url.path",
	"bereq.url.qs",
	"beresp.backend.alternate_ips",
	"beresp.backend.ip",
	"beresp.backend.name",
	"beresp.backend.port",
	"beresp.backend.requests",
	"beresp.backend.src_ip",
	"beresp.brotli",
	"beresp.cacheable",
	"beresp.do_esi",
	"beresp.do_stream",
	"beresp.grace",
	"beresp.gzip",
	"beresp.handshake_time_to_origin_ms",
	"beresp.hipaa",
	"beresp.pci",
	"beresp.proto",
	"beresp.response",
	"beresp.stale_if_error",
	"beresp.stale_while_revalidate",
	"beresp.status",
	"beresp.ttl",
	"beresp.used_alternate_path_to_origin",
	"client.as.name",
	"client.as.number",
	"client.bot.name",
	"client.browser.name",
	"client.browser.version",
	"client.class.bot",
	"client.class.browser",
	"client.class.checker",
	"client.class.downloader",
	"client.class.feedreader",
	"client.class.filter",
	"client.class.masquerading",
	"client.class.spam",
	"client.display.height",
	"client.display.ppi",
	"client.display.touchscreen",
	"client.display.width",
	"client.geo.area_code",
	"client.geo.city",
	"client.geo.city.ascii",
	"client.geo.city.latin1",
	"client.geo.city.utf8",
	"client.geo.conn_speed",
	"client.geo.conn_type",
	"client.geo.continent_code",
	"client.geo.country_code",
	"client.geo.country_code3",
	"client.geo.country_name",
	"client.geo.country_name.ascii",
	"client.geo.country_name.latin1",
	"client.geo.country_name.utf8",
	"client.geo.gmt_offset",
	"client.geo.ip_override",
	"client.geo.latitude",
	"client.geo.longitude",
	"client.geo.metro_code",
	"client.geo.postal_code",
	"client.geo.proxy_description",
	"client.geo.proxy_type",
	"client.geo.region",
	"client.geo.region.ascii",
	"client.geo.region.latin1",
	"client.geo.region.utf8",
	"client.geo.utc_offset",
	"client.identified",
	"client.identity",
	"client.ip",
	"client.os.name",
	"client.os.version",
	"client.platform.ereader",
	"client.platform.gameconsole",
	"client.platform.hwtype",
	"client.platform.mediaplayer",
	"client.platform.mobile",
	"client.platform.smarttv",
	"client.platform.tablet",
	"client.platform.tvplayer",
	"client.port",
	"client.requests",
	"client.socket.congestion_algorithm",
	"client.socket.cwnd",
	"client.socket.nexthop",
	"client.socket.pace",
	"client.socket.ploss",
	"client.socket.tcp_info",
	"client.socket.tcpi_advmss",
	"client.socket.tcpi_bytes_acked",
	"client.socket.tcpi_bytes_received",
	"client.socket.tcpi_data_segs_in",
	"client.socket.tcpi_data_segs_out",
	"client.socket.tcpi_delivery_rate",
	"client.socket.tcpi_delta_retrans",
	"client.socket.tcpi_last_data_sent",
	"client.socket.tcpi_max_pacing_rate",
	"client.socket.tcpi_min_rtt",
	"client.socket.tcpi_notsent_bytes",
	"client.socket.tcpi_pacing_rate",
	"client.socket.tcpi_pmtu",
	"client.socket.tcpi_rcv_mss",
	"client.socket.tcpi_rcv_rtt",
	"client.socket.tcpi_rcv_space",
	"client.socket.tcpi_rcv_ssthresh",
	"client.socket.tcpi_reordering",
	"client.socket.tcpi_rtt",
	"client.socket.tcpi_rttvar",
	"client.socket.tcpi_segs_in",
	"client.socket.tcpi_segs_out",
	"client.socket.tcpi_snd_cwnd",
	"client.socket.tcpi_snd_mss",
	"client.socket.tcpi_snd_ssthresh",
	"client.socket.tcpi_total_retrans",
	"esi.allow_inside_cdata",
	"fastly.error",
	"fastly.ff.visits_this_pop",
	"fastly.ff.visits_this_pop_this_service",
	"fastly.ff.visits_this_service",
	"fastly_info.edge.is_tls",
	"fastly_info.h2.fingerprint",
	"fastly_info.h2.is_push",
	"fastly_info.h2.stream_id",
	"fastly_info.host_header",
	"fastly_info.is_cluster_edge",
	"fastly_info.is_cluster_shield",
	"fastly_info.is_h2",
	"fastly_info.is_h3",
	"fastly_info.state",
	"geoip.area_code",
	"geoip.city",
	"geoip.city.ascii",
	"geoip.city.latin1",
	"geoip.city.utf8",
	"geoip.continent_code",
	"geoip.country_code",
	"geoip.country_code3",
	"geoip.country_name",
	"geoip.country_name.ascii",
	"geoip.country_name.latin1",
	"geoip.country_name.utf8",
	"geoip.ip_override",
	"geoip.latitude",
	"geoip.longitude",
	"geoip.metro_code",
	"geoip.postal_code",
	"geoip.region",
	"geoip.region.ascii",
	"geoip.region.latin1",
	"geoip.region.utf8",
	"geoip.use_x_forwarded_for",
	"math.1_PI",
	"math.2PI",
	"math.2_PI",
	"math.2_SQRTPI",
	"math.E",
	"math.FLOAT_DIG",
	"math.FLOAT_EPSILON",
	"math.FLOAT_MANT_DIG",
	"math.FLOAT_MAX",
	"math.FLOAT_MAX_10_EXP",
	"math.FLOAT_MAX_EXP",
	"math.FLOAT_MIN",
	"math.FLOAT_MIN_10_EXP",
	"math.FLOAT_MIN_EXP",
	"math.FLOAT_RADIX",
	"math.INTEGER_BIT",
	"math.INTEGER_MAX",
	"math.INTEGER_MIN",
	"math.LN10",
	"math.LN2",
	"math.LOG10E",
	"math.LOG2E",
	"math.NAN",
	"math.NEG_HUGE_VAL",
	"math.NEG_INFINITY",
	"math.PHI",
	"math.PI",
	"math.PI_2",
	"math.PI_4",
	"math.POS_HUGE_VAL",
	"math.POS_INFINITY",
	"math.SQRT1_2",
	"math.SQRT2",
	"math.TAU",
	"now",
	"now.sec",
	"obj.age",
	"obj.cacheable",
	"obj.entered",
	"obj.grace",
	"obj.hits",
	"obj.is_pci",
	"obj.lastuse",
	"obj.proto",
	"obj.response",
	"obj.stale_if_error",
	"obj.stale_while_revalidate",
	"obj.status",
	"obj.ttl",
	"quic.cc.cwnd",
	"quic.cc.ssthresh",
	"quic.num_bytes.received",
	"quic.num_bytes.sent",
	"quic.num_packets.ack_received",
	"quic.num_packets.decryption_failed",
	"quic.num_packets.late_acked",
	"quic.num_packets.lost",
	"quic.num_packets.received",
	"quic.num_packets.sent",
	"quic.rtt.latest",
	"quic.rtt.minimum",
	"quic.rtt.smoothed",
	"quic.rtt.variance",
	"req.backend",
	"req.backend.healthy",
	"req.backend.ip",
	"req.backend.is_cluster",
	"req.backend.is_origin",
	"req.backend.is_shield",
	"req.backend.name",
	"req.backend.port",
	"req.body",
	"req.body.base64",
	"req.body_bytes_read",
	"req.bytes_read",
	"req.customer_id",
	"req.digest",
	"req.digest.ratio",
	"req.enable_range_on_pass",
	"req.enable_segmented_caching",
	"req.esi",
	"req.esi_level",
	"req.grace",
	"req.hash",
	"req.hash_always_miss",
	"req.hash_ignore_busy",
	"req.header_bytes_read",
	"req.is_background_fetch",
	"req.is_clustering",
	"req.is_esi_subreq",
	"req.is_ipv6",
	"req.is_purge",
	"req.is_ssl",
	"req.max_stale_if_error",
	"req.max_stale_while_revalidate",
	"req.method",
	"req.postbody",
	"req.proto",
	"req.protocol",
	"req.request",
	"req.restarts",
	"req.service_id",
	"req.topurl",
	"req.url",
	"req.url.basename",
	"req.url.dirname",
	"req.url.ext",
	"req.url.path",
	"req.url.qs",
	"req.vcl",
	"req.vcl.generation",
	"req.vcl.md5",
	"req.vcl.version",
	"req.xid",
	"resp.body_bytes_written",
	"resp.bytes_written",
	"resp.completed",
	"resp.header_bytes_written",
	"resp.is_locally_generated",
	"resp.proto",
	"resp.response",
	"resp.stale",
	"resp.stale.is_error",
	"resp.stale.is_revalidating",
	"resp.status",
	"segmented_caching.autopurged",
	"segmented_caching.block_number",
	"segmented_caching.block_size",
	"segmented_caching.cancelled",
	"segmented_caching.client_req.is_open_ended",
	"segmented_caching.client_req.is_range",
	"segmented_caching.client_req.range_high",
	"segmented_caching.client_req.range_low",
	"segmented_caching.completed",
	"segmented_caching.error",
	"segmented_caching.failed",
	"segmented_caching.is_inner_req",
	"segmented_caching.is_outer_req",
	"segmented_caching.obj.complete_length",
	"segmented_caching.rounded_req.range_high",
	"segmented_caching.rounded_req.range_low",
	"segmented_caching.total_blocks",
	"server.billing_region",
	"server.datacenter",
	"server.hostname",
	"server.identity",
	"server.ip",
	"server.pop",
	"server.port",
	"server.region",
	"stale.exists",
	"time.elapsed",
	"time.elapsed.msec",
	"time.elapsed.msec_frac",
	"time.elapsed.sec",
	"time.elapsed.usec",
	"time.elapsed.usec_frac",
	"time.end",
	"time.end.msec",
	"time.end.msec_frac",
	"time.end.sec",
	"time.end.usec",
	"time.end.usec_frac",
	"time.start",
	"time.start.msec",
	"time.start.msec_frac",
	"time.start.sec",
	"time.start.usec",
	"time.start.usec_frac",
	"time.to_first_byte",
	"tls.client.certificate.dn",
	"tls.client.certificate.is_cert_bad",
	"tls.client.certificate.is_cert_expired",
	"tls.client.certificate.is_cert_missing",
	"tls.client.certificate.is_cert_revoked",
	"tls.client.certificate.is_cert_unknown",
	"tls.client.certificate.is_unknown_ca",
	"tls.client.certificate.is_verified",
	"tls.client.certificate.issuer_dn",
	"tls.client.certificate.not_after",
	"tls.client.certificate.not_before",
	"tls.client.certificate.raw_certificate_b64",
	"tls.client.certificate.serial_number",
	"tls.client.cipher",
	"tls.client.ciphers_list",
	"tls.client.ciphers_list_sha",
	"tls.client.ciphers_list_txt",
	"tls.client.ciphers_sha",
	"tls.client.handshake_sent_bytes",
	"tls.client.iana_chosen_cipher_id",
	"tls.client.ja3_md5",
	"tls.client.protocol",
	"tls.client.servername",
	"tls.client.tlsexts_list",
	"tls.client.tlsexts_list_sha",
	"tls.client.tlsexts_list_txt",
	"tls.client.tlsexts_sha",
	"transport.bw_estimate",
	"transport.type",
	"waf.anomaly_score",
	"waf.blocked",
	"waf.counter",
	"waf.executed",
	"waf.failures",
	"waf.http_violation_score",
	"waf.inbound_anomaly_score",
	"waf.lfi_score",
	"waf.logdata",
	"waf.logged",
	"waf.message",
	"waf.passed",
	"waf.php_injection_score",
	"waf.rce_score",
	"waf.rfi_score",
	"waf.rule_id",
	"waf.session_fixation_score",
	"waf.severity",
	"waf.sql_injection_score",
	"waf.xss_score",
	"workspace.bytes_free",
	"workspace.bytes_total",
	"workspace.overflowed",
}

var Setable = []string{
	"bereq.between_bytes_timeout",
	"bereq.connect_timeout",
	"bereq.first_byte_timeout",
	"bereq.method",
	"bereq.request",
	"bereq.url",
	"beresp.brotli",
	"beresp.cacheable",
	"beresp.do_esi",
	"beresp.do_stream",
	"beresp.grace",
	"beresp.gzip",
	"beresp.hipaa",
	"beresp.pci",
	"beresp.response",
	"beresp.saintmode",
	"beresp.stale_if_error",
	"beresp.stale_while_revalidate",
	"beresp.status",
	"beresp.ttl",
	"client.geo.ip_override",
	"client.identity",
	"client.socket.congestion_algorithm",
	"client.socket.cwnd",
	"client.socket.pace",
	"esi.allow_inside_cdata",
	"geoip.ip_override",
	"geoip.use_x_forwarded_for",
	"obj.grace",
	"obj.response",
	"obj.status",
	"obj.ttl",
	"req.backend",
	"req.enable_range_on_pass",
	"req.enable_segmented_caching",
	"req.esi",
	"req.grace",
	"req.hash",
	"req.hash_always_miss",
	"req.hash_ignore_busy",
	"req.max_stale_if_error",
	"req.max_stale_while_revalidate",
	"req.method",
	"req.request",
	"req.url",
	"resp.response",
	"resp.stale",
	"resp.stale.is_error",
	"resp.stale.is_revalidating",
	"resp.status",
	"segmented_caching.block_size",
	"waf.anomaly_score",
	"waf.blocked",
	"waf.counter",
	"waf.executed",
	"waf.http_violation_score",
	"waf.inbound_anomaly_score",
	"waf.lfi_score",
	"waf.logdata",
	"waf.logged",
	"waf.message",
	"waf.passed",
	"waf.rfi_score",
	"waf.rule_id",
	"waf.session_fixation_score",
	"waf.severity",
	"waf.xss_score",
}

var Unsetable = []string{
	"fastly.error",
}
