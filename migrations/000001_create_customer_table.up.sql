CREATE SCHEMA "customers";
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE "customers"."customer" (
    "customer_id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "firstname" VARCHAR(255) NOT NULL,
    "lastname" varchar(200) NOT NULL,
    "mobile" varchar(200) NOT NULL, 
    "email" varchar(200) NOT NULL,
    "status" VARCHAR(255) NOT NULL DEFAULT 'active',
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" uuid NOT NULL,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_by" uuid NOT NULL,
    "deleted_at" timestamptz,
    "deleted_by" uuid
);
