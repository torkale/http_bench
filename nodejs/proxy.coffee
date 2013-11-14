#!/usr/bin/env coffee

httpProxy = require 'http-proxy'
http = require 'http'

app = httpProxy.createServer((req, res, proxy) ->
  buffer = httpProxy.buffer(req)
  proxy.proxyRequest(req, res, {
    port: 4000
    host: 'localhost'
    buffer: buffer
  })
).listen(2000)
