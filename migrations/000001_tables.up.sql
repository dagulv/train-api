CREATE TABLE IF NOT EXISTS "users" (
    "id" VARCHAR(20) PRIMARY KEY,
    "firstName" VARCHAR(255) NOT NULL,
    "lastName" VARCHAR(255),
    "email" VARCHAR(255) NOT NULL UNIQUE,
    "password" VARCHAR(64),
    "credentialId" VARCHAR(1023),
    "publicKey" TEXT NOT NULL,
    "timeCreated" TIMESTAMP WITH TIME ZONE NOT NULL,
    "timeUpdated" TIMESTAMP WITH TIME ZONE NOT NULL
);

CREATE INDEX IF NOT EXISTS "usersEmail" ON "users" ("email");
