package services_bot_repository

import (
	models_bot "backend/internal/models/bot"
)

func (object *botRepositoryServiceImplementation) Add(data *models_bot.BotModel) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	object.data[data.ID] = data
}

func (object *botRepositoryServiceImplementation) Get(id uint) (*models_bot.BotModel, bool) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	botModel, exists := object.data[id]

	return botModel, exists
}

func (object *botRepositoryServiceImplementation) Remove(id uint) {
	object.mutex.Lock()
	defer object.mutex.Unlock()

	delete(object.data, id)
}
