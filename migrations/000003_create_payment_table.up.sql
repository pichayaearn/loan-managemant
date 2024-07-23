CREATE TABLE "customers"."payment" (
    "id" BIGSERIAL PRIMARY KEY,
    "loan_id" BIGINT NOT NULL,
    "monthly_amount" DECIMAL NOT NULL, 
    "loan_amount" DECIMAL NOT NULL,
    "interest_amount" DECIMAL NOT NULL,
    "loan_balance" DECIMAL NOT NULL,
    "pay_date" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "next_pay_date" timestamptz NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "deleted_at" timestamptz,
    "created_by" uuid NOT NULL,
    "updated_by" uuid NOT NULL,
    "deleted_by" uuid NOT NULL,
    FOREIGN KEY (loan_id) REFERENCES "customers".loan (id)
);