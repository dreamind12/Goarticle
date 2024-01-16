## Langkah Langkah Mengclone Project 

## Clone Backend
- Salin repository dengan menggunakan perintah git clone https://github.com/dreamind12/Goarticle.git
- Instal modul-modul yang diperlukan dari file go.mod dengan menjalankan perintah go mod download.
- Buka file main.go, dan hapus tanda "//" pada baris // database.CreateArticleTable() untuk membuat tabel artikel pada database.
- Sesuaikan atau buatlah database pada file database/database.go. Atur variabel dsn sesuai dengan konfigurasi database Anda. Sebagai contoh, dsn := "root:@tcp(localhost:3306)/article?charset=utf8mb4&parseTime=True&loc=Local" (saya menggunakan MySQL).
- Terakhir, jalankan perintah go run main.go untuk menjalankan aplikasi.

## Documenter API Dari Postman
- https://documenter.getpostman.com/view/25017003/2s9YsQ7V51#a4a75729-624a-4fe0-8667-2fb435f08917

## Mock Server Postman
- https://f97927d8-31b7-4c2b-8057-60f911e1b098.mock.pstmn.io
