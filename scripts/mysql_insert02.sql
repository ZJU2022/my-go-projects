DELIMITER //

CREATE PROCEDURE manage_data()
BEGIN
    DECLARE i INT DEFAULT 1;
    DECLARE j INT DEFAULT 1;
    DECLARE tableName VARCHAR(255);
    DECLARE insertCount INT DEFAULT 0;
    DECLARE deleteCount INT DEFAULT 0;
    
    -- Loop to create 10 tables
    WHILE i <= 1000 DO
        SET tableName = CONCAT('table_', i);
        
        -- Create table
        SET @createTable = CONCAT('CREATE TABLE ', tableName, ' (
            id INT AUTO_INCREMENT PRIMARY KEY,
            data VARCHAR(255),
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )');
        PREPARE stmt FROM @createTable;
        EXECUTE stmt;
        DEALLOCATE PREPARE stmt;
        
        -- Loop to insert 100 rows into each table
        WHILE j <= 100 DO
            SET @insertData = CONCAT('INSERT INTO ', tableName, ' (data) VALUES (\'Sample data ', j, '\')');
            PREPARE stmt FROM @insertData;
            EXECUTE stmt;
            DEALLOCATE PREPARE stmt;
            
            -- Randomly delete a row with a 20% chance
            SET insertCount = insertCount + 1;
            IF RAND() < 0.2 THEN
                SET deleteCount = deleteCount + 1;
                SET @deleteData = CONCAT('DELETE FROM ', tableName, ' ORDER BY RAND() LIMIT 1');
                PREPARE stmt FROM @deleteData;
                EXECUTE stmt;
                DEALLOCATE PREPARE stmt;
            END IF;
            
            SET j = j + 1;
        END WHILE;
        
        SET i = i + 1;
        SET j = 1;
    END WHILE;
END//

DELIMITER ;

-- Call the procedure
CALL manage_data();

