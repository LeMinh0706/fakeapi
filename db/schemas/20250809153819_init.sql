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
  "applied_position" varchar(100) NOT NULL,
  "email" varchar(100) NOT NULL,
  "cv_info" text NOT NULL,
  "cv_path" varchar(500) NOT NULL,
  "original_cv_name" varchar(500),
  "file_hash" varchar(64) NOT NULL,
  "plain_text" text NOT NULL,
  "embedding" vector(768) NOT NULL,
  "created_at" date NOT NULL
);

CREATE TABLE "job_description" (
  "id" uuid PRIMARY KEY,
  "plain_text" text NOT NULL,
  "job_title" varchar(255) NOT NULL,
  "jd_path" varchar(500) NOT NULL,
  "file_hash" varchar(64) NOT NULL,
  "jd_info" text NOT NULL,
  "embedding" vector(768) NOT NULL
);

CREATE TABLE "benchmark_result" (
  "candidate_id" uuid NOT NULL,
  "job_description_id" uuid NOT NULL,
  "benchmark_type" benchmark_ktype NOT NULL
);

CREATE TABLE "benchmark_result_detail" (
  "id" serial PRIMARY KEY,
  "candidate_id" uuid NOT NULL,
  "job_description_id" uuid NOT NULL,
  "benchmark_type" benchmark_ktype NOT NULL,
  "result" text NOT NULL
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
DROP TABLE IF EXISTS "users";
DROP TYPE IF EXISTS "benchmark_ktype";
-- +goose StatementEnd
