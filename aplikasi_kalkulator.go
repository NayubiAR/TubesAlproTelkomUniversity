package main

import (
	"fmt"
)

const A1 int = 1
const A2 int = 101

type User struct {
	Id                 int
	Username, Password string
}

type Pemain struct {
	Id, Kill, Death, Assist, HeadShot int
	NamaPemain, Peran                 string
	Value                             float64
}

var nextID int = 1
var registered bool

var dataPemain [A2]Pemain

var dataUser [A1]User

func main() {
	var u, p string
	fmt.Println("====================================================================================")
	fmt.Println("====               ", "Selamat Datang DiAplikasi Player LeaderBoard", "               ====")
	fmt.Println("====================================================================================")
	login(&u, &p)
}

func login(user, pass *string) { // Melakukan Login
	var validasiReg string
	var validasi, validasi1 int

	if dataUser[0].Id != 1 {
		fmt.Println("====================================================================================")
		fmt.Println("====                             =", "Halaman Login", "=                              ====")
		fmt.Println("====", "Sudah punya akun ?", "                                                        ====")
		fmt.Println("====", "Silahkan Ketik '1' Jika sudah memiliki akun ", "                              ====")
		fmt.Println("====", "Silahkan Ketik '2' Jika belum memiliki akun ", "                              ====")
		fmt.Println("====", "Silahkan Ketik '3' Jika ingin masuk sebagai guest ", "                        ====")
		fmt.Println("====", "Silahkan Ketik '4' Jika ingin keluar dari aplikasi", "                        ====")
		fmt.Print("==== ")
		fmt.Scan(&validasi)
	} else {
		fmt.Println("====================================================================================")
		fmt.Println("====                             =", "Halaman Login", "=                              ====")
		fmt.Println("====", "Silahkan Ketik '1' Jika anda ingin Login ", "       	                       ====")
		fmt.Println("====", "Silahkan Ketik '2' Jika anda ingin masuk sebagai guest ", "                   ====")
		fmt.Println("====", "Silahkan Ketik '3' Jika ingin keluar dari aplikasi", "                        ====")
		fmt.Print("==== ")
		fmt.Scan(&validasi1)
	}

	// Validasi Jika belum memiliki akun
	if validasi == 1 {
		fmt.Println("====================================================================================")
		fmt.Println("====                             =", "Halaman Login", "=                              ====")
		fmt.Print("==== ", "Masukkan Username: ")
		fmt.Scan(user)
		fmt.Print("==== ", "Masukkan Password: ")
		fmt.Scan(pass)
		if validasiLogin(*user, *pass) == 1 {
			fmt.Println("====", "Login berhasil!")
			fmt.Println("====================================================================================")
			// Panggil function tampilan
			tampilan()
		} else {
			fmt.Println("====", "Login gagal! Silakan coba lagi.")
			login(user, pass) // Re-prompt for login
		}
		fmt.Println("====================================================================================")
	} else if validasi == 2 {
		fmt.Println("====================================================================================")
		fmt.Println("====                             =", "Halaman Login", "=                              ====")
		fmt.Println("====", "Ingin Membuat akun ?", "                                                      ====")
		fmt.Println("====", "Ketik '1' Jika ingin Membuat akun.", "                                        ====")
		fmt.Println("====", "Ketik '2' Jika tidak ingin Membuat akun.", "                                  ====")
		fmt.Print("==== ")
		fmt.Scan(&validasiReg)
		if validasiReg == "1" {
			registrasi()
		} else {
			login(user, pass)
		}
		fmt.Println("====================================================================================")
	} else if validasi == 3 {
		// Langsung masuk ke tampilan sebagai guest
		fmt.Println("====================================================================================")
		tampilan()
	} else {
	}

	// Validasi Jika sudah memiliki akun
	// Validasi Jika sudah memiliki akun
	if validasi1 == 1 {
		for {
			fmt.Println("====================================================================================")
			fmt.Println("====                             =", "Halaman Login", "=                              ====")
			fmt.Print("==== ", "Masukkan Username: ")
			fmt.Scan(user)
			fmt.Print("==== ", "Masukkan Password: ")
			fmt.Scan(pass)
			if validasiLogin(*user, *pass) == 1 {
				fmt.Println("====", "Login berhasil!")
				// Panggil function tampilan
				tampilan()
			} else {
				fmt.Println("====", "Login gagal! Silahkan coba lagi.")
				// Lanjutkan loop untuk mencoba login kembali
			}
		}

	} else if validasi1 == 2 {
		// Langsung masuk ke tampilan sebagai guest
		tampilan()
	} else {
	}

}

