DROP TABLE IF EXISTS messages;
CREATE TABLE "messages" (
    "id" bigserial ,
    "sender_id" serial ,
    "room_id" serial ,
    "text" text,
    "time" timestamp
)
