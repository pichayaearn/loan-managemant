CREATE TABLE "customers"."loan" (
    "id" SERIAL PRIMARY KEY,
    "customer_id" uuid NOT NULL,
    "amount" decimal NOT NULL, 
    "interest" decimal NOT NULL,
    "start_date" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "end_date" timestamptz NOT NULL,
    "deb_pay_date" INT NOT NULL DEFAULT 30,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    "created_by" uuid NOT NULL,
    "updated_by" uuid NOT NULL,
    "deleted_by" uuid NOT NULL,
    FOREIGN KEY (customer_id) REFERENCES "customers".customer (customer_id)
);