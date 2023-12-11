package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Repo struct {
	Name     string `json:"name"`
	Language string `json:"language"`
}

var accessToken = "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"

func getRepo(c *gin.Context) {
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

	c.JSON(http.StatusOK, repos)
}
