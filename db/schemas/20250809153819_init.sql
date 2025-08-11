-- +goose Up
-- +goose StatementBegin
CREATE TYPE "benchmark_ktype" AS ENUM (
  'high'
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL
);

CREATE TABLE "candidate" (
  "id" uuid PRIMARY KEY,
  "applied_position" varchar(100),
  "email" varchar(100),
  "cv_info" text,
  "cv_path" varchar(500),
  "original_cv_name" varchar(500),
  "file_hash" varchar(64),
  "plain_text" text,
  "embedding" vector(768),
  "created_at" date
);

CREATE TABLE "job_description" (
  "id" uuid PRIMARY KEY,
  "plain_text" text,
  "job_title" varchar(255),
  "jd_path" varchar(500),
  "file_hash" varchar(64),
  "jd_info" text,
  "embedding" vector(768)
);

CREATE TABLE "benchmark_result" (
  "candidate_id" uuid,
  "job_description_id" uuid,
  "benchmark_type" benchmark_ktype
);

CREATE TABLE "benchmark_result_detail" (
  "id" serial PRIMARY KEY,
  "candidate_id" uuid,
  "job_description_id" uuid,
  "benchmark_type" benchmark_ktype,
  "result" text
);

ALTER TABLE "benchmark_result" ADD FOREIGN KEY ("candidate_id") REFERENCES "candidate" ("id");

ALTER TABLE "benchmark_result" ADD FOREIGN KEY ("job_description_id") REFERENCES "job_description" ("id");

ALTER TABLE "benchmark_result_detail" ADD FOREIGN KEY ("candidate_id") REFERENCES "candidate" ("id");

ALTER TABLE "benchmark_result_detail" ADD FOREIGN KEY ("job_description_id") REFERENCES "job_description" ("id");
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "benchmark_result_detail";
DROP TABLE IF EXISTS "benchmark_result";
DROP TABLE IF EXISTS "job_description";
DROP TABLE IF EXISTS "candidate";
DROP TABLE IF EXISTS "user";
DROP TYPE IF EXISTS "benchmark_ktype";
-- +goose StatementEnd
