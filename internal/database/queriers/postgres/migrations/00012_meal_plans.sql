CREATE TYPE meal_plan_status AS ENUM ('awaiting_votes', 'finalized');

CREATE TABLE IF NOT EXISTS meal_plans (
	"id" CHAR(27) NOT NULL PRIMARY KEY,
	"notes" TEXT NOT NULL,
    "status" meal_plan_status NOT NULL DEFAULT 'awaiting_votes',
    "voting_deadline" BIGINT NOT NULL,
	"starts_at" BIGINT NOT NULL,
	"ends_at" BIGINT NOT NULL,
	"created_on" BIGINT NOT NULL DEFAULT extract(epoch FROM NOW()),
	"last_updated_on" BIGINT DEFAULT NULL,
	"archived_on" BIGINT DEFAULT NULL,
	"belongs_to_household" CHAR(27) NOT NULL REFERENCES households("id") ON DELETE CASCADE
);

CREATE INDEX IF NOT EXISTS meal_plans_belongs_to_household on meal_plans (belongs_to_household);
