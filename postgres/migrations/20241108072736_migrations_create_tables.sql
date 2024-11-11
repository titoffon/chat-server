-- +goose Up
CREATE TABLE chats (
  id SERIAL PRIMARY KEY,
  usernames TEXT[] NOT NULL,
  created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE messages (
    id SERIAL PRIMARY KEY,           -- Уникальный идентификатор сообщения
    chat_id INT NOT NULL,             -- Идентификатор чата, к которому относится сообщение
    sender VARCHAR(255) NOT NULL,     -- Отправитель сообщения (может быть username или user_id)
    text TEXT NOT NULL,               -- Текст сообщения
    timestamp TIMESTAMP NOT NULL DEFAULT NOW(), -- Время отправки сообщения
    FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE -- Связь с таблицей чатов
);


-- +goose StatementBegin
--SELECT 'up SQL query';
-- +goose StatementEnd

-- +goose Down
DROP TABLE messages;
DROP TABLE chats;
-- +goose StatementBegin
--SELECT 'down SQL query';
-- +goose StatementEnd
