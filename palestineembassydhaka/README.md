# Palestine Embassy Dhaka Official Website
Welcome to Embassy of the State of Palestine, Dhaka

# Testing web hosting
> http://91.205.173.170:8085

# Installation Guide
> sudo firewall-cmd --permanent --add-port=8085/tcp

> sudo firewall-cmd --reload 

> sudo firewall-cmd --list-ports

> mkdir /home/mastererp/palestineembassybd.com

> sudo nano /lib/systemd/system/palestine.service

```
[Unit]
Description=Palestine embassy dhaka website
After=network.target

[Service]
Type=simple
Restart=always
RestartSec=5s
WorkingDirectory=/home/mastererp/palestineembassybd.com
ExecStart=/home/mastererp/palestineembassybd.com/palestine

[Install]
WantedBy=multi-user.target
```

> chmod 664 /lib/systemd/system/palestine.service

> systemctl daemon-reload

> sudo chmod +x /home/mastererp/palestineembassybd.com/palestine

> sudo systemctl daemon-reload

# Virtual Hosting Creation

> cp /etc/nginx/conf.d/mostain.net.conf /etc/nginx/conf.d/palestineembassybd.com.conf

> nano /etc/nginx/conf.d/palestineembassybd.com.conf

```
server {
        listen 80;
        server_name palestineembassybd.com www.palestineembassybd.com;
        #expires 1d;

        location / {

          #proxy_cache my_cache;
          #proxy_buffering        on;
          #proxy_cache_valid      200  1d;
          #proxy_cache_use_stale  error timeout invalid_header updating
          #                         http_500 http_502 http_503 http_504;

          proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
          proxy_set_header X-Forwarded-Proto $scheme;
          proxy_set_header X-Real-IP $remote_addr;
          proxy_set_header Host $host;
          proxy_pass  http://localhost:8085;
        }

}
```
> sudo nginx -t

> sudo systemctl restart nginx or service nginx reload

> sudo systemctl status nginx or service nginx status

> nano /etc/hosts
add the following line to the end of the file

```
91.205.173.170 palestineembassybd.com

```
[Tutorial](https://www.digitalocean.com/community/tutorialshow-to-secure-nginx-with-let-s-encrypt-on-centos-8)

# Obtaining a Letsencrypt SSL Certificate

> sudo certbot --nginx -d palestineembassybd.com -d www.palestineembassybd.com
Above command will automatically update the nginx config file /etc/nginx/conf.d/palestineembassybd.com.conf

[SSL Cert Status](https://www.ssllabs.com/ssltest/analyze.html?d=your_domain)

> sudo certbot renew --dry-run
