package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Contributions struct {
	Data struct {
		Viewer struct {
			contributionCollection struct {
				ContributionCalendar struct {
					TotalContributions int `json:"totalContributions"`
				}
			}
		}
	}
}

func contributionStats(c *gin.Context) {
	req, err := http.NewRequest("POST", "https://api.github.com/graphql", nil)
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

	var contributions Contributions
	err = json.Unmarshal(body, &contributions)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erreur lors du parsing de la réponse de l'API GitHub"})
		return
	}

	c.JSON(http.StatusOK, contributions)
}
