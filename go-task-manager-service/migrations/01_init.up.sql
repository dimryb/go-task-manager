CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,               -- Уникальный идентификатор задачи
    title VARCHAR(255) NOT NULL,         -- Название задачи
    description TEXT,                    -- Описание задачи
    status VARCHAR(50) NOT NULL,         -- Статус задачи (pending, in_progress, done)
    priority VARCHAR(50) NOT NULL,       -- Приоритет задачи (low, medium, high)
    due_date TIMESTAMP NOT NULL,         -- Дата завершения задачи
    created_at TIMESTAMP DEFAULT NOW(),  -- Дата создания задачи
    updated_at TIMESTAMP DEFAULT NOW()   -- Дата последнего обновления задачи
);