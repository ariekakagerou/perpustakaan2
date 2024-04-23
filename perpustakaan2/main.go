package main

import (
	"database/sql"
	"log"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

var db *sql.DB

type anggota struct {
	ID            int    `json:"id"`
	Nama          string `json:"nama"`
	Jenis_kelamin string `json:"jenis_kelamin"`
	Alamat        string `json:"alamat"`
	No_telepon    string `json:"no_telepon"`
}

type buku struct {
	ID           int    `json:"id"`
	Judul        string `json:"judul"`
	Tahun_terbit int    `json:"tahun_terbit"`
	Jumlah       int    `json:"jumlah"`
	ISBN         string `json:"isbn"`
	Pengarang_id int    `json:"pengarang_id"`
	Penerbit_id  int    `json:"penerbit_id"`
	Rak_Kode_rak int `json:"rak_kode_rak"`
}

type petugas struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Nama       string `json:"nama"`
	No_telepon string `json:"no_telepon"`
	Alamat     string `json:"alamat"`
}

type peminjaman struct {
    ID              int    `json:"id"`
    Tanggal_pinjam  string `json:"tanggal_pinjam"`
    Tanggal_kembali string `json:"tanggal_kembali"`
    Anggota_id      int    `json:"anggota_id"`
    Petugas_id      int    `json:"petugas_id"`
}

type peminjaman_detail struct {
	Peminjaman_id string `json:"peminjaman_id"`
	Buku_id       string `json:"buku_id"`
}

type penerbit struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	No_telepon string `json:"no_telepon"`
}
type pengarang struct {
	ID         int    `json:"id"`
	Nama       string `json:"nama"`
	Alamat     string `json:"alamat"`
	No_telepon string `json:"no_telepon"`
}
type pengembalian struct {
	ID                   int    `json:"id"`
	Tanggal_pengembalian string `json:"tanggal_pengembalian"`
	Denda                int    `json:"denda"`
	Peminjaman_id        int    `json:"peminjaman_id"`
	Anggota_id           int    `json:"anggota_id"`
	Petugas_id           int    `json:"petugas_id"`
}

type pengembalian_detail struct {
	Pengembalian_id string `json:"pengembalian_id"`
	Buku_id         string `json:"buku_id"`
}

type rak struct {
	Kode_rak int    `json:"kode_rak"`
	Lokasi   string `json:"lokasi"`
}

