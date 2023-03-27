cp cd-server.linux /usr/local/bin

mkdir -p /etc/cd-server

cat << EOF > /etc/cd-server/conf.env
CD_SECRET=123
CD_USER_DATA_PATH=$(pwd)/user-data
EOF

cat << EOF > /etc/systemd/system/cd-server.service
[Unit]
Description=Service for handle CD builds
After=multi-user.target

[Service]
EnvironmentFile=/etc/cd-server/conf.env
ExecStart=/usr/local/bin/cd-server.linux
ExecReload=/usr/local/bin/cd-server.linux
Type=simple
Restart=always


[Install]
WantedBy=default.target
RequiredBy=network.target
EOF

cat << EOF > entry.sh
echo 1
mkdir test
EOF

curl --location 'http://81.163.30.137:8080/secret' \
--form myFiles=@"$(echo $GITHUB_WORKSPACE)/docker-compose.yml" \
--form myFiles=@"$(echo $GITHUB_WORKSPACE)/package.json" \
--form 'secret="123"' \
--form entry=@"$(echo $GITHUB_WORKSPACE)/entry.sh"