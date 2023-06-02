package main

import (
	"fmt"
	"strings"

	"a21hc3NpZ25tZW50/helper"
)

var Students string = "A1234_Aditira_TI, B2131_Dito_TK, A3455_Afis_MI"
var StudentStudyPrograms string = "TI_Teknik Informatika, TK_Teknik Komputer, SI_Sistem Informasi, MI_Manajemen Informasi"

func Login(id string, name string) string {
	// Jika id atau nama kosong
    if name == "" {
        return "ID or Name is Undefined"
    }

	if id == "" {
        return "ID or Name is Undefined"
    }

    // Jika id kurang atau lebih dari 5 karakter
    if len(id) != 5 {
        return "ID must be 5 characters long!"
    }

    // Mencari data mahasiswa dengan id yang diberikan
    studentIndex := strings.Index(Students, id)

    // Jika tidak ditemukan, kembalikan pesan error
    if studentIndex == -1 {
        return "Login gagal: data mahasiswa tidak ditemukan"
    }

    // Mengambil nama mahasiswa dan jurusan
    studentData := Students[studentIndex+len(id)+1:]
    studentNameAndMajor := strings.Split(studentData, "_")

    // Jika nama mahasiswa kosong
    if studentNameAndMajor[0] == "" {
        return "ID or Name is Undefined"
    }

    // Jika data mahasiswa ditemukan, kembalikan pesan sukses
    if studentNameAndMajor[0] == name {
        return "Login berhasil: " + studentNameAndMajor[0]
    }

    // Jika nama mahasiswa tidak sesuai dengan yang diberikan
    return "Login gagal: data mahasiswa tidak ditemukan"
	
}



func Register(id string, name string, major string) string {
	

	if name == "" {
		return "ID or Name is Undefined"
	}
	if id == "" || len(id) != 5 {
		return "ID must be 5 characters long!"
	}
	// Cek apakah ID sudah digunakan
	for _, student := range strings.Split(Students, ", ") {
		if strings.HasPrefix(student, id+"_") {
			return "Registrasi gagal: id sudah digunakan"
		}
	}

	// Tambahkan siswa baru
	Students += ", " + id + "_" + name + "_" + major

	return "Registrasi berhasil: " + name + " (" + major + ")"

	
}

func GetStudyProgram(code string) string {
	// Pisahkan kode program studi dan nama program studi dari StudentStudyPrograms menggunakan split
	studyPrograms := strings.Split(StudentStudyPrograms, ", ")
	for _, studyProgram := range studyPrograms {
		if strings.HasPrefix(studyProgram, code+"_") {
			return studyProgram[strings.Index(studyProgram, "_")+1:]
		}
	}
	return "Code is undefined!"
}

func main() {
	fmt.Println("Selamat datang di Student Portal!")

	for {
		helper.ClearScreen()
		fmt.Println("Students: ", Students)
		fmt.Println("Student Study Programs: ", StudentStudyPrograms)

		fmt.Println("\nPilih menu:")
		fmt.Println("1. Login")
		fmt.Println("2. Register")
		fmt.Println("3. Get Program Study")
		fmt.Println("4. Keluar")

		var pilihan string
		fmt.Print("Masukkan pilihan Anda: ")
		fmt.Scanln(&pilihan)

		switch pilihan {
		case "1":
			helper.ClearScreen()
			var id, name string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)

			fmt.Println(Login(id, name))

			helper.Delay(5)
		case "2":
			helper.ClearScreen()
			var id, name, jurusan string
			fmt.Print("Masukkan id: ")
			fmt.Scan(&id)
			fmt.Print("Masukkan name: ")
			fmt.Scan(&name)
			fmt.Print("Masukkan jurusan: ")
			fmt.Scan(&jurusan)
			fmt.Println(Register(id, name, jurusan))

			helper.Delay(5)
		case "3":
			helper.ClearScreen()
			var kode string
			fmt.Print("Masukkan kode: ")
			fmt.Scan(&kode)

			fmt.Println(GetStudyProgram(kode))
			helper.Delay(5)
		case "4":
			fmt.Println("Terima kasih telah menggunakan Student Portal.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
