
## 1. **Query Operations**

### **`First()`**
```go
u.db.First(&profile, id)
```
- ค้นหาข้อมูลแถวแรกที่ตรงกับเงื่อนไข
- หากไม่พบจะ return `gorm.ErrRecordNotFound`
- ใช้กับ primary key หรือเงื่อนไข WHERE(id เป็นเงื่อนไข สำหรับ primary key,profile เป็นการรับค่า)

### **`Find()`**
```go
u.db.Find(&profiles)
```
- ค้นหาข้อมูลทั้งหมดที่ตรงกับเงื่อนไข
- Return เป็น slice/array
- หากไม่พบจะ return slice ว่าง (ไม่ error)

### **`Where()`**
```go
u.db.Where("name = ?", userName)
u.db.Where("user_id = ?", *user.ID)
```
- เพิ่มเงื่อนไข WHERE clause
- ใช้ placeholder `?` เพื่อป้องกัน SQL injection
- สามารถ chain กับ method อื่นได้
- u.db.Where("name = ?", userName).First(&user).Error

## 2. **Write Operations**

### **`Create()`**
```go
u.db.Create(profile)
```
- สร้างข้อมูลใหม่ในฐานข้อมูล
- Auto-increment primary key จะถูกตั้งค่าอัตโนมัติ
- Return error หากมีปัญหา

### **`Save()`**
```go
u.db.Save(&existingProfile)
```
- บันทึกข้อมูลที่มีอยู่ (UPDATE)
- หากไม่มี primary key จะทำการ INSERT
- Update ทุกฟิลด์ของ struct

### **`Delete()`**
```go
u.db.Delete(&profile)
```
- ลบข้อมูลจากฐานข้อมูล
- ใช้ primary key ในการระบุแถวที่จะลบ
- สามารถใช้ soft delete ได้ (หากมี DeletedAt field)

## 3. **Advanced Query Operations**

### **`Table()`**
```go
u.db.Table("user_profiles")
```
- ระบุชื่อตารางที่ต้องการ query
- ใช้เมื่อต้องการ custom table name หรือ JOIN

### **`Select()`**
```go
Select("user_profiles.*")
```
- ระบุคอลัมน์ที่ต้องการ SELECT
- `*` หมายถึงทุกคอลัมน์
- สามารถระบุคอลัมน์เฉพาะได้

### **`Joins()`**
```go
Joins("JOIN users ON user_profiles.user_id = users.id")
```
- ทำการ JOIN ระหว่างตาราง
- รองรับ INNER JOIN, LEFT JOIN, RIGHT JOIN
- ใช้สำหรับ relationship queries

## 4. **Error Handling**

### **`.Error`**
```go
if err := u.db.First(&profile, id).Error; err != nil
```
- ทุก GORM operation จะมี `.Error` property
- ใช้ตรวจสอบว่ามี error เกิดขึ้นหรือไม่

### **`gorm.ErrRecordNotFound`**
```go
if err == gorm.ErrRecordNotFound {
    return nil, fmt.Errorf("user not found")
}
```
- Error พิเศษที่ GORM return เมื่อไม่พบข้อมูล
- ใช้ในการแยกแยะระหว่าง "ไม่พบข้อมูล" กับ "error อื่นๆ"

## 5. **Method Chaining**

GORM รองรับ method chaining:
```go
u.db.Table("user_profiles").
    Select("user_profiles.*").
    Joins("JOIN users ON user_profiles.user_id = users.id").
    Where("users.name LIKE ?", "%"+name+"%").
    Find(&profiles)
```

## 6. **ตัวอย่างการใช้งานในแต่ละฟังก์ชัน**

- **Create**: ใช้ `First()` ตรวจสอบ + `Create()` สร้างข้อมูล
- **Read**: ใช้ `Where()` + `First()` หาข้อมูลตาม condition
- **Update**: ใช้ `First()` หาข้อมูล + `Save()` บันทึก
- **Delete**: ใช้ `First()` หาข้อมูล + `Delete()` ลบ
- **Search**: ใช้ `Table()` + `Joins()` + `Where()` + `Find()` ค้นหา
- **GetAll**: ใช้ `Find()` ดึงข้อมูลทั้งหมด

รูปแบบนี้ช่วยให้จัดการฐานข้อมูลได้อย่างมีประสิทธิภาพและปลอดภัย!