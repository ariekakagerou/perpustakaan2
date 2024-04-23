package main

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// Member adalah struktur data untuk anggota perpustakaan
type anggota struct {
    ID            int    `json:"id"`
    Nama          string `json:"nama"`
    JenisKelamin  string `json:"jenis_kelamin"`
    Alamat        string `json:"alamat"`
    NoTelepon     string `json:"no_telepon"`
}

// Book adalah struktur data untuk buku perpustakaan
type buku struct {
    ID           int    `json:"id"`
    Judul        string `json:"judul"`
    TahunTerbit  int    `json:"tahun_terbit"`
    Jumlah       int    `json:"jumlah"`
    ISBN         string `json:"isbn"`
    PengarangID  int    `json:"pengarang_id"`
    PenerbitID   int    `json:"penerbit_id"`
    RakKodeRak   string `json:"rak_kode_rak"`
}

// Borrower adalah struktur data untuk peminjam
type peminjaman struct {
    ID             int    `json:"id"`
    TanggalPinjam  string `json:"tanggal_pinjam"`
    TanggalKembali string `json:"tanggal_kembali"`
    AnggotaID      int    `json:"anggota_id"`
    PetugasID      int    `json:"petugas_id"`
}

// BorrowerDetail adalah struktur data untuk detail peminjaman
type peminjamanDetail struct {
    PeminjamanID int `json:"peminjaman_id"`
    BukuID       int `json:"buku_id"`
}

// Publisher adalah struktur data untuk penerbit
type penerbit struct {
    ID         int    `json:"id"`
    Nama       string `json:"nama"`
    Alamat     string `json:"alamat"`
    NoTelepon  string `json:"no_telepon"`
}

// Return adalah struktur data untuk pengembalian
type pengembalian struct {
    ID                   int    `json:"id"`
    TanggalPengembalian  string `json:"tanggal_pengembalian"`
    Denda                int    `json:"denda"`
    PeminjamanID         int    `json:"peminjaman_id"`
    AnggotaID            int    `json:"anggota_id"`
    PetugasID            int    `json:"petugas_id"`
}

// ReturnDetail adalah struktur data untuk detail pengembalian
type returnDetail struct {
    PengembalianID int `json:"pengembalian_id"`
    BukuID         int `json:"buku_id"`
}

// Librarian adalah struktur data untuk petugas
type petugas struct {
    ID         int    `json:"id"`
    Username   string `json:"username"`
    Password   string `json:"password"`
    Nama       string `json:"nama"`
    NoTelepon  string `json:"no_telepon"`
    Alamat     string `json:"alamat"`
}

// Shelf adalah struktur data untuk rak
type rak struct {
    KodeRak int    `json:"kode_rak"`
    Lokasi  string `json:"lokasi"`
}

// Author adalah struktur data untuk pengarang
type pengarang struct {
    ID        int    `json:"id"`
    Nama      string `json:"nama"`
    Alamat    string `json:"alamat"`
    NoTelepon string `json:"no_telepon"`
}

func setupDatabase() {
    // Database connection
    var err error
    db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/db_perpustakaan")
    if err != nil {
        log.Fatal(err)
    }
}
