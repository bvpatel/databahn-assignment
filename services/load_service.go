package service

import (
	dataSource "databahn-api/data_sources"
	util "databahn-api/utils"
	"errors"
	"fmt"
	"log"
	"time"
)

var wg WaitGroupCount

type LoadService struct {
	dataSource dataSource.DataSource
	maxLimit   int64
}

func NewLoadService(dataSource dataSource.DataSource, maxLimit int64) *LoadService {
	return &LoadService{
		dataSource: dataSource,
		maxLimit:   maxLimit,
	}
}

func (ls *LoadService) LoadData(directoryName, templateFileName string, count int) error {
	filePath := fmt.Sprintf("%s/%s/%s", ".", directoryName, templateFileName)
	log.Printf("Filepath: %s, count: %d", filePath, count)
	templateContent, err := util.ReadTemplateFile(filePath)
	if err != nil {
		log.Printf("Error reading template file: count=%d, template=%s, error=%v\n", count, templateFileName, err)
		return err
	}

	log.Printf("WG count: %d, Limit: %d", wg.GetCount(), ls.maxLimit)
	if wg.GetCount() >= ls.maxLimit {
		return errors.New("server is reached in max capacity, Please try again after sometime")
	}

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			renderedContent, err := util.RenderTemplate(templateContent, nil)
			// log.Printf("Render contents: %s", renderedContent)
			if err != nil {
				log.Printf("Error rendering template: count=%d, template=%s, error=%v\n", count, templateFileName, err)
				return
			}
			time.Sleep(10 * time.Second)
			go func() {
				if err := ls.dataSource.PushData(renderedContent); err != nil {
					log.Printf("Error pushing data to data source: count=%d, template=%s, error=%v\n", count, templateFileName, err)
				}
			}()
		}()
	}

	// go func() {
	// 	wg.Wait()
	// 	close(errChan)

	// 	var errors []error
	// 	for err := range errChan {
	// 		errors = append(errors, err)
	// 	}

	// 	if len(errors) > 0 {
	// 		errMsg := fmt.Sprintf("Encountered %d errors: count=%d, template=%s", len(errors), count, templateFileName)
	// 		log.Println(errMsg)
	// 	}

	// }()

	return nil
}
