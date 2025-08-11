# Hướng dẫn kết nối Supabase

## Bước 1: Tạo project trên Supabase
1. Truy cập [https://supabase.com](https://supabase.com)
2. Đăng nhập và tạo project mới
3. Chọn region gần nhất (Singapore cho Việt Nam)
4. Tạo password cho database

## Bước 2: Lấy thông tin kết nối
1. Vào project dashboard
2. Vào Settings > Database
3. Tìm phần "Connection string" và "Connection pooling"
4. Copy connection string có dạng:
   ```
   postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres
   ```

## Bước 3: Cập nhật file .env
Có 2 cách để cấu hình:

### Cách 1: Sử dụng connection string (Khuyến nghị)
```env
# Uncomment và thay thế với connection string của bạn
DB_SOURCE_DSN=postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:5432/postgres

# Comment lại các dòng config riêng lẻ
# DB_SOURCE_HOST=...
# DB_SOURCE_PORT=...
# etc.
```

### Cách 2: Sử dụng config riêng lẻ
```env
DB_SOURCE_HOST=db.[YOUR-PROJECT-REF].supabase.co
DB_SOURCE_PORT=5432
DB_SOURCE_DATABASE=postgres
DB_SOURCE_USERNAME=postgres
DB_SOURCE_PASSWORD=[YOUR-PASSWORD]
DB_SOURCE_SSL_MODE=require
```

## Bước 4: Test kết nối
```bash
make run
# hoặc
go run cmd/server/main.go
```

## Bước 5: Chạy migrations (nếu có)
Nếu bạn có schema migrations, hãy chạy chúng trên Supabase:

1. Vào Supabase Dashboard > SQL Editor
2. Chạy nội dung file `db/schemas/20250809153819_init.sql`
3. Hoặc sử dụng công cụ migration nếu có

## Lưu ý quan trọng:
- **SSL Mode**: Supabase yêu cầu `SSL_MODE=require`
- **Connection pooling**: Supabase có connection pooling built-in
- **Firewall**: Supabase cho phép kết nối từ mọi IP mặc định
- **Row Level Security**: Nên bật RLS cho production

## Troubleshooting:
1. **Connection timeout**: Kiểm tra region của Supabase project
2. **SSL errors**: Đảm bảo `SSL_MODE=require`
3. **Authentication failed**: Kiểm tra lại password
4. **Database not found**: Đảm bảo sử dụng database name `postgres`

## Connection pooling (tùy chọn):
Nếu muốn sử dụng connection pooling của Supabase:
```
postgresql://postgres:[YOUR-PASSWORD]@db.[YOUR-PROJECT-REF].supabase.co:6543/postgres
```
(Chú ý port 6543 thay vì 5432)