func validasiLogin(user, pass string) int { // Mengvalidasi login
	if user == dataUser[0].Username && pass == dataUser[0].Password {
		registered = true
		return 1
	}
	return 0
}

func registrasi() { // Membuat Akun
	var b, c string
	fmt.Println("====================================================================================")
	dataUser[0].Id = 1
	fmt.Println("====                           =", "Halaman Registrasi", "=                           ====")
	fmt.Print("==== ", "Masukkan Username: ")
	fmt.Scan(&dataUser[0].Username)
	fmt.Print("==== ", "Masukkan Password: ")
	fmt.Scan(&dataUser[0].Password)
	fmt.Println("====", "Anda berhasil membuat akun!")
	fmt.Println("====================================================================================")
	login(&b, &c)
}

func tampilan() { // Menampilkan Dashboard Halaman
	var validasi int
	fmt.Println("====================================================================================")
	fmt.Println("====                           =", "Halaman Dashboard", "=                            ====")
	fmt.Println("====", "Silahkan Ketik '1' Jika anda ingin masuk ke Halaman LeaderBoard ", "          ====")
	fmt.Println("====", "Silahkan Ketik '2' Jika anda ingin masuk ke Halaman Tambah data pemain ", "   ====")
	fmt.Println("====", "Silahkan Ketik '3' Jika anda ingin masuk ke Halaman Ubah data pemain ", "     ====")
	fmt.Println("====", "Silahkan Ketik '4' Jika anda ingin masuk ke Halaman Hapus data pemain ", "    ====")
	fmt.Println("====", "Silahkan Ketik '5' Jika anda ingin mendapatkan feedback pemain ", "           ====")
	fmt.Println("====", "Silahkan Ketik '6' Jika anda ingin membandingkan pemain ", "                  ====")
	fmt.Println("====", "Silahkan Ketik '7' Jika ingin keluar dari aplikasi", "                        ====")
	fmt.Print("==== ")
	fmt.Scan(&validasi)
	if validasi == 1 {
		if nextID == 1 {
			fmt.Println("====", "Maaf anda perlu memasukkan data terlebih dahulu untuk masuk ke halaman ini")
			tampilan()
		} else {

			tampilanLeaderBoard()
		}
	} else if validasi == 2 {
		tambahData()
	} else if validasi == 3 {
		if nextID == 1 {
			fmt.Println("====", "Maaf anda perlu memasukkan data terlebih dahulu untuk masuk ke halaman ini")
			tampilan()
		} else {
			ubahData()
		}
	} else if validasi == 4 {
		if nextID == 1 {
			fmt.Println("====", "Maaf anda perlu memasukkan data terlebih dahulu untuk masuk ke halaman ini")
			tampilan()
		} else {
			hapusData()
		}
	} else if validasi == 5 {
		if registered == true {
			if nextID == 1 {
				fmt.Println("====", "Maaf anda perlu memasukkan data terlebih dahulu untuk masuk ke halaman ini")
				tampilan()
			} else {
				feedBack()
			}
		} else {
			fmt.Println("====", "Maaf anda perlu melakukan login terlebih dahulu")
			tampilan()
		}
	} else if validasi == 6 {
		if registered == true {
			if nextID == 1 {
				fmt.Println("====", "Maaf anda perlu memasukkan data terlebih dahulu untuk masuk ke halaman ini")
				tampilan()
			} else {
				comparePlayer()
			}
		} else {
			fmt.Println("====", "Maaf anda perlu melakukan login terlebih dahulu")
			tampilan()
		}

	}
}

