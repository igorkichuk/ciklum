package usecase

import (
	"github.com/igorkichuk/ciklum/internal/models"
	"github.com/igorkichuk/ciklum/internal/provider"
	"math"
)

type article struct {
	articleUrl          string
	contentMarketingUrl string
	ciklumProvider      provider.CiklumProvider
}

func NewArticleUsecase(ciklumProvider provider.CiklumProvider, articleUrl, cmUrl string) Article {
	return &article{
		articleUrl:          articleUrl,
		contentMarketingUrl: cmUrl,
		ciklumProvider:      ciklumProvider,
	}
}

func (u *article) SplitResponses() (models.CiklumResponse, error) {
	articles, err := u.ciklumProvider.GetResponse(u.articleUrl)
	if err != nil {
		return models.CiklumResponse{}, err
	}

	cm, err := u.ciklumProvider.GetResponse(u.contentMarketingUrl)
	if err != nil {
		return models.CiklumResponse{}, err
	}

	articleItems := articles.Response.Items
	cmItems := cm.Response.Items

	var splitedItems []models.Item

	for i := 0; i < len(articleItems); i++ {
		elNumber := i + 1

		if math.Mod(float64(elNumber), 6.0) == 0 && len(cmItems) > 0 {
			splitedItems = append(splitedItems, cmItems[0])
			cmItems = cmItems[1:]
			continue
		}

		if math.Mod(float64(elNumber), 6.0) == 0 && len(cmItems) <= 0 {
			item := models.Item{
				Type: "Ad",
			}
			splitedItems = append(splitedItems, item)
			continue
		}

		splitedItems = append(splitedItems, articleItems[i])
	}

	articles.Response.Items = splitedItems

	return articles, nil
}
