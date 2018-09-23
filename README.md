# Welcome to [VictorSesma.com](https://victorsesma.com) Go Repository

This is the API end point for the [VictorSesma.com](https://victorsesma.com) [React App](https://github.com/leviatan89/victorsesma.com)

## How to Deploy

1. Install Golang (for example: `snap install --classic go`)
2. Clone the repository (for example: `git clone git@github.com:leviatan89/api.victorsesma.com.git`)
3. Make a copy and configure conf.json
4. Run the migrate comand. Something like:
´migrate -source file:///home/user/go/src/github.com/leviatan89/api.victorsesma.com/migrations -database mysql://"user:password@tcp(dbserver.example:3306)/dbname" up 2`
5. Run `go install`
6. Create the following Unit file:

```
[Unit]
Description= instance to serve apiVictorSesma
After=network.target

[Service]
Environment=GOPATH=/home/leviatan89/go
User=root
Group=www-data

ExecStart=/home/leviatan89/go/bin/api.victorsesma.com

[Install]
WantedBy=multi-user.target

** Note that the msyql server should be running **
```
7. sudo systemctl enable apiVictorSesma