func tampilanLeaderBoard() { // Menampilkan Ranking Leaderboard
	var kembali int
	hitungValue()
	// menggunakan Selection Sort Descending ( langsung menampilkan value tertinggi)
	for i := 1; i < nextID-1; i++ {
		maxIdx := i
		for j := i + 1; j < nextID; j++ {
			if dataPemain[j].Value > dataPemain[maxIdx].Value {
				maxIdx = j
			}
		}
		dataPemain[i], dataPemain[maxIdx] = dataPemain[maxIdx], dataPemain[i]
	}

	fmt.Println("====================================================================================")
	fmt.Println("====                        =", "Halaman Ranking LeaderBoard", "=                     ====")
	for i := 1; i < nextID; i++ {
		player := dataPemain[i]
		fmt.Printf("==== %d. %s - Kill: %d, Death: %d, Assist: %d, Persentase HeadShot: %d, Peran: %s, Value: %.2f \n",
			i, player.NamaPemain, player.Kill, player.Death, player.Assist, player.HeadShot, player.Peran, player.Value)
	}
	fmt.Println("====================================================================================")
	fmt.Println("====", "Silahkan Ketik '1' Jika ingin kembali", "                                     ====")
	fmt.Println("====", "Silahkan Ketik '2' Jika ingin mengurutkan berdasarkan nama depan pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '3' Jika ingin mengurutkan berdasarkan nama belakang pemain", "====")
	fmt.Println("====", "Silahkan Ketik '4' Jika ingin mengurutkan berdasarkan value paling rendah pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '5' Jika ingin mengurutkan berdasarkan value paling tinggi pemain", "   ====")
	fmt.Scan(&kembali)
	if kembali == 1 {
		tampilan()
	} else if kembali == 2 {
		urutnamaDepan()
	} else if kembali == 3 {
		urutnamaBelakang()
	} else if kembali == 4 {
		valueAscending()
	} else if kembali == 5 {
		valueDescending()
	}
}

func urutnamaDepan() { // Mengurutkan Data Berdasarkan Huruf Terdepan nama pemain (A-Z)
	var kembali, i, j int
	var key Pemain
	// Insertion Sort
	for i = 2; i < nextID; i++ { // Mulai dari indeks ke-2 karena indeks ke-1 sudah dianggap urut
		key = dataPemain[i]
		j = i - 1
		for j >= 1 && dataPemain[j].NamaPemain[0] > key.NamaPemain[0] {
			dataPemain[j+1] = dataPemain[j]
			j--
		}
		dataPemain[j+1] = key
	}

	fmt.Println("====================================================================================")
	fmt.Println("====       =", "Data Pemain Setelah Diurutkan Berdasarkan Huruf Pertama Nama", "=      ====")
	for i := 1; i < nextID; i++ {
		player := dataPemain[i]
		fmt.Printf("==== %d. %s - Kill: %d, Death: %d, Assist: %d, Persentase HeadShot: %d, Peran: %s, Value: %.2f \n",
			i, player.NamaPemain, player.Kill, player.Death, player.Assist, player.HeadShot, player.Peran, player.Value)
	}
	fmt.Println("====================================================================================")
	fmt.Println("====", "Silahkan Ketik '1' Jika ingin kembali", "                                     ====")
	fmt.Println("====", "Silahkan Ketik '2' Jika ingin mengurutkan berdasarkan nama pemain", "         ====")
	fmt.Println("====", "Silahkan Ketik '3' Jika ingin mengurutkan berdasarkan nama belakang pemain", "         ====")
	fmt.Println("====", "Silahkan Ketik '4' Jika ingin mengurutkan berdasarkan value paling rendah pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '5' Jika ingin mengurutkan berdasarkan value paling tinggi pemain", "   ====")
	fmt.Scan(&kembali)
	if kembali == 1 {
		tampilan()
	} else if kembali == 2 {
		urutnamaDepan()
	} else if kembali == 3 {
		urutnamaBelakang()
	} else if kembali == 4 {
		valueAscending()
	} else if kembali == 5 {
		valueDescending()
	}
}

