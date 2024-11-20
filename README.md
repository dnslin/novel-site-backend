# 书籍数据库

```sql
CREATE TABLE "books" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,   -- 主键id
  "file_name" TEXT NOT NULL,   --文件名
  "title" TEXT NOT NULL,   -- 书名
  "author" TEXT NOT NULL,  --作者
  "file_size" INTEGER NOT NULL,  --文件大小
  "md5" TEXT NOT NULL,   -- 文件md5
  "new_file_name" TEXT NOT NULL,  --新的文件名
  "cover" TEXT,   --封面
  "intro" TEXT,   -- 简介
  "parts" TEXT,   -- 模糊搜索关键字
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP,  --创建时间
  "file_url" TEXT,    --文件下载地址
  "sort" TEXT,   -- 分类
  "type" TEXT,    -- 状态
  "tag" TEXT,    -- 书籍标签 
  UNIQUE ("md5" ASC)  
);

-- 评分类型表
CREATE TABLE "rating_types" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "name" TEXT NOT NULL,        -- 评级名称:仙草/粮草等
  "description" TEXT,          -- 评级描述 
  "level" INTEGER NOT NULL,    -- 评级等级:5/4/3/2/1
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
);

-- 匿名评价表
CREATE TABLE "book_ratings" (
  "id" INTEGER PRIMARY KEY AUTOINCREMENT,
  "book_id" INTEGER NOT NULL REFERENCES books(id),
  "rating_type_id" INTEGER NOT NULL REFERENCES rating_types(id), -- 评价类型
  "comment" TEXT,              -- 评价内容
  "ip" TEXT,                   -- 评价者IP(可选)
  "created_at" DATETIME DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO rating_types (name, description, level) VALUES
('仙草', '非常好看,值得反复阅读', 5),
('粮草', '好看,值得一读', 4), 
('干草', '一般,可以打发时间', 3),
('枯草', '不好看,不推荐', 2),
('毒草', '极差,不建议阅读', 1);

```