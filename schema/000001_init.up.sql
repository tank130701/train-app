CREATE TABLE users
(
    id serial primary key,
    username varchar(255) not null,
    password_hash varchar(255) not null,
    created_at timestamp
);

CREATE TABLE users_data
(
    user_id int references users(id) on delete cascade not null,
    name varchar(255) not null,
    age int not null,
    sex varchar(255) not null,
    weight int not null,
    height int not null,
    goal varchar(255) not null,
    place varchar(255) not null,
    calories_intake int not null
);

CREATE TABLE sessions
(
    id serial primary key,
    session varchar(255) not null,
    user_id int references users(id) on delete cascade not null,
    created_at timestamp
);

CREATE TABLE workouts (
  id SERIAL PRIMARY KEY,
  date TIMESTAMP NOT NULL,
  title VARCHAR(255) NOT NULL
);

CREATE TABLE exercises (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    sets INTEGER NOT NULL,
    reps INTEGER[] NOT NULL,
    weights INTEGER[],
    calories INTEGER NOT NULL
);

CREATE TABLE workout_exercises (
    workout_id INTEGER REFERENCES workouts (id),
    exercise_id INTEGER REFERENCES exercises (id),
    PRIMARY KEY (workout_id, exercise_id)
);

CREATE TABLE workouts_archive
(
    id serial primary key,
    user_id int references users(id),
    workout_id int references workouts(id)
);