func main() {
	var err error
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/db_perpustakaan")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Service API!")
	})

	// Handler for GET /anggota
	e.GET("/anggota", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM anggota")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var ANGGOTA []anggota
		for res.Next() {
			var m anggota
			err := res.Scan(&m.ID, &m.Nama, &m.Jenis_kelamin, &m.Alamat, &m.No_telepon)
			if err != nil {
				log.Fatal(err)
				return err
			}
			ANGGOTA = append(ANGGOTA, m)
		}

		return c.JSON(http.StatusOK, ANGGOTA)
	})

	// PUT /anggota/:id
	e.PUT("/anggota/:id", func(c echo.Context) error {
		// Get ID from URL parameter
		id := c.Param("id")

		// Bind data from request
		var m anggota
		if err := c.Bind(&m); err != nil {
			return err
		}

		// Run query to update
		sqlStatement := "UPDATE anggota SET nama = ?, jenis_kelamin = ?, alamat = ?, no_telepon = ? WHERE id = ?"
		_, err := db.Exec(sqlStatement, m.Nama, m.Jenis_kelamin, m.Alamat, m.No_telepon, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusOK, m)
	})

	// POST /anggota
	e.POST("/anggota", func(c echo.Context) error {
		// Bind data from request
		var m anggota
		if err := c.Bind(&m); err != nil {
			return err
		}

		// Run query to insert
		sqlStatement := "INSERT INTO anggota(Nama, Jenis_kelamin, Alamat, No_telepon) VALUES (?, ?, ?, ?)"
		_, err := db.Exec(sqlStatement, m.Nama, m.Jenis_kelamin, m.Alamat, m.No_telepon)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusCreated, m)
	})

	// DELETE /anggota/:id
	e.DELETE("/anggota/:id", func(c echo.Context) error {

		// Get ID from URL parameter
		id := c.Param("id")

		// Run query to delete
		sqlStatement := "DELETE FROM anggota WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", id)
    return c.String(http.StatusOK, message)
	})

	// Handler untuk route GET /buku
	e.GET("/buku", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM buku")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var books []buku
		for res.Next() {
			var b buku
			err := res.Scan(&b.ID, &b.Judul, &b.Tahun_terbit, &b.Jumlah, &b.ISBN, &b.Pengarang_id, &b.Penerbit_id, &b.Rak_Kode_rak)
			if err != nil {
				log.Fatal(err)
				return err
			}
			books = append(books, b)
		}

		return c.JSON(http.StatusOK, books)
	})

	// POST /buku
	e.POST("/buku", func(c echo.Context) error {
		var b buku
		if err := c.Bind(&b); err != nil {
			return err
		}

		sqlStatement := "INSERT INTO buku (judul, tahun_terbit, jumlah, isbn, pengarang_id, penerbit_id, rak_kode_rak) VALUES (?, ?, ?, ?, ?, ?, ?)"
		_, err := db.Exec(sqlStatement, b.Judul, b.Tahun_terbit, b.Jumlah, b.ISBN, b.Pengarang_id, b.Penerbit_id, b.Rak_Kode_rak)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusCreated, b)
	})

	//update buku
	e.PUT("/buku/:id", func(c echo.Context) error {
		// Get ID from URL parameter
		id := c.Param("id")
	
		// Bind data from request
		var b buku
		if err := c.Bind(&b); err != nil {
			return err
		}
	
		// Run query to update
		sqlStatement := "UPDATE buku SET judul = ?, tahun_terbit = ?, jumlah = ?, isbn = ?, pengarang_id = ?, penerbit_id =?, rak_kode_rak = ? WHERE id = ?"
		_, err := db.Exec(sqlStatement, b.Judul, b.Tahun_terbit, b.Jumlah, b.ISBN, b.Pengarang_id, b.Penerbit_id, b.Rak_Kode_rak, id)
		if err != nil {
			log.Fatal(err)
			return err
		}
	
		return c.JSON(http.StatusOK, b)
	})
	
	// DELETE /buku/:id
	e.DELETE("/buku/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM buku WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", id)
    return c.String(http.StatusOK, message)
	})

	// Handler untuk route GET /petugas
	e.GET("/petugas", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM petugas")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var librarians []petugas
		for res.Next() {
			var l petugas
			err := res.Scan(&l.ID, &l.Username, &l.Password, &l.Nama, &l.No_telepon, &l.Alamat)
			if err != nil {
				log.Fatal(err)
				return err
			}
			librarians = append(librarians, l)
		}

		return c.JSON(http.StatusOK, librarians)
	})

	// POST /petugas
	e.POST("/petugas", func(c echo.Context) error {
		var l petugas
		if err := c.Bind(&l); err != nil {
			return err
		}

		sqlStatement := "INSERT INTO petugas (username, password, nama, no_telepon, alamat) VALUES (?, ?, ?, ?, ?)"
		_, err := db.Exec(sqlStatement, l.Username, l.Password, l.Nama, l.No_telepon, l.Alamat)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusCreated, l)
	})

	// update petugas
	e.PUT("/petugas/:id", func(c echo.Context) error {
		// Get ID from URL parameter
		id := c.Param("id")

		// Bind data from request
		var m petugas
		if err := c.Bind(&m); err != nil {
			return err
		}

		// Run query to update
		sqlStatement := "UPDATE petugas SET username = ?, password = ?, nama = ? , no_telepon = ?, alamat = ? WHERE id = ?"
		_, err := db.Exec(sqlStatement, m.Username, m.Password, m.Nama, m.No_telepon, m.Alamat, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusOK, m)
	})
	
	// DELETE /petugas/:id
	e.DELETE("/petugas/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM petugas WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}
		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", id)
		return c.String(http.StatusOK, message)
	})

	// Handler untuk route GET /peminjaman
	e.GET("/peminjaman", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM peminjaman")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var borrowings []peminjaman
		for res.Next() {
			var p peminjaman
			err := res.Scan(&p.ID, &p.Tanggal_pinjam, &p.Tanggal_kembali, &p.Anggota_id, &p.Petugas_id)
			if err != nil {
				log.Fatal(err)
				return err
			}
			borrowings = append(borrowings, p)
		}

		return c.JSON(http.StatusOK, borrowings)
	})

	// update peminjaman
	e.PUT("/peminjaman/:id", func(c echo.Context) error {
			// Get ID from URL parameter
			id := c.Param("id")
	
			// Bind data from request
			var m peminjaman
			if err := c.Bind(&m); err != nil {
				return err
			}
	
			// Run query to update
			sqlStatement := "UPDATE peminjaman SET tanggal_pinjam = ?, tanggal_kembali = ?, anggota_id = ? , petugas_id= ? WHERE id = ?"
			_, err := db.Exec(sqlStatement, m.Tanggal_pinjam, m.Tanggal_kembali, m.Anggota_id, m.Petugas_id, id)
			if err != nil {
				log.Fatal(err)
				return err
			}
	
			return c.JSON(http.StatusOK, m)
		})

	// POST /peminjaman
	e.POST("/peminjaman", func(c echo.Context) error {
		var p peminjaman
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Menjalankan query untuk menambahkan peminjaman baru
		sqlStatement := "INSERT INTO peminjaman (tanggal_pinjam, tanggal_kembali, anggota_id, petugas_id) VALUES (?, ?, ?, ?)"
		result, err := db.Exec(sqlStatement, p.Tanggal_pinjam, p.Tanggal_kembali, p.Anggota_id, p.Petugas_id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		// Mendapatkan ID peminjaman yang baru saja ditambahkan
		id, _ := result.LastInsertId()
		p.ID = int(id)

		// Mengembalikan data peminjaman yang baru saja dimasukkan dalam bentuk JSON
		return c.JSON(http.StatusCreated, p)
	})
	
	// DELETE /peminjaman/:id
	e.DELETE("/peminjaman/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM peminjaman WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", id)
    return c.String(http.StatusOK, message)
	})

	// Handler untuk route GET /penerbit
	e.GET("/penerbit", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM penerbit")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var publishers []penerbit
		for res.Next() {
			var p penerbit
			err := res.Scan(&p.ID, &p.Nama, &p.Alamat, &p.No_telepon)
			if err != nil {
				log.Fatal(err)
				return err
			}
			publishers = append(publishers, p)
		}

		return c.JSON(http.StatusOK, publishers)
	})

	// POST /penerbit
	e.POST("/penerbit", func(c echo.Context) error {
		var p penerbit
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Execute the SQL statement to insert the data into the database
		sqlStatement := "INSERT INTO penerbit (nama, alamat, no_telepon) VALUES (?, ?, ?)"
		_, err := db.Exec(sqlStatement, p.Nama, p.Alamat, p.No_telepon)
		if err != nil {
			log.Fatal(err)
			return err
		}

		// Return the newly added 'pengarang' data in JSON format
		return c.JSON(http.StatusCreated, p)
	})
	
	// update penerbit
	e.PUT("/penerbit/:id", func(c echo.Context) error {
		// Get ID from URL parameter
		id := c.Param("id")

		// Bind data from request
		var m penerbit
		if err := c.Bind(&m); err != nil {
			return err
		}

		// Run query to update
		sqlStatement := "UPDATE penerbit SET nama = ?, alamat = ?, no_telepon = ?  WHERE id = ?"
		_, err := db.Exec(sqlStatement, m.Nama, m.Alamat, m.No_telepon, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusOK, m)
	})

	// DELETE /penerbit/:id
	e.DELETE("/penerbit/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM penerbit WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", id)
    return c.String(http.StatusOK, message)
	})

	// Handler untuk route GET /pengarang
	e.GET("/pengarang", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM pengarang")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var authors []pengarang
		for res.Next() {
			var p pengarang
			err := res.Scan(&p.ID, &p.Nama, &p.Alamat, &p.No_telepon)
			if err != nil {
				log.Fatal(err)
				return err
			}
			authors = append(authors, p)
		}

		return c.JSON(http.StatusOK, authors)
	})

	// POST /pengarang
	e.POST("/pengarang", func(c echo.Context) error {
		var p pengarang
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Execute the SQL statement to insert the data into the database
		sqlStatement := "INSERT INTO pengarang (nama, alamat, no_telepon) VALUES (?, ?, ?)"
		_, err := db.Exec(sqlStatement, p.Nama, p.Alamat, p.No_telepon)
		if err != nil {
			log.Fatal(err)
			return err
		}

		// Return the newly added 'pengarang' data in JSON format
		return c.JSON(http.StatusCreated, p)
	})
	
	// update pengarang
	e.PUT("/pengarang/:id", func(c echo.Context) error {
		// Get ID from URL parameter
		id := c.Param("id")
	
		// Bind data from request
		var m pengarang
		if err := c.Bind(&m); err != nil {
			return err
		}
	
		// Run query to update
		sqlStatement := "UPDATE pengarang SET nama = ?, alamat = ?, no_telepon = ? WHERE id = ?"
		_, err := db.Exec(sqlStatement, m.Nama, m.Alamat, m.No_telepon, id)
		if err != nil {
			log.Fatal(err)
			return err
		}
	
		return c.JSON(http.StatusOK, m)
	})
	
	// DELETE /pengarang/:id
	e.DELETE("/pengarang/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM pengarang WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", id)
    return c.String(http.StatusOK, message)
	})

	// Handler untuk route GET /detail_pengembalian
	e.GET("/pengembalian_detail", func(c echo.Context) error {
		// Pastikan Anda telah terhubung ke database sebelum menggunakan operasi SQL
		res, err := db.Query("SELECT * FROM pengembalian_detail")
		if err != nil {
			log.Fatal(err)
			return c.String(http.StatusInternalServerError, "Gagal mengambil data pengembalian_detail")
		}
		defer res.Close()
	
		var returnDetails []pengembalian_detail
		for res.Next() {
			var d pengembalian_detail
			err := res.Scan(&d.Pengembalian_id, &d.Buku_id)
			if err != nil {
				log.Fatal(err)
				return c.String(http.StatusInternalServerError, "Gagal memindai data pengembalian_detail")
			}
			returnDetails = append(returnDetails, d)
		}
	
		return c.JSON(http.StatusOK, returnDetails)
	})
	
	// POST /detail_pengembalian
	e.POST("/pengembalian_detail", func(c echo.Context) error {
		var d pengembalian_detail
		if err := c.Bind(&d); err != nil {
			return c.String(http.StatusBadRequest, "Data yang diberikan tidak valid")
		}
	
		sqlStatement := "INSERT INTO pengembalian_detail (pengembalian_id, buku_id) VALUES (?, ?)"
		_, err := db.Exec(sqlStatement, d.Pengembalian_id, d.Buku_id)
		if err != nil {
			log.Fatal(err)
			return c.String(http.StatusInternalServerError, "Gagal menambahkan data pengembalian_detail")
		}
	
		return c.JSON(http.StatusCreated, d)
	})
	
	// DELETE /detail_pengembalian/:pengembalian_id/:buku_id
	e.DELETE("/pengembalian_detail/:pengembalian_id/:buku_id", func(c echo.Context) error {
		pengembalianID := c.Param("pengembalian_id")
		bukuID := c.Param("buku_id")
	
		sqlStatement := "DELETE FROM pengembalian_detail WHERE pengembalian_id = ? AND buku_id = ?"
		_, err := db.Exec(sqlStatement, pengembalianID, bukuID)
		if err != nil {
			log.Fatal(err)
			return c.String(http.StatusInternalServerError, "Gagal menghapus data pengembalian_detail")
		}
	
		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus", bukuID)
		return c.String(http.StatusOK, message)
	})
	
	// Handler untuk route GET /pengembalian
	e.GET("/pengembalian", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM pengembalian")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var returns []pengembalian
		for res.Next() {
			var p pengembalian
			err := res.Scan(&p.ID, &p.Tanggal_pengembalian, &p.Denda, &p.Peminjaman_id, &p.Anggota_id, &p.Petugas_id)
			if err != nil {
				log.Fatal(err)
				return err
			}
			returns = append(returns, p)
		}

		return c.JSON(http.StatusOK, returns)
	})

	// POST /pengembalian
	e.POST("/pengembalian", func(c echo.Context) error {
		var p pengembalian
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Execute the SQL statement to insert the data into the database
		sqlStatement := "INSERT INTO pengembalian (tanggal_pengembalian, denda, peminjaman_id, anggota_id, petugas_id) VALUES (?, ?, ?, ?, ?)"
		_, err := db.Exec(sqlStatement, p.Tanggal_pengembalian, p.Denda, p.Peminjaman_id, p.Anggota_id, p.Petugas_id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		// Return the newly added 'pengembalian' data in JSON format
		return c.JSON(http.StatusCreated, p)
	})
	// update pengembalian
	e.PUT("/pengembalian/:id", func(c echo.Context) error {
		// Get ID from URL parameter
		id := c.Param("id")

		// Bind data from request
		var p pengembalian
		if err := c.Bind(&p); err != nil {
			return err
		}

		// Run query to update
		sqlStatement := "UPDATE pengembalian SET tanggal_pengembalian = ?, denda = ?, peminjaman_id = ?, anggota_id = ?, petugas_id = ?  WHERE id = ?"
		_, err := db.Exec(sqlStatement, p.Tanggal_pengembalian, p.Denda, p.Peminjaman_id, p.Anggota_id, p.Petugas_id, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusOK, p)
	})


	// DELETE /pengembalian/:id
	e.DELETE("/pengembalian/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM pengembalian WHERE id = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		
		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus",  id)
    return c.String(http.StatusOK, message)
	})


