CREATE TABLE IF NOT EXISTS items (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category ENUM('時計', 'バッグ', 'ジュエリー', '靴', 'その他') NOT NULL,
    brand VARCHAR(100) NOT NULL,
    purchase_price INT NOT NULL,
    purchase_date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

INSERT INTO items (name, category, brand, purchase_price, purchase_date) VALUES
('ロレックス デイトナ', '時計', 'ROLEX', 1500000, '2023-01-15'),
('エルメス バーキン', 'バッグ', 'HERMÈS', 2000000, '2023-02-20'),
('ティファニー ネックレス', 'ジュエリー', 'TIFFANY & Co.', 500000, '2023-03-10'),
('ルブタン パンプス', '靴', 'Christian Louboutin', 120000, '2023-04-05'),
('アップルウォッチ', 'その他', 'Apple', 60000, '2023-05-12');