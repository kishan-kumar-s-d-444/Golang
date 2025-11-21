package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

//Model for course - file

type Course struct {
	CourseId    string  `json:"courseid"`
	CourseName  string  `json:"coursename"`
	CoursePrice int     `json:"price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// middleware , helper -file
func (c *Course) IsEmpty() bool {
	//return c.CourseId == "" && c.CourseName == ""
	return c.CourseName == ""
}

func main() {
	fmt.Println("API")
	r:=mux.NewRouter()

	//seeding
	courses = append(courses, Course{CourseId: "2",CourseName: "React",CoursePrice:299,Author: &Author{Fullname:"kishan",Website:".com"} })
	courses = append(courses, Course{CourseId: "4",CourseName: "Mern",CoursePrice:199,Author: &Author{Fullname:"kishan",Website:".gov"} })

	//routing
	r.HandleFunc("/",serveHome).Methods("GET")
	r.HandleFunc("/courses",getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}",getOneCourse).Methods("GET")
	r.HandleFunc("/course",createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}",updateOneCourse).Methods("PUT")
	r.HandleFunc("/courses/{id}",deleteOneCourse).Methods("DELETE")



	//listen to port
	log.Fatal(http.ListenAndServe(":4000",r))

}

//controllers - file

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>API</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all Courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)

}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	//grab id from request
	params := mux.Vars(r)

	//loop through courses , find matching id and return response

	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("No course found")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	//what if : Body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Send data")
	}
	// if data is like {}
	var course Course
	_ = json.NewDecoder(r.Body).Decode(&course)
	if course.IsEmpty() {
		json.NewEncoder(w).Encode("Send data")
		return
	}

	//generate a unique id, convert them to string
	//append new course into Courses
	rand.Seed(time.Now().UnixNano())
	course.CourseId = strconv.Itoa(rand.Intn(100))
	courses = append(courses, course)
	json.NewEncoder(w).Encode(course)
	return
}

func updateOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("update one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params:=mux.Vars(r)

	//loop through values once found id remove from that and again add that with updated value

	for ind,course :=range courses {
		if course.CourseId==params["id"]{
			courses =append(courses[:ind],courses[ind+1:]...)
			var course Course
			_=json.NewDecoder(r.Body).Decode(&course)
			course.CourseId=params["id"]
			courses=append(courses,course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	// TO DO :send a response when id is not found
}

func deleteOneCourse(w http.ResponseWriter,r *http.Request){
	fmt.Println("delete one course")
	w.Header().Set("Content-Type", "application/json")

	//first - grab id from req
	params:=mux.Vars(r)

	//loop through values once found id remove from that
	for ind,course :=range courses {
		if course.CourseId==params["id"]{
			courses =append(courses[:ind],courses[ind+1:]...)
			break
		}
	}
}