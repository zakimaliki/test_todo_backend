#!/bin/bash

# Pull docker image oracle-xe latest version
docker pull gvenzl/oracle-xe:latest

# Membuat container oracle-xe
docker run -d \
  --name oracle-xe \
  -p 1521:1521 \
  -e ORACLE_PASSWORD=yourpassword \
  gvenzl/oracle-xe

# Menunggu database siap digunakan (sekitar 2-3 menit)
echo "Menunggu database siap..."
sleep 180

# Cek status container
docker ps | grep oracle-xe

echo "Oracle XE sudah terinstall dan berjalan"
echo "Koneksi database:"
echo "Hostname: localhost"
echo "Port: 1521"
echo "Service Name: XEPDB1"