func urutnamaBelakang() { // // Mengurutkan Data Berdasarkan Huruf Terdepan nama pemain (Z - A)
	var kembali, i, j int
	var key Pemain
	// Insertion Sort
	for i = 2; i < nextID; i++ { // Mulai dari indeks ke-2 karena indeks ke-1 sudah dianggap urut
		key = dataPemain[i]
		j = i - 1
		for j >= 1 && dataPemain[j].NamaPemain[0] < key.NamaPemain[0] {
			dataPemain[j+1] = dataPemain[j]
			j--
		}
		dataPemain[j+1] = key
	}

	fmt.Println("====================================================================================")
	fmt.Println("====       =", "Data Pemain Setelah Diurutkan Berdasarkan Huruf Pertama Nama (Descending)", "=      ====")
	for i := 1; i < nextID; i++ {
		player := dataPemain[i]
		fmt.Printf("==== %d. %s - Kill: %d, Death: %d, Assist: %d, Persentase HeadShot: %d, Peran: %s, Value: %.2f \n",
			i, player.NamaPemain, player.Kill, player.Death, player.Assist, player.HeadShot, player.Peran, player.Value)
	}
	fmt.Println("====================================================================================")
	fmt.Println("====", "Silahkan Ketik '1' Jika ingin kembali", "                                     ====")
	fmt.Println("====", "Silahkan Ketik '2' Jika ingin mengurutkan berdasarkan nama depan pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '3' Jika ingin mengurutkan berdasarkan nama belakang pemain", "         ====")
	fmt.Println("====", "Silahkan Ketik '4' Jika ingin mengurutkan berdasarkan value paling rendah pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '5' Jika ingin mengurutkan berdasarkan value paling tinggi pemain", "   ====")
	fmt.Scan(&kembali)
	if kembali == 1 {
		tampilan()
	} else if kembali == 2 {
		urutnamaDepan()
	} else if kembali == 3 {
		urutnamaBelakang()
	} else if kembali == 4 {
		valueAscending()
	} else if kembali == 5 {
		valueDescending()
	}
}

func valueAscending() { // Mengurutkan Data berdasarkan Value Terendah
	var kembali int
	// Selection Sort
	for i := 1; i < nextID-1; i++ {
		minIdx := i
		for j := i + 1; j < nextID; j++ {
			if dataPemain[j].Value < dataPemain[minIdx].Value {
				minIdx = j
			}
		}
		dataPemain[i], dataPemain[minIdx] = dataPemain[minIdx], dataPemain[i]
	}

	fmt.Println("====================================================================================")
	fmt.Println("====       =", "Data Pemain Setelah Diurutkan Berdasarkan Nilai (Ascending)", "=      ====")
	for i := 1; i < nextID; i++ {
		player := dataPemain[i]
		fmt.Printf("==== %d. %s - Kill: %d, Death: %d, Assist: %d, Persentase HeadShot: %d, Peran: %s, Value: %.2f \n",
			i, player.NamaPemain, player.Kill, player.Death, player.Assist, player.HeadShot, player.Peran, player.Value)
	}
	fmt.Println("====================================================================================")
	fmt.Println("====", "Silahkan Ketik '1' Jika ingin kembali", "                                     ====")
	fmt.Println("====", "Silahkan Ketik '2' Jika ingin mengurutkan berdasarkan nama depan pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '3' Jika ingin mengurutkan berdasarkan nama belakang pemain", "         ====")
	fmt.Println("====", "Silahkan Ketik '4' Jika ingin mengurutkan berdasarkan value paling rendah pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '5' Jika ingin mengurutkan berdasarkan value paling tinggi pemain", "   ====")
	fmt.Scan(&kembali)
	if kembali == 1 {
		tampilan()
	} else if kembali == 2 {
		urutnamaDepan()
	} else if kembali == 3 {
		urutnamaBelakang()
	} else if kembali == 4 {
		valueAscending()
	} else if kembali == 5 {
		valueDescending()
	}
}

