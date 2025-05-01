package services_bot

import (
	models_bot "backend/internal/models/bot"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type apiResponse struct {
	Success bool `json:"success"`
}

func (object *botServiceImplementation) sendApi(botModel *models_bot.BotModel) error {
	var triggerStart float64 = 0
	var stopTime int64 = 0
	var stopPercent float64 = 0

	if botModel.Param.StopTime > 0 {
		stopTime = botModel.Param.StopTime
	}

	if botModel.Param.StopPercent > 0 {
		stopPercent = botModel.Param.StopPercent
	}

	if botModel.Param.PercentIn >= 1 {
		triggerStart = 1
	} else if botModel.Param.PercentIn > 0.5 {
		triggerStart = 0.5
	} else {
		triggerStart = botModel.Param.PercentIn - 0.1

		if triggerStart < 0 {
			triggerStart = 0
		}
	}

	request := models_bot.ApiRequestModel{
		Deposit:        15,
		IsReal:         true,
		Symbol:         botModel.Symbol,
		Window:         botModel.Window,
		Interval:       botModel.Interval,
		TradeDirection: botModel.TradeDirection,
		Bind:           botModel.Param.Bind,
		PercentIn:      botModel.Param.PercentIn,
		PercentOut:     botModel.Param.PercentOut,
		StopTime:       stopTime,
		StopPercent:    stopPercent,
		TriggerStart:   triggerStart,
		LimitQuotes:    1,
		IsCalculator:   true,
	}

	payload, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal ParamModel: %w", err)
	}

	req, err := http.NewRequest("POST", object.configService().GetConfig().Api.Url, bytes.NewBuffer(payload))
	if err != nil {
		return fmt.Errorf("failed to create POST request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send POST request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var response apiResponse

	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("failed to decode API response: %w", err)
	}

	if !response.Success {
		return fmt.Errorf("API returned success = false")
	}

	return nil
}
