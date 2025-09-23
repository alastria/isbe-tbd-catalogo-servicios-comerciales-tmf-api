package repository

const CreateTMFTableSQL = `
CREATE TABLE IF NOT EXISTS tmf_object (
	"id" TEXT NOT NULL,
	"type" TEXT NOT NULL,
	"version" TEXT DEFAULT '1.0',
	"api_version" TEXT DEFAULT 'v4',
	"seller" TEXT DEFAULT '',
	"buyer" TEXT DEFAULT '',
	"last_update" TEXT,
	"content" BLOB NOT NULL,
	"random" INTEGER DEFAULT 0,
	"created_at" DATETIME NOT NULL,
	"updated_at" DATETIME NOT NULL,
	PRIMARY KEY ("id", "type", "version")
);
`
