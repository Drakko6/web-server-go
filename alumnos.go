package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)


var Materias map[string]map[string]int64
var Alumnos map[string]map[string]int64

var misAlumnos AdminAlumnos

func cargarHtml(a string) string {
	html, _ := ioutil.ReadFile(a)

	return string(html)
}

func form(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		cargarHtml("form.html"),
	)
}

type Alumno struct {
	Nombre string
	Materia string
	Calificacion int64
}

type AdminAlumnos struct {

}


func (alumnos *AdminAlumnos) Agregar(alumno Alumno) {
	// alumnos.Alumnos = append(alumnos.Alumnos, alumno)

	// creación de un alumno - calificacion
	alumnoNuevo := make(map[string]int64)
	alumnoNuevo[alumno.Nombre] = alumno.Calificacion

	var existe bool
	//comprobar si existe la materia
	for mat := range Materias {
		if mat == alumno.Materia {
			Materias[alumno.Materia][alumno.Nombre] = alumno.Calificacion
			
			existe = true
			break
		}
	}
	if !existe{
		
			Materias[alumno.Materia] = alumnoNuevo
	}


	// creacion materia - calificacion  
	materiaNueva := make(map[string]int64)
	materiaNueva[alumno.Materia] = alumno.Calificacion

	// creación de un alumno
	var existe2 bool
	//comprobar si existe la materia
	for alum := range Alumnos {
		if alum == alumno.Nombre {
			Alumnos[alumno.Nombre][alumno.Materia] = alumno.Calificacion
			existe2 = true
			break
		}
	}
	if !existe2{
		
			Alumnos[alumno.Nombre] = materiaNueva
	}


}

func (alumnos * AdminAlumnos) ObtenerPromedio(nombre string) string{

	var suma int64
	var cont int64
	for _, calificacion := range Alumnos[nombre] {
		
		suma += calificacion
		cont++
	}


 return strconv.Itoa(int(suma/cont))


}


func (alumnos *AdminAlumnos) String() string {
	var html string

	for alumno := range Alumnos {

		for materia, calificacion := range Alumnos[alumno] {
		
			html += "<tr>" +
			"<td>" + alumno + "</td>" +
			"<td>" + materia + "</td>" +
			"<td>" + strconv.Itoa(int(calificacion)) + "</td>" +

			"</tr>"
			
		}
		
	}


	return html
}


func promedioAlumno(res http.ResponseWriter, req *http.Request) {

	switch req.Method {

		case "GET":
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("formPromedio.html"),
				misAlumnos.FormAlumnos(),
			)

		
		case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		// fmt.Println(req.PostForm)

		promedio := misAlumnos.ObtenerPromedio(req.FormValue("nombre"))

		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedio.html"),
			req.FormValue("nombre"),
			promedio,
		)


	}

}

func (alumnos *AdminAlumnos) FormAlumnos() string {
	var html string

	for alumno := range  Alumnos {

		html += "<option value='"+ alumno +"'>"+ alumno + "</option>"
		
	}
	return html

}


func alumnos(res http.ResponseWriter, req *http.Request) {
	// fmt.Println(req.Method)
	switch req.Method {
	case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		// fmt.Println(req.PostForm)

		n, _ := strconv.ParseInt(req.FormValue("calificacion"), 10, 64)

		alumno := Alumno{Nombre: req.FormValue("alumno"),
		 Materia: req.FormValue("materia"),
		  Calificacion: n }

		misAlumnos.Agregar(alumno)
		// fmt.Println(misAlumnos)
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("respuesta.html"),
			alumno.Nombre,
		)
	case "GET":
		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("tabla.html"),
			misAlumnos.String(),
		)
	}
}

func (alumnos *AdminAlumnos) ObtenerPromedioMateria(materia string) string {

	var suma int64
	var cont int64
	for _, calificacion := range Materias[materia] {
		
		suma += calificacion
		cont++
	}

	return strconv.Itoa(int(suma/cont))

}

func (alumnos *AdminAlumnos) FormMateria() string {
	var html string

	for materia := range  Materias {

		html += "<option value='"+ materia +"'>"+ materia + "</option>"
		
	}
	return html

}

func materia(res http.ResponseWriter, req *http.Request) {

	switch req.Method {

		case "GET":
			res.Header().Set(
				"Content-Type",
				"text/html",
			)
			fmt.Fprintf(
				res,
				cargarHtml("formMateria.html"),
				misAlumnos.FormMateria(),
			)

		
		case "POST":
		if err := req.ParseForm(); err != nil {
			fmt.Fprintf(res, "ParseForm() error %v", err)
			return
		}
		// fmt.Println(req.PostForm)

		promedio := misAlumnos.ObtenerPromedioMateria(req.FormValue("materia"))

		res.Header().Set(
			"Content-Type",
			"text/html",
		)
		fmt.Fprintf(
			res,
			cargarHtml("promedioMateria.html"),
			req.FormValue("materia"),
			promedio,
		)


	}

}

func (alumnos *AdminAlumnos) ObtenerPromedioGeneral() string {
	var html string

	var suma int64
	var cont int64
	for alumno := range Alumnos {

		for _, calificacion := range Alumnos[alumno] {
		
			suma += calificacion
			cont++
		}
		
	}

	html += "<a href='/'>Ir a Inicio</a> <hr> <br> <h3> El promedio general es: " + strconv.Itoa(int(suma/cont)) +"</h3>" 
	 

	return html

}

func promedioGeneral(res http.ResponseWriter, req *http.Request) {

	res.Header().Set(
		"Content-Type",
		"text/html",
	)
	fmt.Fprintf(
		res,
		misAlumnos.ObtenerPromedioGeneral(),
	
	)


}
func main() {

	Materias = make(map[string]map[string]int64)
	Alumnos = make(map[string]map[string]int64)

	http.HandleFunc("/", form)
	http.HandleFunc("/alumno", promedioAlumno)

	http.HandleFunc("/promedio-general", promedioGeneral )

	http.HandleFunc("/materia", materia )

	http.HandleFunc("/alumnos", alumnos)
	fmt.Println("Corriendo servidor de calificaciones...")
	http.ListenAndServe(":9000", nil)
}
