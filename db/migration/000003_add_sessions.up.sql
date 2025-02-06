CREATE TABLE sessions (
  "id" uuid PRIMARY KEY,
  "customer_id" bigint NOT NULL REFERENCES customers(id) ON DELETE CASCADE,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);