// GET /peminjaman_detail
e.GET("/peminjaman_detail", func(c echo.Context) error {
	// Pastikan Anda telah terhubung ke database sebelum menggunakan operasi SQL
	res, err := db.Query("SELECT * FROM peminjaman_detail")
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "Gagal mengambil data peminjaman_detail")
	}
	defer res.Close()

	var returnDetails []peminjaman_detail
	for res.Next() {
		var d peminjaman_detail
		err := res.Scan(&d.Peminjaman_id, &d.Buku_id)
		if err != nil {
			log.Fatal(err)
			return c.String(http.StatusInternalServerError, "Gagal memindai data peminjaman_detail")
		}
		returnDetails = append(returnDetails, d)
	}

	return c.JSON(http.StatusOK, returnDetails)
})

// POST /peminjaman_detail
e.POST("/peminjaman_detail", func(c echo.Context) error {
	var d peminjaman_detail
	if err := c.Bind(&d); err != nil {
		return c.String(http.StatusBadRequest, "Data yang diberikan tidak valid")
	}

	sqlStatement := "INSERT INTO peminjaman_detail (peminjaman_id, buku_id) VALUES (?, ?)"
	_, err := db.Exec(sqlStatement, d.Peminjaman_id, d.Buku_id)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "Gagal menambahkan data peminjaman_id")
	}

	return c.JSON(http.StatusCreated, d)
})

