package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"text/template"
)

type Repo struct {
	Name     string `json:"name"`
	Language string `json:"language"`
	Owner    struct {
		Login string `json:"login"`
	}
}

type Value struct {
	Couleur     string
	Pourcentage int
}

var accessToken = "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"
var languageCouleur = map[string]string{
	"Go":         "#00ADD8",
	"Python":     "#3572A5",
	"JavaScript": "#F1E05A",
	"HTML":       "#E34C26",
	"CSS":        "#563D7C",
	"Shell":      "#89E051",
	"Java":       "#B07219",
	"PHP":        "#4F5D95",
	"C":          "#555555",
	"C++":        "#f34b7d",
	"C#":         "#178600",
	"TypeScript": "#2b7489",
	"Ruby":       "#701516",
	"Rust":       "#dea584",
	"Kotlin":     "#F18E33",
	"Swift":      "#ffac45",
	"Svelte":     "#ff3e00",
}

func langageStats(c *gin.Context) {
	req, err := http.NewRequest("GET", "https://api.github.com/user/repos", nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la création de la requête"})
		return
	}
	req.Header.Set("Authorization", "token "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de l'envoi de la requête à l'API GitHub"})
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors de la lecture de la réponse de l'API GitHub"})
		return
	}

	var repos []Repo
	err = json.Unmarshal(body, &repos)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du parsing de la réponse de l'API GitHub"})
		return
	}

	result := make(map[string]int)
	var totalRepo int
	languagePourcentage := make(map[string]Value)

	for _, repo := range repos {
		if repo.Owner.Login == "GoRoutine" || repo.Language == "" {
			continue
		}
		result[repo.Language] += 1
		totalRepo++
	}

	for key, value := range result {
		languagePourcentage[key] = Value{
			Couleur:     languageCouleur[key],
			Pourcentage: value * 100 / totalRepo,
		}
	}
	c.JSON(http.StatusOK, languagePourcentage)

	t, err := template.New("language.svg").ParseFiles("language.svg")
	if err != nil {
		panic(err)
	}

	file, err := os.Create("language2.svg")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = t.Execute(file, languagePourcentage)
	if err != nil {
		panic(err)
	}
}
