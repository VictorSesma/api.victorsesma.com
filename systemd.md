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