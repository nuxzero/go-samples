-- Drop the "owner_currency_key" constraint if it exists
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "owner_currency_key";

-- Drop the "accounts_owner_fkey" foreign key if it exists
ALTER TABLE IF EXISTS "accounts" DROP CONSTRAINT IF EXISTS "accounts_owner_fkey";

-- Drop the "users" table if it exists
DROP TABLE IF EXISTS "users";
