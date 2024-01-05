-- name: FindUserByID :one
SELECT * FROM "users" WHERE id = $1 LIMIT 1 OFFSET 0;

-- name: ListUsers :many
SELECT * FROM "users" LIMIT $1 OFFSET $2;

-- name: ListUserConnectionsID :many
SELECT id FROM "user_connections" WHERE user_1 = $1 OR user_2 = $1;

-- name: FindUserConnectionsByID :one
SELECT * FROM "user_connections" WHERE id = $1 LIMIT 1 OFFSET 0;

-- name: ListUserConnectWith :many
SELECT user_1 as contacts
FROM "user_connections" WHERE user_2 = $1
UNION
SELECT user_2 as contacts
FROM "user_connections" WHERE user_1 = $1;

-- name: FindUserConnectionsWith :one
SELECT * FROM "user_connections" 
WHERE user_1 = $1 and user_2 = $2 
OR user_1 = $2 and user_2 = $1
LIMIT 1 OFFSET 0;

-- name: CountUserConnections :one
SELECT COUNT(user_1 = $1 or user_2 = $1) as cout_contacts
FROM "user_connections";

-- name: ListUserMessages :many
SELECT * FROM messages
WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- name: ListUserMessagesByConetent :many
SELECT * FROM messages
WHERE message_content LIKE '%'|| $1 || '%';

-- name: FindMessageByConnectionID :many
SELECT * FROM messages
WHERE connection_id = $1
ORDER BY created_at ASC LIMIT $2 OFFSET $3;

-- name: FindMessageByID :one
SELECT * FROM messages
WHERE id = $1 LIMIT 1 OFFSET 0;

-- name: FindAttachmentByID :one
SELECT * FROM attachment
WHERE id = $1 LIMIT 1 OFFSET 0;

-- name: LisAttachmentByMessageID :many
SELECT * FROM attachment
WHERE id_messages = $1;

-- name: ListAttachmentByUserCon :many
SELECT * FROM attachment
WHERE connection_id = $1 LIMIT $2 OFFSET $3;

-- name: ListUserGroups :many
SELECT group_id from group_participants
WHERE user_id = $1 LIMIT $2 OFFSET $3;

-- name: ListGroupParticipants :many
SELECT user_id as participants from group_participants
WHERE group_id = $1;

-- nama: FindGroupByID :one
SELECT * FROM groups
WHERE id = $1 LIMIT 1 OFFSET 0;

-- name: ListGroupsByName :many
SELECT * FROM groups
WHERE name LIKE '%' || $1 || '%';

-- name: ListGroupMessages :many
SELECT * FROM group_messages
WHERE group_id = $1 LIMIT $2 OFFSET $3;

-- nama: FindGroupMessageByID :one
SELECT * FROM group_messages
WHERE id = $1 LIMIT 1;

-- name: ListGroupAttachments :many
SELECT * FROM group_attachment
WHERE group_id = $1 LIMIT $2 OFFSET $3;

-- name: ListMessageAttachments :many
SELECT * FROM group_attachment
WHERE id_message = $1;
