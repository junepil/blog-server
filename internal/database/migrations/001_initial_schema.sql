-- Create the 'posts' table
CREATE TABLE posts (
    post_id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    content TEXT NOT NULL,
    published_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'tags' table
CREATE TABLE tags (
    tag_id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

-- Create the 'post_tags' join table
CREATE TABLE post_tags (
    post_id INTEGER REFERENCES posts(post_id) ON DELETE CASCADE,
    tag_id INTEGER REFERENCES tags(tag_id) ON DELETE CASCADE,
    PRIMARY KEY (post_id, tag_id)
);

-- Create the 'comments' table
CREATE TABLE comments (
    comment_id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(post_id) ON DELETE CASCADE,
    author_name VARCHAR(100) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);

-- Create the 'post_traffic' table
CREATE TABLE post_traffic (
    view_id SERIAL PRIMARY KEY,
    post_id INTEGER REFERENCES posts(post_id) ON DELETE CASCADE,
    viewed_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
    ip_address VARCHAR(45)
);
