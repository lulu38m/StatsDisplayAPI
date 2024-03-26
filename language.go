package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"sort"
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
	Nom         string
}

var accessToken = os.Getenv("GITHUB_TOKEN")
var languageCouleur = map[string]string{
	"Go":         "#00ADD8",
	"Python":     "#3572A5",
	"JavaScript": "#F1E05A",
	"HTML":       "#E34C26",
	"CSS":        "#563D7C",
	"Shell":      "#89E051",
	"Java":       "#B07219",
	"PHP":        "#a7c957",
	"C":          "#555555",
	"C++":        "#f34b7d",
	"C#":         "#178600",
	"TypeScript": "#2b7489",
	"Ruby":       "#701516",
	"Rust":       "#dea584",
	"Kotlin":     "#F18E33",
	"Swift":      "#ffac45",
	"Svelte":     "#ff006e",
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
			Nom:         key,
		}
	}

	type kv struct {
		Key   string
		Value Value
	}
	var value []kv
	for lang, count := range languagePourcentage {
		value = append(value, kv{lang, count})
	}
	sort.Slice(value, func(i, j int) bool {
		return value[i].Value.Pourcentage > value[j].Value.Pourcentage
	})

	const maxLanguages = 3
	var totalTopLanguages int
	topLanguages := make([]Value, 0)
	otherLanguages := 0

	for i, kv := range value {
		if i < maxLanguages {
			topLanguages = append(topLanguages, kv.Value)
			totalTopLanguages += kv.Value.Pourcentage
		} else {
			otherLanguages += kv.Value.Pourcentage
		}
	}

	if len(value) > maxLanguages {
		topLanguages = append(topLanguages, Value{
			Couleur:     "#FFF",
			Pourcentage: otherLanguages,
			Nom:         "Autre",
		})
	} else {
		for _, kv := range value {
			topLanguages = append(topLanguages, kv.Value)
			totalTopLanguages += kv.Value.Pourcentage
		}
	}

	t, err := template.New("language.svg").ParseFiles("language.svg")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": " Erreur lors de la lecture du fichier language.svg"})
		return
	}

	c.Writer.Header().Set("Content-Type", "image/svg+xml")
	err = t.Execute(c.Writer, topLanguages)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": " Erreur lors de l'execution du template"})
		return
	}
}