func valueDescending() { // Mengurutkan Data berdasarkan Value Tertinggi
	var kembali int
	//Selection Sort
	for i := 1; i < nextID-1; i++ {
		minIdx := i
		for j := i + 1; j < nextID; j++ {
			if dataPemain[j].Value > dataPemain[minIdx].Value {
				minIdx = j
			}
		}
		dataPemain[i], dataPemain[minIdx] = dataPemain[minIdx], dataPemain[i]
	}

	fmt.Println("====================================================================================")
	fmt.Println("====       =", "Data Pemain Setelah Diurutkan Berdasarkan Nilai (Ascending)", "=      ====")
	for i := 1; i < nextID; i++ {
		player := dataPemain[i]
		fmt.Printf("==== %d. %s - Kill: %d, Death: %d, Assist: %d, Persentase HeadShot: %d, Peran: %s, Value: %.2f \n",
			i, player.NamaPemain, player.Kill, player.Death, player.Assist, player.HeadShot, player.Peran, player.Value)
	}
	fmt.Println("====================================================================================")
	fmt.Println("====", "Silahkan Ketik '1' Jika ingin kembali", "                                     ====")
	fmt.Println("====", "Silahkan Ketik '2' Jika ingin mengurutkan berdasarkan nama depan pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '3' Jika ingin mengurutkan berdasarkan nama belakang pemain", "         ====")
	fmt.Println("====", "Silahkan Ketik '4' Jika ingin mengurutkan berdasarkan value paling rendah pemain", "   ====")
	fmt.Println("====", "Silahkan Ketik '5' Jika ingin mengurutkan berdasarkan value paling tinggi pemain", "   ====")
	fmt.Scan(&kembali)
	if kembali == 1 {
		tampilan()
	} else if kembali == 2 {
		urutnamaDepan()
	} else if kembali == 3 {
		urutnamaBelakang()
	} else if kembali == 4 {
		valueAscending()
	} else if kembali == 5 {
		valueDescending()
	}
}

func tambahData() { // Menambahkan data
	var n, validasi int

	fmt.Println("====================================================================================")
	fmt.Println("====                       =", "Halaman Tambah Data Pemain", "=                       ====")

	fmt.Println("====", "Berapa banyak pemain yang ingin anda tambah kan?")
	fmt.Println("====", "Maksimal pemain yang ditambah kan 100 !")
	fmt.Print("==== ")
	fmt.Scan(&n)
	validasi = n + nextID
	fmt.Println("====================================================================================")

	if validasi > A2 {
		fmt.Println("==== Maaf data yang anda masukkan sudah melebihi maksimal pemain")
		tampilan()
	} else {
		for i := 1; i <= n; i++ {
			fmt.Println("====", "Silahkan masukkan data pemain !")

			dataPemain[nextID].Id = nextID
			fmt.Print("==== ", "Masukkan Nama Pemain: ")
			fmt.Scan(&dataPemain[nextID].NamaPemain)
			fmt.Print("==== ", "Masukkan Jumlah Kill Pemain: ")
			fmt.Scan(&dataPemain[nextID].Kill)
			fmt.Print("==== ", "Masukkan Jumlah Death Pemain: ")
			fmt.Scan(&dataPemain[nextID].Death)
			fmt.Print("==== ", "Masukkan Jumlah Assist Pemain: ")
			fmt.Scan(&dataPemain[nextID].Assist)
			fmt.Print("==== ", "Masukkan Persentase HeadShot Pemain: ")
			fmt.Scan(&dataPemain[nextID].HeadShot)
			fmt.Print("==== ", "Masukkan Peran Pemain: ")
			fmt.Scan(&dataPemain[nextID].Peran)
			nextID++ // Increment ke ID untuk id pemain dan array nya

			fmt.Print("==== ")
			fmt.Println("Berhasil memasukkan", i, "data Pemain")
			fmt.Println("====================================================================================")
		}
	}
	tampilan()
}

