whatsmyip.go
============

A dumb HTTP server that waits for HTTP requests and outputs the contents of the X-Real-IP header (provided by Nginx or some other reverse proxy). Sample Nginx configuration is provided below.  
  
*NOTE: The default configuration of this app makes it useless when accessed directly, unless the client provides the necessary header to the app, which could be very dangerous. For this reason, I do NOT recommend using this application unless it's behind some proxy. If you really must run this app without a proxy in front, please modify the app to output `req.RemoteAddr` instead and not accept any header input. The format of `req.RemoteAddr` is `IP:Port`, so you will need to discard the `:Port` side of that string to produce the same result.*

Usage
-----
Usage is really just as simple as starting the application:
```
go run whatsmyip.go
```
However, most people will want to start this application once and have it continue running, so it's recommended to run using supervisord (example config below), upstart, in screen, with nohup, etc.

HTTP Daemon Configuration
-------------------------
There are two things in the app you may wish to customize for your own purposes. The first is the bind address/port. You can change the default `127.0.0.1:8999` to suit your needs, but it's not recommended to bind to all interfaces, unless the server it runs on is not accessible to the internet.  
  
The second thing you may want to customize is the header the remote IP is pulled from. For example, some admins prefer to add `X-Forwarded-For` headers instead of the default `X-Real-IP`, and of course, you may change this to whatever you like.

Example Supervisord config file
-------------------------------
```
[program:whatsmyip]
command=/usr/local/bin/whatsmyip
autorestart=true
user=www-data
```
In this example, I've installed the `whatsmyip` binary to `/usr/local/bin`, and I'm starting it as the `www-data` user, which is the default web server user on most Debian-like distributions. You can change this user to whatever you like. I've also enabled the `autorestart=true` option, as I'd like it to restart even on a proper exit - I always want the little daemon running.

Example Nginx configuration
---------------------------
```
    location /ip {
        proxy_set_header X-Real-IP $proxy_add_x_forwarded_for;
        proxy_redirect off;
        proxy_http_version 1.1;

        proxy_pass http://127.0.0.1:8999;
    }
```
In my example, I've used `/ip/` as the location to proxy to the application. Of course, you're free to use whatever URI you prefer, or you can even forward an entire subdomain by creating a new `server` block with `location /` instead of the above.

Performance
-----------
You'll generally get the best performance by compiling this script into a binary first, which is as simple as running: `go build whatsmyip.go`  
You can then run `./whatsmyip` or simply `whatsmyip` if you copy the binary to `~/bin` or `/usr/bin`.
