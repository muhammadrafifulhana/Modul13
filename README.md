## **Soal Latihan: 1**  
### **Program Pengurutan dan Pengecekan Jarak Data**

**Penjelasan Program:**  
1. **Tujuan**: Program ini digunakan untuk mengurutkan data yang dimasukkan menggunakan algoritma **Insertion Sort** dan memeriksa apakah data yang telah diurutkan memiliki jarak yang tetap antar elemen.
2. **Alur Program**:  
   - Program menerima input angka satu per satu hingga angka negatif dimasukkan, yang menandakan akhir input.
   - Data yang telah dimasukkan akan diurutkan menggunakan algoritma **Insertion Sort**.
   - Program kemudian memeriksa apakah jarak antar elemen dalam data yang terurut tetap sama. Jika tidak, program akan menginformasikan bahwa jaraknya tidak tetap.
3. **Output**:  

**Contoh Input:**
```
31 13 25 43 1 7 19 37 -5
```

**Contoh Output:**
```
1 7 13 19 25 31 37 43
Data berjarak 6
```

---

## **Soal Latihan: 2**  
### **Program Manajemen Pustaka Buku**

**Penjelasan Program:**  
1. **Tujuan**: Program ini digunakan untuk mengelola pustaka buku. Pengguna dapat menambahkan buku, mencetak buku dengan rating tertinggi, mengurutkan buku berdasarkan rating, mencetak 5 buku terbaru, dan mencari buku berdasarkan rating.
2. **Fungsi Utama**:  
   - **`DaftarkanBuku`**: Fungsi untuk mendaftarkan buku baru ke dalam pustaka.
   - **`CetakTerfavorit`**: Fungsi untuk mencetak buku dengan rating tertinggi.
   - **`UrutBuku`**: Fungsi untuk mengurutkan buku berdasarkan rating secara menurun.
   - **`Cetak5Terbaru`**: Fungsi untuk mencetak 5 buku terbaru berdasarkan urutan yang telah diurutkan.
   - **`CariBuku`**: Fungsi untuk mencari buku dengan rating tertentu menggunakan pencarian biner.
3. **Alur Program**:  
   - Pengguna memasukkan data buku (ID, judul, penulis, penerbit, jumlah eksemplar, tahun terbit, dan rating).
   - Program akan mencetak buku dengan rating tertinggi, mengurutkan buku berdasarkan rating, mencetak 5 buku terbaru, dan mencari buku dengan rating tertentu menggunakan metode pencarian biner.
4. **Output**:  

**Contoh Input:**
```
3
B001
Buku A
Penulis A
Penerbit A
10
2020
4
B002
Buku B
Penulis B
Penerbit B
5
2021
5
B003
Buku C
Penulis C
Penerbit C
2
2019
3
```

**Contoh Output:**
```
Buku B Penulis B Penerbit B 2021
Buku B
Buku A
Buku C
Buku B
Buku A
```