func ubahData() { // Mengubah / Mengedit data
	var cariNama, namaBaru, peranBaru string
	var killBaru, deathBaru, assistBaru, headshotBaru int

	fmt.Println("====================================================================================")
	fmt.Println("====                       =", "Halaman Ubah Data Pemain", "=                         ====")
	fmt.Print("====", "Silahkan masukkan nama pemain! ")
	fmt.Scan(&cariNama)
	// Sequential search
	for i := 1; i < nextID; i++ {
		if cariNama == dataPemain[i].NamaPemain {
			fmt.Print("==== Buat Nama Baru: ")
			fmt.Scan(&namaBaru)
			dataPemain[i].NamaPemain = namaBaru
			fmt.Print("==== Buat Jumlah Kill Baru: ")
			fmt.Scan(&killBaru)
			dataPemain[i].Kill = killBaru
			fmt.Print("==== Buat Jumlah Death Baru: ")
			fmt.Scan(&deathBaru)
			dataPemain[i].Death = deathBaru
			fmt.Print("==== Buat Jumlah Assist Baru: ")
			fmt.Scan(&assistBaru)
			dataPemain[i].Assist = assistBaru
			fmt.Print("==== Buat Jumlah HeadShot Baru: ")
			fmt.Scan(&headshotBaru)
			dataPemain[i].HeadShot = headshotBaru
			fmt.Print("==== Buat Peran Baru: ")
			fmt.Scan(&peranBaru)
			dataPemain[i].Peran = peranBaru
		}
	}
	tampilan()
}

