package answer1

import (
	"database/sql"
	"fmt"
	"log"
)

func Migrate(db *sql.DB) bool {
	createJenisKainTable := `
    IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='JenisKain' AND xtype='U')
    BEGIN
        CREATE TABLE JenisKain (
            id INT PRIMARY KEY IDENTITY,
            jenis_kain VARCHAR(50)
        );
    END
    `

	createNamaKainTable := `
    IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='NamaKain' AND xtype='U')
    BEGIN
        CREATE TABLE NamaKain (
            id INT PRIMARY KEY IDENTITY,
            nama_kain VARCHAR(50),
            jenis_kain_id INT,
            FOREIGN KEY (jenis_kain_id) REFERENCES JenisKain(id)
        );
    END
    `

	createKualitasTable := `
    IF NOT EXISTS (SELECT * FROM sysobjects WHERE name='Kualitas' AND xtype='U')
    BEGIN
        CREATE TABLE Kualitas (
            id INT PRIMARY KEY IDENTITY,
            kualitas INT,
            nama_kualitas VARCHAR(50),
            harga VARCHAR(50),
            nama_kain_id INT,
            FOREIGN KEY (nama_kain_id) REFERENCES NamaKain(id)
        );
    END
    `
	_, err := db.Exec(createJenisKainTable)
	if err != nil {
		log.Fatal("Failed to create JenisKain table:", err)
	}

	_, err = db.Exec(createNamaKainTable)
	if err != nil {
		log.Fatal("Failed to create NamaKain table:", err)
	}

	_, err = db.Exec(createKualitasTable)
	if err != nil {
		log.Fatal("Failed to create Kualitas table:", err)
	}
	return true
}

func Read(db *sql.DB) {
	query := `
        SELECT jk.jenis_kain, nk.nama_kain, k.kualitas, k.nama_kualitas, k.harga
        FROM JenisKain jk
        JOIN NamaKain nk ON jk.id = nk.jenis_kain_id
        JOIN Kualitas k ON nk.id = k.nama_kain_id
        ORDER BY jk.jenis_kain, nk.nama_kain, k.kualitas;
    `

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Failed to execute query:", err)
	}
	defer rows.Close()

	var jenisKain, namaKain, namaKualitas, harga string
	var kualitas int

	fmt.Printf("%-10s %-10s %-10s %-20s %-10s\n", "Jenis Kain", "Nama Kain", "Kualitas", "Nama Kualitas", "Harga")
	for rows.Next() {
		err := rows.Scan(&jenisKain, &namaKain, &kualitas, &namaKualitas, &harga)
		if err != nil {
			log.Fatal("Failed to scan row:", err)
		}
		fmt.Printf("%-10s %-10s %-10d %-20s %-10s\n", jenisKain, namaKain, kualitas, namaKualitas, harga)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error iterating over rows:", err)
	}
}

func Create(db *sql.DB) bool {
	insertJenisKain := `
    INSERT INTO JenisKain (jenis_kain) VALUES ('STB'), ('NTB');
    `

	insertNamaKain := `
    INSERT INTO NamaKain (nama_kain, jenis_kain_id) VALUES ('Sutra', 1), ('Katun', 2);
    `

	insertKualitas := `
    INSERT INTO Kualitas (kualitas, nama_kualitas, harga, nama_kain_id) 
    VALUES 
    (1, 'Sangat Bagus', 'Rp 10000000', 1),
    (2, 'Bagus', 'Rp 9000000', 1),
    (3, 'Cukup Bagus', 'Rp 8000000', 1),
    (1, 'Sangat Bagus', 'Rp 10000000', 2),
    (2, 'Bagus', 'Rp 10000000', 2),
    (3, 'Cukup Bagus', 'Rp 10000000', 2);
    `
	_, err := db.Exec(insertJenisKain)
	if err != nil {
		log.Fatal("Failed to insert into JenisKain:", err)
		return false
	}

	_, err = db.Exec(insertNamaKain)
	if err != nil {
		log.Fatal("Failed to insert into NamaKain:", err)
		return false
	}

	_, err = db.Exec(insertKualitas)
	if err != nil {
		log.Fatal("Failed to insert into Kualitas:", err)
		return false
	}
	return true
}

func InsertJenisKain(db *sql.DB, input string) bool {
	insertJenisKain := `INSERT INTO JenisKain (jenis_kain) VALUES (?)`

	_, err := db.Exec(insertJenisKain, input)
	if err != nil {
		log.Fatal("Failed to insert into JenisKain:", err)
		return false
	}

	return true
}

func InsertNamaKain(db *sql.DB, nama string, id int) bool {
	insertNamaKain := `
    INSERT INTO NamaKain (nama_kain, jenis_kain_id) VALUES (?,?);
    `
	_, err := db.Exec(insertNamaKain, nama, id)
	if err != nil {
		log.Fatal("Failed to insert into NamaKain:", err)
		return false
	}

	return true
}

func InsertKualitasKain(db *sql.DB, kualitas int, namakualitas string, harga string, namaid int) bool {
	insertKualitas := `
    INSERT INTO Kualitas (kualitas, nama_kualitas, harga, nama_kain_id) 
    VALUES (?,?,?,?)`
	_, err := db.Exec(insertKualitas, kualitas, namakualitas, harga, namaid)
	if err != nil {
		log.Fatal("Failed to insert into KualitasKain:", err)
		return false
	}
	return true
}
