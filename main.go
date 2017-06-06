package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	//Term table creation.
	termSQL = "create table if not exists terms(id serial primary key, description text, code text, starttime text, endtime text, current text)"

	//Course table creation.
	courseSQL = "create table if not exists courses(id serial primary key, crn text, waitlistpos text, registrationstatus text, registrationdescription text, departmentcode text, departmentdescription text, coursetitle text, coursedescription text, termcode text, subjectcode text, subjectnumber text, credit text, section text)"

	//Meeting table creation.
	meetingSQL = "create table if not exists meetings(id serial primary key, crn text, startdate text, enddate text, starttime text, endtime text, coursetype text, coursetypecode text, buildingroom text, campus text, meetday text)"

	//meetingSQL = "create table if not exists meeting(id serial primary key, crn text, startdate text, enddate text, starttime text, endtime text, coursetype text, coursetypecode text, buildingroom text, campus text, meetdays text, starthour text, startminutes text, startmonth text, startyear text, startdayofmonth text, startdayofweek text, startweekofmonth text, endhour text, endminutes text, endmonth text, endyear text, enddayofmonth text, enddayofweek text, endweekofmonth text)"
	//Instructor table creation.
	instructorSQL = "create table if not exists instructors(id serial primary key, crn text, firstname text, lastname text, office text, email text)"

	//Grade table creation.
	gradeSQL = "create table if not exists grades(id serial primary key, credit text, grade text, crn text)"

	//Overall credit and grade table creation.
	gpaSQL = "create table if not exists gpa(id serial primary key, level text, credits text, gpa text)"
)

// Configurations for the database
type config struct {
	URL      string
	Username string
	Password string
	Dbname   string
}

// Term represents the current time period that will be displayed (e.g Winter Semester 2018, Fall Semester 1989).
type Term struct {
	ID          int
	Description string `json:"description"`
	Code        string `json:"code"`
	Start       string `json:"start"`
	End         string `json:"end"`
	Current     string `json:"current"`
}

// Course represents the class that is assosciated with a given term (e.g Introduction to Golang).
type Course struct {
	ID                            int
	Crn                           string       `json:"crn"`
	WaitlistPos                   string       `json:"waitlistPost"`
	RegistrationStatus            string       `json:"registrationStatus"`
	RegistrationStatusDescription string       `json:"registrationStatusDescription"`
	DepartmentCode                string       `json:"departmentCode"`
	DepartmentDescription         string       `json:"departmentDescription"`
	CourseTitle                   string       `json:"courseTitle"`
	CourseDescription             string       `json:"courseDescription"`
	TermCode                      string       `json:"termCode"`
	SubjectCode                   string       `json:"subjectCode"`
	SubjectNumber                 string       `json:"subjectNumber"`
	Section                       string       `json:"section"`
	Credit                        string       `json:"credit"`
	Meetings                      []Meeting    `json:"meetings"`
	Instructors                   []Instructor `json:"instructors"`
	Grade                         Grade        `json:"grade"`
}

// lanStrings represents internationalization
type languageStrings struct {
	Section       string `json:"section"`
	CRN           string `json:"crn"`
	Credits       string `json:"credits"`
	CourseDetails string `json:"courseDetails"`
	CourseTitle   string `json:"courseTitle"`
	Department    string `json:"department"`
	Grade         string `json:"grade"`
	Description   string `json:"description"`
	Close         string `json:"close"`
	Courses       string `json:"courses"`
	Calendar      string `json:"calendar"`
	Grades        string `json:"grades"`
}

// Meeting represents the times and locations a course will gather (e.g. Monday at 1:47 PM).
type Meeting struct {
	ID             int
	Crn            string `json:"crn"`
	StartDate      string `json:"startDate"`
	EndDate        string `json:"endDate"`
	StartTime      string `json:"startTime"`
	EndTime        string `json:"endTime"`
	CourseType     string `json:"courseType"`
	CourseTypeCode string `json:"courseTypeCode"`
	BuildingRoom   string `json:"buildingRoom"`
	Campus         string `json:"campus"`
	MeetDays       string `json:"meetDays"`
}

