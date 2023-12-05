ALTER TABLE
    "attachment" DROP CONSTRAINT IF EXISTS "attachment_id_messages_foreign";

ALTER TABLE
    "group_attachment" DROP CONSTRAINT IF EXISTS "group_attachment_group_id_foreign";

ALTER TABLE
    "group_attachment" DROP CONSTRAINT IF EXISTS "group_attachment_id_message_foreign";

ALTER TABLE
    "group_messages" DROP CONSTRAINT IF EXISTS "group_messages_group_id_foreign";

ALTER TABLE
    "user_connections" DROP CONSTRAINT IF EXISTS "user_connections_user_2_foreign";

ALTER TABLE
    "attachment" DROP CONSTRAINT IF EXISTS "attachment_connection_id_foreign";

ALTER TABLE
    "group_participants" DROP CONSTRAINT IF EXISTS "group_participants_group_id_foreign";

ALTER TABLE
    "messages" DROP CONSTRAINT IF EXISTS "messages_user_id_foreign";

ALTER TABLE
    "group_participants" DROP CONSTRAINT IF EXISTS "group_participants_user_id_foreign";

ALTER TABLE
    "user_connections" DROP CONSTRAINT IF EXISTS "user_connections_user_1_foreign";

ALTER TABLE
    "messages" DROP CONSTRAINT IF EXISTS "messages_connection_id_foreign";

DROP INDEX IF EXISTS "user_connections_user_1_index";

DROP INDEX IF EXISTS "user_connections_user_2_index";

DROP TABLE IF EXISTS "attachment";

DROP TABLE IF EXISTS "messages";

DROP TABLE IF EXISTS "user_connections";

DROP TABLE IF EXISTS "group_participants";

DROP TABLE IF EXISTS "groups";

DROP TABLE IF EXISTS "group_messages";

DROP TABLE IF EXISTS "group_attachment";

DROP TYPE IF EXISTS status;

DROP TABLE IF EXISTS "users" 