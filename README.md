# go-http-logger
Simple server that logs every http message it recieves.

Conveniently this also logs in quoted text so you can fire it at a server using netcat without any escaping issues from your terminal:

```bash
echo="...." ; cat $echo | nc domain port
```

This was written to circumvent issues I was having enabling loggin in a python project using requests/urllib3/httplib.