func hapusData() { // Menghapus / Mengdelete data
	var cariNama string
	var found bool

	fmt.Println("====================================================================================")
	fmt.Println("====                       =", "Halaman Hapus Data Pemain", "=                        ====")
	fmt.Print("====", "Silahkan masukkan nama pemain yang ingin dihapus: ")
	fmt.Scan(&cariNama)

	// Binary search
	low, high := 1, nextID-1
	for low <= high {
		mid := (low + high) / 2
		if dataPemain[mid].NamaPemain == cariNama {
			found = true

			for j := mid; j < nextID-1; j++ {
				dataPemain[j] = dataPemain[j+1]
			}

			nextID--
			fmt.Println("====", "Data pemain berhasil dihapus!")
			// Memperbarui nilai high agar loop berhenti setelah data ditemukan
			high = mid - 1
		} else if dataPemain[mid].NamaPemain < cariNama {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if !found {
		fmt.Println("====", "Pemain dengan nama tersebut tidak ditemukan.")
	}
	fmt.Println("====================================================================================")
	tampilan()
}

func hitungValue() { // Menghitung Value pemain untuk digunakan saat  mengurutkan pemain
	for i := 1; i < nextID; i++ {
		if dataPemain[i].Death != 0 {
			dataPemain[i].Value = float64(dataPemain[i].Kill+dataPemain[i].Assist-dataPemain[i].Death) / 10.0
		} else {
			dataPemain[i].Value = float64(dataPemain[i].Kill+dataPemain[i].Assist) / 10.0
		}
	}
}

func feedBack() { // Mengembalikan Feedback dari program jika user sudah melakukan registrasi
	var cariNama string
	fmt.Println("====================================================================================")
	fmt.Println("====                        =", "Halaman Feedback Pemain", "=                     ====")
	fmt.Print("==== ", "Silahkan masukkan nama pemain! ")
	fmt.Scan(&cariNama)
	for i := 1; i < nextID; i++ {
		if cariNama == dataPemain[i].NamaPemain {
			if dataPemain[i].Value < 1.00 {
				if dataPemain[i].HeadShot > 50 && dataPemain[i].Peran == "rifler" {
					fmt.Print("==== Tingkat kan lagi kill dan assist anda berlindung disemak-semak mungkin akan membantu anda")
				} else if dataPemain[i].HeadShot < 50 && dataPemain[i].Peran == "rifler" {
					fmt.Print("==== Mengubah Role anda menjadi sniper mungkin akan membantu ====")
				} else {
					fmt.Print("==== Tingkat kan lagi kill dan assist anda berlindung disemak-semak mungkin akan membantu anda")
				}

			} else if dataPemain[i].Value > 1.00 && dataPemain[i].Value <= 2.00 {
				if dataPemain[i].HeadShot > 50 && dataPemain[i].Peran == "rifler" {
					fmt.Print("==== Tingkat kan lagi kill dan assist anda berlindung disemak-semak mungkin akan membantu anda")
				} else if dataPemain[i].HeadShot < 50 && dataPemain[i].Peran == "rifler" {
					fmt.Print("==== Mengubah Role anda menjadi sniper mungkin akan membantu ====")
				} else {
					fmt.Print("==== Tingkat kan lagi kill dan assist anda berlindung disemak-semak mungkin akan membantu anda")
				}

			} else if dataPemain[i].Value > 2.00 && dataPemain[i].Value <= 3.00 {
				if dataPemain[i].HeadShot > 50 && dataPemain[i].Peran == "rifler" {
					fmt.Print("==== Tingkat kan lagi kill dan assist anda berlindung disemak-semak mungkin akan membantu anda")
				} else if dataPemain[i].HeadShot < 50 && dataPemain[i].Peran == "rifler" {
					fmt.Print("==== Mengubah Role anda menjadi sniper mungkin akan membantu ====")
				} else {
					fmt.Print("==== Tingkat kan lagi kill dan assist anda berlindung disemak-semak mungkin akan membantu anda")
				}
			} else {
				fmt.Print("==== KDA anda sudah sangat bagus tetap pertahankan")
			}
		}
	}
	fmt.Println("====================================================================================")
	tampilan()
}

func comparePlayer() { // melakukan perbandingan antar pemain jika user sudah melakukan registrasi
	var cariNama, cariNama1 string
	var pemain1, pemain2 Pemain

	fmt.Print("====", "Silahkan masukkan nama pemain pertama! ")
	fmt.Scan(&cariNama)
	fmt.Print("====", "Silahkan masukkan nama pemain kedua! ")
	fmt.Scan(&cariNama1)
	for i := 1; i < nextID; i++ {
		if cariNama == dataPemain[i].NamaPemain {
			pemain1 = dataPemain[i]
		}
		if cariNama1 == dataPemain[i].NamaPemain {
			pemain2 = dataPemain[i]
		}
		if pemain1.Value > pemain2.Value {
			fmt.Println("Pemain", pemain1.NamaPemain, "memiliki value yang lebih besar dari pada pemain", pemain2.NamaPemain)
		} else if pemain1.Value < pemain2.Value {
			fmt.Println("Pemain", pemain2.NamaPemain, "memiliki value yang lebih besar dari pada pemain", pemain1.NamaPemain)
		} else {
			fmt.Println("Pemain", pemain1.NamaPemain, "memiliki value yang sama dengan", pemain2.NamaPemain)
		}
	}
	fmt.Println("====================================================================================")
	tampilan()
}
