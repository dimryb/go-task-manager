CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,                  -- Уникальный идентификатор пользователя
    username VARCHAR(255) UNIQUE NOT NULL,  -- Уникальное имя пользователя
    password TEXT NOT NULL,                 -- Захешированный пароль
    created_at TIMESTAMP DEFAULT NOW(),     -- Дата создания пользователя
    updated_at TIMESTAMP DEFAULT NOW(),     -- Дата последнего обновления пользователя
    deleted_at TIMESTAMP NULL
    );

CREATE TABLE IF NOT EXISTS tasks (
    id SERIAL PRIMARY KEY,               -- Уникальный идентификатор задачи
    title VARCHAR(255) NOT NULL,         -- Название задачи
    description TEXT,                    -- Описание задачи
    status VARCHAR(50) NOT NULL,         -- Статус задачи (pending, in_progress, done)
    priority VARCHAR(50) NOT NULL,       -- Приоритет задачи (low, medium, high)
    due_date TIMESTAMP NOT NULL,         -- Дата завершения задачи
    created_at TIMESTAMP DEFAULT NOW(),  -- Дата создания задачи
    updated_at TIMESTAMP DEFAULT NOW(),  -- Дата последнего обновления задачи
    user_id INT,                         -- ID пользователя, связанный с задачей
    CONSTRAINT fk_tasks_users FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL
    );
