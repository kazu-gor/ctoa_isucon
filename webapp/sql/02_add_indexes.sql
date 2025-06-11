CREATE INDEX idx_post_id_created_at ON comments (post_id, created_at);
CREATE INDEX idx_user_id_created_at ON posts (user_id, created_at);
CREATE INDEX idx_account_name ON users (account_name);
-- さらに先ほど追加したインデックスも追記しておくと良い
CREATE INDEX idx_created_at ON posts (created_at);
CREATE INDEX idx_user_id ON comments (user_id);
