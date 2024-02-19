domain_names = ["launchexample.nttdata.com"]
a_records = {
  "first" = {
    ttl     = 300
    records = ["20.0.0.1"]
    tags    = { "test" = "false" }
  },
  "second" = {
    ttl     = 300
    records = ["20.0.0.2", "20.0.0.3"]
    tags    = { "test" = "true" }
  }
}

cname_records = {
  "first" = {
    ttl    = 100
    record = "www"
    tags   = { "test" = "www" }
  },
  "second" = {
    ttl    = 100
    record = "www.test.com"
    tags   = { "test" = "www.test.com" }
  }
}

ns_records = {
  "first" = {
    ttl     = 100
    records = ["ns1-32.azure-dns.com."]
    tags    = { "test" = "ns1-32.azure-dns.com." }
  },
  "second" = {
    ttl     = 100
    records = ["ns3-32.azure-dns.org."]
    tags    = { "test" = "ns3-32.azure-dns.org." }
  }
}

txt_records = {
  "first" = {
    ttl     = 100
    records = ["random-text"]
    tags    = { "test" = "random-text" }
  },
  "second" = {
    ttl     = 100
    records = ["random-text-2", "random-text-3"]
    tags    = { "test" = "random-tag" }
  }
}
