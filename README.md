# Perfume Quiz Backend

Backend API cho hệ thống quiz nước hoa sử dụng Golang và Gin framework.

## Setup

1. Cài đặt dependencies:
```bash
go mod tidy
```

2. Cấu hình database trong file `.env`:
```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=perfume_quiz
PORT=8080
```

3. Tạo database MySQL:
```sql
CREATE DATABASE perfume_quiz;
```

4. Chạy server:
```bash
go run main.go
```

## API Endpoints

### Cho khách hàng:
- `GET /api/questions` - Lấy tất cả câu hỏi
- `POST /api/submit-quiz` - Nộp bài quiz

### Cho admin:
- `GET /api/admin/results` - Lấy tất cả kết quả
- `GET /api/admin/stats` - Thống kê kết quả
- `DELETE /api/admin/results/:id` - Xóa kết quả

## Cấu trúc dữ liệu

### Submit Quiz Request:
```json
{
  "name": "Nguyễn Văn A",
  "phone": "0901234567", 
  "gender": "male", // "male", "female", "unisex"
  "answers": ["A", "B", "C", "D", "E"]
}
```

### Quiz Result Response:
```json
{
  "result": "The Maestro",
  "description": "Mô tả về kết quả...",
  "scores": {
    "M": 2,
    "S": 1, 
    "P": 0,
    "B": 0,
    "A": 1
  },
  "id": 1
}
```