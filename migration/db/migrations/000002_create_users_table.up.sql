CREATE TABLE IF NOT EXISTS account.users(
  id SERIAL PRIMARY KEY,
  email VARCHAR(100) UNIQUE NOT NULL,
  password VARCHAR(100) NOT NULL,
  created_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITHOUT TIME ZONE DEFAULT CURRENT_TIMESTAMP
);
CREATE INDEX IF NOT EXISTS idx_account_users_email on account.users(email);

INSERT INTO account.users(
  email,
  password
)
VALUES (
  'user@example.com', '$2a$10$sDiOoY6dwVqd2ispv360O.vedUD/rNqznW.UMd.k.z.PC2bznZfd2'
);
