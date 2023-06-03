CREATE TABLE IF NOT EXISTS user_ingredient_preferences (
    "id" TEXT NOT NULL PRIMARY KEY,
    "ingredient" TEXT NOT NULL REFERENCES valid_ingredients("id") ON DELETE CASCADE ON UPDATE CASCADE,
    "rating" SMALLINT NOT NULL DEFAULT 0,
    "notes" TEXT NOT NULL DEFAULT '',
    "allergy" BOOLEAN NOT NULL DEFAULT 'false',
    "created_at" TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    "last_updated_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    "archived_at" TIMESTAMP WITH TIME ZONE DEFAULT NULL,
    "belongs_to_user" TEXT NOT NULL REFERENCES users("id") ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE("belongs_to_user", "ingredient")
);