// DELETE /peminjaman_detail/:id
e.DELETE("/peminjaman_detail/:peminjaman_id/:buku_id", func(c echo.Context) error {
	peminjamanID := c.Param("peminjaman_id")
	bukuID := c.Param("buku_id")

	sqlStatement := "DELETE FROM peminjaman_detail WHERE peminjaman_id = ? AND buku_id = ?"
	_, err := db.Exec(sqlStatement, peminjamanID, bukuID)
	if err != nil {
		log.Fatal(err)
		return c.String(http.StatusInternalServerError, "Gagal menghapus data peminjaman_detail")
	}

	message := fmt.Sprintf("Data dengan  buku ID %s berhasil dihapus", bukuID)
	return c.String(http.StatusOK, message)
})

	// GET /rak
	e.GET("/rak", func(c echo.Context) error {
		res, err := db.Query("SELECT * FROM rak")

		if err != nil {
			log.Fatal(err)
			return err
		}
		defer res.Close()

		var racks []rak
		for res.Next() {
			var r rak
			err := res.Scan(&r.Kode_rak, &r.Lokasi)
			if err != nil {
				log.Fatal(err)
				return err
			}
			racks = append(racks, r)
		}

		return c.JSON(http.StatusOK, racks)
	})

	// POST /rak
	e.POST("/rak", func(c echo.Context) error {
		var r rak
		if err := c.Bind(&r); err != nil {
			return err
		}

		sqlStatement := "INSERT INTO rak (kode_rak, lokasi) VALUES (?, ?)"
		_, err := db.Exec(sqlStatement, r.Kode_rak, r.Lokasi)
		if err != nil {
			log.Fatal(err)
			return err
		}

		return c.JSON(http.StatusCreated, r)
	})

	// DELETE /rak/:id
	e.DELETE("/rak/:id", func(c echo.Context) error {
		id := c.Param("id")

		sqlStatement := "DELETE FROM rak WHERE kode_rak = ?"
		_, err := db.Exec(sqlStatement, id)
		if err != nil {
			log.Fatal(err)
			return err
		}

		
		message := fmt.Sprintf("Data dengan ID %s berhasil dihapus",  id)
    return c.String(http.StatusOK, message)
	})

	// Start server
	e.Logger.Fatal(e.Start(":8081"))
}
