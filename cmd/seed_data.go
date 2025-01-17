// seed_data.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"my-gin-mongo/config" // pastikan path import sesuai

	"google.golang.org/api/iterator"
)

// Jalankan: go run seed_data.go
func main() {
	// Inisialisasi Firebase dan Firestore
	if err := config.InitFirestore(); err != nil {
		log.Fatalf("Gagal menginisialisasi Firestore: %v", err)
	}

	// Jalankan fungsi-fungsi seeding
	seedProfile()
	seedEducation()
	seedExperiences()
	seedCertifications()
	seedHonors()

	fmt.Println("=== SEEDING DATA SELESAI ===")
}

// ---------------------------------------------------------------------
// 1) SEED PROFILE
// ---------------------------------------------------------------------
func seedProfile() {
	coll := config.GetCollection("profiles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cek apakah sudah ada data Profile
	iter := coll.Limit(1).Documents(ctx)
	defer iter.Stop()
	_, err := iter.Next()
	if err == nil {
		log.Println("[seedProfile] Data profil sudah ada, skip insert.")
		return
	} else if err != iterator.Done {
		log.Printf("[seedProfile] Error saat cek data: %v", err)
		return
	}

	// Data profil (sesuaikan dengan kebutuhan)
	profile := map[string]interface{}{
		"full_name":           "Muhammad Fuad Fakhruzzaki",
		"headline":            "Software Engineer",
		"about":               "With a strong passion for technology and innovation, I thrive on exploring new advancements...",
		"profile_picture_url": "https://example.com/my-profile-pic.png", // Ganti jika perlu
		"cv_url":              "https://example.com/my-cv.pdf",            // Ganti jika perlu
		"updated_at":          time.Now().Format(time.RFC3339),
		"location":            "Semarang, Indonesia",
		"contact_number":      "081392302787",
		"contact_email":       "mfuadfakhruzzaki@students.undip.ac.id",
		"website":             "fuadfakhruz.id",
	}

	// Tambah data ke Firestore
	_, _, err = coll.Add(ctx, profile)
	if err != nil {
		log.Printf("[seedProfile] Gagal insert profil: %v", err)
		return
	}

	log.Println("[seedProfile] Berhasil insert Profile.")
}

// ---------------------------------------------------------------------
// 2) SEED EDUCATION
// ---------------------------------------------------------------------
func seedEducation() {
	coll := config.GetCollection("educations")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cek apakah data pendidikan dari "Diponegoro University" sudah ada
	q := coll.Where("institution", "==", "Diponegoro University").Limit(1)
	iter := q.Documents(ctx)
	defer iter.Stop()
	_, err := iter.Next()
	if err == nil {
		log.Println("[seedEducation] Data pendidikan dari Diponegoro University sudah ada, skip.")
		return
	} else if err != iterator.Done {
		log.Printf("[seedEducation] Error saat cek data: %v", err)
		return
	}

	edu := map[string]interface{}{
		"institution":    "Diponegoro University",
		"degree":         "Computer Engineering / GPA: 3.51",
		"field_of_study": "Computer Engineering",
		"location":       "Semarang, Indonesia",
		"start_date":     "August 2024",
		"end_date":       "Present",
		"description": `At Diponegoro University, majoring in Computer Engineering, I am building a solid foundation in software development,
cloud computing, and artificial intelligence. I aim to contribute to innovative projects that solve real-world problems.`,
	}

	_, _, err = coll.Add(ctx, edu)
	if err != nil {
		log.Printf("[seedEducation] Gagal insert pendidikan: %v", err)
		return
	}

	log.Println("[seedEducation] Berhasil insert pendidikan.")
}

// ---------------------------------------------------------------------
// 3) SEED EXPERIENCES (Work Experience + Related Experiences)
// ---------------------------------------------------------------------
func seedExperiences() {
	coll := config.GetCollection("experiences")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Contoh data Work Experience: Telkomsel
	telkomsel := map[string]interface{}{
		"title":          "Intern IT Ops Departemen Jateng dan DIY",
		"company":        "Telkomsel",
		"location":       "Semarang, Indonesia",
		"start_date":     "July 2024",
		"end_date":       "August 2024",
		"description":    `During my internship at Telkomsel in the IT Operations Department, I gained experience using Google Cloud Platform services...`,
		"experience_type": "Work Experience",
		"date_range":      "July 2024 - August 2024",
	}

	// Related Experiences:
	bem := map[string]interface{}{
		"title":          "Junior Staff in the Media and Information Office",
		"company":        "Student Executive Board Of The Engineering Faculty (BEM FT)",
		"location":       "Diponegoro University, Semarang, Indonesia",
		"start_date":     "April 2023",
		"end_date":       "February 2024",
		"description":    `I honed my skills in Adobe Illustrator, video production, and media management.`,
		"experience_type": "Related Experience",
		"date_range":      "April 2023 - February 2024",
	}
	senate := map[string]interface{}{
		"title":          "Staff Member in the Event Division for the 2022 Faculty Election",
		"company":        "Senate Of The Faculty Of Engineering",
		"location":       "Diponegoro University, Semarang, Indonesia",
		"start_date":     "November 2022",
		"end_date":       "February 2023",
		"description":    `I developed strong project management and event planning skills.`,
		"experience_type": "Related Experience",
		"date_range":      "November 2022 - February 2023",
	}
	himaskom1 := map[string]interface{}{
		"title":          "Junior Staff in the Student Resource Development Division",
		"company":        "Himpunan Mahasiswa Teknik Komputer (HIMASKOM)",
		"location":       "Diponegoro University, Semarang, Indonesia",
		"start_date":     "October 2023",
		"end_date":       "April 2024",
		"description":    "Organized training programs and seminars to improve leadership and teamwork skills.",
		"experience_type": "Related Experience",
		"date_range":      "October 2023 - April 2024",
	}
	himaskom2 := map[string]interface{}{
		"title":          "PIC Media Partner for The ACE 2024",
		"company":        "Himpunan Mahasiswa Teknik Komputer (HIMASKOM)",
		"location":       "Diponegoro University, Semarang, Indonesia",
		"start_date":     "August 2024",
		"end_date":       "October 2024",
		"description":    "Coordinated partnerships with media outlets to promote events.",
		"experience_type": "Related Experience",
		"date_range":      "August 2024 - October 2024",
	}
	himaskom3 := map[string]interface{}{
		"title":          "Head of the Student Resource Development Division",
		"company":        "Himpunan Mahasiswa Teknik Komputer (HIMASKOM)",
		"location":       "Diponegoro University, Semarang, Indonesia",
		"start_date":     "May 2024",
		"end_date":       "Present",
		"description":    "Designed and managed self-development programs for students.",
		"experience_type": "Related Experience",
		"date_range":      "May 2024 - Present",
	}

	data := []interface{}{telkomsel, bem, senate, himaskom1, himaskom2, himaskom3}

	// Tambahkan data experiences secara bersamaan
	_, _, err := coll.Add(ctx, map[string]interface{}{
		// Untuk mempermudah, kita masukkan masing-masing experience sebagai dokumen terpisah.
		// Jika ingin memasukkan sekaligus, Anda perlu mengulang Add() di dalam loop.
		"experiences": data,
	})
	if err != nil {
		log.Printf("[seedExperiences] Gagal menambahkan data experiences: %v", err)
		return
	}

	log.Printf("[seedExperiences] Berhasil memasukkan data experiences.")
}

// ---------------------------------------------------------------------
// 4) SEED CERTIFICATIONS
// ---------------------------------------------------------------------
func seedCertifications() {
	coll := config.GetCollection("certifications")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cek apakah sertifikasi sudah ada dengan memeriksa salah satu dokumen
	q := coll.Where("name", "==", "CCNA: Switching, Routing, and Wireless Essentials (CISCO)").Limit(1)
	iter := q.Documents(ctx)
	defer iter.Stop()
	_, err := iter.Next()
	if err == nil {
		log.Println("[seedCertifications] Data sertifikasi sudah ada, skip.")
		return
	} else if err != iterator.Done {
		log.Printf("[seedCertifications] Error saat cek data: %v", err)
		return
	}

	certs := []interface{}{
		map[string]interface{}{
			"name":                 "CCNA: Switching, Routing, and Wireless Essentials (CISCO)",
			"issuing_organization": "CISCO",
		},
		map[string]interface{}{
			"name":                 "AI Fundamentals with IBM SkillsBuild (CISCO)",
			"issuing_organization": "CISCO",
		},
		map[string]interface{}{
			"name":                 "Artificial Intelligence Fundamentals (IBM)",
			"issuing_organization": "IBM",
		},
		map[string]interface{}{
			"name":                 "Data Analytics Essentials (CISCO)",
			"issuing_organization": "CISCO",
		},
		map[string]interface{}{
			"name":                 "Introduction to Cybersecurity (CISCO)",
			"issuing_organization": "CISCO",
		},
		map[string]interface{}{
			"name":                 "Database Design (ORACLE ACADEMY)",
			"issuing_organization": "ORACLE ACADEMY",
		},
		map[string]interface{}{
			"name":                 "Database Foundation (ORACLE ACADEMY)",
			"issuing_organization": "ORACLE ACADEMY",
		},
		map[string]interface{}{
			"name":                 "CCNA: Introduction to Networks (CISCO)",
			"issuing_organization": "CISCO",
		},
		map[string]interface{}{
			"name":                 "Introduction to IoT (CISCO)",
			"issuing_organization": "CISCO",
		},
		map[string]interface{}{
			"name":                 "IT Essentials (CISCO)",
			"issuing_organization": "CISCO",
		},
	}

	// Karena Firestore tidak memiliki operasi InsertMany secara langsung,
	// kita gunakan loop untuk memasukkan setiap sertifikasi.
	for _, cert := range certs {
		_, _, err := coll.Add(ctx, cert)
		if err != nil {
			log.Printf("[seedCertifications] Gagal menambahkan sertifikasi: %v", err)
		}
	}

	log.Println("[seedCertifications] Berhasil memasukkan data sertifikasi.")
}

// ---------------------------------------------------------------------
// 5) SEED HONORS / AWARDS
// ---------------------------------------------------------------------
func seedHonors() {
	coll := config.GetCollection("honors")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Cek apakah award tertentu sudah ada
	q := coll.Where("title", "==", "Bakti BCA Scholarship Awardee 2024").Limit(1)
	iter := q.Documents(ctx)
	defer iter.Stop()
	_, err := iter.Next()
	if err == nil {
		log.Println("[seedHonors] 'Bakti BCA Scholarship Awardee 2024' sudah ada, skip.")
		return
	} else if err != iterator.Done {
		log.Printf("[seedHonors] Error saat cek data: %v", err)
		return
	}

	honor := map[string]interface{}{
		"title":        "Bakti BCA Scholarship Awardee 2024",
		"issuer":       "Bank BCA",
		"date_awarded": "2024",
		"description":  "Scholarship awarded to outstanding students.",
	}

	_, _, err = coll.Add(ctx, honor)
	if err != nil {
		log.Printf("[seedHonors] Gagal memasukkan award: %v", err)
		return
	}
	log.Println("[seedHonors] Berhasil memasukkan data award.")
}
