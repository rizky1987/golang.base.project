## - Installations
- clone repository dan rename main folder menjadi **example** atau sesuaikan dengan nama project jika digunakan untuk suatu project.
- branch master adalah branch referensi untuk experiment jadi sebelum merge atau push ke branch ini harus melalui **persetujuan sidang isbat**.
- gunakan branch **baseProject** untuk referensi awal dalam membuat project baru dengan default database **SQL Server**. **baseProjectMongoDB** untuk default database **mongoDB** dan seterunsnya. 
- dalam branch ini ada folder documentations. isi nya adalah catetan yang kira2 diperlukan dalam project ini.
- Jalankan command line berikut untuk install swaggo, jika sebelumnya belum pernah menginstall swaggo.
```sh
go get -u github.com/swaggo/swag/cmd/swag

go install github.com/swaggo/swag/cmd/swag@latest
```

## - Struktur Folder
## CommonHelpers
   Di dalam folder ini akan berisi helper - helper yang akan membantu dalam mempermudah proses development yang bersifat umum. Contoh : helper untuk mengubah data **datetime** menjadi **string** dan fungsi - fungsi lainnya yang diperlukan. Untuk setiap fungsi akan dikelompokkan di masing - masing file berdasarkan inputan. Contoh jika ada suatu fungsi untuk mengubah **integer** menjadi **string** maka fungsi tersebut harus dimasukan ke file **integer.go**

## Config
 Folder ini akan berisi data - data yg bersifat untuk configurasi saja masing - masing sudah dikelompokkan berdasarkan nama foldernya.
 
 - boot 
   folder ini berisi hanya 1 file bootstrap.go. fungsi - fungsi yang mungkin lebih banyak terpakai disini adalah untuk mendaftarkan routing baru yg dibuat. untuk koneksi ke database mana yg akan kita gunakan. 
- env
  berisi file configurasi untuk masing - masing environment yang akan kita gunakan. Jika kita memerlukan suatu data atau variable yang sifatnya global bisi diisi disini. Contoh **BaseURL untuk SAP, Database Configuration Value**
- middleware
  Sesuai namanya berisi fungsi - fungsi untuk menghandle middleware. Folder ini akan sangan jarang sekali diedit dalam proses development
- routes
  Folder ini berisi file - file yang digunakan untuk menuliskan routing yang akan kita pakai diaplikasi. masing - masing file akan berisi satu main endpoint. Contoh : **cart_route.go** maka isinya harus berisi semua endpoint dan swagger anotation untuk **cart** tidak boleh yang lainnya. 
## Databases
 - connection 
   berisi file - file untuk mengkoneksikan ke masing - masing database.
 - entities
 -berisi file - file entity. Entity sebagai representasi untuk masing - masing tabel di dalam data base. file - file tersebut harus dikelompokkan berdasarkan folder dari database yg digunakan. Contoh : Jika ada foldel **sql** dalam folder **entities** ini, maka isi dalam folder **sql** ini adalah entity dari database **sql** saja. Jika misal kita koneksi ke 2 database dalam 1 aplikasi maka kita harus membuat folder **mongo** untuk menyimpan file - file entity yg digunakan khusus untuk **mongoDB**
## Docs
 folder ini bersisi file - file yang telah digenerate oleh swagger generator. **KITA TIDAK MENGEDIT FILE INI SECARA LANGSUNG, MELAINKAN HARUS GENERATE!! INGA INGA TIIIIIIING**
## Documentations
 Berisi file - file yang kira2 diperlukan dalam menunjang keberhasilan project. Misal file untuk cara melakukan deployment ke staging RL dan lainnya.
## Http
  Didalam folder inilah fungsi - fingsi utama akan di kodingkan. Kemungkinan besar setiap developer akan banyak berinteraksi atau melakukan perubahan atau penambahan kodingan di dalam folder ini.

- helpers 
  fungsi - fungsi untuk mempermudah menghasilkan return api. **Setiap perubahan di folder ini harus melalui sidang isbat** karena akan jika ada peubahan di folder ini maka **baseProject** yang sudah kita buat harus di update lagi.
- interfaces
 berisi masing - masing interfaces untuk setiap satu entity. misal cart_interface.go yang ada di folder ini harus bersisi interface - interface yang memeng berhubungan dengan **cart** saja tidak boleh ada yang lain.
- repositories
  implementor untuk setiap interface yang sudah dibuat file nya di folder interface.
- requests
 berisi file - file untuk mendefinisikan setiap inputan dari user. Contoh : dalam melakukan pencarian data cart ada inputan limit, offest dan nama cart. inputan itu harus disimpan di folder ini dengan nama **cart_request.go**
- responses
 Berisi file - file response untuk masing - masing endpoint. satu file untuk satu entity.
- services
 berisi file handler atau service dimana disinilah kita akan menuliskan logic pemograman yang kita buat. 
## Logs
 berisi file .txt. File ini digunakan untuk trace error d lokal. Untuk mempermudah developer jika terjadi kesalahan baik dalam kodingan ataupun inputan. cara membaca  baris error disini ada di dalam folder documentation.
## Utils
 berisi file - file atau folder yang digunakan untuk settingan jika kita menggunakan jasa pihak ke tiga. Misal jika kita menggunakan application insight, maka kodingan untuk bisa connect ke application insight itu harus disimpan disini. karena biasanya kodingan untuk pihak ke tidak akan sangat jarang sekali diubah jika sudah jalan.