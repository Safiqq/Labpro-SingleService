# Single-Service API
Nama: Syafiq Ziyadul Arifin
NIM: 18221048

## How to Use
1. `sudo docker build -t single-service .`
2. `sudo docker run -p 3002:3002 -it --rm --name single-service single-service`

Additional:
* sudo docker rm --force <id>
* sudo docker logs <id>

## Design Pattern
1. Model-View-Controller (MVC)
    
    Design pattern ini memisahkan peran dan tugas di dalam aplikasi agar lebih rapi dan mudah dimengerti. Misalnya, bagian yang mengurus data, bagian yang menampilkan tampilan, dan bagian yang mengatur logika penggunaan data dan tampilan.
2. Dependency Injection

    Dengan menggunakan design pattern ini, kita dapat membuat kode menjadi lebih fleksibel dan mudah diuji. Ketergantungan antara bagian-bagian di dalam aplikasi diurutkan dengan baik, sehingga jika ada perubahan, kita tidak perlu mengubah banyak bagian lain.
3. Singleton

    Design pattern ini memastikan hanya ada satu contoh (instance) dari suatu objek. Jadi, ketika kita butuh objek tersebut, kita tahu pasti bahwa kita akan menggunakan objek yang sama setiap kali, sehingga tidak terjadi pemborosan sumber daya.
4. Middleware

    Design pattern ini digunakan untuk memisahkan peran dari berbagai bagian dalam aplikasi. Misalnya, bagian yang menangani izin, bagian yang mencatat aktivitas, atau bagian yang mengurus izin tampilan. Dengan memisahkan tugas ini, kita bisa menambahkan fitur baru dengan mudah tanpa mengganggu bagian lain.
5. Factory

    Design pattern ini digunakan untuk membuat objek dengan cara yang fleksibel. Jadi, kita bisa membuat objek dengan mudah tanpa perlu tahu detail cara pembuatannya.
6. Decorator

    Design pattern ini digunakan untuk menambahkan fitur tambahan ke objek secara dinamis. Kita bisa menambahkan fitur baru ke objek tanpa harus mengubah struktur kelas yang sudah ada.

## Tech Stacks
1. Golang v1.20
2. Gin v1.9.1
3. GORM v1.25.2
4. PostgreSQL

## Endpoints
* **POST**
    * **/login**: Endpoint untuk melakukan proses login.
    * **/barang**: Endpoint untuk membuat data barang baru (Membutuhkan otentikasi).
    * **/perusahaan**: Endpoint untuk membuat data perusahaan baru (Membutuhkan otentikasi).
* **GET**
    * **/self**: Endpoint untuk mendapatkan informasi diri sendiri (Membutuhkan otentikasi).
    * **/barang**: Endpoint untuk mendapatkan daftar barang.
    * **/barang/:id**: Endpoint untuk mendapatkan detail barang berdasarkan ID.
    * **/perusahaan**: Endpoint untuk mendapatkan daftar perusahaan.
    * **/perusahaan/:id**: Endpoint untuk mendapatkan detail perusahaan berdasarkan ID.
* **DELETE**
    * **/barang/:id**: Endpoint untuk menghapus data barang berdasarkan ID (Membutuhkan otentikasi).
    * **/perusahaan/:id**: Endpoint untuk menghapus data perusahaan berdasarkan ID (Membutuhkan otentikasi).
* **UPDATE**
    * **/barang/:id**: Endpoint untuk memperbarui data barang berdasarkan ID (Membutuhkan otentikasi).
    * **/perusahaan/:id**: Endpoint untuk memperbarui data perusahaan berdasarkan ID (Membutuhkan otentikasi).