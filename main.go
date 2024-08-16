package main

import (
	"bufio"
	"fmt"
	"hangman/Package"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var word string

func main() {
	http.Handle("/videos/", http.StripPrefix("/videos", http.FileServer(http.Dir("./videos"))))
	http.Handle("/images/", http.StripPrefix("/images", http.FileServer(http.Dir("./images"))))
	http.Handle("/audio/", http.StripPrefix("/audio", http.FileServer(http.Dir("./audio"))))
	http.Handle("/JS/", http.StripPrefix("/JS", http.FileServer(http.Dir("./JS"))))
	http.Handle("/CSS/", http.StripPrefix("/CSS", http.FileServer(http.Dir("./CSS"))))
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	http.HandleFunc("/won", Won)
	http.HandleFunc("/lost", Lost)
	http.HandleFunc("/sport.html", sport)
	http.HandleFunc("/flags.html", flags)
	http.HandleFunc("/food.html", food)
	http.HandleFunc("/cars.html", cars)
	http.HandleFunc("/review.html", review)
	http.HandleFunc("/theme.html", theme)
	http.HandleFunc("/gameplay.html", Gameplay)
	http.HandleFunc("/", home)
	fmt.Println("Serving @ : ", "http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func Mottheme(theme string) string {
	file, err := os.Open(theme) //ouvre le  theme  et stocke le fichier dans la variable "file".Et s'il y'a une err elle est stockée dans la variable "err".
	if err != nil {
		fmt.Println("Erreur dans l'ouverture du fichier de la liste des mots :", err)
		return ""
	}
	defer file.Close() //le fichier sera fermé à la fin de l'exécution de la fonction

	var words []string
	scanner := bufio.NewScanner(file) //scanner sera utilisé pour parcourir le fichier ligne par ligne.
	for scanner.Scan() {              // parcourt le fichier
		words = append(words, scanner.Text()) //ajoute chaque ligne  dans le tableau "words" à l'aide de la fonction append().
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Erreur de lecture du fichier de la liste des mots :", err)
		return ""
	}

	return words[rand.Intn(len(words)-1)] // renvoie un mot aléatoire de la liste  "words".
}

func Won(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/won.html")) //utilise template.ParseFiles() pour lire le contenu du fichier
	tmpl.Execute(rw, nil)                                              //exécute le template stocké dans  "tmpl" en passant l'objet http.ResponseWriter "rw" en tant que sortie de la réponse HTTP
}
func Lost(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/lost.html"))
	tmpl.Execute(rw, nil)
}

func theme(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/theme.html"))
	if r.Method == "POST" { //vérifie si la méthode de la requête HTTP est POST
		body, err := ioutil.ReadAll(r.Body) // lit la requête HTTP à l'aide de ioutil.ReadAll()
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}
		letter := string(body)                                                                //converti le contenu de la requête HTTP en une chaîne de caractères,
		fmt.Println("Letter received:", letter)                                               //affichent cette chaîne sur la console
		fmt.Fprintf(w, "Letter received: %s\n", letter+" | "+time.Now().Format(time.RFC3339)) //renvoient une réponse HTTP avec la chaîne convertie

		// crée une chaînes de caractères contenant les chemins des fichiers de texte contenant les mots pour chaque thème.
		textFiles := []string{"./themes/flags.txt", "./themes/cars.txt", "./themes/sports.txt", "./themes/food.txt"}

		//vérifie la lettre envoyée dans la requête HTTP
		//attribue un thème et un fichier de mots à "word" en fonction de cette lettre
		if letter == "F" {
			theme := 1
			word = Mottheme(textFiles[theme-1])
			fmt.Println(word)
		} else if letter == "C" {
			theme := 2
			word = Mottheme(textFiles[theme-1])
			fmt.Println(word)
		} else if letter == "S" {
			theme := 3
			word = Mottheme(textFiles[theme-1])
			fmt.Println(word)
		} else {
			theme := 4
			word = Mottheme(textFiles[theme-1])
			fmt.Println(word)
		}
	}
	//exécute tmpl et écrire sa sortie dans w. Cela permet de renvoyer une réponse HTTP
	tmpl.Execute(w, nil)
}

type GameResponse struct {
	Hangman Package.GameState
	Message string
}

func sport(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/sport.html")) //utilise template.ParseFiles() pour lire le contenu du fichier
	tmpl.Execute(rw, nil)                                                //exécute le template stocké dans  "tmpl" en passant l'objet http.ResponseWriter "rw" en tant que sortie de la réponse HTTP
}

func flags(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/flags.html"))
	tmpl.Execute(rw, nil)
}

func food(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/food.html"))
	tmpl.Execute(rw, nil)
}

func cars(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/cars.html"))
	tmpl.Execute(rw, nil)
}

func review(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/review.html"))
	tmpl.Execute(rw, nil)
}

func home(rw http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/home.html"))
	tmpl.Execute(rw, nil)
}

func Gameplay(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("./templates/gameplay.html"))
	if r.Method == "GET" { // verifie si la requête HTTP est "GET"
		if Package.Won() || Package.Lost() { //vérifie si le joueur a gagné ou perdu
			Package.GameStart(word) //Package.GameStart(word) est appelée pour recommencer le jeu avec un nouveau mot à deviner.
		}
	}
	if r.Method == "POST" { //verifie si la méthode de la requête HTTP est "POST"
		letter := r.FormValue("letter")              //récupère la lettre soumise par l'utilisateur en utilisant  r.FormValue("letter")
		Package.GuessLetter(strings.ToUpper(letter)) //lettre est ensuite est mise dans Package.GuessLetter() pour qu'elle soit utilisée dans le jeu.

		//vérifie si le joueur a gagné ou perdu. Elle redirige l'utilisateur vers les pages "won" ou "lost" en utilisant http.Redirect().
		if Package.Won() {
			http.Redirect(w, r, "/won", http.StatusFound)
		}
		if Package.Lost() {
			http.Redirect(w, r, "/lost", http.StatusFound)
		}
	}
	//exécuter tmpl et écrit sa sortie dans w, en utilisant l'état actuel du jeu stocké dans Package.Hangman
	tmpl.Execute(w, Package.Hangman)
}
