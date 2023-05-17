package services

import (
	"errors"
	"log"
	"path/filepath"
	"practice2sem/transactionsServer/database"
	"practice2sem/transactionsServer/models"
	"time"

	"github.com/briiC/docxplate"
)

func GetAllProviders() ([]models.Provider, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	rows, err := db.GetAllProviders()
	if err != nil {
		return nil, errors.New("failed to get providers: " + err.Error())
	}
	providersArr := make([]models.Provider, 0, 256)
	i := 0
	for rows.Next() {
		providersArr = append(providersArr, models.Provider{})
		err = rows.Scan(
			&providersArr[i].Id, &providersArr[i].Name,
			&providersArr[i].Address, &providersArr[i].Phone)
		if err != nil {
			return nil, errors.New("failed to scan providers" + err.Error())
		}
		i++
	}
	return providersArr, nil
}

func GetAllStorages() ([]models.Storage, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return nil, errors.New("failed to connect to postgresql: " + err.Error())
	}
	rows, err := db.GetAllStorages()
	if err != nil {
		return nil, errors.New("failed to get storages: " + err.Error())
	}
	storagesArr := make([]models.Storage, 0, 256)
	i := 0
	for rows.Next() {
		storagesArr = append(storagesArr, models.Storage{})
		err = rows.Scan(
			&storagesArr[i].Id, &storagesArr[i].Name,
			&storagesArr[i].Address)
		if err != nil {
			return nil, errors.New("failed to scan storages" + err.Error())
		}
		i++
	}
	return storagesArr, nil
}

// ! не работает сука
func CreateNewDelivery(deliveryRequest models.DeliveryRequest) (string, error) {
	db, err := database.GetPostgresql()
	if err != nil {
		return "", errors.New("failed to connect to postgresql: " + err.Error())
	}
	deliveryData, err := db.CreateNewDelivery(deliveryRequest)
	if err != nil {
		return "", errors.New("failed to create new delivery: " + err.Error())
	}

	tdoc, err := docxplate.OpenTemplate(filepath.Join(".", "templates", "testLibre.docx"))
	if err != nil {
		return "", errors.New("failed to open template: " + err.Error())
	}
	log.Println(tdoc.Placeholders())
	tdoc.Params(deliveryData)

	/* replaceMap := docx.PlaceholderMap{
		"ProviderName":   deliveryData.ProviderName,
		"Date":           deliveryData.Date,
		"Items":          deliveryData.Items,
		"StorageAddress": deliveryData.StorageAddress,
	}
	doc, err := docx.Open(filepath.Join(".", "templates", "test.docx"))
	if err != nil {
		return "", errors.New("failed to open template: " + err.Error())
	}
	err = doc.ReplaceAll(replaceMap)
	if err != nil {
		return "", errors.New("failed to replace placeholders: " + err.Error())
	} */
	/* template, err := docxt.OpenTemplate(filepath.Join(".", "templates", "testLibre.docx"))
	if err != nil {
		return "", errors.New("failed to open template: " + err.Error())
	}

	err = template.RenderTemplate(&deliveryData)
	if err != nil {
		return "", errors.New("failed to render template: " + err.Error())
	}
	*/
	log.Println(tdoc.Placeholders())
	tdoc.Params(struct {
		Prov  string
		Date  string
		Adrs  string
		Items []models.ItemTemp
	}{Prov: "provider", Date: deliveryData.Date, Adrs: "address", Items: deliveryData.Items})
	log.Println(tdoc.Placeholders())

	fileName := "delivery_" + time.Now().Format("2006-01-02_15-04-05") + ".docx"
	tdoc.ExportDocx(filepath.Join(".", "files", fileName))

	//err = doc.WriteToFile(filepath.Join(".", "files", fileName))
	/* fileName := "delivery_" + time.Now().Format("2006-01-02_15-04-05") + ".docx"
	err = template.Save(filepath.Join(".", "files", fileName))
	if err != nil {
		return "", errors.New("failed to save file: " + err.Error())
	} */
	return filepath.Join(".", "files", fileName), nil
}

func CreateNewProvider(providerRequest models.Provider) error {
	db, err := database.GetPostgresql()
	if err != nil {
		return errors.New("failed to connect to postgresql: " + err.Error())
	}
	err = db.CreateNewProvider(providerRequest)
	if err != nil {
		return errors.New("failed to create new provider: " + err.Error())
	}
	return nil
}