//Instructor represents the person(s) who will teach a course.
type Instructor struct {
	ID        int
	Crn       string `json:"crn"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Office    string `json:"office"`
	Email     string `json:"email"`
}

// Grade represents the score the student received for a course.
type Grade struct {
	ID     int
	Credit string `json:"credit"`
	Grade  string `json:"grade"`
	Crn    string `json:"crn"`
}

//GPA represents the overall credits and grades a student has
type GPA struct {
	ID      int
	Level   string `json:"level"`
	Credits string `json:"credits"`
	GPA     string `json:"gpa"`
}

// Student represents the academic information about a student
type Student struct {
	ClassStanding        string `json:"classStanding"`
	Major1               string `json:"major1"`
	DegreeType           string `json:"degreeType"`
	College              string `json:"college"`
	Level                string `json:"level"`
	Major1concentration1 string `json:"major1concentration1"`
	Major1concentration2 string `json:"major1concentration2"`
	Major1concentration3 string `json:"major1concentration3"`
	Major2               string `json:"major2"`
	Major2Department     string `json:"major2Department"`
	Major2concentration1 string `json:"major2concentration1"`
	Major2concentration2 string `json:"major2concentration2"`
	Major2concentration3 string `json:"major2concentration3"`
	Minor1               string `json:"minor1"`
	Minor2               string `json:"minor2"`
}

// Person represents the personal information about a student
type Person struct {
	Address       string `json:"address"`
	Email         string `json:"email"`
	Gid           string `json:"gid"`
	LegalName     string `json:"legalName"`
	PhoneNumber   string `json:"phoneNumber"`
	Pidm          string `json:"pidm"`
	PrefFirstName string `json:"prefFirstName"`
}

func lang(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	vars := mux.Vars(r)
	var data languageStrings
	switch vars["lng"] {
	case "ar":
		data =
			languageStrings{
				"الجزء",
				"CRN",
				"قروض",
				"تفاصيل الدورة",
				"عنوان الدورة",
				"قسم",
				"درجة",
				"وصف",
				"أغلق",
				"الدورات",
				"التقويم",
				"درجات",
			}
	case "de":
		data = languageStrings{
			"Abschnitt",
			"CRN",
			"Gutschriften",
			"Kursdetails",
			"Kursname",
			"Abteilung",
			"Klasse",
			"Beshreibung",
			"Schließsen",
			"Kurse",
			"Kalender",
			"Noten",
		}
	case "en":
		data = languageStrings{
			"Section",
			"CRN",
			"Credits",
			"Course Details",
			"Course Title",
			"Department",
			"Grade",
			"Description",
			"Close",
			"Courses",
			"Calendar",
			"Grades",
		}
	case "en-US":
		data = languageStrings{
			"Section",
			"CRN",
			"Credits",
			"Course Details",
			"Course Title",
			"Department",
			"Grade",
			"Description",
			"Close",
			"Courses",
			"Calendar",
			"Grades",
		}
	case "sp":
		data = languageStrings{
			"Sección",
			"CRN",
			"Créditos",
			"Detalles del courso",
			"Título del curso",
			"Departmento",
			"Grado",
			"Descripción",
			"Conclur",
			"Cursos",
			"Calendario",
			"Grados",
		}
	case "fr":
		data = languageStrings{
			"Section",
			"CRN",
			"Crédits",
			"Course Détails",
			"Titre de cours",
			"Départment",
			"Qualité",
			"La description",
			"Conclure",
			"Cours",
			"Calendrier",
			"Les notes",
		}
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

func person(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	data := Person{
		"318 Meadow Brook Rd, Rochester, MI 48309",
		"grizz@oakland.edu",
		"G00000000",
		"Grizz OU",
		"(248) 370-2100",
		"111111",
		"Grizz",
	}

	person := struct {
		Person Person `json:"person"`
	}{
		data,
	}

	if err := json.NewEncoder(w).Encode(person); err != nil {
		panic(err)
	}
}

func mydetails(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var students []Student
	data := Student{"Senior", "computer science", "Bach of Sci", "School CS", "Undergrad", "", "", "", "", "", "", "", "", "", ""}
	students = append(students, data)

	student := struct {
		Student []Student `json:"studentDetails"`
	}{
		students,
	}

	if err := json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}
}

func credits(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var credits []GPA
	data := GPA{2, "Undergraduate", "88", "3.2"}
	credits = append(credits, data)

	student := struct {
		GPA []GPA `json:"gpa"`
	}{
		credits,
	}

	if err := json.NewEncoder(w).Encode(student); err != nil {
		panic(err)
	}
}

func courses(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var courses []Course

	s := struct {
		Term string `json:"code"`
	}{}

	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	termCode := s.Term

	rows, err := db.Query("select * from courses where termcode = $1", termCode)
	if err != nil {
		fmt.Printf("Failed on courses for termcode %s\n", termCode)
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var c Course
		if err := rows.Scan(&c.ID, &c.Crn, &c.WaitlistPos, &c.RegistrationStatus, &c.RegistrationStatusDescription, &c.DepartmentCode, &c.DepartmentDescription, &c.CourseTitle, &c.CourseDescription, &c.TermCode, &c.SubjectCode, &c.SubjectNumber, &c.Credit, &c.Section); err != nil {
			fmt.Println("error on coruses")
			panic(err)
		} else {
			instructorRows, err := db.Query("select * from instructors where crn = $1", c.Crn)
			if err != nil {
				fmt.Printf("Failed on instructors for crn %s\n", c.Crn)
				panic(err)
			}
			defer instructorRows.Close()

			for instructorRows.Next() {
				var i Instructor
				if err := instructorRows.Scan(&i.ID, &i.Crn, &i.FirstName, &i.LastName, &i.Office, &i.Email); err != nil {
					fmt.Println("error on instructors")
					panic(err)
				} else {
					c.Instructors = append(c.Instructors, i)
				}
			}

			meetingRows, err := db.Query("select * from meetings where crn = $1", c.Crn)
			if err != nil {
				fmt.Printf("Failed on meetings for crn %s\n", c.Crn)
				panic(err)
			}
			defer meetingRows.Close()

			for meetingRows.Next() {
				var m Meeting
				if err := meetingRows.Scan(&m.ID, &m.Crn, &m.StartDate, &m.EndDate, &m.StartTime, &m.EndTime, &m.CourseType, &m.CourseTypeCode, &m.BuildingRoom, &m.Campus, &m.MeetDays); err != nil {
					fmt.Println("error on meetings")
					panic(err)
				} else {
					c.Meetings = append(c.Meetings, m)
				}
			}

			var g Grade
			err = db.QueryRow("select * from grades where crn = $1", c.Crn).Scan(&g.ID, &g.Credit, &g.Grade, &g.Crn)
			if err != nil {
				fmt.Printf("Failed on grades for crn %s\n", c.Crn)
				panic(err)
			}

			c.Grade = g
			courses = append(courses, c)
		}
	}

	course := struct {
		Course []Course `json:"courses"`
	}{
		courses,
	}

	if err := json.NewEncoder(w).Encode(course); err != nil {
		panic(err)
	}
}

func terms(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var terms []Term
	rows, err := db.Query("select * from terms")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var t Term
		if err := rows.Scan(&t.ID, &t.Description, &t.Code, &t.Start, &t.End, &t.Current); err != nil {
			fmt.Println("error on coruses")
			panic(err)
		}
		terms = append(terms, t)

	}

	term := struct {
		Term []Term `json:"terms"`
	}{
		terms,
	}

	if err := json.NewEncoder(w).Encode(term); err != nil {
		panic(err)
	}
}

func init() {
	var c config
	file, err := os.Open("database.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&c)
	if err != nil {
		panic(err)
	}

	fmt.Println(c)
	dbURL := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", c.Username, c.Password, c.URL, c.Dbname)

	db, err = sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(termSQL)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(courseSQL)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(meetingSQL)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(instructorSQL)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(gradeSQL)
	if err != nil {
		panic(err)
	}

	_, err = db.Query(gpaSQL)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/locales/{lng}/{ns}", lang)
	r.HandleFunc("/api/person", person)
	r.HandleFunc("/api/mydetails", mydetails)
	r.HandleFunc("/api/courses", courses)
	r.HandleFunc("/api/terms", terms)
  r.HandleFunc("/api/credits", credits)
	http.Handle("/", r)
	http.ListenAndServe(":8082", nil)
}
