Soal Nomor 1.

- Membuat docker file untuk build dan program backend dari exercise 1
  Menggunakan base image go version 1.23.2 untuk menyesuaikan versi go di go.mod

- Mendownload dependency dengan menggunakan syntax go mod download dan go mod tidy

- Melakukan build aplikasi golang, dimana main.go terletak di folder cmd
- Melakukan expose port 8080 menyesuaikan konfigurasi di main.go

Cara saya melakukan build image dan running container

- docker build -t exercise01 .
- docker run -d --name exercise01 -p 8080:8080 exercise01
- cek di localhost:8080/api

Soal nomor 2.

- Melakukan analisa node version, ternyata di package-lock.json terdapat requirement menggunakan versi node >= 14
- Melakukan analisa package.json, sudah ada command build dan preview untuk deployment production
- Melakukan analisa vite.config.js, menambahkan host: "0.0.0.0" dan port: 5173, untuk deployment
- membuat docker file dengan base node 18
- melakukan install node_module dengan npm install, lalu melakukan build npm run build
- Menjalankan aplikasi dengan cmd npm run preview

Cara saya melakukan build image, running container, dan menyimpan output dari buld ke local

- docker build -t exercise02 .
- docker run -d --name exercise02 -p 5173:5173 exercise02
- docker cp exercise02:/app/dist ./dist

Soal nomor 3.

- Membuat konfigurasi default.conf dari nginx dengan tujuan mereverse proxy /api ke container exercise01 port 8080 (backend service), dan yang lain atau "/" diarahkan ke folder hasil build yang dipindahkan ke /usr/share/nginx/html

- Membuat docker-compose.yml untuk menjalankan semua service

Cara saya melakukan running docker-compose :

- docker-compose up --build
