package service

import (
	"gin-restfull-api/data/request"
	"gin-restfull-api/data/response"
	"gin-restfull-api/helper"
	"gin-restfull-api/model"
	"gin-restfull-api/repository"
)

type TagsServiceImpl struct {
	TagsRepository repository.TagsRepository
}

func NewTagsServiceImpl(tagRepository repository.TagsRepository) TagsService {
	return &TagsServiceImpl{
		TagsRepository: tagRepository,
	}
}


func (t *TagsServiceImpl) Create(tags request.CreateTagsRequest) {
	tagModel := model.Tags{
		Name: tags.Name,
	}
	t.TagsRepository.Save(tagModel)
}


func (t *TagsServiceImpl) Delete(tagsId int) {
	t.TagsRepository.Delete(tagsId)
}


func (t *TagsServiceImpl) FindAll() []response.TagsResponse {
	result := t.TagsRepository.FindAll()

	var tags []response.TagsResponse
	for _, value := range result {
		tag := response.TagsResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tags = append(tags, tag)
	}

	return tags
}


func (t *TagsServiceImpl) FindById(tagsId int) response.TagsResponse {
	tagData, err := t.TagsRepository.FindById(tagsId)
	helper.ErrorPanic(err)

	tagResponse := response.TagsResponse{
		Id:   tagData.Id,
		Name: tagData.Name,
	}
	return tagResponse
}


func (t *TagsServiceImpl) Update(tags request.UpdateTagsRequest) {
	tagData, err := t.TagsRepository.FindById(tags.Id)
	helper.ErrorPanic(err)
	tagData.Name = tags.Name
	t.TagsRepository.Update(tagData)
}
