CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(50) UNIQUE NOT NULL,
  password_hash VARCHAR(255) NOT NULL,
  role VARCHAR(50) DEFAULT 'user'
);

CREATE TABLE tasks (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  status VARCHAR(50) DEFAULT 'To-Do' NOT NULL,
  priority VARCHAR(50) NOT NULL, 
  due_date TIMESTAMP,
  user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE SET NULL
);

CREATE TABLE categories (
  id SERIAL PRIMARY KEY,
  name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE task_categories (
  task_id INTEGER REFERENCES tasks(id),
  category_id INTEGER REFERENCES categories(id),
  PRIMARY KEY (task_id, category_id)
);
