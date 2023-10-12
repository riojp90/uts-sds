package models

type Student struct {
	ID        uint     `gorm:"primaryKey" json:"id"`
	Name      string   `json:"name"`
	StudentID string   `json:"student_id"`
	Major     string   `json:"major"`
	Email     string   `json:"email"`
	Courses   []Course `gorm:"many2many:student_courses" json:"courses"`
}

type Professor struct {
	ID         uint     `gorm:"primaryKey" json:"id"`
	Name       string   `json:"name"`
	EmployeeID string   `json:"employee_id"`
	Department string   `json:"department"`
	Email      string   `json:"email"`
	Courses    []Course `gorm:"many2many:professor_courses" json:"courses"`
}

type Course struct {
	ID          uint        `gorm:"primaryKey" json:"id"`
	Code        string      `json:"code"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	Students    []Student   `gorm:"many2many:student_courses" json:"students"`
	Professors  []Professor `gorm:"many2many:professor_courses" json:"professors"`
}
