-- init db
CREATE TABLE IF NOT EXISTS items (
    ID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(255)
);

-- populate data
INSERT IGNORE INTO items (ID, Name) VALUES (1, 'Lipstick');
INSERT IGNORE INTO items (ID, Name) VALUES (2, 'Pineapple');
INSERT IGNORE INTO items (ID, Name) VALUES (3, 'Sunglasses');