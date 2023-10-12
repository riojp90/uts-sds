package main

import (
	"log"
	"uts/database"
	"uts/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Inisialisasi koneksi ke database
	database.Connect()

	// Membuat instansi aplikasi Fiber
	app := fiber.New()

	// Kumpulan Route

	// Endpoint untuk menambahkan data
	app.Post("/insert", route.InsertData)

	// Endpoint untuk mengambil semua data pengguna
	app.Get("/getData", route.GetAllData)

	// Endpoint untuk mengambil data pengguna berdasarkan ID
	app.Get("/getDataUser/:id_user", route.GetUserByid)

	// Endpoint untuk menghapus data pengguna berdasarkan ID
	app.Get("/delete/:id_user", route.Delete)

	// Endpoint untuk memperbarui data pengguna berdasarkan ID
	app.Put("/update/:id_user", route.Update)

	// Jalankan server Fiber pada port tertentu
	log.Fatal(app.Listen(":3000"))
}
