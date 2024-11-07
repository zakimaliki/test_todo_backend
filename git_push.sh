#!/bin/bash

# Menambahkan semua perubahan ke staging area
git add .

# Memeriksa apakah ada argumen untuk pesan commit
if [ -z "$1" ]; then
  echo "Pesan commit tidak diberikan. Menggunakan pesan default."
  git commit -m "Update"
else
  git commit -m "$1"
fi

# Melakukan push ke branch saat ini
git push origin master