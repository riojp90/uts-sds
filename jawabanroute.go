package route

import (
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// InsertData menangani endpoint untuk menambahkan data ke dalam database.
func InsertData(c *fiber.Ctx) error {
	data := new(models.User)
	if err := c.BodyParser(data); err != nil {
		return err
	}

	database.DB.Create(data)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil ditambahkan",
	})
}

// GetAllData mengambil semua data pengguna dari database.
func GetAllData(c *fiber.Ctx) error {
	var users []models.User
	database.DB.Find(&users)

	return c.JSON(fiber.Map{
		"data": users,
	})
}

// GetUserByid mengambil data pengguna berdasarkan id_user yang diberikan sebagai parameter.
func GetUserByid(c *fiber.Ctx) error {
	id_user := c.Params("id_user")
	var user models.User
	database.DB.Where("id_user = ?", id_user).First(&user)

	return c.JSON(fiber.Map{
		"data": user,
	})
}

// Delete menghapus data pengguna berdasarkan id_user yang diberikan sebagai parameter.
func Delete(c *fiber.Ctx) error {
	id_user := c.Params("id_user")
	database.DB.Where("id_user = ?", id_user).Delete(&models.User{})

	return c.JSON(fiber.Map{
		"Pesan": "Data telah dihapus",
	})
}

// Update memperbarui data pengguna berdasarkan id_user yang diberikan sebagai parameter.
func Update(c *fiber.Ctx) error {
	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	updateData := map[string]interface{}{
		"Nama":     data["nama"],
		"Email":    data["email"],
		"Password": data["password"],
	}

	database.DB.Model(&models.User{}).Where("id_user = ?", id_user).Updates(updateData)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah diperbarui",
	})
}
