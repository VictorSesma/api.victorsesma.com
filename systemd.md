## systemd service


`sudo vim /etc/systemd/system/apiVictorSesma.service`

```
[Unit]
Description= instance to serve apiVictorSesma
After=network.target
 
[Service]
User=root
Group=www-data
 
ExecStart=/path/to/binary
 
[Install]
WantedBy=multi-user.target
```


Inspired by [this article](https://kenyaappexperts.com/blog/how-to-deploy-golang-to-production-step-by-step/)