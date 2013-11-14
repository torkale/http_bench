request = require 'request'
http = require 'http'

app = http.createServer((req, res) ->
  req.pipe(request('http://localhost:4000'))
  res.end()
).listen(2000)
