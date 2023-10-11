//@scope: recv,deliver,log
sub custom_logger {
  declare local var.IsMatched BOOL;
  if (client.ip ~ example_acl) {
    set var.IsMatched = true;
  }
  else if (std.str2ip("192.168.0.1", "10.0.0.0") ~ example_acl) {
    set var.IsMatched = true;
  }
  log {" syslog example_service fastly-logs :: "} {"custom logger executed on "} var.IsMatched;
}
