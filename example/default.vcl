include "include";

backend example_com {
  .connect_timeout = 1s;
  .dynamic = true;
  .port = "443";
  .host = "example.com";
  .first_byte_timeout = 20s;
  .max_connections = 500;
  .between_bytes_timeout = 20s;
  .share_key = "xei5lohleex3Joh5ie5uy7du";
  .ssl = true;
  .ssl_sni_hostname = "example.com";
  .ssl_cert_hostname = "example.com";
  .ssl_check_cert = always;
  .min_tls_version = "1.2";
  .max_tls_version = "1.2";
  .bypass_local_route_table = false;
  .probe = {
    .request = "GET / HTTP/1.1" "Host: example.com" "Connection: close";
    .dummy = true;
    .threshold = 1;
    .window = 2;
    .timeout = 5s;
    .initial = 1;
    .expected_response = 200;
    .interval = 10s;
  }
}

acl example_acl {
  "10.0.0.0"/32;
}

table example_table STRING {
  "foo": "bar",
}

director example_director random {
  { .backend =  example_com; .weight = 10; }
}

sub vcl_recv {
  #Fastly recv
  set req.backend = example_com;
  set req.http.Foo = {" foo bar baz "};
  if (req.http.Foo ~ "foo") {
    call custom_logger;
  } else {
    set req.http.Bar = "bar";
  }
  return (pass);
}

sub vcl_deliver {

  #Fastly deliver
  set resp.http.X-Custom-Header = "Custom Header";
  call custom_logger;
  return (deliver);
}

sub vcl_fetch {

  #Fastly fetch
  call custom_logger;
  return(deliver);
}
